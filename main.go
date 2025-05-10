package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/fatih/color"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/rs/zerolog"
	"golang.org/x/sync/errgroup"

	"kostjc/conf"
	"kostjc/domain"
	"kostjc/model"
	"kostjc/model/xMailer"
	"kostjc/presentation"
)

var VERSION = ``
var log *zerolog.Logger

func main() {
	conf.VERSION = VERSION

	// note: set instance id when there's multiple instance
	// lexid.Config.Separator = `~THE_INSTANCE_ID`

	fmt.Println(color.GreenString(conf.PROJECT_NAME + ` ` + S.IfEmpty(VERSION, `local-dev`)))

	log = conf.InitLogger()
	conf.LoadEnv()

	args := os.Args
	if len(args) < 2 {
		L.Print(`must start with: run, web, cron, migrate, or config as first argument`)
		L.Print(args)
		return
	}

	if args[1] == `config` {
		L.Describe(M.SX{
			`web`:        conf.EnvWebConf(),
			`tarantool`:  conf.EnvTarantool(),
			`clickhouse`: conf.EnvClickhouse(),
			`mailhog`:    conf.EnvMailhog(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	eg, _ := errgroup.WithContext(ctx)
	var closers []func() error

	// mailer
	var mailer xMailer.Mailer
	eg.Go(func() error {
		mailer.Conf = conf.EnvMailer()
		switch mailer.Conf.DefaultMailer {
		case `dockermailserver`:
			dms := xMailer.Dockermailserver{DockermailserverConf: conf.EnvDockermailserver()}
			L.PanicIf(dms.Connect(), `Dockermailserver.Connect`)
			mailer.SendMailFunc = dms.SendEmail
		case `mailhog`:
			mh, err := xMailer.NewMailhog(conf.EnvMailhog())
			L.PanicIf(err, `NewMailhog`)
			mailer.SendMailFunc = mh.SendEmail
		default:
			L.Panic(`unknown DefaultMailer`)
		}
		return nil
	})

	var err error

	// connect to tarantool
	var tConn *Tt.Adapter
	eg.Go(func() error {
		tConf := conf.EnvTarantool()
		tConn, err = tConf.Connect()
		if tConn != nil {
			closers = append(closers, tConn.Close)
			fmt.Println(color.BlueString("[CONNECTED]") + ` Tarantool: ` + tConf.DebugStr())
		}
		return err
	})

	// connect to clickhouse
	var cConn *Ch.Adapter
	eg.Go(func() error {
		cConf := conf.EnvClickhouse()
		cConn, err = cConf.Connect()
		if cConn != nil {
			closers = append(closers, cConn.Close)
			fmt.Println(color.BlueString("[CONNECTED]") + ` ClickHouse: ` + cConf.DebugStr())
		}
		return err
	})

	L.PanicIf(eg.Wait(), `eg.Wait`) // if error, make sure no error on: docker compose up
	for _, closer := range closers {
		closer := closer
		defer closer()
	}

	// create domain object
	d := &domain.Domain{
		AuthOltp: tConn,
		AuthOlap: cConn,
		PropOltp: tConn,
		Mailer:   mailer,
		IsBgSvc:  false,
		Log:      log,

		Superadmins: conf.EnvSuperAdmins(),

		WebCfg: conf.EnvWebConf(),
	}
	d.InitTimedBuffer()
	defer d.CloseTimedBuffer()

	mode := S.ToLower(os.Args[1])

	// check table existence
	if mode != `migrate` {
		L.Print(`verifying table schema, if failed, run: go run main.go migrate`)
		model.VerifyTables(tConn, cConn, tConn)
	}

	// start
	switch mode {
	case `web`:
		ws := &presentation.WebServer{
			Domain: d,
		}
		ws.Start(log)
	case `cli`:
		cli := &presentation.CLI{
			Domain: d,
		}
		cli.Run(os.Args[2:], log)
	case `cron`:
		cron := &presentation.Cron{
			Domain: d,
		}
		cron.Start(log)
	case `migrate`:
		model.RunMigration(log, tConn, cConn)
	case `backup`:
		model.BackupDatabase(tConn)
	case `restore`:
		if len(os.Args) < 3 {
			L.LOG.Error(`must start with: restore <datetime>`)
			return
		}
		dateTime := os.Args[2]
		fmt.Println(color.BlueString(`Restoring database from: ` + dateTime))
		model.RestoreDatabase(tConn, dateTime)
	case `truncate`:
		model.TruncateDatabase(tConn)
	default:
		log.Error().Str(`mode`, mode).Msg(`unknown mode`)
	}

}
