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
var afterJsonDummy = AfterJson{}
var beforeJsonDummy = BeforeJson{}
var Preparators = map[Ch.TableName]chBuffer.Preparator{
	mAuth.TableActionLogs: func(tx *sql.Tx) *sql.Stmt {
		query := actionLogsDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
	mAuth.TableAfterJson: func(tx *sql.Tx) *sql.Stmt {
		query := afterJsonDummy.SqlInsert()
		stmt, err := tx.Prepare(query)
		L.IsError(err, `failed to tx.Prepare: `+query)
		return stmt
	},
	mAuth.TableBeforeJson: func(tx *sql.Tx) *sql.Stmt {
		query := beforeJsonDummy.SqlInsert()
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

type AfterJson struct {
	Adapter   *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	CreatedAt time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	ActorId   uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	Action    string      `json:"action" form:"action" query:"action" long:"action" msg:"action"`
	DataJson  string      `json:"dataJson" form:"dataJson" query:"dataJson" long:"dataJson" msg:"dataJson"`
	Table     string      `json:"table" form:"table" query:"table" long:"table" msg:"table"`
}

func NewAfterJson(adapter *Ch.Adapter) *AfterJson {
	return &AfterJson{Adapter: adapter}
}

// AfterJsonFieldTypeMap returns key value of field name and key
var AfterJsonFieldTypeMap = map[string]Ch.DataType{ //nolint:dupl false positive
	`createdAt`: Ch.DateTime,
	`actorId`:   Ch.UInt64,
	`action`:    Ch.String,
	`dataJson`:  Ch.String,
	`table`:     Ch.String,
}

func (a *AfterJson) TableName() Ch.TableName { //nolint:dupl false positive
	return mAuth.TableAfterJson
}

func (a *AfterJson) SqlTableName() string { //nolint:dupl false positive
	return `"afterJson"`
}

func (a *AfterJson) ScanRowAllCols(rows *sql.Rows) (err error) { //nolint:dupl false positive
	return rows.Scan(
		&a.CreatedAt,
		&a.ActorId,
		&a.Action,
		&a.DataJson,
		&a.Table,
	)
}

func (a *AfterJson) ScanRowsAllCols(rows *sql.Rows, estimateRows int) (res []AfterJson, err error) { //nolint:dupl false positive
	res = make([]AfterJson, 0, estimateRows)
	defer rows.Close()
	for rows.Next() {
		var row AfterJson
		err = row.ScanRowAllCols(rows)
		if err != nil {
			return
		}
		res = append(res, row)
	}
	return
}

// insert, error if exists
func (a *AfterJson) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + a.SqlTableName() + `(` + a.SqlAllFields() + `) VALUES (?,?,?,?,?)`
}

func (a *AfterJson) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + a.SqlTableName()
}

func (a *AfterJson) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, actorId
	, action
	, dataJson
	, table
	`
}

func (a *AfterJson) SqlAllFields() string { //nolint:dupl false positive
	return `createdAt, actorId, action, dataJson, table`
}

func (a AfterJson) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		a.CreatedAt, // 0
		a.ActorId,   // 1
		a.Action,    // 2
		a.DataJson,  // 3
		a.Table,     // 4
	}
}

func (a *AfterJson) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (a *AfterJson) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (a *AfterJson) IdxActorId() int { //nolint:dupl false positive
	return 1
}

func (a *AfterJson) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (a *AfterJson) IdxAction() int { //nolint:dupl false positive
	return 2
}

func (a *AfterJson) SqlAction() string { //nolint:dupl false positive
	return `action`
}

func (a *AfterJson) IdxDataJson() int { //nolint:dupl false positive
	return 3
}

func (a *AfterJson) SqlDataJson() string { //nolint:dupl false positive
	return `dataJson`
}

func (a *AfterJson) IdxTable() int { //nolint:dupl false positive
	return 4
}

func (a *AfterJson) SqlTable() string { //nolint:dupl false positive
	return `table`
}

func (a *AfterJson) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		a.CreatedAt, // 0
		a.ActorId,   // 1
		a.Action,    // 2
		a.DataJson,  // 3
		a.Table,     // 4
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go

type BeforeJson struct {
	Adapter   *Ch.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	CreatedAt time.Time   `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	ActorId   uint64      `json:"actorId,string" form:"actorId" query:"actorId" long:"actorId" msg:"actorId"`
	Action    string      `json:"action" form:"action" query:"action" long:"action" msg:"action"`
	DataJson  string      `json:"dataJson" form:"dataJson" query:"dataJson" long:"dataJson" msg:"dataJson"`
	Table     string      `json:"table" form:"table" query:"table" long:"table" msg:"table"`
}

func NewBeforeJson(adapter *Ch.Adapter) *BeforeJson {
	return &BeforeJson{Adapter: adapter}
}

// BeforeJsonFieldTypeMap returns key value of field name and key
var BeforeJsonFieldTypeMap = map[string]Ch.DataType{ //nolint:dupl false positive
	`createdAt`: Ch.DateTime,
	`actorId`:   Ch.UInt64,
	`action`:    Ch.String,
	`dataJson`:  Ch.String,
	`table`:     Ch.String,
}

func (b *BeforeJson) TableName() Ch.TableName { //nolint:dupl false positive
	return mAuth.TableBeforeJson
}

func (b *BeforeJson) SqlTableName() string { //nolint:dupl false positive
	return `"beforeJson"`
}

func (b *BeforeJson) ScanRowAllCols(rows *sql.Rows) (err error) { //nolint:dupl false positive
	return rows.Scan(
		&b.CreatedAt,
		&b.ActorId,
		&b.Action,
		&b.DataJson,
		&b.Table,
	)
}

func (b *BeforeJson) ScanRowsAllCols(rows *sql.Rows, estimateRows int) (res []BeforeJson, err error) { //nolint:dupl false positive
	res = make([]BeforeJson, 0, estimateRows)
	defer rows.Close()
	for rows.Next() {
		var row BeforeJson
		err = row.ScanRowAllCols(rows)
		if err != nil {
			return
		}
		res = append(res, row)
	}
	return
}

// insert, error if exists
func (b *BeforeJson) SqlInsert() string { //nolint:dupl false positive
	return `INSERT INTO ` + b.SqlTableName() + `(` + b.SqlAllFields() + `) VALUES (?,?,?,?,?)`
}

func (b *BeforeJson) SqlCount() string { //nolint:dupl false positive
	return `SELECT COUNT(*) FROM ` + b.SqlTableName()
}

func (b *BeforeJson) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` createdAt
	, actorId
	, action
	, dataJson
	, table
	`
}

func (b *BeforeJson) SqlAllFields() string { //nolint:dupl false positive
	return `createdAt, actorId, action, dataJson, table`
}

func (b BeforeJson) SqlInsertParam() []any { //nolint:dupl false positive
	return []any{
		b.CreatedAt, // 0
		b.ActorId,   // 1
		b.Action,    // 2
		b.DataJson,  // 3
		b.Table,     // 4
	}
}

func (b *BeforeJson) IdxCreatedAt() int { //nolint:dupl false positive
	return 0
}

func (b *BeforeJson) SqlCreatedAt() string { //nolint:dupl false positive
	return `createdAt`
}

func (b *BeforeJson) IdxActorId() int { //nolint:dupl false positive
	return 1
}

func (b *BeforeJson) SqlActorId() string { //nolint:dupl false positive
	return `actorId`
}

func (b *BeforeJson) IdxAction() int { //nolint:dupl false positive
	return 2
}

func (b *BeforeJson) SqlAction() string { //nolint:dupl false positive
	return `action`
}

func (b *BeforeJson) IdxDataJson() int { //nolint:dupl false positive
	return 3
}

func (b *BeforeJson) SqlDataJson() string { //nolint:dupl false positive
	return `dataJson`
}

func (b *BeforeJson) IdxTable() int { //nolint:dupl false positive
	return 4
}

func (b *BeforeJson) SqlTable() string { //nolint:dupl false positive
	return `table`
}

func (b *BeforeJson) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		b.CreatedAt, // 0
		b.ActorId,   // 1
		b.Action,    // 2
		b.DataJson,  // 3
		b.Table,     // 4
	}
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/Ch/clickhouse_orm_generator.go
