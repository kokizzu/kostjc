package saAuth

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

import (
	"database/sql"
	"net"
	"time"

	"kostjc/model/mAuth"

	_ "github.com/ClickHouse/clickhouse-go/v2"
	chBuffer "github.com/kokizzu/ch-timed-buffer"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/L"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file saAuth__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type saAuth__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type saAuth__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type saAuth__ORM.GEN.go
// go:generate msgp -tests=false -file saAuth__ORM.GEN.go -o saAuth__MSG.GEN.go

var actionLogsDummy = ActionLogs{}
var tenantLogsDummy = TenantLogs{}
var userLogsDummy = UserLogs{}
var Preparators = map[Ch.TableName]chBuffer.Preparator{
	mAuth.TableActionLogs: func(tx *sql.Tx) *sql.Stmt {
		query := actionLogsDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
	mAuth.TableTenantLogs: func(tx *sql.Tx) *sql.Stmt {
		query := tenantLogsDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
	mAuth.TableUserLogs: func(tx *sql.Tx) *sql.Stmt {
		query := userLogsDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
}

type ActionLogs struct {
	Adapter    *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	CreatedAt  time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	RequestId  string      `json:"requestId,string" form:"requestId" query:"requestId" long:"requestId" msg:"requestId"`
	ActorId    uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	Action     string      `json:"action" form:"action" query:"action" long:"action" msg:"action"`
	StatusCode int16       `json:"statusCode" form:"statusCode" query:"statusCode" long:"statusCode" msg:"statusCode"`
	Traces     string      `json:"traces" form:"traces" query:"traces" long:"traces" msg:"traces"`
	Error      string      `json:"error" form:"error" query:"error" long:"error" msg:"error"`
	IpAddr4    net.IP      `json:"ipAddr4" form:"ipAddr4" query:"ipAddr4" long:"ipAddr4" msg:"ipAddr4"`
	IpAddr6    net.IP      `json:"ipAddr6" form:"ipAddr6" query:"ipAddr6" long:"ipAddr6" msg:"ipAddr6"`
	UserAgent  string      `json:"userAgent" form:"userAgent" query:"userAgent" long:"userAgent" msg:"userAgent"`
	Lat        float64     `json:"lat" form:"lat" query:"lat" long:"lat" msg:"lat"`
	Long       float64     `json:"long" form:"long" query:"long" long:"long" msg:"long"`
	Latency    float64     `json:"latency" form:"latency" query:"latency" long:"latency" msg:"latency"`
	RefId      uint64      `json:"refId,string" form:"refId" query:"refId" long:"refId" msg:"refId"`
}

func NewActionLogs(adapter *Ch.Adapter) *ActionLogs {
	return &ActionLogs{Adapter: adapter}
}

// ActionLogsFieldTypeMap returns key value of field name and key
var ActionLogsFieldTypeMap = map[string]Ch.DataType{ //nolint:dupl false positive
	`createdAt`:  Ch.DateTime,
	`requestId`:  Ch.String,
	`actorId`:    Ch.UInt64,
	`action`:     Ch.String,
	`statusCode`: Ch.Int16,
	`traces`:     Ch.String,
	`error`:      Ch.String,
	`ipAddr4`:    Ch.IPv4,
	`ipAddr6`:    Ch.IPv6,
	`userAgent`:  Ch.String,
	`lat`:        Ch.Float64,
	`long`:       Ch.Float64,
	`latency`:    Ch.Float64,
	`refId`:      Ch.UInt64,
}

func (a *ActionLogs) TableName() Ch.TableName { //nolint:dupl false positive
	return mAuth.TableActionLogs
}

func (a *ActionLogs) SqlTableName() string { //nolint:dupl false positive
	return `"actionLogs"`
}

func (a *ActionLogs) ScanRowAllCols(rows *sql.Rows) (err error) { //nolint:dupl false positive
	return rows.Scan(
		&a.CreatedAt,
		&a.RequestId,
		&a.ActorId,
		&a.Action,
		&a.StatusCode,
		&a.Traces,
		&a.Error,
		&a.IpAddr4,
		&a.IpAddr6,
		&a.UserAgent,
		&a.Lat,
		&a.Long,
		&a.Latency,
		&a.RefId,
	)
}

func (a *ActionLogs) ScanRowsAllCols(rows *sql.Rows, estimateRows int) (res []ActionLogs, err error) { //nolint:dupl false positive
	res = make([]ActionLogs, 0, estimateRows)
	defer rows.Close()
	for rows.Next() {
		var row ActionLogs
		err = row.ScanRowAllCols(rows)
		if err != nil {
			return
		}
		res = append(res, row)
	}
	return
}

// insert, error if exists
func (a *ActionLogs) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + a.SqlTableName() + `(` + a.SqlAllFields() + `) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
}

func (a *ActionLogs) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + a.SqlTableName()
}

func (a *ActionLogs) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, requestId
	, actorId
	, action
	, statusCode
	, traces
	, error
	, ipAddr4
	, ipAddr6
	, userAgent
	, lat
	, long
	, latency
	, refId
	`
}

func (a *ActionLogs) SqlAllFields() string { //nolint:dupl false positive
	return `createdAt, requestId, actorId, action, statusCode, traces, error, ipAddr4, ipAddr6, userAgent, lat, long, latency, refId`
}

func (a ActionLogs) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		a.CreatedAt,  // 0
		a.RequestId,  // 1
		a.ActorId,    // 2
		a.Action,     // 3
		a.StatusCode, // 4
		a.Traces,     // 5
		a.Error,      // 6
		a.IpAddr4,    // 7
		a.IpAddr6,    // 8
		a.UserAgent,  // 9
		a.Lat,        // 10
		a.Long,       // 11
		a.Latency,    // 12
		a.RefId,      // 13
	}
}

func (a *ActionLogs) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (a *ActionLogs) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (a *ActionLogs) IdxRequestId() int { //nolint:dupl false positive
	return 1
}

func (a *ActionLogs) SqlRequestId() string { //nolint:dupl false positive
	return `requestId`
}

func (a *ActionLogs) IdxActorId() int { //nolint:dupl false positive
	return 2
}

func (a *ActionLogs) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (a *ActionLogs) IdxAction() int { //nolint:dupl false positive
	return 3
}

func (a *ActionLogs) SqlAction() string { //nolint:dupl false positive
	return `action`
}

func (a *ActionLogs) IdxStatusCode() int { //nolint:dupl false positive
	return 4
}

func (a *ActionLogs) SqlStatusCode() string { //nolint:dupl false positive
	return `statusCode`
}

func (a *ActionLogs) IdxTraces() int { //nolint:dupl false positive
	return 5
}

func (a *ActionLogs) SqlTraces() string { //nolint:dupl false positive
	return `traces`
}

func (a *ActionLogs) IdxError() int { //nolint:dupl false positive
	return 6
}

func (a *ActionLogs) SqlError() string { //nolint:dupl false positive
	return `error`
}

func (a *ActionLogs) IdxIpAddr4() int { //nolint:dupl false positive
	return 7
}

func (a *ActionLogs) SqlIpAddr4() string { //nolint:dupl false positive
	return `ipAddr4`
}

func (a *ActionLogs) IdxIpAddr6() int { //nolint:dupl false positive
	return 8
}

func (a *ActionLogs) SqlIpAddr6() string { //nolint:dupl false positive
	return `ipAddr6`
}

func (a *ActionLogs) IdxUserAgent() int { //nolint:dupl false positive
	return 9
}

func (a *ActionLogs) SqlUserAgent() string { //nolint:dupl false positive
	return `userAgent`
}

func (a *ActionLogs) IdxLat() int { //nolint:dupl false positive
	return 10
}

func (a *ActionLogs) SqlLat() string { //nolint:dupl false positive
	return `lat`
}

func (a *ActionLogs) IdxLong() int { //nolint:dupl false positive
	return 11
}

func (a *ActionLogs) SqlLong() string { //nolint:dupl false positive
	return `long`
}

func (a *ActionLogs) IdxLatency() int { //nolint:dupl false positive
	return 12
}

func (a *ActionLogs) SqlLatency() string { //nolint:dupl false positive
	return `latency`
}

func (a *ActionLogs) IdxRefId() int { //nolint:dupl false positive
	return 13
}

func (a *ActionLogs) SqlRefId() string { //nolint:dupl false positive
	return `refId`
}

func (a *ActionLogs) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		a.CreatedAt,  // 0
		a.RequestId,  // 1
		a.ActorId,    // 2
		a.Action,     // 3
		a.StatusCode, // 4
		a.Traces,     // 5
		a.Error,      // 6
		a.IpAddr4,    // 7
		a.IpAddr6,    // 8
		a.UserAgent,  // 9
		a.Lat,        // 10
		a.Long,       // 11
		a.Latency,    // 12
		a.RefId,      // 13
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

type TenantLogs struct {
	Adapter    *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	CreatedAt  time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	ActorId    uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	BeforeJson string      `json:"beforeJson" form:"beforeJson" query:"beforeJson" long:"beforeJson" msg:"beforeJson"`
	AfterJson  string      `json:"afterJson" form:"afterJson" query:"afterJson" long:"afterJson" msg:"afterJson"`
}

func NewTenantLogs(adapter *Ch.Adapter) *TenantLogs {
	return &TenantLogs{Adapter: adapter}
}

// TenantLogsFieldTypeMap returns key value of field name and key
var TenantLogsFieldTypeMap = map[string]Ch.DataType{ //nolint:dupl false positive
	`createdAt`:  Ch.DateTime,
	`actorId`:    Ch.UInt64,
	`beforeJson`: Ch.String,
	`afterJson`:  Ch.String,
}

func (t *TenantLogs) TableName() Ch.TableName { //nolint:dupl false positive
	return mAuth.TableTenantLogs
}

func (t *TenantLogs) SqlTableName() string { //nolint:dupl false positive
	return `"tenantLogs"`
}

func (t *TenantLogs) ScanRowAllCols(rows *sql.Rows) (err error) { //nolint:dupl false positive
	return rows.Scan(
		&t.CreatedAt,
		&t.ActorId,
		&t.BeforeJson,
		&t.AfterJson,
	)
}

func (t *TenantLogs) ScanRowsAllCols(rows *sql.Rows, estimateRows int) (res []TenantLogs, err error) { //nolint:dupl false positive
	res = make([]TenantLogs, 0, estimateRows)
	defer rows.Close()
	for rows.Next() {
		var row TenantLogs
		err = row.ScanRowAllCols(rows)
		if err != nil {
			return
		}
		res = append(res, row)
	}
	return
}

// insert, error if exists
func (t *TenantLogs) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + t.SqlTableName() + `(` + t.SqlAllFields() + `) VALUES (?,?,?,?)`
}

func (t *TenantLogs) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + t.SqlTableName()
}

func (t *TenantLogs) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, actorId
	, beforeJson
	, afterJson
	`
}

func (t *TenantLogs) SqlAllFields() string { //nolint:dupl false positive
	return `createdAt, actorId, beforeJson, afterJson`
}

func (t TenantLogs) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		t.CreatedAt,  // 0
		t.ActorId,    // 1
		t.BeforeJson, // 2
		t.AfterJson,  // 3
	}
}

func (t *TenantLogs) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (t *TenantLogs) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (t *TenantLogs) IdxActorId() int { //nolint:dupl false positive
	return 1
}

func (t *TenantLogs) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (t *TenantLogs) IdxBeforeJson() int { //nolint:dupl false positive
	return 2
}

func (t *TenantLogs) SqlBeforeJson() string { //nolint:dupl false positive
	return `beforeJson`
}

func (t *TenantLogs) IdxAfterJson() int { //nolint:dupl false positive
	return 3
}

func (t *TenantLogs) SqlAfterJson() string { //nolint:dupl false positive
	return `afterJson`
}

func (t *TenantLogs) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		t.CreatedAt,  // 0
		t.ActorId,    // 1
		t.BeforeJson, // 2
		t.AfterJson,  // 3
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

type UserLogs struct {
	Adapter    *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	CreatedAt  time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	ActorId    uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	BeforeJson string      `json:"beforeJson" form:"beforeJson" query:"beforeJson" long:"beforeJson" msg:"beforeJson"`
	AfterJson  string      `json:"afterJson" form:"afterJson" query:"afterJson" long:"afterJson" msg:"afterJson"`
}

func NewUserLogs(adapter *Ch.Adapter) *UserLogs {
	return &UserLogs{Adapter: adapter}
}

// UserLogsFieldTypeMap returns key value of field name and key
var UserLogsFieldTypeMap = map[string]Ch.DataType{ //nolint:dupl false positive
	`createdAt`:  Ch.DateTime,
	`actorId`:    Ch.UInt64,
	`beforeJson`: Ch.String,
	`afterJson`:  Ch.String,
}

func (u *UserLogs) TableName() Ch.TableName { //nolint:dupl false positive
	return mAuth.TableUserLogs
}

func (u *UserLogs) SqlTableName() string { //nolint:dupl false positive
	return `"userLogs"`
}

func (u *UserLogs) ScanRowAllCols(rows *sql.Rows) (err error) { //nolint:dupl false positive
	return rows.Scan(
		&u.CreatedAt,
		&u.ActorId,
		&u.BeforeJson,
		&u.AfterJson,
	)
}

func (u *UserLogs) ScanRowsAllCols(rows *sql.Rows, estimateRows int) (res []UserLogs, err error) { //nolint:dupl false positive
	res = make([]UserLogs, 0, estimateRows)
	defer rows.Close()
	for rows.Next() {
		var row UserLogs
		err = row.ScanRowAllCols(rows)
		if err != nil {
			return
		}
		res = append(res, row)
	}
	return
}

// insert, error if exists
func (u *UserLogs) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + u.SqlTableName() + `(` + u.SqlAllFields() + `) VALUES (?,?,?,?)`
}

func (u *UserLogs) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + u.SqlTableName()
}

func (u *UserLogs) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, actorId
	, beforeJson
	, afterJson
	`
}

func (u *UserLogs) SqlAllFields() string { //nolint:dupl false positive
	return `createdAt, actorId, beforeJson, afterJson`
}

func (u UserLogs) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		u.CreatedAt,  // 0
		u.ActorId,    // 1
		u.BeforeJson, // 2
		u.AfterJson,  // 3
	}
}

func (u *UserLogs) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (u *UserLogs) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (u *UserLogs) IdxActorId() int { //nolint:dupl false positive
	return 1
}

func (u *UserLogs) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (u *UserLogs) IdxBeforeJson() int { //nolint:dupl false positive
	return 2
}

func (u *UserLogs) SqlBeforeJson() string { //nolint:dupl false positive
	return `beforeJson`
}

func (u *UserLogs) IdxAfterJson() int { //nolint:dupl false positive
	return 3
}

func (u *UserLogs) SqlAfterJson() string { //nolint:dupl false positive
	return `afterJson`
}

func (u *UserLogs) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		u.CreatedAt,  // 0
		u.ActorId,    // 1
		u.BeforeJson, // 2
		u.AfterJson,  // 3
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go
