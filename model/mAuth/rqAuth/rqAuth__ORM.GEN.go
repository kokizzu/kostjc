package rqAuth

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"kostjc/model/mAuth"

	"github.com/tarantool/go-tarantool"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

// Sessions DAO reader/query struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqAuth__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqAuth__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqAuth__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqAuth__ORM.GEN.go
type Sessions struct {
	Adapter      *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	SessionToken string      `json:"sessionToken" form:"sessionToken" query:"sessionToken" long:"sessionToken" msg:"sessionToken"`
	UserId       uint64      `json:"userId,string" form:"userId" query:"userId" long:"userId" msg:"userId"`
	ExpiredAt    int64       `json:"expiredAt" form:"expiredAt" query:"expiredAt" long:"expiredAt" msg:"expiredAt"`
	Device       string      `json:"device" form:"device" query:"device" long:"device" msg:"device"`
	LoginAt      int64       `json:"loginAt" form:"loginAt" query:"loginAt" long:"loginAt" msg:"loginAt"`
	LoginIPs     string      `json:"loginIPs" form:"loginIPs" query:"loginIPs" long:"loginIPs" msg:"loginIPs"`
}

// NewSessions create new ORM reader/query object
func NewSessions(adapter *Tt.Adapter) *Sessions {
	return &Sessions{Adapter: adapter}
}

// SpaceName returns full package and table name
func (s *Sessions) SpaceName() string { //nolint:dupl false positive
	return string(mAuth.TableSessions) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (s *Sessions) SqlTableName() string { //nolint:dupl false positive
	return `"sessions"`
}

// UniqueIndexSessionToken return unique index name
func (s *Sessions) UniqueIndexSessionToken() string { //nolint:dupl false positive
	return `sessionToken`
}

// FindBySessionToken Find one by SessionToken
func (s *Sessions) FindBySessionToken() bool { //nolint:dupl false positive
	res, err := s.Adapter.Select(s.SpaceName(), s.UniqueIndexSessionToken(), 0, 1, tarantool.IterEq, A.X{s.SessionToken})
	if L.IsError(err, `Sessions.FindBySessionToken failed: `+s.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		s.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (s *Sessions) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "sessionToken"
	, "userId"
	, "expiredAt"
	, "device"
	, "loginAt"
	, "loginIPs"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (s *Sessions) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "sessionToken"
	, "userId"
	, "expiredAt"
	, "device"
	, "loginAt"
	, "loginIPs"
	`
}

// ToUpdateArray generate slice of update command
func (s *Sessions) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, s.SessionToken},
		A.X{`=`, 1, s.UserId},
		A.X{`=`, 2, s.ExpiredAt},
		A.X{`=`, 3, s.Device},
		A.X{`=`, 4, s.LoginAt},
		A.X{`=`, 5, s.LoginIPs},
	}
}

// IdxSessionToken return name of the index
func (s *Sessions) IdxSessionToken() int { //nolint:dupl false positive
	return 0
}

// SqlSessionToken return name of the column being indexed
func (s *Sessions) SqlSessionToken() string { //nolint:dupl false positive
	return `"sessionToken"`
}

// IdxUserId return name of the index
func (s *Sessions) IdxUserId() int { //nolint:dupl false positive
	return 1
}

// SqlUserId return name of the column being indexed
func (s *Sessions) SqlUserId() string { //nolint:dupl false positive
	return `"userId"`
}

// IdxExpiredAt return name of the index
func (s *Sessions) IdxExpiredAt() int { //nolint:dupl false positive
	return 2
}

// SqlExpiredAt return name of the column being indexed
func (s *Sessions) SqlExpiredAt() string { //nolint:dupl false positive
	return `"expiredAt"`
}

// IdxDevice return name of the index
func (s *Sessions) IdxDevice() int { //nolint:dupl false positive
	return 3
}

// SqlDevice return name of the column being indexed
func (s *Sessions) SqlDevice() string { //nolint:dupl false positive
	return `"device"`
}

// IdxLoginAt return name of the index
func (s *Sessions) IdxLoginAt() int { //nolint:dupl false positive
	return 4
}

// SqlLoginAt return name of the column being indexed
func (s *Sessions) SqlLoginAt() string { //nolint:dupl false positive
	return `"loginAt"`
}

// IdxLoginIPs return name of the index
func (s *Sessions) IdxLoginIPs() int { //nolint:dupl false positive
	return 5
}

// SqlLoginIPs return name of the column being indexed
func (s *Sessions) SqlLoginIPs() string { //nolint:dupl false positive
	return `"loginIPs"`
}

// ToArray receiver fields to slice
func (s *Sessions) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		s.SessionToken, // 0
		s.UserId,       // 1
		s.ExpiredAt,    // 2
		s.Device,       // 3
		s.LoginAt,      // 4
		s.LoginIPs,     // 5
	}
}

// FromArray convert slice to receiver fields
func (s *Sessions) FromArray(a A.X) *Sessions { //nolint:dupl false positive
	s.SessionToken = X.ToS(a[0])
	s.UserId = X.ToU(a[1])
	s.ExpiredAt = X.ToI(a[2])
	s.Device = X.ToS(a[3])
	s.LoginAt = X.ToI(a[4])
	s.LoginIPs = X.ToS(a[5])
	return s
}

// FromUncensoredArray convert slice to receiver fields
func (s *Sessions) FromUncensoredArray(a A.X) *Sessions { //nolint:dupl false positive
	s.SessionToken = X.ToS(a[0])
	s.UserId = X.ToU(a[1])
	s.ExpiredAt = X.ToI(a[2])
	s.Device = X.ToS(a[3])
	s.LoginAt = X.ToI(a[4])
	s.LoginIPs = X.ToS(a[5])
	return s
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (s *Sessions) FindOffsetLimit(offset, limit uint32, idx string) []Sessions { //nolint:dupl false positive
	var rows []Sessions
	res, err := s.Adapter.Select(s.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Sessions.FindOffsetLimit failed: `+s.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Sessions{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (s *Sessions) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := s.Adapter.Select(s.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Sessions.FindOffsetLimit failed: `+s.SpaceName()) {
		return rows, Tt.QueryMetaFrom(res, err)
	}
	tuples := res.Tuples()
	rows = make([]A.X, len(tuples))
	for z, row := range tuples {
		rows[z] = row
	}
	return rows, Tt.QueryMetaFrom(res, nil)
}

// Total count number of rows
func (s *Sessions) Total() int64 { //nolint:dupl false positive
	rows := s.Adapter.CallBoxSpace(s.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// SessionsFieldTypeMap returns key value of field name and key
var SessionsFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`sessionToken`: Tt.String,
	`userId`:       Tt.Unsigned,
	`expiredAt`:    Tt.Integer,
	`device`:       Tt.String,
	`loginAt`:      Tt.Integer,
	`loginIPs`:     Tt.String,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Tenants DAO reader/query struct
type Tenants struct {
	Adapter          *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id               uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	TenantName       string      `json:"tenantName" form:"tenantName" query:"tenantName" long:"tenantName" msg:"tenantName"`
	KtpRegion        string      `json:"ktpRegion" form:"ktpRegion" query:"ktpRegion" long:"ktpRegion" msg:"ktpRegion"`
	KtpNumber        string      `json:"ktpNumber" form:"ktpNumber" query:"ktpNumber" long:"ktpNumber" msg:"ktpNumber"`
	KtpName          string      `json:"ktpName" form:"ktpName" query:"ktpName" long:"ktpName" msg:"ktpName"`
	KtpPlaceBirth    string      `json:"ktpPlaceBirth" form:"ktpPlaceBirth" query:"ktpPlaceBirth" long:"ktpPlaceBirth" msg:"ktpPlaceBirth"`
	KtpDateBirth     string      `json:"ktpDateBirth" form:"ktpDateBirth" query:"ktpDateBirth" long:"ktpDateBirth" msg:"ktpDateBirth"`
	KtpGender        string      `json:"ktpGender" form:"ktpGender" query:"ktpGender" long:"ktpGender" msg:"ktpGender"`
	KtpAddress       string      `json:"ktpAddress" form:"ktpAddress" query:"ktpAddress" long:"ktpAddress" msg:"ktpAddress"`
	KtpRtRw          string      `json:"ktpRtRw" form:"ktpRtRw" query:"ktpRtRw" long:"ktpRtRw" msg:"ktpRtRw"`
	KtpKelurahanDesa string      `json:"ktpKelurahanDesa" form:"ktpKelurahanDesa" query:"ktpKelurahanDesa" long:"ktpKelurahanDesa" msg:"ktpKelurahanDesa"`
	KtpKecamatan     string      `json:"ktpKecamatan" form:"ktpKecamatan" query:"ktpKecamatan" long:"ktpKecamatan" msg:"ktpKecamatan"`
	KtpReligion      string      `json:"ktpReligion" form:"ktpReligion" query:"ktpReligion" long:"ktpReligion" msg:"ktpReligion"`
	KtpMaritalStatus string      `json:"ktpMaritalStatus" form:"ktpMaritalStatus" query:"ktpMaritalStatus" long:"ktpMaritalStatus" msg:"ktpMaritalStatus"`
	KtpCitizenship   string      `json:"ktpCitizenship" form:"ktpCitizenship" query:"ktpCitizenship" long:"ktpCitizenship" msg:"ktpCitizenship"`
	TelegramUsername string      `json:"telegramUsername" form:"telegramUsername" query:"telegramUsername" long:"telegramUsername" msg:"telegramUsername"`
	WhatsappNumber   string      `json:"whatsappNumber" form:"whatsappNumber" query:"whatsappNumber" long:"whatsappNumber" msg:"whatsappNumber"`
	CreatedAt        int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy        uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt        int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy        uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt        int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy        uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy       uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
	KtpOccupation    string      `json:"ktpOccupation" form:"ktpOccupation" query:"ktpOccupation" long:"ktpOccupation" msg:"ktpOccupation"`
	WaAddedAt        string      `json:"waAddedAt" form:"waAddedAt" query:"waAddedAt" long:"waAddedAt" msg:"waAddedAt"`
	TeleAddedAt      string      `json:"teleAddedAt" form:"teleAddedAt" query:"teleAddedAt" long:"teleAddedAt" msg:"teleAddedAt"`
}

// NewTenants create new ORM reader/query object
func NewTenants(adapter *Tt.Adapter) *Tenants {
	return &Tenants{Adapter: adapter}
}

// SpaceName returns full package and table name
func (t *Tenants) SpaceName() string { //nolint:dupl false positive
	return string(mAuth.TableTenants) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (t *Tenants) SqlTableName() string { //nolint:dupl false positive
	return `"tenants"`
}

func (t *Tenants) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (t *Tenants) FindById() bool { //nolint:dupl false positive
	res, err := t.Adapter.Select(t.SpaceName(), t.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{t.Id})
	if L.IsError(err, `Tenants.FindById failed: `+t.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		t.FromArray(rows[0])
		return true
	}
	return false
}

// UniqueIndexKtpNumber return unique index name
func (t *Tenants) UniqueIndexKtpNumber() string { //nolint:dupl false positive
	return `ktpNumber`
}

// FindByKtpNumber Find one by KtpNumber
func (t *Tenants) FindByKtpNumber() bool { //nolint:dupl false positive
	res, err := t.Adapter.Select(t.SpaceName(), t.UniqueIndexKtpNumber(), 0, 1, tarantool.IterEq, A.X{t.KtpNumber})
	if L.IsError(err, `Tenants.FindByKtpNumber failed: `+t.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		t.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (t *Tenants) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantName"
	, "ktpRegion"
	, "ktpNumber"
	, "ktpName"
	, "ktpPlaceBirth"
	, "ktpDateBirth"
	, "ktpGender"
	, "ktpAddress"
	, "ktpRtRw"
	, "ktpKelurahanDesa"
	, "ktpKecamatan"
	, "ktpReligion"
	, "ktpMaritalStatus"
	, "ktpCitizenship"
	, "telegramUsername"
	, "whatsappNumber"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "ktpOccupation"
	, "waAddedAt"
	, "teleAddedAt"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (t *Tenants) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantName"
	, "ktpRegion"
	, "ktpNumber"
	, "ktpName"
	, "ktpPlaceBirth"
	, "ktpDateBirth"
	, "ktpGender"
	, "ktpAddress"
	, "ktpRtRw"
	, "ktpKelurahanDesa"
	, "ktpKecamatan"
	, "ktpReligion"
	, "ktpMaritalStatus"
	, "ktpCitizenship"
	, "telegramUsername"
	, "whatsappNumber"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "ktpOccupation"
	, "waAddedAt"
	, "teleAddedAt"
	`
}

// ToUpdateArray generate slice of update command
func (t *Tenants) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, t.Id},
		A.X{`=`, 1, t.TenantName},
		A.X{`=`, 2, t.KtpRegion},
		A.X{`=`, 3, t.KtpNumber},
		A.X{`=`, 4, t.KtpName},
		A.X{`=`, 5, t.KtpPlaceBirth},
		A.X{`=`, 6, t.KtpDateBirth},
		A.X{`=`, 7, t.KtpGender},
		A.X{`=`, 8, t.KtpAddress},
		A.X{`=`, 9, t.KtpRtRw},
		A.X{`=`, 10, t.KtpKelurahanDesa},
		A.X{`=`, 11, t.KtpKecamatan},
		A.X{`=`, 12, t.KtpReligion},
		A.X{`=`, 13, t.KtpMaritalStatus},
		A.X{`=`, 14, t.KtpCitizenship},
		A.X{`=`, 15, t.TelegramUsername},
		A.X{`=`, 16, t.WhatsappNumber},
		A.X{`=`, 17, t.CreatedAt},
		A.X{`=`, 18, t.CreatedBy},
		A.X{`=`, 19, t.UpdatedAt},
		A.X{`=`, 20, t.UpdatedBy},
		A.X{`=`, 21, t.DeletedAt},
		A.X{`=`, 22, t.DeletedBy},
		A.X{`=`, 23, t.RestoredBy},
		A.X{`=`, 24, t.KtpOccupation},
		A.X{`=`, 25, t.WaAddedAt},
		A.X{`=`, 26, t.TeleAddedAt},
	}
}

// IdxId return name of the index
func (t *Tenants) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (t *Tenants) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxTenantName return name of the index
func (t *Tenants) IdxTenantName() int { //nolint:dupl false positive
	return 1
}

// SqlTenantName return name of the column being indexed
func (t *Tenants) SqlTenantName() string { //nolint:dupl false positive
	return `"tenantName"`
}

// IdxKtpRegion return name of the index
func (t *Tenants) IdxKtpRegion() int { //nolint:dupl false positive
	return 2
}

// SqlKtpRegion return name of the column being indexed
func (t *Tenants) SqlKtpRegion() string { //nolint:dupl false positive
	return `"ktpRegion"`
}

// IdxKtpNumber return name of the index
func (t *Tenants) IdxKtpNumber() int { //nolint:dupl false positive
	return 3
}

// SqlKtpNumber return name of the column being indexed
func (t *Tenants) SqlKtpNumber() string { //nolint:dupl false positive
	return `"ktpNumber"`
}

// IdxKtpName return name of the index
func (t *Tenants) IdxKtpName() int { //nolint:dupl false positive
	return 4
}

// SqlKtpName return name of the column being indexed
func (t *Tenants) SqlKtpName() string { //nolint:dupl false positive
	return `"ktpName"`
}

// IdxKtpPlaceBirth return name of the index
func (t *Tenants) IdxKtpPlaceBirth() int { //nolint:dupl false positive
	return 5
}

// SqlKtpPlaceBirth return name of the column being indexed
func (t *Tenants) SqlKtpPlaceBirth() string { //nolint:dupl false positive
	return `"ktpPlaceBirth"`
}

// IdxKtpDateBirth return name of the index
func (t *Tenants) IdxKtpDateBirth() int { //nolint:dupl false positive
	return 6
}

// SqlKtpDateBirth return name of the column being indexed
func (t *Tenants) SqlKtpDateBirth() string { //nolint:dupl false positive
	return `"ktpDateBirth"`
}

// IdxKtpGender return name of the index
func (t *Tenants) IdxKtpGender() int { //nolint:dupl false positive
	return 7
}

// SqlKtpGender return name of the column being indexed
func (t *Tenants) SqlKtpGender() string { //nolint:dupl false positive
	return `"ktpGender"`
}

// IdxKtpAddress return name of the index
func (t *Tenants) IdxKtpAddress() int { //nolint:dupl false positive
	return 8
}

// SqlKtpAddress return name of the column being indexed
func (t *Tenants) SqlKtpAddress() string { //nolint:dupl false positive
	return `"ktpAddress"`
}

// IdxKtpRtRw return name of the index
func (t *Tenants) IdxKtpRtRw() int { //nolint:dupl false positive
	return 9
}

// SqlKtpRtRw return name of the column being indexed
func (t *Tenants) SqlKtpRtRw() string { //nolint:dupl false positive
	return `"ktpRtRw"`
}

// IdxKtpKelurahanDesa return name of the index
func (t *Tenants) IdxKtpKelurahanDesa() int { //nolint:dupl false positive
	return 10
}

// SqlKtpKelurahanDesa return name of the column being indexed
func (t *Tenants) SqlKtpKelurahanDesa() string { //nolint:dupl false positive
	return `"ktpKelurahanDesa"`
}

// IdxKtpKecamatan return name of the index
func (t *Tenants) IdxKtpKecamatan() int { //nolint:dupl false positive
	return 11
}

// SqlKtpKecamatan return name of the column being indexed
func (t *Tenants) SqlKtpKecamatan() string { //nolint:dupl false positive
	return `"ktpKecamatan"`
}

// IdxKtpReligion return name of the index
func (t *Tenants) IdxKtpReligion() int { //nolint:dupl false positive
	return 12
}

// SqlKtpReligion return name of the column being indexed
func (t *Tenants) SqlKtpReligion() string { //nolint:dupl false positive
	return `"ktpReligion"`
}

// IdxKtpMaritalStatus return name of the index
func (t *Tenants) IdxKtpMaritalStatus() int { //nolint:dupl false positive
	return 13
}

// SqlKtpMaritalStatus return name of the column being indexed
func (t *Tenants) SqlKtpMaritalStatus() string { //nolint:dupl false positive
	return `"ktpMaritalStatus"`
}

// IdxKtpCitizenship return name of the index
func (t *Tenants) IdxKtpCitizenship() int { //nolint:dupl false positive
	return 14
}

// SqlKtpCitizenship return name of the column being indexed
func (t *Tenants) SqlKtpCitizenship() string { //nolint:dupl false positive
	return `"ktpCitizenship"`
}

// IdxTelegramUsername return name of the index
func (t *Tenants) IdxTelegramUsername() int { //nolint:dupl false positive
	return 15
}

// SqlTelegramUsername return name of the column being indexed
func (t *Tenants) SqlTelegramUsername() string { //nolint:dupl false positive
	return `"telegramUsername"`
}

// IdxWhatsappNumber return name of the index
func (t *Tenants) IdxWhatsappNumber() int { //nolint:dupl false positive
	return 16
}

// SqlWhatsappNumber return name of the column being indexed
func (t *Tenants) SqlWhatsappNumber() string { //nolint:dupl false positive
	return `"whatsappNumber"`
}

// IdxCreatedAt return name of the index
func (t *Tenants) IdxCreatedAt() int { //nolint:dupl false positive
	return 17
}

// SqlCreatedAt return name of the column being indexed
func (t *Tenants) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (t *Tenants) IdxCreatedBy() int { //nolint:dupl false positive
	return 18
}

// SqlCreatedBy return name of the column being indexed
func (t *Tenants) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (t *Tenants) IdxUpdatedAt() int { //nolint:dupl false positive
	return 19
}

// SqlUpdatedAt return name of the column being indexed
func (t *Tenants) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (t *Tenants) IdxUpdatedBy() int { //nolint:dupl false positive
	return 20
}

// SqlUpdatedBy return name of the column being indexed
func (t *Tenants) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (t *Tenants) IdxDeletedAt() int { //nolint:dupl false positive
	return 21
}

// SqlDeletedAt return name of the column being indexed
func (t *Tenants) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (t *Tenants) IdxDeletedBy() int { //nolint:dupl false positive
	return 22
}

// SqlDeletedBy return name of the column being indexed
func (t *Tenants) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (t *Tenants) IdxRestoredBy() int { //nolint:dupl false positive
	return 23
}

// SqlRestoredBy return name of the column being indexed
func (t *Tenants) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// IdxKtpOccupation return name of the index
func (t *Tenants) IdxKtpOccupation() int { //nolint:dupl false positive
	return 24
}

// SqlKtpOccupation return name of the column being indexed
func (t *Tenants) SqlKtpOccupation() string { //nolint:dupl false positive
	return `"ktpOccupation"`
}

// IdxWaAddedAt return name of the index
func (t *Tenants) IdxWaAddedAt() int { //nolint:dupl false positive
	return 25
}

// SqlWaAddedAt return name of the column being indexed
func (t *Tenants) SqlWaAddedAt() string { //nolint:dupl false positive
	return `"waAddedAt"`
}

// IdxTeleAddedAt return name of the index
func (t *Tenants) IdxTeleAddedAt() int { //nolint:dupl false positive
	return 26
}

// SqlTeleAddedAt return name of the column being indexed
func (t *Tenants) SqlTeleAddedAt() string { //nolint:dupl false positive
	return `"teleAddedAt"`
}

// ToArray receiver fields to slice
func (t *Tenants) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if t.Id != 0 {
		id = t.Id
	}
	return A.X{
		id,
		t.TenantName,       // 1
		t.KtpRegion,        // 2
		t.KtpNumber,        // 3
		t.KtpName,          // 4
		t.KtpPlaceBirth,    // 5
		t.KtpDateBirth,     // 6
		t.KtpGender,        // 7
		t.KtpAddress,       // 8
		t.KtpRtRw,          // 9
		t.KtpKelurahanDesa, // 10
		t.KtpKecamatan,     // 11
		t.KtpReligion,      // 12
		t.KtpMaritalStatus, // 13
		t.KtpCitizenship,   // 14
		t.TelegramUsername, // 15
		t.WhatsappNumber,   // 16
		t.CreatedAt,        // 17
		t.CreatedBy,        // 18
		t.UpdatedAt,        // 19
		t.UpdatedBy,        // 20
		t.DeletedAt,        // 21
		t.DeletedBy,        // 22
		t.RestoredBy,       // 23
		t.KtpOccupation,    // 24
		t.WaAddedAt,        // 25
		t.TeleAddedAt,      // 26
	}
}

// FromArray convert slice to receiver fields
func (t *Tenants) FromArray(a A.X) *Tenants { //nolint:dupl false positive
	t.Id = X.ToU(a[0])
	t.TenantName = X.ToS(a[1])
	t.KtpRegion = X.ToS(a[2])
	t.KtpNumber = X.ToS(a[3])
	t.KtpName = X.ToS(a[4])
	t.KtpPlaceBirth = X.ToS(a[5])
	t.KtpDateBirth = X.ToS(a[6])
	t.KtpGender = X.ToS(a[7])
	t.KtpAddress = X.ToS(a[8])
	t.KtpRtRw = X.ToS(a[9])
	t.KtpKelurahanDesa = X.ToS(a[10])
	t.KtpKecamatan = X.ToS(a[11])
	t.KtpReligion = X.ToS(a[12])
	t.KtpMaritalStatus = X.ToS(a[13])
	t.KtpCitizenship = X.ToS(a[14])
	t.TelegramUsername = X.ToS(a[15])
	t.WhatsappNumber = X.ToS(a[16])
	t.CreatedAt = X.ToI(a[17])
	t.CreatedBy = X.ToU(a[18])
	t.UpdatedAt = X.ToI(a[19])
	t.UpdatedBy = X.ToU(a[20])
	t.DeletedAt = X.ToI(a[21])
	t.DeletedBy = X.ToU(a[22])
	t.RestoredBy = X.ToU(a[23])
	t.KtpOccupation = X.ToS(a[24])
	t.WaAddedAt = X.ToS(a[25])
	t.TeleAddedAt = X.ToS(a[26])
	return t
}

// FromUncensoredArray convert slice to receiver fields
func (t *Tenants) FromUncensoredArray(a A.X) *Tenants { //nolint:dupl false positive
	t.Id = X.ToU(a[0])
	t.TenantName = X.ToS(a[1])
	t.KtpRegion = X.ToS(a[2])
	t.KtpNumber = X.ToS(a[3])
	t.KtpName = X.ToS(a[4])
	t.KtpPlaceBirth = X.ToS(a[5])
	t.KtpDateBirth = X.ToS(a[6])
	t.KtpGender = X.ToS(a[7])
	t.KtpAddress = X.ToS(a[8])
	t.KtpRtRw = X.ToS(a[9])
	t.KtpKelurahanDesa = X.ToS(a[10])
	t.KtpKecamatan = X.ToS(a[11])
	t.KtpReligion = X.ToS(a[12])
	t.KtpMaritalStatus = X.ToS(a[13])
	t.KtpCitizenship = X.ToS(a[14])
	t.TelegramUsername = X.ToS(a[15])
	t.WhatsappNumber = X.ToS(a[16])
	t.CreatedAt = X.ToI(a[17])
	t.CreatedBy = X.ToU(a[18])
	t.UpdatedAt = X.ToI(a[19])
	t.UpdatedBy = X.ToU(a[20])
	t.DeletedAt = X.ToI(a[21])
	t.DeletedBy = X.ToU(a[22])
	t.RestoredBy = X.ToU(a[23])
	t.KtpOccupation = X.ToS(a[24])
	t.WaAddedAt = X.ToS(a[25])
	t.TeleAddedAt = X.ToS(a[26])
	return t
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (t *Tenants) FindOffsetLimit(offset, limit uint32, idx string) []Tenants { //nolint:dupl false positive
	var rows []Tenants
	res, err := t.Adapter.Select(t.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Tenants.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Tenants{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (t *Tenants) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := t.Adapter.Select(t.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Tenants.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows, Tt.QueryMetaFrom(res, err)
	}
	tuples := res.Tuples()
	rows = make([]A.X, len(tuples))
	for z, row := range tuples {
		rows[z] = row
	}
	return rows, Tt.QueryMetaFrom(res, nil)
}

// Total count number of rows
func (t *Tenants) Total() int64 { //nolint:dupl false positive
	rows := t.Adapter.CallBoxSpace(t.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// TenantsFieldTypeMap returns key value of field name and key
var TenantsFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:               Tt.Unsigned,
	`tenantName`:       Tt.String,
	`ktpRegion`:        Tt.String,
	`ktpNumber`:        Tt.String,
	`ktpName`:          Tt.String,
	`ktpPlaceBirth`:    Tt.String,
	`ktpDateBirth`:     Tt.String,
	`ktpGender`:        Tt.String,
	`ktpAddress`:       Tt.String,
	`ktpRtRw`:          Tt.String,
	`ktpKelurahanDesa`: Tt.String,
	`ktpKecamatan`:     Tt.String,
	`ktpReligion`:      Tt.String,
	`ktpMaritalStatus`: Tt.String,
	`ktpCitizenship`:   Tt.String,
	`telegramUsername`: Tt.String,
	`whatsappNumber`:   Tt.String,
	`createdAt`:        Tt.Integer,
	`createdBy`:        Tt.Unsigned,
	`updatedAt`:        Tt.Integer,
	`updatedBy`:        Tt.Unsigned,
	`deletedAt`:        Tt.Integer,
	`deletedBy`:        Tt.Unsigned,
	`restoredBy`:       Tt.Unsigned,
	`ktpOccupation`:    Tt.String,
	`waAddedAt`:        Tt.String,
	`teleAddedAt`:      Tt.String,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Users DAO reader/query struct
type Users struct {
	Adapter       *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id            uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	Email         string      `json:"email" form:"email" query:"email" long:"email" msg:"email"`
	Password      string      `json:"password" form:"password" query:"password" long:"password" msg:"password"`
	CreatedAt     int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy     uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt     int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy     uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt     int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	PasswordSetAt int64       `json:"passwordSetAt" form:"passwordSetAt" query:"passwordSetAt" long:"passwordSetAt" msg:"passwordSetAt"`
	SecretCode    string      `json:"secretCode" form:"secretCode" query:"secretCode" long:"secretCode" msg:"secretCode"`
	SecretCodeAt  int64       `json:"secretCodeAt" form:"secretCodeAt" query:"secretCodeAt" long:"secretCodeAt" msg:"secretCodeAt"`
	VerifiedAt    int64       `json:"verifiedAt" form:"verifiedAt" query:"verifiedAt" long:"verifiedAt" msg:"verifiedAt"`
	LastLoginAt   int64       `json:"lastLoginAt" form:"lastLoginAt" query:"lastLoginAt" long:"lastLoginAt" msg:"lastLoginAt"`
	FullName      string      `json:"fullName" form:"fullName" query:"fullName" long:"fullName" msg:"fullName"`
	UserName      string      `json:"userName" form:"userName" query:"userName" long:"userName" msg:"userName"`
	Role          string      `json:"role" form:"role" query:"role" long:"role" msg:"role"`
}

// NewUsers create new ORM reader/query object
func NewUsers(adapter *Tt.Adapter) *Users {
	return &Users{Adapter: adapter}
}

// SpaceName returns full package and table name
func (u *Users) SpaceName() string { //nolint:dupl false positive
	return string(mAuth.TableUsers) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (u *Users) SqlTableName() string { //nolint:dupl false positive
	return `"users"`
}

func (u *Users) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (u *Users) FindById() bool { //nolint:dupl false positive
	res, err := u.Adapter.Select(u.SpaceName(), u.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{u.Id})
	if L.IsError(err, `Users.FindById failed: `+u.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		u.FromArray(rows[0])
		return true
	}
	return false
}

// UniqueIndexEmail return unique index name
func (u *Users) UniqueIndexEmail() string { //nolint:dupl false positive
	return `email`
}

// FindByEmail Find one by Email
func (u *Users) FindByEmail() bool { //nolint:dupl false positive
	res, err := u.Adapter.Select(u.SpaceName(), u.UniqueIndexEmail(), 0, 1, tarantool.IterEq, A.X{u.Email})
	if L.IsError(err, `Users.FindByEmail failed: `+u.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		u.FromArray(rows[0])
		return true
	}
	return false
}

// UniqueIndexUserName return unique index name
func (u *Users) UniqueIndexUserName() string { //nolint:dupl false positive
	return `userName`
}

// FindByUserName Find one by UserName
func (u *Users) FindByUserName() bool { //nolint:dupl false positive
	res, err := u.Adapter.Select(u.SpaceName(), u.UniqueIndexUserName(), 0, 1, tarantool.IterEq, A.X{u.UserName})
	if L.IsError(err, `Users.FindByUserName failed: `+u.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		u.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (u *Users) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "email"
	, "password"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "passwordSetAt"
	, "secretCode"
	, "secretCodeAt"
	, "verifiedAt"
	, "lastLoginAt"
	, "fullName"
	, "userName"
	, "role"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (u *Users) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "email"
	, "password"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "passwordSetAt"
	, "secretCode"
	, "secretCodeAt"
	, "verifiedAt"
	, "lastLoginAt"
	, "fullName"
	, "userName"
	, "role"
	`
}

// ToUpdateArray generate slice of update command
func (u *Users) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, u.Id},
		A.X{`=`, 1, u.Email},
		A.X{`=`, 2, u.Password},
		A.X{`=`, 3, u.CreatedAt},
		A.X{`=`, 4, u.CreatedBy},
		A.X{`=`, 5, u.UpdatedAt},
		A.X{`=`, 6, u.UpdatedBy},
		A.X{`=`, 7, u.DeletedAt},
		A.X{`=`, 8, u.PasswordSetAt},
		A.X{`=`, 9, u.SecretCode},
		A.X{`=`, 10, u.SecretCodeAt},
		A.X{`=`, 11, u.VerifiedAt},
		A.X{`=`, 12, u.LastLoginAt},
		A.X{`=`, 13, u.FullName},
		A.X{`=`, 14, u.UserName},
		A.X{`=`, 15, u.Role},
	}
}

// IdxId return name of the index
func (u *Users) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (u *Users) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxEmail return name of the index
func (u *Users) IdxEmail() int { //nolint:dupl false positive
	return 1
}

// SqlEmail return name of the column being indexed
func (u *Users) SqlEmail() string { //nolint:dupl false positive
	return `"email"`
}

// IdxPassword return name of the index
func (u *Users) IdxPassword() int { //nolint:dupl false positive
	return 2
}

// SqlPassword return name of the column being indexed
func (u *Users) SqlPassword() string { //nolint:dupl false positive
	return `"password"`
}

// IdxCreatedAt return name of the index
func (u *Users) IdxCreatedAt() int { //nolint:dupl false positive
	return 3
}

// SqlCreatedAt return name of the column being indexed
func (u *Users) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (u *Users) IdxCreatedBy() int { //nolint:dupl false positive
	return 4
}

// SqlCreatedBy return name of the column being indexed
func (u *Users) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (u *Users) IdxUpdatedAt() int { //nolint:dupl false positive
	return 5
}

// SqlUpdatedAt return name of the column being indexed
func (u *Users) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (u *Users) IdxUpdatedBy() int { //nolint:dupl false positive
	return 6
}

// SqlUpdatedBy return name of the column being indexed
func (u *Users) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (u *Users) IdxDeletedAt() int { //nolint:dupl false positive
	return 7
}

// SqlDeletedAt return name of the column being indexed
func (u *Users) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxPasswordSetAt return name of the index
func (u *Users) IdxPasswordSetAt() int { //nolint:dupl false positive
	return 8
}

// SqlPasswordSetAt return name of the column being indexed
func (u *Users) SqlPasswordSetAt() string { //nolint:dupl false positive
	return `"passwordSetAt"`
}

// IdxSecretCode return name of the index
func (u *Users) IdxSecretCode() int { //nolint:dupl false positive
	return 9
}

// SqlSecretCode return name of the column being indexed
func (u *Users) SqlSecretCode() string { //nolint:dupl false positive
	return `"secretCode"`
}

// IdxSecretCodeAt return name of the index
func (u *Users) IdxSecretCodeAt() int { //nolint:dupl false positive
	return 10
}

// SqlSecretCodeAt return name of the column being indexed
func (u *Users) SqlSecretCodeAt() string { //nolint:dupl false positive
	return `"secretCodeAt"`
}

// IdxVerifiedAt return name of the index
func (u *Users) IdxVerifiedAt() int { //nolint:dupl false positive
	return 11
}

// SqlVerifiedAt return name of the column being indexed
func (u *Users) SqlVerifiedAt() string { //nolint:dupl false positive
	return `"verifiedAt"`
}

// IdxLastLoginAt return name of the index
func (u *Users) IdxLastLoginAt() int { //nolint:dupl false positive
	return 12
}

// SqlLastLoginAt return name of the column being indexed
func (u *Users) SqlLastLoginAt() string { //nolint:dupl false positive
	return `"lastLoginAt"`
}

// IdxFullName return name of the index
func (u *Users) IdxFullName() int { //nolint:dupl false positive
	return 13
}

// SqlFullName return name of the column being indexed
func (u *Users) SqlFullName() string { //nolint:dupl false positive
	return `"fullName"`
}

// IdxUserName return name of the index
func (u *Users) IdxUserName() int { //nolint:dupl false positive
	return 14
}

// SqlUserName return name of the column being indexed
func (u *Users) SqlUserName() string { //nolint:dupl false positive
	return `"userName"`
}

// IdxRole return name of the index
func (u *Users) IdxRole() int { //nolint:dupl false positive
	return 15
}

// SqlRole return name of the column being indexed
func (u *Users) SqlRole() string { //nolint:dupl false positive
	return `"role"`
}

// CensorFields remove sensitive fields for output
func (u *Users) CensorFields() { //nolint:dupl false positive
	u.Password = ``
	u.SecretCode = ``
	u.SecretCodeAt = 0
}

// ToArray receiver fields to slice
func (u *Users) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if u.Id != 0 {
		id = u.Id
	}
	return A.X{
		id,
		u.Email,         // 1
		u.Password,      // 2
		u.CreatedAt,     // 3
		u.CreatedBy,     // 4
		u.UpdatedAt,     // 5
		u.UpdatedBy,     // 6
		u.DeletedAt,     // 7
		u.PasswordSetAt, // 8
		u.SecretCode,    // 9
		u.SecretCodeAt,  // 10
		u.VerifiedAt,    // 11
		u.LastLoginAt,   // 12
		u.FullName,      // 13
		u.UserName,      // 14
		u.Role,          // 15
	}
}

// FromArray convert slice to receiver fields
func (u *Users) FromArray(a A.X) *Users { //nolint:dupl false positive
	u.Id = X.ToU(a[0])
	u.Email = X.ToS(a[1])
	u.Password = X.ToS(a[2])
	u.CreatedAt = X.ToI(a[3])
	u.CreatedBy = X.ToU(a[4])
	u.UpdatedAt = X.ToI(a[5])
	u.UpdatedBy = X.ToU(a[6])
	u.DeletedAt = X.ToI(a[7])
	u.PasswordSetAt = X.ToI(a[8])
	u.SecretCode = X.ToS(a[9])
	u.SecretCodeAt = X.ToI(a[10])
	u.VerifiedAt = X.ToI(a[11])
	u.LastLoginAt = X.ToI(a[12])
	u.FullName = X.ToS(a[13])
	u.UserName = X.ToS(a[14])
	u.Role = X.ToS(a[15])
	return u
}

// FromUncensoredArray convert slice to receiver fields
func (u *Users) FromUncensoredArray(a A.X) *Users { //nolint:dupl false positive
	u.Id = X.ToU(a[0])
	u.Email = X.ToS(a[1])
	u.CreatedAt = X.ToI(a[3])
	u.CreatedBy = X.ToU(a[4])
	u.UpdatedAt = X.ToI(a[5])
	u.UpdatedBy = X.ToU(a[6])
	u.DeletedAt = X.ToI(a[7])
	u.PasswordSetAt = X.ToI(a[8])
	u.VerifiedAt = X.ToI(a[11])
	u.LastLoginAt = X.ToI(a[12])
	u.FullName = X.ToS(a[13])
	u.UserName = X.ToS(a[14])
	u.Role = X.ToS(a[15])
	return u
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (u *Users) FindOffsetLimit(offset, limit uint32, idx string) []Users { //nolint:dupl false positive
	var rows []Users
	res, err := u.Adapter.Select(u.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Users.FindOffsetLimit failed: `+u.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Users{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (u *Users) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := u.Adapter.Select(u.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Users.FindOffsetLimit failed: `+u.SpaceName()) {
		return rows, Tt.QueryMetaFrom(res, err)
	}
	tuples := res.Tuples()
	rows = make([]A.X, len(tuples))
	for z, row := range tuples {
		rows[z] = row
	}
	return rows, Tt.QueryMetaFrom(res, nil)
}

// Total count number of rows
func (u *Users) Total() int64 { //nolint:dupl false positive
	rows := u.Adapter.CallBoxSpace(u.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// UsersFieldTypeMap returns key value of field name and key
var UsersFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:            Tt.Unsigned,
	`email`:         Tt.String,
	`password`:      Tt.String,
	`createdAt`:     Tt.Integer,
	`createdBy`:     Tt.Unsigned,
	`updatedAt`:     Tt.Integer,
	`updatedBy`:     Tt.Unsigned,
	`deletedAt`:     Tt.Integer,
	`passwordSetAt`: Tt.Integer,
	`secretCode`:    Tt.String,
	`secretCodeAt`:  Tt.Integer,
	`verifiedAt`:    Tt.Integer,
	`lastLoginAt`:   Tt.Integer,
	`fullName`:      Tt.String,
	`userName`:      Tt.String,
	`role`:          Tt.String,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
