package saCafe

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

import (
	"database/sql"
	"time"

	"kostjc/model/mCafe"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	chBuffer "github.com/kokizzu/ch-timed-buffer"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/L"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file saCafe__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type saCafe__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type saCafe__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type saCafe__ORM.GEN.go
// go:generate msgp -tests=false -file saCafe__ORM.GEN.go -o saCafe__MSG.GEN.go

var menuLogsDummy = MenuLogs{}
var saleLogsDummy = SaleLogs{}
var Preparators = map[Ch.TableName]chBuffer.Preparator{
	mCafe.TableMenuLogs: func(tx *sql.Tx) *sql.Stmt {
		query := menuLogsDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
	mCafe.TableSaleLogs: func(tx *sql.Tx) *sql.Stmt {
		query := saleLogsDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
}

type MenuLogs struct {
	Adapter    *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	CreatedAt  time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	ActorId    uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	BeforeJson string      `json:"beforeJson" form:"beforeJson" query:"beforeJson" long:"beforeJson" msg:"beforeJson"`
	AfterJson  string      `json:"afterJson" form:"afterJson" query:"afterJson" long:"afterJson" msg:"afterJson"`
}

func NewMenuLogs(adapter *Ch.Adapter) *MenuLogs {
	return &MenuLogs{Adapter: adapter}
}

// MenuLogsFieldTypeMap returns key value of field name and key
var MenuLogsFieldTypeMap = map[string]Ch.DataType{ //nolint:dupl false positive
	`createdAt`:  Ch.DateTime,
	`actorId`:    Ch.UInt64,
	`beforeJson`: Ch.String,
	`afterJson`:  Ch.String,
}

func (m *MenuLogs) TableName() Ch.TableName { //nolint:dupl false positive
	return mCafe.TableMenuLogs
}

func (m *MenuLogs) SqlTableName() string { //nolint:dupl false positive
	return `"menuLogs"`
}

func (m *MenuLogs) ScanRowAllCols(rows *sql.Rows) (err error) { //nolint:dupl false positive
	return rows.Scan(
		&m.CreatedAt,
		&m.ActorId,
		&m.BeforeJson,
		&m.AfterJson,
	)
}

func (m *MenuLogs) ScanRowsAllCols(rows *sql.Rows, estimateRows int) (res []MenuLogs, err error) { //nolint:dupl false positive
	res = make([]MenuLogs, 0, estimateRows)
	defer rows.Close()
	for rows.Next() {
		var row MenuLogs
		err = row.ScanRowAllCols(rows)
		if err != nil {
			return
		}
		res = append(res, row)
	}
	return
}

// insert, error if exists
func (m *MenuLogs) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + m.SqlTableName() + `(` + m.SqlAllFields() + `) VALUES (?,?,?,?)`
}

func (m *MenuLogs) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + m.SqlTableName()
}

func (m *MenuLogs) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, actorId
	, beforeJson
	, afterJson
	`
}

func (m *MenuLogs) SqlAllFields() string { //nolint:dupl false positive
	return `createdAt, actorId, beforeJson, afterJson`
}

func (m MenuLogs) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		m.CreatedAt,  // 0
		m.ActorId,    // 1
		m.BeforeJson, // 2
		m.AfterJson,  // 3
	}
}

func (m *MenuLogs) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (m *MenuLogs) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (m *MenuLogs) IdxActorId() int { //nolint:dupl false positive
	return 1
}

func (m *MenuLogs) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (m *MenuLogs) IdxBeforeJson() int { //nolint:dupl false positive
	return 2
}

func (m *MenuLogs) SqlBeforeJson() string { //nolint:dupl false positive
	return `beforeJson`
}

func (m *MenuLogs) IdxAfterJson() int { //nolint:dupl false positive
	return 3
}

func (m *MenuLogs) SqlAfterJson() string { //nolint:dupl false positive
	return `afterJson`
}

func (m *MenuLogs) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		m.CreatedAt,  // 0
		m.ActorId,    // 1
		m.BeforeJson, // 2
		m.AfterJson,  // 3
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

type SaleLogs struct {
	Adapter    *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	CreatedAt  time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	ActorId    uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	BeforeJson string      `json:"beforeJson" form:"beforeJson" query:"beforeJson" long:"beforeJson" msg:"beforeJson"`
	AfterJson  string      `json:"afterJson" form:"afterJson" query:"afterJson" long:"afterJson" msg:"afterJson"`
}

func NewSaleLogs(adapter *Ch.Adapter) *SaleLogs {
	return &SaleLogs{Adapter: adapter}
}

// SaleLogsFieldTypeMap returns key value of field name and key
var SaleLogsFieldTypeMap = map[string]Ch.DataType{ //nolint:dupl false positive
	`createdAt`:  Ch.DateTime,
	`actorId`:    Ch.UInt64,
	`beforeJson`: Ch.String,
	`afterJson`:  Ch.String,
}

func (s *SaleLogs) TableName() Ch.TableName { //nolint:dupl false positive
	return mCafe.TableSaleLogs
}

func (s *SaleLogs) SqlTableName() string { //nolint:dupl false positive
	return `"saleLogs"`
}

func (s *SaleLogs) ScanRowAllCols(rows *sql.Rows) (err error) { //nolint:dupl false positive
	return rows.Scan(
		&s.CreatedAt,
		&s.ActorId,
		&s.BeforeJson,
		&s.AfterJson,
	)
}

func (s *SaleLogs) ScanRowsAllCols(rows *sql.Rows, estimateRows int) (res []SaleLogs, err error) { //nolint:dupl false positive
	res = make([]SaleLogs, 0, estimateRows)
	defer rows.Close()
	for rows.Next() {
		var row SaleLogs
		err = row.ScanRowAllCols(rows)
		if err != nil {
			return
		}
		res = append(res, row)
	}
	return
}

// insert, error if exists
func (s *SaleLogs) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + s.SqlTableName() + `(` + s.SqlAllFields() + `) VALUES (?,?,?,?)`
}

func (s *SaleLogs) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + s.SqlTableName()
}

func (s *SaleLogs) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, actorId
	, beforeJson
	, afterJson
	`
}

func (s *SaleLogs) SqlAllFields() string { //nolint:dupl false positive
	return `createdAt, actorId, beforeJson, afterJson`
}

func (s SaleLogs) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		s.CreatedAt,  // 0
		s.ActorId,    // 1
		s.BeforeJson, // 2
		s.AfterJson,  // 3
	}
}

func (s *SaleLogs) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (s *SaleLogs) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (s *SaleLogs) IdxActorId() int { //nolint:dupl false positive
	return 1
}

func (s *SaleLogs) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (s *SaleLogs) IdxBeforeJson() int { //nolint:dupl false positive
	return 2
}

func (s *SaleLogs) SqlBeforeJson() string { //nolint:dupl false positive
	return `beforeJson`
}

func (s *SaleLogs) IdxAfterJson() int { //nolint:dupl false positive
	return 3
}

func (s *SaleLogs) SqlAfterJson() string { //nolint:dupl false positive
	return `afterJson`
}

func (s *SaleLogs) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		s.CreatedAt,  // 0
		s.ActorId,    // 1
		s.BeforeJson, // 2
		s.AfterJson,  // 3
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go
