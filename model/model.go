package model

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/rs/zerolog"

	"kostjc/model/mAuth/wcAuth"

	"kostjc/model/mAuth"
)

type Migrator struct {
	AuthOltp     *Tt.Adapter
	AuthOlap     *Ch.Adapter
	PropOltp     *Tt.Adapter
	PropOlap     *Ch.Adapter
	BusinessOltp *Tt.Adapter
	StorOltp     *Tt.Adapter
}

func RunMigration(logger *zerolog.Logger, authOltp *Tt.Adapter, authOlap *Ch.Adapter) {
	Tt.DEBUG = true
	Ch.DEBUG = true
	L.Print(`run migration..`)
	m := Migrator{
		AuthOltp: authOltp,
		AuthOlap: authOlap,
	}
	mAuth.TarantoolTables[mAuth.TableUsers].PreUnique1MigrationHook = wcAuth.UniqueUsernameMigration
	m.AuthOltp.MigrateTables(mAuth.TarantoolTables)
	m.AuthOlap.MigrateTables(mAuth.ClickhouseTables)
}

// VerifyTables function to check whether tables are there or not
// go run main.go migrate
func VerifyTables(
	authOltp *Tt.Adapter,
	authOlap *Ch.Adapter,
) {
	Ch.CheckClickhouseTables(authOlap, mAuth.ClickhouseTables)
	Tt.CheckTarantoolTables(authOltp, mAuth.TarantoolTables)
}
