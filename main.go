package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"kostjc/model/mAuth/wcAuth"
	"kostjc/model/mProperty/rqProperty"

	"github.com/fatih/color"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
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

type roomExport struct {
	Name        string
	ImageURL    string
	Size        string
	AvailableAt string
	NormalPrice int64
}

func formatJSInt64(v int64) string {
	s := fmt.Sprintf("%d", v)
	if len(s) <= 3 {
		return s
	}

	var b strings.Builder
	prefixLen := len(s) % 3
	if prefixLen > 0 {
		b.WriteString(s[:prefixLen])
		if len(s) > prefixLen {
			b.WriteByte('_')
		}
	}
	for i := prefixLen; i < len(s); i += 3 {
		b.WriteString(s[i : i+3])
		if i+3 < len(s) {
			b.WriteByte('_')
		}
	}
	return b.String()
}

func loadExistingRoomImageURLs(roomsJSPath string) map[string]string {
	content, err := os.ReadFile(roomsJSPath)
	if err != nil {
		return map[string]string{}
	}

	res := map[string]string{}
	var roomName string
	for _, line := range strings.Split(string(content), "\n") {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, `name: "`) {
			roomName = strings.TrimSuffix(strings.TrimPrefix(trimmed, `name: "`), `",`)
			roomName = strings.ToUpper(strings.TrimSpace(roomName))
			continue
		}
		if roomName == `` || !strings.HasPrefix(trimmed, `image_url: "`) {
			continue
		}

		imageURL := strings.TrimSuffix(strings.TrimPrefix(trimmed, `image_url: "`), `",`)
		imageURL = strings.TrimSpace(imageURL)
		if imageURL != `` {
			res[roomName] = imageURL
		}
		roomName = ``
	}
	return res
}

func loadPreferredRoomImageURLs(outputFile string) map[string]string {
	imageURLs := loadExistingRoomImageURLs(filepath.Clean(`../benalu.dev/src/rooms.js`))
	if len(imageURLs) == 0 {
		return loadExistingRoomImageURLs(outputFile)
	}

	for roomName, imageURL := range loadExistingRoomImageURLs(outputFile) {
		if _, ok := imageURLs[roomName]; !ok {
			imageURLs[roomName] = imageURL
		}
	}
	return imageURLs
}

func generateRoomsJS(tConn *Tt.Adapter, outputDir string) error {
	const query = `
SELECT
	r."roomName",
	r."basePriceIDR",
	r."roomSize",
	r."imageUrl",
	COALESCE(MAX(CASE WHEN b."deletedAt" = 0 THEN b."dateEnd" END), '') AS "availableAt"
FROM "rooms" r
LEFT JOIN "bookings" b ON b."roomId" = r."id"
WHERE r."deletedAt" = 0
GROUP BY r."id"
ORDER BY r."roomName" ASC`

	groupedRooms := map[string][]roomExport{}
	for _, building := range []string{`A`, `B`, `C`, `D`, `E`, `F`, `G`, `H`, `I`, `J`} {
		groupedRooms[building] = []roomExport{}
	}

	outputFile := filepath.Join(outputDir, `rooms.js`)
	existingImageURLs := loadPreferredRoomImageURLs(outputFile)

	rooms := rqProperty.NewRooms(tConn)
	today := time.Now().Format(time.DateOnly)
	rooms.Adapter.QuerySql(query, func(row []any) {
		if len(row) != 5 {
			return
		}

		roomName := strings.ToUpper(strings.TrimSpace(X.ToS(row[0])))
		if roomName == `` {
			return
		}
		building := roomName[:1]
		if _, ok := groupedRooms[building]; !ok {
			groupedRooms[building] = []roomExport{}
		}

		size := strings.TrimSpace(X.ToS(row[2]))
		if size == `` {
			size = `3x4`
		}

		imageURL := strings.TrimSpace(existingImageURLs[roomName])
		if imageURL == `` {
			imageURL = strings.TrimSpace(X.ToS(row[3]))
		}
		if imageURL == `` {
			imageURL = `/placeholder-image.webp`
		}

		availableAt := strings.TrimSpace(X.ToS(row[4]))
		if availableAt == `` {
			availableAt = today
		}

		groupedRooms[building] = append(groupedRooms[building], roomExport{
			Name:        roomName,
			ImageURL:    imageURL,
			Size:        size,
			AvailableAt: availableAt,
			NormalPrice: X.ToI(row[1]),
		})
	})

	var b strings.Builder
	for idx, building := range []string{`A`, `B`, `C`, `D`, `E`, `F`, `G`, `H`, `I`, `J`} {
		rooms := groupedRooms[building]
		sort.Slice(rooms, func(i, j int) bool {
			return rooms[i].Name < rooms[j].Name
		})

		b.WriteString(`export const rooms`)
		b.WriteString(building)
		b.WriteString(` = [` + "\n")
		for _, room := range rooms {
			b.WriteString("  {\n")
			b.WriteString(fmt.Sprintf("    name: %q,\n", room.Name))
			b.WriteString(fmt.Sprintf("    image_url: %q,\n", room.ImageURL))
			b.WriteString(fmt.Sprintf("    size: %q,\n", room.Size))
			b.WriteString(fmt.Sprintf("    availableAt: %q,\n", room.AvailableAt))
			b.WriteString(`    normalPrice: ` + formatJSInt64(room.NormalPrice) + ",\n")
			b.WriteString("    facilities: [],\n")
			b.WriteString("  },\n")
		}
		b.WriteString("]\n")
		if idx < 9 {
			b.WriteString("\n")
		}
	}

	if err := os.MkdirAll(outputDir, 0o755); err != nil {
		return err
	}
	return os.WriteFile(outputFile, []byte(b.String()), 0o644)
}

func main() {
	conf.VERSION = VERSION

	// note: set instance id when there's multiple instance
	// lexid.Config.Separator = `~THE_INSTANCE_ID`

	fmt.Println(color.GreenString(conf.PROJECT_NAME + ` ` + S.IfEmpty(VERSION, `local-dev`)))

	log = conf.InitLogger()
	conf.LoadEnv()

	args := os.Args
	if len(args) < 2 {
		L.Print(`must start with: web, cli, cron, migrate, config, gen-rooms.js, or other supported mode as first argument`)
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
		PropOlap: cConn,
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
		model.VerifyTables(tConn, cConn)
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

		fmt.Println(color.BlueString(`Truncate database`))
		model.TruncateDatabase(tConn)

		fmt.Println(color.BlueString(`Restoring database from: ` + dateTime))
		model.RestoreDatabase(tConn, dateTime)
	case `truncate`:
		model.TruncateDatabase(tConn)
	case `resetpass`:
		if len(os.Args) < 4 {
			L.LOG.Error(`must have argument: email and password`)
			return
		}
		usr := wcAuth.NewUsersMutator(tConn)
		usr.Email = os.Args[2]
		if !usr.FindByEmail() {
			L.LOG.Error(`cannot find user with email: ` + os.Args[2])
			return
		}
		usr.SetEncryptedPassword(os.Args[3], time.Now().Unix())
		if !usr.DoUpdateById() {
			L.LOG.Error(`cannot update password for email: ` + os.Args[2])
			return
		}
		L.LOG.Infof(`reset password success`)
	case `gen-rooms.js`:
		outputDir := `../benalu.dev/src/`
		if len(os.Args) >= 3 && strings.TrimSpace(os.Args[2]) != `` {
			outputDir = os.Args[2]
		}
		if err := generateRoomsJS(tConn, outputDir); err != nil {
			L.LOG.Error(`failed generating rooms.js: `, err, ` outputDir: `, outputDir)
			return
		}
		L.LOG.Info(`generated rooms.js: `, filepath.Join(outputDir, `rooms.js`))
	default:
		log.Error().Str(`mode`, mode).Msg(`unknown mode`)
	}

}
