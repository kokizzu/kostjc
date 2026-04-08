package domain

import (
	"context"
	"database/sql"
	"kostjc/conf"
	"kostjc/model"
	"kostjc/model/mAuth"
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mAuth/wcAuth"
	"kostjc/model/xMailer"
	"os"
	"testing"

	"github.com/kokizzu/gotro/D"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/lexid"
	"github.com/kpango/fastime"
	"github.com/ory/dockertest/v3"
	"github.com/tarantool/go-tarantool"
	"golang.org/x/sync/errgroup"
)

// create dockertest instance

var testTt *Tt.Adapter
var testCh *Ch.Adapter
var testMailer xMailer.Mailer
var testTime = fastime.Now()
var testSuperAdminSessionToken string
var testStaffSessionToken string
var testUserSessionToken string
var testAdmin *rqAuth.Users
var testStaff *rqAuth.Users
var testUser *rqAuth.Users

const (
	testSuperAdminEmail    = `admin@localhost`
	testSuperAdminUserName = `admin1`
	testStaffEmail         = `staff@localhost`
	testStaffUserName      = `staff1`
	testUserEmail          = `user@localhost`
	testUserUserName       = `user1`
)

func TestMain(m *testing.M) {
	if os.Getenv(`USE_COMPOSE`) != `` {
		// use local docker compose
		conf.LoadEnv()
		testMailer = xMailer.Mailer{
			SendMailFunc: func(toEmailName map[string]string, subject, text, html string) error {
				return nil
			},
		}

		var err error
		eg := errgroup.Group{}
		eg.Go(func() error {
			chConf := conf.EnvClickhouse()
			testCh, err = chConf.Connect()
			return err
		})
		eg.Go(func() error {
			ttConf := conf.EnvTarantool()
			testTt, err = ttConf.Connect()
			return err
		})
		err = eg.Wait()
		L.PanicIf(err, `eg.Wait`)

	} else {
		// setup dockertest instance
		dockerPool := D.InitDockerTest("")
		defer dockerPool.Cleanup()
		testMailer = xMailer.Mailer{
			SendMailFunc: func(toEmailName map[string]string, subject, text, html string) error {
				return nil
			},
		}

		eg := errgroup.Group{}

		// attach tarantool
		eg.Go(func() error {
			tdt := &Tt.TtDockerTest{
				User:     "testT",
				Password: "passT",
			}
			img := tdt.ImageVersion(dockerPool, ``)
			dockerPool.Spawn(img, func(res *dockertest.Resource) error {
				t, err := tdt.ConnectCheck(res)
				if err != nil {
					return err
				}
				testTt = &Tt.Adapter{
					Connection: t,
					Reconnect: func() *tarantool.Connection {
						t, err := tdt.ConnectCheck(res)
						L.IsError(err, `tdt.ConnectCheck`)
						return t
					},
				}
				return nil
			})
			return nil
		})

		// attach clickhouse
		eg.Go(func() error {
			cdt := &Ch.ChDockerTest{
				User:     "testC",
				Password: "passC",
				Database: "default",
			}
			img := cdt.ImageLatest(dockerPool)
			dockerPool.Spawn(img, func(res *dockertest.Resource) error {
				c, err := cdt.ConnectCheck(res)
				if err != nil {
					return err
				}
				testCh = &Ch.Adapter{
					DB: c,
					Reconnect: func() *sql.DB {
						c, err := cdt.ConnectCheck(res)
						L.IsError(err, `cdt.ConnectCheck`)
						return c
					},
				}
				return nil
			})
			return nil
		})

		err := eg.Wait()
		L.PanicIf(err, `eg.Wait`)
	}

	// run migration
	model.RunMigration(nil, testTt, testCh)

	// run tests
	m.Run()

	// teardown dockertest instance
}

func testDomain() (*Domain, func()) {
	log := conf.InitLogger()

	d := &Domain{
		AuthOltp: testTt,
		AuthOlap: testCh,
		PropOltp: testTt,
		PropOlap: testCh,

		Mailer:  xMailer.Mailer{SendMailFunc: testMailer.SendMailFunc},
		IsBgSvc: false,

		Log: log,

		Superadmins: M.SB{testSuperAdminEmail: true},
	}
	d.InitTimedBuffer()

	// create admin
	admin := wcAuth.NewUsersMutator(testTt)
	admin.Email = testSuperAdminEmail
	admin.UserName = testSuperAdminUserName
	admin.Role = mAuth.RoleAdmin
	if !admin.FindByEmail() {
		admin.DoInsert()
	}
	testAdmin = &admin.Users
	testAdmin.Adapter = nil // prevent modification

	// create session
	session := wcAuth.NewSessionsMutator(testTt)
	session.UserId = admin.Id
	sess := &Session{
		UserId:    admin.Id,
		ExpiredAt: testTime.AddDate(0, 0, conf.CookieDays).Unix(),
		Email:     admin.Email,
	}
	testSuperAdminSessionToken = sess.Encrypt(``) // empty user agent to simplify testing
	session.SessionToken = testSuperAdminSessionToken
	session.ExpiredAt = sess.ExpiredAt
	if !session.FindBySessionToken() {
		session.DoInsert()
	}

	staff := wcAuth.NewUsersMutator(testTt)
	staff.Email = testStaffEmail
	staff.UserName = testStaffUserName
	staff.Role = mAuth.RoleStaff
	if !staff.FindByEmail() {
		staff.DoInsert()
	}
	testStaff = &staff.Users
	testStaff.Adapter = nil

	staffSession := wcAuth.NewSessionsMutator(testTt)
	staffSession.UserId = staff.Id
	staffSess := &Session{
		UserId:    staff.Id,
		ExpiredAt: testTime.AddDate(0, 0, conf.CookieDays).Unix(),
		Email:     staff.Email,
	}
	testStaffSessionToken = staffSess.Encrypt(``)
	staffSession.SessionToken = testStaffSessionToken
	staffSession.ExpiredAt = staffSess.ExpiredAt
	if !staffSession.FindBySessionToken() {
		staffSession.DoInsert()
	}

	user := wcAuth.NewUsersMutator(testTt)
	user.Email = testUserEmail
	user.UserName = testUserUserName
	user.Role = mAuth.RoleUser
	if !user.FindByEmail() {
		user.DoInsert()
	}
	testUser = &user.Users
	testUser.Adapter = nil

	userSession := wcAuth.NewSessionsMutator(testTt)
	userSession.UserId = user.Id
	userSess := &Session{
		UserId:    user.Id,
		ExpiredAt: testTime.AddDate(0, 0, conf.CookieDays).Unix(),
		Email:     user.Email,
	}
	testUserSessionToken = userSess.Encrypt(``)
	userSession.SessionToken = testUserSessionToken
	userSession.ExpiredAt = userSess.ExpiredAt
	if !userSession.FindBySessionToken() {
		userSession.DoInsert()
	}

	return d, func() {
		go d.authLogs.Close()
		d.WaitTimedBufferFinalFlush()
	}
}

func testAdminRequestCommon(action string) RequestCommon {
	return RequestCommon{
		TracerContext: context.Background(),
		RequestId:     lexid.ID(),
		SessionToken:  testSuperAdminSessionToken,
		UserAgent:     "",
		IpAddress:     "127.0.2.1",
		Debug:         true,
		Host:          "localhost:1234",
		Action:        action,
		Lat:           -1,
		Long:          -2,
		now:           testTime.Unix(),
		start:         testTime,
	}
}

func testStaffRequestCommon(action string) RequestCommon {
	return RequestCommon{
		TracerContext: context.Background(),
		RequestId:     lexid.ID(),
		SessionToken:  testStaffSessionToken,
		UserAgent:     "",
		IpAddress:     "127.0.2.2",
		Debug:         true,
		Host:          "localhost:1234",
		Action:        action,
		Lat:           -1,
		Long:          -2,
		now:           testTime.Unix(),
		start:         testTime,
	}
}

func testUserRequestCommon(action string) RequestCommon {
	return RequestCommon{
		TracerContext: context.Background(),
		RequestId:     lexid.ID(),
		SessionToken:  testUserSessionToken,
		UserAgent:     "",
		IpAddress:     "127.0.2.3",
		Debug:         true,
		Host:          "localhost:1234",
		Action:        action,
		Lat:           -1,
		Long:          -2,
		now:           testTime.Unix(),
		start:         testTime,
	}
}
