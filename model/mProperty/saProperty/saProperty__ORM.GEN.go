package saProperty

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

import (
	"database/sql"
	"time"

	"kostjc/model/mProperty"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	chBuffer "github.com/kokizzu/ch-timed-buffer"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/L"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file saProperty__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type saProperty__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type saProperty__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type saProperty__ORM.GEN.go
// go:generate msgp -tests=false -file saProperty__ORM.GEN.go -o saProperty__MSG.GEN.go

var bookingLogsDummy = BookingLogs{}
var buildingLogsDummy = BuildingLogs{}
var facilityLogsDummy = FacilityLogs{}
var locationLogsDummy = LocationLogs{}
var paymentLogsDummy = PaymentLogs{}
var roomLogsDummy = RoomLogs{}
var stockLogsDummy = StockLogs{}
var Preparators = map[Ch.TableName]chBuffer.Preparator{
	mProperty.TableBookingLogs: func(tx *sql.Tx) *sql.Stmt {
		query := bookingLogsDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
	mProperty.TableBuildingLogs: func(tx *sql.Tx) *sql.Stmt {
		query := buildingLogsDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
	mProperty.TableFacilityLogs: func(tx *sql.Tx) *sql.Stmt {
		query := facilityLogsDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
	mProperty.TableLocationLogs: func(tx *sql.Tx) *sql.Stmt {
		query := locationLogsDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
	mProperty.TablePaymentLogs: func(tx *sql.Tx) *sql.Stmt {
		query := paymentLogsDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
	mProperty.TableRoomLogs: func(tx *sql.Tx) *sql.Stmt {
		query := roomLogsDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
	mProperty.TableStockLogs: func(tx *sql.Tx) *sql.Stmt {
		query := stockLogsDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
}

type BookingLogs struct {
	Adapter    *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	CreatedAt  time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	ActorId    uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	BeforeJson string      `json:"beforeJson" form:"beforeJson" query:"beforeJson" long:"beforeJson" msg:"beforeJson"`
	AfterJson  string      `json:"afterJson" form:"afterJson" query:"afterJson" long:"afterJson" msg:"afterJson"`
}

func NewBookingLogs(adapter *Ch.Adapter) *BookingLogs {
	return &BookingLogs{Adapter: adapter}
}

// BookingLogsFieldTypeMap returns key value of field name and key
var BookingLogsFieldTypeMap = map[string]Ch.DataType{ //nolint:dupl false positive
	`createdAt`:  Ch.DateTime,
	`actorId`:    Ch.UInt64,
	`beforeJson`: Ch.String,
	`afterJson`:  Ch.String,
}

func (b *BookingLogs) TableName() Ch.TableName { //nolint:dupl false positive
	return mProperty.TableBookingLogs
}

func (b *BookingLogs) SqlTableName() string { //nolint:dupl false positive
	return `"bookingLogs"`
}

func (b *BookingLogs) ScanRowAllCols(rows *sql.Rows) (err error) { //nolint:dupl false positive
	return rows.Scan(
		&b.CreatedAt,
		&b.ActorId,
		&b.BeforeJson,
		&b.AfterJson,
	)
}

func (b *BookingLogs) ScanRowsAllCols(rows *sql.Rows, estimateRows int) (res []BookingLogs, err error) { //nolint:dupl false positive
	res = make([]BookingLogs, 0, estimateRows)
	defer rows.Close()
	for rows.Next() {
		var row BookingLogs
		err = row.ScanRowAllCols(rows)
		if err != nil {
			return
		}
		res = append(res, row)
	}
	return
}

// insert, error if exists
func (b *BookingLogs) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + b.SqlTableName() + `(` + b.SqlAllFields() + `) VALUES (?,?,?,?)`
}

func (b *BookingLogs) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + b.SqlTableName()
}

func (b *BookingLogs) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, actorId
	, beforeJson
	, afterJson
	`
}

func (b *BookingLogs) SqlAllFields() string { //nolint:dupl false positive
	return `createdAt, actorId, beforeJson, afterJson`
}

func (b BookingLogs) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		b.CreatedAt,  // 0
		b.ActorId,    // 1
		b.BeforeJson, // 2
		b.AfterJson,  // 3
	}
}

func (b *BookingLogs) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (b *BookingLogs) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (b *BookingLogs) IdxActorId() int { //nolint:dupl false positive
	return 1
}

func (b *BookingLogs) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (b *BookingLogs) IdxBeforeJson() int { //nolint:dupl false positive
	return 2
}

func (b *BookingLogs) SqlBeforeJson() string { //nolint:dupl false positive
	return `beforeJson`
}

func (b *BookingLogs) IdxAfterJson() int { //nolint:dupl false positive
	return 3
}

func (b *BookingLogs) SqlAfterJson() string { //nolint:dupl false positive
	return `afterJson`
}

func (b *BookingLogs) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		b.CreatedAt,  // 0
		b.ActorId,    // 1
		b.BeforeJson, // 2
		b.AfterJson,  // 3
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

type BuildingLogs struct {
	Adapter    *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	CreatedAt  time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	ActorId    uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	BeforeJson string      `json:"beforeJson" form:"beforeJson" query:"beforeJson" long:"beforeJson" msg:"beforeJson"`
	AfterJson  string      `json:"afterJson" form:"afterJson" query:"afterJson" long:"afterJson" msg:"afterJson"`
}

func NewBuildingLogs(adapter *Ch.Adapter) *BuildingLogs {
	return &BuildingLogs{Adapter: adapter}
}

// BuildingLogsFieldTypeMap returns key value of field name and key
var BuildingLogsFieldTypeMap = map[string]Ch.DataType{ //nolint:dupl false positive
	`createdAt`:  Ch.DateTime,
	`actorId`:    Ch.UInt64,
	`beforeJson`: Ch.String,
	`afterJson`:  Ch.String,
}

func (b *BuildingLogs) TableName() Ch.TableName { //nolint:dupl false positive
	return mProperty.TableBuildingLogs
}

func (b *BuildingLogs) SqlTableName() string { //nolint:dupl false positive
	return `"buildingLogs"`
}

func (b *BuildingLogs) ScanRowAllCols(rows *sql.Rows) (err error) { //nolint:dupl false positive
	return rows.Scan(
		&b.CreatedAt,
		&b.ActorId,
		&b.BeforeJson,
		&b.AfterJson,
	)
}

func (b *BuildingLogs) ScanRowsAllCols(rows *sql.Rows, estimateRows int) (res []BuildingLogs, err error) { //nolint:dupl false positive
	res = make([]BuildingLogs, 0, estimateRows)
	defer rows.Close()
	for rows.Next() {
		var row BuildingLogs
		err = row.ScanRowAllCols(rows)
		if err != nil {
			return
		}
		res = append(res, row)
	}
	return
}

// insert, error if exists
func (b *BuildingLogs) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + b.SqlTableName() + `(` + b.SqlAllFields() + `) VALUES (?,?,?,?)`
}

func (b *BuildingLogs) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + b.SqlTableName()
}

func (b *BuildingLogs) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, actorId
	, beforeJson
	, afterJson
	`
}

func (b *BuildingLogs) SqlAllFields() string { //nolint:dupl false positive
	return `createdAt, actorId, beforeJson, afterJson`
}

func (b BuildingLogs) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		b.CreatedAt,  // 0
		b.ActorId,    // 1
		b.BeforeJson, // 2
		b.AfterJson,  // 3
	}
}

func (b *BuildingLogs) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (b *BuildingLogs) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (b *BuildingLogs) IdxActorId() int { //nolint:dupl false positive
	return 1
}

func (b *BuildingLogs) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (b *BuildingLogs) IdxBeforeJson() int { //nolint:dupl false positive
	return 2
}

func (b *BuildingLogs) SqlBeforeJson() string { //nolint:dupl false positive
	return `beforeJson`
}

func (b *BuildingLogs) IdxAfterJson() int { //nolint:dupl false positive
	return 3
}

func (b *BuildingLogs) SqlAfterJson() string { //nolint:dupl false positive
	return `afterJson`
}

func (b *BuildingLogs) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		b.CreatedAt,  // 0
		b.ActorId,    // 1
		b.BeforeJson, // 2
		b.AfterJson,  // 3
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

type FacilityLogs struct {
	Adapter    *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	CreatedAt  time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	ActorId    uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	BeforeJson string      `json:"beforeJson" form:"beforeJson" query:"beforeJson" long:"beforeJson" msg:"beforeJson"`
	AfterJson  string      `json:"afterJson" form:"afterJson" query:"afterJson" long:"afterJson" msg:"afterJson"`
}

func NewFacilityLogs(adapter *Ch.Adapter) *FacilityLogs {
	return &FacilityLogs{Adapter: adapter}
}

// FacilityLogsFieldTypeMap returns key value of field name and key
var FacilityLogsFieldTypeMap = map[string]Ch.DataType{ //nolint:dupl false positive
	`createdAt`:  Ch.DateTime,
	`actorId`:    Ch.UInt64,
	`beforeJson`: Ch.String,
	`afterJson`:  Ch.String,
}

func (f *FacilityLogs) TableName() Ch.TableName { //nolint:dupl false positive
	return mProperty.TableFacilityLogs
}

func (f *FacilityLogs) SqlTableName() string { //nolint:dupl false positive
	return `"facilityLogs"`
}

func (f *FacilityLogs) ScanRowAllCols(rows *sql.Rows) (err error) { //nolint:dupl false positive
	return rows.Scan(
		&f.CreatedAt,
		&f.ActorId,
		&f.BeforeJson,
		&f.AfterJson,
	)
}

func (f *FacilityLogs) ScanRowsAllCols(rows *sql.Rows, estimateRows int) (res []FacilityLogs, err error) { //nolint:dupl false positive
	res = make([]FacilityLogs, 0, estimateRows)
	defer rows.Close()
	for rows.Next() {
		var row FacilityLogs
		err = row.ScanRowAllCols(rows)
		if err != nil {
			return
		}
		res = append(res, row)
	}
	return
}

// insert, error if exists
func (f *FacilityLogs) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + f.SqlTableName() + `(` + f.SqlAllFields() + `) VALUES (?,?,?,?)`
}

func (f *FacilityLogs) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + f.SqlTableName()
}

func (f *FacilityLogs) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, actorId
	, beforeJson
	, afterJson
	`
}

func (f *FacilityLogs) SqlAllFields() string { //nolint:dupl false positive
	return `createdAt, actorId, beforeJson, afterJson`
}

func (f FacilityLogs) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		f.CreatedAt,  // 0
		f.ActorId,    // 1
		f.BeforeJson, // 2
		f.AfterJson,  // 3
	}
}

func (f *FacilityLogs) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (f *FacilityLogs) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (f *FacilityLogs) IdxActorId() int { //nolint:dupl false positive
	return 1
}

func (f *FacilityLogs) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (f *FacilityLogs) IdxBeforeJson() int { //nolint:dupl false positive
	return 2
}

func (f *FacilityLogs) SqlBeforeJson() string { //nolint:dupl false positive
	return `beforeJson`
}

func (f *FacilityLogs) IdxAfterJson() int { //nolint:dupl false positive
	return 3
}

func (f *FacilityLogs) SqlAfterJson() string { //nolint:dupl false positive
	return `afterJson`
}

func (f *FacilityLogs) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		f.CreatedAt,  // 0
		f.ActorId,    // 1
		f.BeforeJson, // 2
		f.AfterJson,  // 3
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

type LocationLogs struct {
	Adapter    *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	CreatedAt  time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	ActorId    uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	BeforeJson string      `json:"beforeJson" form:"beforeJson" query:"beforeJson" long:"beforeJson" msg:"beforeJson"`
	AfterJson  string      `json:"afterJson" form:"afterJson" query:"afterJson" long:"afterJson" msg:"afterJson"`
}

func NewLocationLogs(adapter *Ch.Adapter) *LocationLogs {
	return &LocationLogs{Adapter: adapter}
}

// LocationLogsFieldTypeMap returns key value of field name and key
var LocationLogsFieldTypeMap = map[string]Ch.DataType{ //nolint:dupl false positive
	`createdAt`:  Ch.DateTime,
	`actorId`:    Ch.UInt64,
	`beforeJson`: Ch.String,
	`afterJson`:  Ch.String,
}

func (l *LocationLogs) TableName() Ch.TableName { //nolint:dupl false positive
	return mProperty.TableLocationLogs
}

func (l *LocationLogs) SqlTableName() string { //nolint:dupl false positive
	return `"locationLogs"`
}

func (l *LocationLogs) ScanRowAllCols(rows *sql.Rows) (err error) { //nolint:dupl false positive
	return rows.Scan(
		&l.CreatedAt,
		&l.ActorId,
		&l.BeforeJson,
		&l.AfterJson,
	)
}

func (l *LocationLogs) ScanRowsAllCols(rows *sql.Rows, estimateRows int) (res []LocationLogs, err error) { //nolint:dupl false positive
	res = make([]LocationLogs, 0, estimateRows)
	defer rows.Close()
	for rows.Next() {
		var row LocationLogs
		err = row.ScanRowAllCols(rows)
		if err != nil {
			return
		}
		res = append(res, row)
	}
	return
}

// insert, error if exists
func (l *LocationLogs) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + l.SqlTableName() + `(` + l.SqlAllFields() + `) VALUES (?,?,?,?)`
}

func (l *LocationLogs) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + l.SqlTableName()
}

func (l *LocationLogs) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, actorId
	, beforeJson
	, afterJson
	`
}

func (l *LocationLogs) SqlAllFields() string { //nolint:dupl false positive
	return `createdAt, actorId, beforeJson, afterJson`
}

func (l LocationLogs) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		l.CreatedAt,  // 0
		l.ActorId,    // 1
		l.BeforeJson, // 2
		l.AfterJson,  // 3
	}
}

func (l *LocationLogs) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (l *LocationLogs) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (l *LocationLogs) IdxActorId() int { //nolint:dupl false positive
	return 1
}

func (l *LocationLogs) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (l *LocationLogs) IdxBeforeJson() int { //nolint:dupl false positive
	return 2
}

func (l *LocationLogs) SqlBeforeJson() string { //nolint:dupl false positive
	return `beforeJson`
}

func (l *LocationLogs) IdxAfterJson() int { //nolint:dupl false positive
	return 3
}

func (l *LocationLogs) SqlAfterJson() string { //nolint:dupl false positive
	return `afterJson`
}

func (l *LocationLogs) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		l.CreatedAt,  // 0
		l.ActorId,    // 1
		l.BeforeJson, // 2
		l.AfterJson,  // 3
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

type PaymentLogs struct {
	Adapter    *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	CreatedAt  time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	ActorId    uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	BeforeJson string      `json:"beforeJson" form:"beforeJson" query:"beforeJson" long:"beforeJson" msg:"beforeJson"`
	AfterJson  string      `json:"afterJson" form:"afterJson" query:"afterJson" long:"afterJson" msg:"afterJson"`
}

func NewPaymentLogs(adapter *Ch.Adapter) *PaymentLogs {
	return &PaymentLogs{Adapter: adapter}
}

// PaymentLogsFieldTypeMap returns key value of field name and key
var PaymentLogsFieldTypeMap = map[string]Ch.DataType{ //nolint:dupl false positive
	`createdAt`:  Ch.DateTime,
	`actorId`:    Ch.UInt64,
	`beforeJson`: Ch.String,
	`afterJson`:  Ch.String,
}

func (p *PaymentLogs) TableName() Ch.TableName { //nolint:dupl false positive
	return mProperty.TablePaymentLogs
}

func (p *PaymentLogs) SqlTableName() string { //nolint:dupl false positive
	return `"paymentLogs"`
}

func (p *PaymentLogs) ScanRowAllCols(rows *sql.Rows) (err error) { //nolint:dupl false positive
	return rows.Scan(
		&p.CreatedAt,
		&p.ActorId,
		&p.BeforeJson,
		&p.AfterJson,
	)
}

func (p *PaymentLogs) ScanRowsAllCols(rows *sql.Rows, estimateRows int) (res []PaymentLogs, err error) { //nolint:dupl false positive
	res = make([]PaymentLogs, 0, estimateRows)
	defer rows.Close()
	for rows.Next() {
		var row PaymentLogs
		err = row.ScanRowAllCols(rows)
		if err != nil {
			return
		}
		res = append(res, row)
	}
	return
}

// insert, error if exists
func (p *PaymentLogs) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + p.SqlTableName() + `(` + p.SqlAllFields() + `) VALUES (?,?,?,?)`
}

func (p *PaymentLogs) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + p.SqlTableName()
}

func (p *PaymentLogs) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, actorId
	, beforeJson
	, afterJson
	`
}

func (p *PaymentLogs) SqlAllFields() string { //nolint:dupl false positive
	return `createdAt, actorId, beforeJson, afterJson`
}

func (p PaymentLogs) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		p.CreatedAt,  // 0
		p.ActorId,    // 1
		p.BeforeJson, // 2
		p.AfterJson,  // 3
	}
}

func (p *PaymentLogs) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (p *PaymentLogs) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (p *PaymentLogs) IdxActorId() int { //nolint:dupl false positive
	return 1
}

func (p *PaymentLogs) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (p *PaymentLogs) IdxBeforeJson() int { //nolint:dupl false positive
	return 2
}

func (p *PaymentLogs) SqlBeforeJson() string { //nolint:dupl false positive
	return `beforeJson`
}

func (p *PaymentLogs) IdxAfterJson() int { //nolint:dupl false positive
	return 3
}

func (p *PaymentLogs) SqlAfterJson() string { //nolint:dupl false positive
	return `afterJson`
}

func (p *PaymentLogs) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		p.CreatedAt,  // 0
		p.ActorId,    // 1
		p.BeforeJson, // 2
		p.AfterJson,  // 3
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

type RoomLogs struct {
	Adapter    *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	CreatedAt  time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	ActorId    uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	BeforeJson string      `json:"beforeJson" form:"beforeJson" query:"beforeJson" long:"beforeJson" msg:"beforeJson"`
	AfterJson  string      `json:"afterJson" form:"afterJson" query:"afterJson" long:"afterJson" msg:"afterJson"`
}

func NewRoomLogs(adapter *Ch.Adapter) *RoomLogs {
	return &RoomLogs{Adapter: adapter}
}

// RoomLogsFieldTypeMap returns key value of field name and key
var RoomLogsFieldTypeMap = map[string]Ch.DataType{ //nolint:dupl false positive
	`createdAt`:  Ch.DateTime,
	`actorId`:    Ch.UInt64,
	`beforeJson`: Ch.String,
	`afterJson`:  Ch.String,
}

func (r *RoomLogs) TableName() Ch.TableName { //nolint:dupl false positive
	return mProperty.TableRoomLogs
}

func (r *RoomLogs) SqlTableName() string { //nolint:dupl false positive
	return `"roomLogs"`
}

func (r *RoomLogs) ScanRowAllCols(rows *sql.Rows) (err error) { //nolint:dupl false positive
	return rows.Scan(
		&r.CreatedAt,
		&r.ActorId,
		&r.BeforeJson,
		&r.AfterJson,
	)
}

func (r *RoomLogs) ScanRowsAllCols(rows *sql.Rows, estimateRows int) (res []RoomLogs, err error) { //nolint:dupl false positive
	res = make([]RoomLogs, 0, estimateRows)
	defer rows.Close()
	for rows.Next() {
		var row RoomLogs
		err = row.ScanRowAllCols(rows)
		if err != nil {
			return
		}
		res = append(res, row)
	}
	return
}

// insert, error if exists
func (r *RoomLogs) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + r.SqlTableName() + `(` + r.SqlAllFields() + `) VALUES (?,?,?,?)`
}

func (r *RoomLogs) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + r.SqlTableName()
}

func (r *RoomLogs) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, actorId
	, beforeJson
	, afterJson
	`
}

func (r *RoomLogs) SqlAllFields() string { //nolint:dupl false positive
	return `createdAt, actorId, beforeJson, afterJson`
}

func (r RoomLogs) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		r.CreatedAt,  // 0
		r.ActorId,    // 1
		r.BeforeJson, // 2
		r.AfterJson,  // 3
	}
}

func (r *RoomLogs) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (r *RoomLogs) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (r *RoomLogs) IdxActorId() int { //nolint:dupl false positive
	return 1
}

func (r *RoomLogs) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (r *RoomLogs) IdxBeforeJson() int { //nolint:dupl false positive
	return 2
}

func (r *RoomLogs) SqlBeforeJson() string { //nolint:dupl false positive
	return `beforeJson`
}

func (r *RoomLogs) IdxAfterJson() int { //nolint:dupl false positive
	return 3
}

func (r *RoomLogs) SqlAfterJson() string { //nolint:dupl false positive
	return `afterJson`
}

func (r *RoomLogs) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		r.CreatedAt,  // 0
		r.ActorId,    // 1
		r.BeforeJson, // 2
		r.AfterJson,  // 3
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

type StockLogs struct {
	Adapter    *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	CreatedAt  time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	ActorId    uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	BeforeJson string      `json:"beforeJson" form:"beforeJson" query:"beforeJson" long:"beforeJson" msg:"beforeJson"`
	AfterJson  string      `json:"afterJson" form:"afterJson" query:"afterJson" long:"afterJson" msg:"afterJson"`
}

func NewStockLogs(adapter *Ch.Adapter) *StockLogs {
	return &StockLogs{Adapter: adapter}
}

// StockLogsFieldTypeMap returns key value of field name and key
var StockLogsFieldTypeMap = map[string]Ch.DataType{ //nolint:dupl false positive
	`createdAt`:  Ch.DateTime,
	`actorId`:    Ch.UInt64,
	`beforeJson`: Ch.String,
	`afterJson`:  Ch.String,
}

func (s *StockLogs) TableName() Ch.TableName { //nolint:dupl false positive
	return mProperty.TableStockLogs
}

func (s *StockLogs) SqlTableName() string { //nolint:dupl false positive
	return `"stockLogs"`
}

func (s *StockLogs) ScanRowAllCols(rows *sql.Rows) (err error) { //nolint:dupl false positive
	return rows.Scan(
		&s.CreatedAt,
		&s.ActorId,
		&s.BeforeJson,
		&s.AfterJson,
	)
}

func (s *StockLogs) ScanRowsAllCols(rows *sql.Rows, estimateRows int) (res []StockLogs, err error) { //nolint:dupl false positive
	res = make([]StockLogs, 0, estimateRows)
	defer rows.Close()
	for rows.Next() {
		var row StockLogs
		err = row.ScanRowAllCols(rows)
		if err != nil {
			return
		}
		res = append(res, row)
	}
	return
}

// insert, error if exists
func (s *StockLogs) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + s.SqlTableName() + `(` + s.SqlAllFields() + `) VALUES (?,?,?,?)`
}

func (s *StockLogs) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + s.SqlTableName()
}

func (s *StockLogs) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, actorId
	, beforeJson
	, afterJson
	`
}

func (s *StockLogs) SqlAllFields() string { //nolint:dupl false positive
	return `createdAt, actorId, beforeJson, afterJson`
}

func (s StockLogs) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		s.CreatedAt,  // 0
		s.ActorId,    // 1
		s.BeforeJson, // 2
		s.AfterJson,  // 3
	}
}

func (s *StockLogs) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (s *StockLogs) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (s *StockLogs) IdxActorId() int { //nolint:dupl false positive
	return 1
}

func (s *StockLogs) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (s *StockLogs) IdxBeforeJson() int { //nolint:dupl false positive
	return 2
}

func (s *StockLogs) SqlBeforeJson() string { //nolint:dupl false positive
	return `beforeJson`
}

func (s *StockLogs) IdxAfterJson() int { //nolint:dupl false positive
	return 3
}

func (s *StockLogs) SqlAfterJson() string { //nolint:dupl false positive
	return `afterJson`
}

func (s *StockLogs) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		s.CreatedAt,  // 0
		s.ActorId,    // 1
		s.BeforeJson, // 2
		s.AfterJson,  // 3
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go
