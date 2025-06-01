package model

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/rs/zerolog"

	"kostjc/model/mAuth/wcAuth"
	"kostjc/model/mCafe"
	"kostjc/model/mProperty"

	"kostjc/model/mAuth"
)

type Migrator struct {
	OltpConn *Tt.Adapter
	OlapConn *Ch.Adapter
}

func RunMigration(logger *zerolog.Logger, oltpConn *Tt.Adapter, olapConn *Ch.Adapter) {
	Tt.DEBUG = true
	Ch.DEBUG = true
	L.Print(`run migration..`)
	m := Migrator{
		OltpConn: oltpConn,
		OlapConn: olapConn,
	}
	mAuth.TarantoolTables[mAuth.TableUsers].PreUnique1MigrationHook = wcAuth.UniqueUsernameMigration

	m.OltpConn.MigrateTables(mAuth.TarantoolTables)
	m.OlapConn.MigrateTables(mAuth.ClickhouseTables)
	m.OltpConn.MigrateTables(mProperty.TarantoolTables)
	m.OlapConn.MigrateTables(mProperty.ClickhouseTables)
	m.OltpConn.MigrateTables(mCafe.TarantoolTables)
}

// VerifyTables function to check whether tables are there or not
// go run main.go migrate
func VerifyTables(
	oltpConn *Tt.Adapter,
	olapConn *Ch.Adapter,
) {
	Tt.CheckTarantoolTables(oltpConn, mAuth.TarantoolTables)
	Ch.CheckClickhouseTables(olapConn, mAuth.ClickhouseTables)
	Tt.CheckTarantoolTables(oltpConn, mProperty.TarantoolTables)
	Ch.CheckClickhouseTables(olapConn, mProperty.ClickhouseTables)
	Tt.CheckTarantoolTables(oltpConn, mCafe.TarantoolTables)
}
