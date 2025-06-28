package rqProperty

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"kostjc/model/mProperty"

	"github.com/tarantool/go-tarantool"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

// Bookings DAO reader/query struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqProperty__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqProperty__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqProperty__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqProperty__ORM.GEN.go
type Bookings struct {
	Adapter       *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id            uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	DateStart     string      `json:"dateStart" form:"dateStart" query:"dateStart" long:"dateStart" msg:"dateStart"`
	DateEnd       string      `json:"dateEnd" form:"dateEnd" query:"dateEnd" long:"dateEnd" msg:"dateEnd"`
	BasePriceIDR  int64       `json:"basePriceIDR" form:"basePriceIDR" query:"basePriceIDR" long:"basePriceIDR" msg:"basePriceIDR"`
	FacilitiesObj string      `json:"facilitiesObj" form:"facilitiesObj" query:"facilitiesObj" long:"facilitiesObj" msg:"facilitiesObj"`
	TotalPriceIDR int64       `json:"totalPriceIDR" form:"totalPriceIDR" query:"totalPriceIDR" long:"totalPriceIDR" msg:"totalPriceIDR"`
	PaidAt        string      `json:"paidAt" form:"paidAt" query:"paidAt" long:"paidAt" msg:"paidAt"`
	TenantId      uint64      `json:"tenantId,string" form:"tenantId" query:"tenantId" long:"tenantId" msg:"tenantId"`
	CreatedAt     int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy     uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt     int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy     uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt     int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy     uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy    uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
	ExtraTenants  []any       `json:"extraTenants" form:"extraTenants" query:"extraTenants" long:"extraTenants" msg:"extraTenants"`
	RoomId        uint64      `json:"roomId,string" form:"roomId" query:"roomId" long:"roomId" msg:"roomId"`
}

// NewBookings create new ORM reader/query object
func NewBookings(adapter *Tt.Adapter) *Bookings {
	return &Bookings{Adapter: adapter}
}

// SpaceName returns full package and table name
func (b *Bookings) SpaceName() string { //nolint:dupl false positive
	return string(mProperty.TableBookings) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (b *Bookings) SqlTableName() string { //nolint:dupl false positive
	return `"bookings"`
}

func (b *Bookings) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (b *Bookings) FindById() bool { //nolint:dupl false positive
	res, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{b.Id})
	if L.IsError(err, `Bookings.FindById failed: `+b.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		b.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (b *Bookings) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "dateStart"
	, "dateEnd"
	, "basePriceIDR"
	, "facilitiesObj"
	, "totalPriceIDR"
	, "paidAt"
	, "tenantId"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "extraTenants"
	, "roomId"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (b *Bookings) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "dateStart"
	, "dateEnd"
	, "basePriceIDR"
	, "facilitiesObj"
	, "totalPriceIDR"
	, "paidAt"
	, "tenantId"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "extraTenants"
	, "roomId"
	`
}

// ToUpdateArray generate slice of update command
func (b *Bookings) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, b.Id},
		A.X{`=`, 1, b.DateStart},
		A.X{`=`, 2, b.DateEnd},
		A.X{`=`, 3, b.BasePriceIDR},
		A.X{`=`, 4, b.FacilitiesObj},
		A.X{`=`, 5, b.TotalPriceIDR},
		A.X{`=`, 6, b.PaidAt},
		A.X{`=`, 7, b.TenantId},
		A.X{`=`, 8, b.CreatedAt},
		A.X{`=`, 9, b.CreatedBy},
		A.X{`=`, 10, b.UpdatedAt},
		A.X{`=`, 11, b.UpdatedBy},
		A.X{`=`, 12, b.DeletedAt},
		A.X{`=`, 13, b.DeletedBy},
		A.X{`=`, 14, b.RestoredBy},
		A.X{`=`, 15, b.ExtraTenants},
		A.X{`=`, 16, b.RoomId},
	}
}

// IdxId return name of the index
func (b *Bookings) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (b *Bookings) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxDateStart return name of the index
func (b *Bookings) IdxDateStart() int { //nolint:dupl false positive
	return 1
}

// SqlDateStart return name of the column being indexed
func (b *Bookings) SqlDateStart() string { //nolint:dupl false positive
	return `"dateStart"`
}

// IdxDateEnd return name of the index
func (b *Bookings) IdxDateEnd() int { //nolint:dupl false positive
	return 2
}

// SqlDateEnd return name of the column being indexed
func (b *Bookings) SqlDateEnd() string { //nolint:dupl false positive
	return `"dateEnd"`
}

// IdxBasePriceIDR return name of the index
func (b *Bookings) IdxBasePriceIDR() int { //nolint:dupl false positive
	return 3
}

// SqlBasePriceIDR return name of the column being indexed
func (b *Bookings) SqlBasePriceIDR() string { //nolint:dupl false positive
	return `"basePriceIDR"`
}

// IdxFacilitiesObj return name of the index
func (b *Bookings) IdxFacilitiesObj() int { //nolint:dupl false positive
	return 4
}

// SqlFacilitiesObj return name of the column being indexed
func (b *Bookings) SqlFacilitiesObj() string { //nolint:dupl false positive
	return `"facilitiesObj"`
}

// IdxTotalPriceIDR return name of the index
func (b *Bookings) IdxTotalPriceIDR() int { //nolint:dupl false positive
	return 5
}

// SqlTotalPriceIDR return name of the column being indexed
func (b *Bookings) SqlTotalPriceIDR() string { //nolint:dupl false positive
	return `"totalPriceIDR"`
}

// IdxPaidAt return name of the index
func (b *Bookings) IdxPaidAt() int { //nolint:dupl false positive
	return 6
}

// SqlPaidAt return name of the column being indexed
func (b *Bookings) SqlPaidAt() string { //nolint:dupl false positive
	return `"paidAt"`
}

// IdxTenantId return name of the index
func (b *Bookings) IdxTenantId() int { //nolint:dupl false positive
	return 7
}

// SqlTenantId return name of the column being indexed
func (b *Bookings) SqlTenantId() string { //nolint:dupl false positive
	return `"tenantId"`
}

// IdxCreatedAt return name of the index
func (b *Bookings) IdxCreatedAt() int { //nolint:dupl false positive
	return 8
}

// SqlCreatedAt return name of the column being indexed
func (b *Bookings) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (b *Bookings) IdxCreatedBy() int { //nolint:dupl false positive
	return 9
}

// SqlCreatedBy return name of the column being indexed
func (b *Bookings) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (b *Bookings) IdxUpdatedAt() int { //nolint:dupl false positive
	return 10
}

// SqlUpdatedAt return name of the column being indexed
func (b *Bookings) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (b *Bookings) IdxUpdatedBy() int { //nolint:dupl false positive
	return 11
}

// SqlUpdatedBy return name of the column being indexed
func (b *Bookings) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (b *Bookings) IdxDeletedAt() int { //nolint:dupl false positive
	return 12
}

// SqlDeletedAt return name of the column being indexed
func (b *Bookings) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (b *Bookings) IdxDeletedBy() int { //nolint:dupl false positive
	return 13
}

// SqlDeletedBy return name of the column being indexed
func (b *Bookings) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (b *Bookings) IdxRestoredBy() int { //nolint:dupl false positive
	return 14
}

// SqlRestoredBy return name of the column being indexed
func (b *Bookings) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// IdxExtraTenants return name of the index
func (b *Bookings) IdxExtraTenants() int { //nolint:dupl false positive
	return 15
}

// SqlExtraTenants return name of the column being indexed
func (b *Bookings) SqlExtraTenants() string { //nolint:dupl false positive
	return `"extraTenants"`
}

// IdxRoomId return name of the index
func (b *Bookings) IdxRoomId() int { //nolint:dupl false positive
	return 16
}

// SqlRoomId return name of the column being indexed
func (b *Bookings) SqlRoomId() string { //nolint:dupl false positive
	return `"roomId"`
}

// ToArray receiver fields to slice
func (b *Bookings) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if b.Id != 0 {
		id = b.Id
	}
	return A.X{
		id,
		b.DateStart,     // 1
		b.DateEnd,       // 2
		b.BasePriceIDR,  // 3
		b.FacilitiesObj, // 4
		b.TotalPriceIDR, // 5
		b.PaidAt,        // 6
		b.TenantId,      // 7
		b.CreatedAt,     // 8
		b.CreatedBy,     // 9
		b.UpdatedAt,     // 10
		b.UpdatedBy,     // 11
		b.DeletedAt,     // 12
		b.DeletedBy,     // 13
		b.RestoredBy,    // 14
		b.ExtraTenants,  // 15
		b.RoomId,        // 16
	}
}

// FromArray convert slice to receiver fields
func (b *Bookings) FromArray(a A.X) *Bookings { //nolint:dupl false positive
	b.Id = X.ToU(a[0])
	b.DateStart = X.ToS(a[1])
	b.DateEnd = X.ToS(a[2])
	b.BasePriceIDR = X.ToI(a[3])
	b.FacilitiesObj = X.ToS(a[4])
	b.TotalPriceIDR = X.ToI(a[5])
	b.PaidAt = X.ToS(a[6])
	b.TenantId = X.ToU(a[7])
	b.CreatedAt = X.ToI(a[8])
	b.CreatedBy = X.ToU(a[9])
	b.UpdatedAt = X.ToI(a[10])
	b.UpdatedBy = X.ToU(a[11])
	b.DeletedAt = X.ToI(a[12])
	b.DeletedBy = X.ToU(a[13])
	b.RestoredBy = X.ToU(a[14])
	b.ExtraTenants = X.ToArr(a[15])
	b.RoomId = X.ToU(a[16])
	return b
}

// FromUncensoredArray convert slice to receiver fields
func (b *Bookings) FromUncensoredArray(a A.X) *Bookings { //nolint:dupl false positive
	b.Id = X.ToU(a[0])
	b.DateStart = X.ToS(a[1])
	b.DateEnd = X.ToS(a[2])
	b.BasePriceIDR = X.ToI(a[3])
	b.FacilitiesObj = X.ToS(a[4])
	b.TotalPriceIDR = X.ToI(a[5])
	b.PaidAt = X.ToS(a[6])
	b.TenantId = X.ToU(a[7])
	b.CreatedAt = X.ToI(a[8])
	b.CreatedBy = X.ToU(a[9])
	b.UpdatedAt = X.ToI(a[10])
	b.UpdatedBy = X.ToU(a[11])
	b.DeletedAt = X.ToI(a[12])
	b.DeletedBy = X.ToU(a[13])
	b.RestoredBy = X.ToU(a[14])
	b.ExtraTenants = X.ToArr(a[15])
	b.RoomId = X.ToU(a[16])
	return b
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (b *Bookings) FindOffsetLimit(offset, limit uint32, idx string) []Bookings { //nolint:dupl false positive
	var rows []Bookings
	res, err := b.Adapter.Select(b.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Bookings.FindOffsetLimit failed: `+b.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Bookings{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (b *Bookings) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := b.Adapter.Select(b.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Bookings.FindOffsetLimit failed: `+b.SpaceName()) {
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
func (b *Bookings) Total() int64 { //nolint:dupl false positive
	rows := b.Adapter.CallBoxSpace(b.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// BookingsFieldTypeMap returns key value of field name and key
var BookingsFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:            Tt.Unsigned,
	`dateStart`:     Tt.String,
	`dateEnd`:       Tt.String,
	`basePriceIDR`:  Tt.Integer,
	`facilitiesObj`: Tt.String,
	`totalPriceIDR`: Tt.Integer,
	`paidAt`:        Tt.String,
	`tenantId`:      Tt.Unsigned,
	`createdAt`:     Tt.Integer,
	`createdBy`:     Tt.Unsigned,
	`updatedAt`:     Tt.Integer,
	`updatedBy`:     Tt.Unsigned,
	`deletedAt`:     Tt.Integer,
	`deletedBy`:     Tt.Unsigned,
	`restoredBy`:    Tt.Unsigned,
	`extraTenants`:  Tt.Array,
	`roomId`:        Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Buildings DAO reader/query struct
type Buildings struct {
	Adapter      *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id           uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	BuildingName string      `json:"buildingName" form:"buildingName" query:"buildingName" long:"buildingName" msg:"buildingName"`
	LocationId   uint64      `json:"locationId,string" form:"locationId" query:"locationId" long:"locationId" msg:"locationId"`
	Facilities   []any       `json:"facilities" form:"facilities" query:"facilities" long:"facilities" msg:"facilities"`
	CreatedAt    int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy    uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt    int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy    uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt    int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy    uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy   uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
}

// NewBuildings create new ORM reader/query object
func NewBuildings(adapter *Tt.Adapter) *Buildings {
	return &Buildings{Adapter: adapter}
}

// SpaceName returns full package and table name
func (b *Buildings) SpaceName() string { //nolint:dupl false positive
	return string(mProperty.TableBuildings) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (b *Buildings) SqlTableName() string { //nolint:dupl false positive
	return `"buildings"`
}

func (b *Buildings) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (b *Buildings) FindById() bool { //nolint:dupl false positive
	res, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{b.Id})
	if L.IsError(err, `Buildings.FindById failed: `+b.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		b.FromArray(rows[0])
		return true
	}
	return false
}

// UniqueIndexBuildingName return unique index name
func (b *Buildings) UniqueIndexBuildingName() string { //nolint:dupl false positive
	return `buildingName`
}

// FindByBuildingName Find one by BuildingName
func (b *Buildings) FindByBuildingName() bool { //nolint:dupl false positive
	res, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexBuildingName(), 0, 1, tarantool.IterEq, A.X{b.BuildingName})
	if L.IsError(err, `Buildings.FindByBuildingName failed: `+b.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		b.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (b *Buildings) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "buildingName"
	, "locationId"
	, "facilities"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (b *Buildings) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "buildingName"
	, "locationId"
	, "facilities"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// ToUpdateArray generate slice of update command
func (b *Buildings) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, b.Id},
		A.X{`=`, 1, b.BuildingName},
		A.X{`=`, 2, b.LocationId},
		A.X{`=`, 3, b.Facilities},
		A.X{`=`, 4, b.CreatedAt},
		A.X{`=`, 5, b.CreatedBy},
		A.X{`=`, 6, b.UpdatedAt},
		A.X{`=`, 7, b.UpdatedBy},
		A.X{`=`, 8, b.DeletedAt},
		A.X{`=`, 9, b.DeletedBy},
		A.X{`=`, 10, b.RestoredBy},
	}
}

// IdxId return name of the index
func (b *Buildings) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (b *Buildings) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxBuildingName return name of the index
func (b *Buildings) IdxBuildingName() int { //nolint:dupl false positive
	return 1
}

// SqlBuildingName return name of the column being indexed
func (b *Buildings) SqlBuildingName() string { //nolint:dupl false positive
	return `"buildingName"`
}

// IdxLocationId return name of the index
func (b *Buildings) IdxLocationId() int { //nolint:dupl false positive
	return 2
}

// SqlLocationId return name of the column being indexed
func (b *Buildings) SqlLocationId() string { //nolint:dupl false positive
	return `"locationId"`
}

// IdxFacilities return name of the index
func (b *Buildings) IdxFacilities() int { //nolint:dupl false positive
	return 3
}

// SqlFacilities return name of the column being indexed
func (b *Buildings) SqlFacilities() string { //nolint:dupl false positive
	return `"facilities"`
}

// IdxCreatedAt return name of the index
func (b *Buildings) IdxCreatedAt() int { //nolint:dupl false positive
	return 4
}

// SqlCreatedAt return name of the column being indexed
func (b *Buildings) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (b *Buildings) IdxCreatedBy() int { //nolint:dupl false positive
	return 5
}

// SqlCreatedBy return name of the column being indexed
func (b *Buildings) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (b *Buildings) IdxUpdatedAt() int { //nolint:dupl false positive
	return 6
}

// SqlUpdatedAt return name of the column being indexed
func (b *Buildings) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (b *Buildings) IdxUpdatedBy() int { //nolint:dupl false positive
	return 7
}

// SqlUpdatedBy return name of the column being indexed
func (b *Buildings) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (b *Buildings) IdxDeletedAt() int { //nolint:dupl false positive
	return 8
}

// SqlDeletedAt return name of the column being indexed
func (b *Buildings) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (b *Buildings) IdxDeletedBy() int { //nolint:dupl false positive
	return 9
}

// SqlDeletedBy return name of the column being indexed
func (b *Buildings) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (b *Buildings) IdxRestoredBy() int { //nolint:dupl false positive
	return 10
}

// SqlRestoredBy return name of the column being indexed
func (b *Buildings) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// ToArray receiver fields to slice
func (b *Buildings) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if b.Id != 0 {
		id = b.Id
	}
	return A.X{
		id,
		b.BuildingName, // 1
		b.LocationId,   // 2
		b.Facilities,   // 3
		b.CreatedAt,    // 4
		b.CreatedBy,    // 5
		b.UpdatedAt,    // 6
		b.UpdatedBy,    // 7
		b.DeletedAt,    // 8
		b.DeletedBy,    // 9
		b.RestoredBy,   // 10
	}
}

// FromArray convert slice to receiver fields
func (b *Buildings) FromArray(a A.X) *Buildings { //nolint:dupl false positive
	b.Id = X.ToU(a[0])
	b.BuildingName = X.ToS(a[1])
	b.LocationId = X.ToU(a[2])
	b.Facilities = X.ToArr(a[3])
	b.CreatedAt = X.ToI(a[4])
	b.CreatedBy = X.ToU(a[5])
	b.UpdatedAt = X.ToI(a[6])
	b.UpdatedBy = X.ToU(a[7])
	b.DeletedAt = X.ToI(a[8])
	b.DeletedBy = X.ToU(a[9])
	b.RestoredBy = X.ToU(a[10])
	return b
}

// FromUncensoredArray convert slice to receiver fields
func (b *Buildings) FromUncensoredArray(a A.X) *Buildings { //nolint:dupl false positive
	b.Id = X.ToU(a[0])
	b.BuildingName = X.ToS(a[1])
	b.LocationId = X.ToU(a[2])
	b.Facilities = X.ToArr(a[3])
	b.CreatedAt = X.ToI(a[4])
	b.CreatedBy = X.ToU(a[5])
	b.UpdatedAt = X.ToI(a[6])
	b.UpdatedBy = X.ToU(a[7])
	b.DeletedAt = X.ToI(a[8])
	b.DeletedBy = X.ToU(a[9])
	b.RestoredBy = X.ToU(a[10])
	return b
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (b *Buildings) FindOffsetLimit(offset, limit uint32, idx string) []Buildings { //nolint:dupl false positive
	var rows []Buildings
	res, err := b.Adapter.Select(b.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Buildings.FindOffsetLimit failed: `+b.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Buildings{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (b *Buildings) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := b.Adapter.Select(b.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Buildings.FindOffsetLimit failed: `+b.SpaceName()) {
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
func (b *Buildings) Total() int64 { //nolint:dupl false positive
	rows := b.Adapter.CallBoxSpace(b.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// BuildingsFieldTypeMap returns key value of field name and key
var BuildingsFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:           Tt.Unsigned,
	`buildingName`: Tt.String,
	`locationId`:   Tt.Unsigned,
	`facilities`:   Tt.Array,
	`createdAt`:    Tt.Integer,
	`createdBy`:    Tt.Unsigned,
	`updatedAt`:    Tt.Integer,
	`updatedBy`:    Tt.Unsigned,
	`deletedAt`:    Tt.Integer,
	`deletedBy`:    Tt.Unsigned,
	`restoredBy`:   Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Facilities DAO reader/query struct
type Facilities struct {
	Adapter        *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id             uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	FacilityName   string      `json:"facilityName" form:"facilityName" query:"facilityName" long:"facilityName" msg:"facilityName"`
	ExtraChargeIDR int64       `json:"extraChargeIDR" form:"extraChargeIDR" query:"extraChargeIDR" long:"extraChargeIDR" msg:"extraChargeIDR"`
	CreatedAt      int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy      uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt      int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy      uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt      int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy      uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy     uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
	FacilityType   string      `json:"facilityType" form:"facilityType" query:"facilityType" long:"facilityType" msg:"facilityType"`
	DescriptionEN  string      `json:"descriptionEN" form:"descriptionEN" query:"descriptionEN" long:"descriptionEN" msg:"descriptionEN"`
}

// NewFacilities create new ORM reader/query object
func NewFacilities(adapter *Tt.Adapter) *Facilities {
	return &Facilities{Adapter: adapter}
}

// SpaceName returns full package and table name
func (f *Facilities) SpaceName() string { //nolint:dupl false positive
	return string(mProperty.TableFacilities) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (f *Facilities) SqlTableName() string { //nolint:dupl false positive
	return `"facilities"`
}

func (f *Facilities) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (f *Facilities) FindById() bool { //nolint:dupl false positive
	res, err := f.Adapter.Select(f.SpaceName(), f.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{f.Id})
	if L.IsError(err, `Facilities.FindById failed: `+f.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		f.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (f *Facilities) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "facilityName"
	, "extraChargeIDR"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "facilityType"
	, "descriptionEN"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (f *Facilities) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "facilityName"
	, "extraChargeIDR"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "facilityType"
	, "descriptionEN"
	`
}

// ToUpdateArray generate slice of update command
func (f *Facilities) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, f.Id},
		A.X{`=`, 1, f.FacilityName},
		A.X{`=`, 2, f.ExtraChargeIDR},
		A.X{`=`, 3, f.CreatedAt},
		A.X{`=`, 4, f.CreatedBy},
		A.X{`=`, 5, f.UpdatedAt},
		A.X{`=`, 6, f.UpdatedBy},
		A.X{`=`, 7, f.DeletedAt},
		A.X{`=`, 8, f.DeletedBy},
		A.X{`=`, 9, f.RestoredBy},
		A.X{`=`, 10, f.FacilityType},
		A.X{`=`, 11, f.DescriptionEN},
	}
}

// IdxId return name of the index
func (f *Facilities) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (f *Facilities) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxFacilityName return name of the index
func (f *Facilities) IdxFacilityName() int { //nolint:dupl false positive
	return 1
}

// SqlFacilityName return name of the column being indexed
func (f *Facilities) SqlFacilityName() string { //nolint:dupl false positive
	return `"facilityName"`
}

// IdxExtraChargeIDR return name of the index
func (f *Facilities) IdxExtraChargeIDR() int { //nolint:dupl false positive
	return 2
}

// SqlExtraChargeIDR return name of the column being indexed
func (f *Facilities) SqlExtraChargeIDR() string { //nolint:dupl false positive
	return `"extraChargeIDR"`
}

// IdxCreatedAt return name of the index
func (f *Facilities) IdxCreatedAt() int { //nolint:dupl false positive
	return 3
}

// SqlCreatedAt return name of the column being indexed
func (f *Facilities) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (f *Facilities) IdxCreatedBy() int { //nolint:dupl false positive
	return 4
}

// SqlCreatedBy return name of the column being indexed
func (f *Facilities) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (f *Facilities) IdxUpdatedAt() int { //nolint:dupl false positive
	return 5
}

// SqlUpdatedAt return name of the column being indexed
func (f *Facilities) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (f *Facilities) IdxUpdatedBy() int { //nolint:dupl false positive
	return 6
}

// SqlUpdatedBy return name of the column being indexed
func (f *Facilities) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (f *Facilities) IdxDeletedAt() int { //nolint:dupl false positive
	return 7
}

// SqlDeletedAt return name of the column being indexed
func (f *Facilities) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (f *Facilities) IdxDeletedBy() int { //nolint:dupl false positive
	return 8
}

// SqlDeletedBy return name of the column being indexed
func (f *Facilities) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (f *Facilities) IdxRestoredBy() int { //nolint:dupl false positive
	return 9
}

// SqlRestoredBy return name of the column being indexed
func (f *Facilities) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// IdxFacilityType return name of the index
func (f *Facilities) IdxFacilityType() int { //nolint:dupl false positive
	return 10
}

// SqlFacilityType return name of the column being indexed
func (f *Facilities) SqlFacilityType() string { //nolint:dupl false positive
	return `"facilityType"`
}

// IdxDescriptionEN return name of the index
func (f *Facilities) IdxDescriptionEN() int { //nolint:dupl false positive
	return 11
}

// SqlDescriptionEN return name of the column being indexed
func (f *Facilities) SqlDescriptionEN() string { //nolint:dupl false positive
	return `"descriptionEN"`
}

// ToArray receiver fields to slice
func (f *Facilities) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if f.Id != 0 {
		id = f.Id
	}
	return A.X{
		id,
		f.FacilityName,   // 1
		f.ExtraChargeIDR, // 2
		f.CreatedAt,      // 3
		f.CreatedBy,      // 4
		f.UpdatedAt,      // 5
		f.UpdatedBy,      // 6
		f.DeletedAt,      // 7
		f.DeletedBy,      // 8
		f.RestoredBy,     // 9
		f.FacilityType,   // 10
		f.DescriptionEN,  // 11
	}
}

// FromArray convert slice to receiver fields
func (f *Facilities) FromArray(a A.X) *Facilities { //nolint:dupl false positive
	f.Id = X.ToU(a[0])
	f.FacilityName = X.ToS(a[1])
	f.ExtraChargeIDR = X.ToI(a[2])
	f.CreatedAt = X.ToI(a[3])
	f.CreatedBy = X.ToU(a[4])
	f.UpdatedAt = X.ToI(a[5])
	f.UpdatedBy = X.ToU(a[6])
	f.DeletedAt = X.ToI(a[7])
	f.DeletedBy = X.ToU(a[8])
	f.RestoredBy = X.ToU(a[9])
	f.FacilityType = X.ToS(a[10])
	f.DescriptionEN = X.ToS(a[11])
	return f
}

// FromUncensoredArray convert slice to receiver fields
func (f *Facilities) FromUncensoredArray(a A.X) *Facilities { //nolint:dupl false positive
	f.Id = X.ToU(a[0])
	f.FacilityName = X.ToS(a[1])
	f.ExtraChargeIDR = X.ToI(a[2])
	f.CreatedAt = X.ToI(a[3])
	f.CreatedBy = X.ToU(a[4])
	f.UpdatedAt = X.ToI(a[5])
	f.UpdatedBy = X.ToU(a[6])
	f.DeletedAt = X.ToI(a[7])
	f.DeletedBy = X.ToU(a[8])
	f.RestoredBy = X.ToU(a[9])
	f.FacilityType = X.ToS(a[10])
	f.DescriptionEN = X.ToS(a[11])
	return f
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (f *Facilities) FindOffsetLimit(offset, limit uint32, idx string) []Facilities { //nolint:dupl false positive
	var rows []Facilities
	res, err := f.Adapter.Select(f.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Facilities.FindOffsetLimit failed: `+f.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Facilities{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (f *Facilities) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := f.Adapter.Select(f.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Facilities.FindOffsetLimit failed: `+f.SpaceName()) {
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
func (f *Facilities) Total() int64 { //nolint:dupl false positive
	rows := f.Adapter.CallBoxSpace(f.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// FacilitiesFieldTypeMap returns key value of field name and key
var FacilitiesFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:             Tt.Unsigned,
	`facilityName`:   Tt.String,
	`extraChargeIDR`: Tt.Integer,
	`createdAt`:      Tt.Integer,
	`createdBy`:      Tt.Unsigned,
	`updatedAt`:      Tt.Integer,
	`updatedBy`:      Tt.Unsigned,
	`deletedAt`:      Tt.Integer,
	`deletedBy`:      Tt.Unsigned,
	`restoredBy`:     Tt.Unsigned,
	`facilityType`:   Tt.String,
	`descriptionEN`:  Tt.String,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Locations DAO reader/query struct
type Locations struct {
	Adapter      *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id           uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	Name         string      `json:"name" form:"name" query:"name" long:"name" msg:"name"`
	Address      string      `json:"address" form:"address" query:"address" long:"address" msg:"address"`
	GmapLocation string      `json:"gmapLocation" form:"gmapLocation" query:"gmapLocation" long:"gmapLocation" msg:"gmapLocation"`
	CreatedAt    int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy    uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt    int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy    uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt    int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy    uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy   uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
}

// NewLocations create new ORM reader/query object
func NewLocations(adapter *Tt.Adapter) *Locations {
	return &Locations{Adapter: adapter}
}

// SpaceName returns full package and table name
func (l *Locations) SpaceName() string { //nolint:dupl false positive
	return string(mProperty.TableLocations) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (l *Locations) SqlTableName() string { //nolint:dupl false positive
	return `"locations"`
}

func (l *Locations) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (l *Locations) FindById() bool { //nolint:dupl false positive
	res, err := l.Adapter.Select(l.SpaceName(), l.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{l.Id})
	if L.IsError(err, `Locations.FindById failed: `+l.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		l.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (l *Locations) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "name"
	, "address"
	, "gmapLocation"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (l *Locations) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "name"
	, "address"
	, "gmapLocation"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// ToUpdateArray generate slice of update command
func (l *Locations) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, l.Id},
		A.X{`=`, 1, l.Name},
		A.X{`=`, 2, l.Address},
		A.X{`=`, 3, l.GmapLocation},
		A.X{`=`, 4, l.CreatedAt},
		A.X{`=`, 5, l.CreatedBy},
		A.X{`=`, 6, l.UpdatedAt},
		A.X{`=`, 7, l.UpdatedBy},
		A.X{`=`, 8, l.DeletedAt},
		A.X{`=`, 9, l.DeletedBy},
		A.X{`=`, 10, l.RestoredBy},
	}
}

// IdxId return name of the index
func (l *Locations) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (l *Locations) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxName return name of the index
func (l *Locations) IdxName() int { //nolint:dupl false positive
	return 1
}

// SqlName return name of the column being indexed
func (l *Locations) SqlName() string { //nolint:dupl false positive
	return `"name"`
}

// IdxAddress return name of the index
func (l *Locations) IdxAddress() int { //nolint:dupl false positive
	return 2
}

// SqlAddress return name of the column being indexed
func (l *Locations) SqlAddress() string { //nolint:dupl false positive
	return `"address"`
}

// IdxGmapLocation return name of the index
func (l *Locations) IdxGmapLocation() int { //nolint:dupl false positive
	return 3
}

// SqlGmapLocation return name of the column being indexed
func (l *Locations) SqlGmapLocation() string { //nolint:dupl false positive
	return `"gmapLocation"`
}

// IdxCreatedAt return name of the index
func (l *Locations) IdxCreatedAt() int { //nolint:dupl false positive
	return 4
}

// SqlCreatedAt return name of the column being indexed
func (l *Locations) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (l *Locations) IdxCreatedBy() int { //nolint:dupl false positive
	return 5
}

// SqlCreatedBy return name of the column being indexed
func (l *Locations) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (l *Locations) IdxUpdatedAt() int { //nolint:dupl false positive
	return 6
}

// SqlUpdatedAt return name of the column being indexed
func (l *Locations) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (l *Locations) IdxUpdatedBy() int { //nolint:dupl false positive
	return 7
}

// SqlUpdatedBy return name of the column being indexed
func (l *Locations) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (l *Locations) IdxDeletedAt() int { //nolint:dupl false positive
	return 8
}

// SqlDeletedAt return name of the column being indexed
func (l *Locations) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (l *Locations) IdxDeletedBy() int { //nolint:dupl false positive
	return 9
}

// SqlDeletedBy return name of the column being indexed
func (l *Locations) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (l *Locations) IdxRestoredBy() int { //nolint:dupl false positive
	return 10
}

// SqlRestoredBy return name of the column being indexed
func (l *Locations) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// ToArray receiver fields to slice
func (l *Locations) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if l.Id != 0 {
		id = l.Id
	}
	return A.X{
		id,
		l.Name,         // 1
		l.Address,      // 2
		l.GmapLocation, // 3
		l.CreatedAt,    // 4
		l.CreatedBy,    // 5
		l.UpdatedAt,    // 6
		l.UpdatedBy,    // 7
		l.DeletedAt,    // 8
		l.DeletedBy,    // 9
		l.RestoredBy,   // 10
	}
}

// FromArray convert slice to receiver fields
func (l *Locations) FromArray(a A.X) *Locations { //nolint:dupl false positive
	l.Id = X.ToU(a[0])
	l.Name = X.ToS(a[1])
	l.Address = X.ToS(a[2])
	l.GmapLocation = X.ToS(a[3])
	l.CreatedAt = X.ToI(a[4])
	l.CreatedBy = X.ToU(a[5])
	l.UpdatedAt = X.ToI(a[6])
	l.UpdatedBy = X.ToU(a[7])
	l.DeletedAt = X.ToI(a[8])
	l.DeletedBy = X.ToU(a[9])
	l.RestoredBy = X.ToU(a[10])
	return l
}

// FromUncensoredArray convert slice to receiver fields
func (l *Locations) FromUncensoredArray(a A.X) *Locations { //nolint:dupl false positive
	l.Id = X.ToU(a[0])
	l.Name = X.ToS(a[1])
	l.Address = X.ToS(a[2])
	l.GmapLocation = X.ToS(a[3])
	l.CreatedAt = X.ToI(a[4])
	l.CreatedBy = X.ToU(a[5])
	l.UpdatedAt = X.ToI(a[6])
	l.UpdatedBy = X.ToU(a[7])
	l.DeletedAt = X.ToI(a[8])
	l.DeletedBy = X.ToU(a[9])
	l.RestoredBy = X.ToU(a[10])
	return l
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (l *Locations) FindOffsetLimit(offset, limit uint32, idx string) []Locations { //nolint:dupl false positive
	var rows []Locations
	res, err := l.Adapter.Select(l.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Locations.FindOffsetLimit failed: `+l.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Locations{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (l *Locations) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := l.Adapter.Select(l.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Locations.FindOffsetLimit failed: `+l.SpaceName()) {
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
func (l *Locations) Total() int64 { //nolint:dupl false positive
	rows := l.Adapter.CallBoxSpace(l.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// LocationsFieldTypeMap returns key value of field name and key
var LocationsFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:           Tt.Unsigned,
	`name`:         Tt.String,
	`address`:      Tt.String,
	`gmapLocation`: Tt.String,
	`createdAt`:    Tt.Integer,
	`createdBy`:    Tt.Unsigned,
	`updatedAt`:    Tt.Integer,
	`updatedBy`:    Tt.Unsigned,
	`deletedAt`:    Tt.Integer,
	`deletedBy`:    Tt.Unsigned,
	`restoredBy`:   Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Payments DAO reader/query struct
type Payments struct {
	Adapter       *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id            uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	BookingId     uint64      `json:"bookingId,string" form:"bookingId" query:"bookingId" long:"bookingId" msg:"bookingId"`
	PaymentAt     string      `json:"paymentAt" form:"paymentAt" query:"paymentAt" long:"paymentAt" msg:"paymentAt"`
	PaidIDR       int64       `json:"paidIDR" form:"paidIDR" query:"paidIDR" long:"paidIDR" msg:"paidIDR"`
	PaymentMethod string      `json:"paymentMethod" form:"paymentMethod" query:"paymentMethod" long:"paymentMethod" msg:"paymentMethod"`
	PaymentStatus string      `json:"paymentStatus" form:"paymentStatus" query:"paymentStatus" long:"paymentStatus" msg:"paymentStatus"`
	Note          string      `json:"note" form:"note" query:"note" long:"note" msg:"note"`
	CreatedAt     int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy     uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt     int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy     uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt     int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy     uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy    uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
}

// NewPayments create new ORM reader/query object
func NewPayments(adapter *Tt.Adapter) *Payments {
	return &Payments{Adapter: adapter}
}

// SpaceName returns full package and table name
func (p *Payments) SpaceName() string { //nolint:dupl false positive
	return string(mProperty.TablePayments) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (p *Payments) SqlTableName() string { //nolint:dupl false positive
	return `"payments"`
}

func (p *Payments) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (p *Payments) FindById() bool { //nolint:dupl false positive
	res, err := p.Adapter.Select(p.SpaceName(), p.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{p.Id})
	if L.IsError(err, `Payments.FindById failed: `+p.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		p.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (p *Payments) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "bookingId"
	, "paymentAt"
	, "paidIDR"
	, "paymentMethod"
	, "paymentStatus"
	, "note"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (p *Payments) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "bookingId"
	, "paymentAt"
	, "paidIDR"
	, "paymentMethod"
	, "paymentStatus"
	, "note"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// ToUpdateArray generate slice of update command
func (p *Payments) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, p.Id},
		A.X{`=`, 1, p.BookingId},
		A.X{`=`, 2, p.PaymentAt},
		A.X{`=`, 3, p.PaidIDR},
		A.X{`=`, 4, p.PaymentMethod},
		A.X{`=`, 5, p.PaymentStatus},
		A.X{`=`, 6, p.Note},
		A.X{`=`, 7, p.CreatedAt},
		A.X{`=`, 8, p.CreatedBy},
		A.X{`=`, 9, p.UpdatedAt},
		A.X{`=`, 10, p.UpdatedBy},
		A.X{`=`, 11, p.DeletedAt},
		A.X{`=`, 12, p.DeletedBy},
		A.X{`=`, 13, p.RestoredBy},
	}
}

// IdxId return name of the index
func (p *Payments) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (p *Payments) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxBookingId return name of the index
func (p *Payments) IdxBookingId() int { //nolint:dupl false positive
	return 1
}

// SqlBookingId return name of the column being indexed
func (p *Payments) SqlBookingId() string { //nolint:dupl false positive
	return `"bookingId"`
}

// IdxPaymentAt return name of the index
func (p *Payments) IdxPaymentAt() int { //nolint:dupl false positive
	return 2
}

// SqlPaymentAt return name of the column being indexed
func (p *Payments) SqlPaymentAt() string { //nolint:dupl false positive
	return `"paymentAt"`
}

// IdxPaidIDR return name of the index
func (p *Payments) IdxPaidIDR() int { //nolint:dupl false positive
	return 3
}

// SqlPaidIDR return name of the column being indexed
func (p *Payments) SqlPaidIDR() string { //nolint:dupl false positive
	return `"paidIDR"`
}

// IdxPaymentMethod return name of the index
func (p *Payments) IdxPaymentMethod() int { //nolint:dupl false positive
	return 4
}

// SqlPaymentMethod return name of the column being indexed
func (p *Payments) SqlPaymentMethod() string { //nolint:dupl false positive
	return `"paymentMethod"`
}

// IdxPaymentStatus return name of the index
func (p *Payments) IdxPaymentStatus() int { //nolint:dupl false positive
	return 5
}

// SqlPaymentStatus return name of the column being indexed
func (p *Payments) SqlPaymentStatus() string { //nolint:dupl false positive
	return `"paymentStatus"`
}

// IdxNote return name of the index
func (p *Payments) IdxNote() int { //nolint:dupl false positive
	return 6
}

// SqlNote return name of the column being indexed
func (p *Payments) SqlNote() string { //nolint:dupl false positive
	return `"note"`
}

// IdxCreatedAt return name of the index
func (p *Payments) IdxCreatedAt() int { //nolint:dupl false positive
	return 7
}

// SqlCreatedAt return name of the column being indexed
func (p *Payments) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (p *Payments) IdxCreatedBy() int { //nolint:dupl false positive
	return 8
}

// SqlCreatedBy return name of the column being indexed
func (p *Payments) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (p *Payments) IdxUpdatedAt() int { //nolint:dupl false positive
	return 9
}

// SqlUpdatedAt return name of the column being indexed
func (p *Payments) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (p *Payments) IdxUpdatedBy() int { //nolint:dupl false positive
	return 10
}

// SqlUpdatedBy return name of the column being indexed
func (p *Payments) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (p *Payments) IdxDeletedAt() int { //nolint:dupl false positive
	return 11
}

// SqlDeletedAt return name of the column being indexed
func (p *Payments) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (p *Payments) IdxDeletedBy() int { //nolint:dupl false positive
	return 12
}

// SqlDeletedBy return name of the column being indexed
func (p *Payments) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (p *Payments) IdxRestoredBy() int { //nolint:dupl false positive
	return 13
}

// SqlRestoredBy return name of the column being indexed
func (p *Payments) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// ToArray receiver fields to slice
func (p *Payments) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if p.Id != 0 {
		id = p.Id
	}
	return A.X{
		id,
		p.BookingId,     // 1
		p.PaymentAt,     // 2
		p.PaidIDR,       // 3
		p.PaymentMethod, // 4
		p.PaymentStatus, // 5
		p.Note,          // 6
		p.CreatedAt,     // 7
		p.CreatedBy,     // 8
		p.UpdatedAt,     // 9
		p.UpdatedBy,     // 10
		p.DeletedAt,     // 11
		p.DeletedBy,     // 12
		p.RestoredBy,    // 13
	}
}

// FromArray convert slice to receiver fields
func (p *Payments) FromArray(a A.X) *Payments { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.BookingId = X.ToU(a[1])
	p.PaymentAt = X.ToS(a[2])
	p.PaidIDR = X.ToI(a[3])
	p.PaymentMethod = X.ToS(a[4])
	p.PaymentStatus = X.ToS(a[5])
	p.Note = X.ToS(a[6])
	p.CreatedAt = X.ToI(a[7])
	p.CreatedBy = X.ToU(a[8])
	p.UpdatedAt = X.ToI(a[9])
	p.UpdatedBy = X.ToU(a[10])
	p.DeletedAt = X.ToI(a[11])
	p.DeletedBy = X.ToU(a[12])
	p.RestoredBy = X.ToU(a[13])
	return p
}

// FromUncensoredArray convert slice to receiver fields
func (p *Payments) FromUncensoredArray(a A.X) *Payments { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.BookingId = X.ToU(a[1])
	p.PaymentAt = X.ToS(a[2])
	p.PaidIDR = X.ToI(a[3])
	p.PaymentMethod = X.ToS(a[4])
	p.PaymentStatus = X.ToS(a[5])
	p.Note = X.ToS(a[6])
	p.CreatedAt = X.ToI(a[7])
	p.CreatedBy = X.ToU(a[8])
	p.UpdatedAt = X.ToI(a[9])
	p.UpdatedBy = X.ToU(a[10])
	p.DeletedAt = X.ToI(a[11])
	p.DeletedBy = X.ToU(a[12])
	p.RestoredBy = X.ToU(a[13])
	return p
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (p *Payments) FindOffsetLimit(offset, limit uint32, idx string) []Payments { //nolint:dupl false positive
	var rows []Payments
	res, err := p.Adapter.Select(p.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Payments.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Payments{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (p *Payments) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := p.Adapter.Select(p.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Payments.FindOffsetLimit failed: `+p.SpaceName()) {
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
func (p *Payments) Total() int64 { //nolint:dupl false positive
	rows := p.Adapter.CallBoxSpace(p.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// PaymentsFieldTypeMap returns key value of field name and key
var PaymentsFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:            Tt.Unsigned,
	`bookingId`:     Tt.Unsigned,
	`paymentAt`:     Tt.String,
	`paidIDR`:       Tt.Integer,
	`paymentMethod`: Tt.String,
	`paymentStatus`: Tt.String,
	`note`:          Tt.String,
	`createdAt`:     Tt.Integer,
	`createdBy`:     Tt.Unsigned,
	`updatedAt`:     Tt.Integer,
	`updatedBy`:     Tt.Unsigned,
	`deletedAt`:     Tt.Integer,
	`deletedBy`:     Tt.Unsigned,
	`restoredBy`:    Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Rooms DAO reader/query struct
type Rooms struct {
	Adapter         *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id              uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	RoomName        string      `json:"roomName" form:"roomName" query:"roomName" long:"roomName" msg:"roomName"`
	BasePriceIDR    int64       `json:"basePriceIDR" form:"basePriceIDR" query:"basePriceIDR" long:"basePriceIDR" msg:"basePriceIDR"`
	CurrentTenantId uint64      `json:"currentTenantId,string" form:"currentTenantId" query:"currentTenantId" long:"currentTenantId" msg:"currentTenantId"`
	FirstUseAt      string      `json:"firstUseAt" form:"firstUseAt" query:"firstUseAt" long:"firstUseAt" msg:"firstUseAt"`
	CreatedAt       int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy       uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt       int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy       uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt       int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy       uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy      uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
	BuildingId      uint64      `json:"buildingId,string" form:"buildingId" query:"buildingId" long:"buildingId" msg:"buildingId"`
	LastUseAt       string      `json:"lastUseAt" form:"lastUseAt" query:"lastUseAt" long:"lastUseAt" msg:"lastUseAt"`
	RoomSize        string      `json:"roomSize" form:"roomSize" query:"roomSize" long:"roomSize" msg:"roomSize"`
	ImageUrl        string      `json:"imageUrl" form:"imageUrl" query:"imageUrl" long:"imageUrl" msg:"imageUrl"`
}

// NewRooms create new ORM reader/query object
func NewRooms(adapter *Tt.Adapter) *Rooms {
	return &Rooms{Adapter: adapter}
}

// SpaceName returns full package and table name
func (r *Rooms) SpaceName() string { //nolint:dupl false positive
	return string(mProperty.TableRooms) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (r *Rooms) SqlTableName() string { //nolint:dupl false positive
	return `"rooms"`
}

func (r *Rooms) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (r *Rooms) FindById() bool { //nolint:dupl false positive
	res, err := r.Adapter.Select(r.SpaceName(), r.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{r.Id})
	if L.IsError(err, `Rooms.FindById failed: `+r.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		r.FromArray(rows[0])
		return true
	}
	return false
}

// UniqueIndexRoomName return unique index name
func (r *Rooms) UniqueIndexRoomName() string { //nolint:dupl false positive
	return `roomName`
}

// FindByRoomName Find one by RoomName
func (r *Rooms) FindByRoomName() bool { //nolint:dupl false positive
	res, err := r.Adapter.Select(r.SpaceName(), r.UniqueIndexRoomName(), 0, 1, tarantool.IterEq, A.X{r.RoomName})
	if L.IsError(err, `Rooms.FindByRoomName failed: `+r.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		r.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (r *Rooms) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "roomName"
	, "basePriceIDR"
	, "currentTenantId"
	, "firstUseAt"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "buildingId"
	, "lastUseAt"
	, "roomSize"
	, "imageUrl"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (r *Rooms) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "roomName"
	, "basePriceIDR"
	, "currentTenantId"
	, "firstUseAt"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "buildingId"
	, "lastUseAt"
	, "roomSize"
	, "imageUrl"
	`
}

// ToUpdateArray generate slice of update command
func (r *Rooms) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, r.Id},
		A.X{`=`, 1, r.RoomName},
		A.X{`=`, 2, r.BasePriceIDR},
		A.X{`=`, 3, r.CurrentTenantId},
		A.X{`=`, 4, r.FirstUseAt},
		A.X{`=`, 5, r.CreatedAt},
		A.X{`=`, 6, r.CreatedBy},
		A.X{`=`, 7, r.UpdatedAt},
		A.X{`=`, 8, r.UpdatedBy},
		A.X{`=`, 9, r.DeletedAt},
		A.X{`=`, 10, r.DeletedBy},
		A.X{`=`, 11, r.RestoredBy},
		A.X{`=`, 12, r.BuildingId},
		A.X{`=`, 13, r.LastUseAt},
		A.X{`=`, 14, r.RoomSize},
		A.X{`=`, 15, r.ImageUrl},
	}
}

// IdxId return name of the index
func (r *Rooms) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (r *Rooms) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxRoomName return name of the index
func (r *Rooms) IdxRoomName() int { //nolint:dupl false positive
	return 1
}

// SqlRoomName return name of the column being indexed
func (r *Rooms) SqlRoomName() string { //nolint:dupl false positive
	return `"roomName"`
}

// IdxBasePriceIDR return name of the index
func (r *Rooms) IdxBasePriceIDR() int { //nolint:dupl false positive
	return 2
}

// SqlBasePriceIDR return name of the column being indexed
func (r *Rooms) SqlBasePriceIDR() string { //nolint:dupl false positive
	return `"basePriceIDR"`
}

// IdxCurrentTenantId return name of the index
func (r *Rooms) IdxCurrentTenantId() int { //nolint:dupl false positive
	return 3
}

// SqlCurrentTenantId return name of the column being indexed
func (r *Rooms) SqlCurrentTenantId() string { //nolint:dupl false positive
	return `"currentTenantId"`
}

// IdxFirstUseAt return name of the index
func (r *Rooms) IdxFirstUseAt() int { //nolint:dupl false positive
	return 4
}

// SqlFirstUseAt return name of the column being indexed
func (r *Rooms) SqlFirstUseAt() string { //nolint:dupl false positive
	return `"firstUseAt"`
}

// IdxCreatedAt return name of the index
func (r *Rooms) IdxCreatedAt() int { //nolint:dupl false positive
	return 5
}

// SqlCreatedAt return name of the column being indexed
func (r *Rooms) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (r *Rooms) IdxCreatedBy() int { //nolint:dupl false positive
	return 6
}

// SqlCreatedBy return name of the column being indexed
func (r *Rooms) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (r *Rooms) IdxUpdatedAt() int { //nolint:dupl false positive
	return 7
}

// SqlUpdatedAt return name of the column being indexed
func (r *Rooms) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (r *Rooms) IdxUpdatedBy() int { //nolint:dupl false positive
	return 8
}

// SqlUpdatedBy return name of the column being indexed
func (r *Rooms) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (r *Rooms) IdxDeletedAt() int { //nolint:dupl false positive
	return 9
}

// SqlDeletedAt return name of the column being indexed
func (r *Rooms) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (r *Rooms) IdxDeletedBy() int { //nolint:dupl false positive
	return 10
}

// SqlDeletedBy return name of the column being indexed
func (r *Rooms) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (r *Rooms) IdxRestoredBy() int { //nolint:dupl false positive
	return 11
}

// SqlRestoredBy return name of the column being indexed
func (r *Rooms) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// IdxBuildingId return name of the index
func (r *Rooms) IdxBuildingId() int { //nolint:dupl false positive
	return 12
}

// SqlBuildingId return name of the column being indexed
func (r *Rooms) SqlBuildingId() string { //nolint:dupl false positive
	return `"buildingId"`
}

// IdxLastUseAt return name of the index
func (r *Rooms) IdxLastUseAt() int { //nolint:dupl false positive
	return 13
}

// SqlLastUseAt return name of the column being indexed
func (r *Rooms) SqlLastUseAt() string { //nolint:dupl false positive
	return `"lastUseAt"`
}

// IdxRoomSize return name of the index
func (r *Rooms) IdxRoomSize() int { //nolint:dupl false positive
	return 14
}

// SqlRoomSize return name of the column being indexed
func (r *Rooms) SqlRoomSize() string { //nolint:dupl false positive
	return `"roomSize"`
}

// IdxImageUrl return name of the index
func (r *Rooms) IdxImageUrl() int { //nolint:dupl false positive
	return 15
}

// SqlImageUrl return name of the column being indexed
func (r *Rooms) SqlImageUrl() string { //nolint:dupl false positive
	return `"imageUrl"`
}

// ToArray receiver fields to slice
func (r *Rooms) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if r.Id != 0 {
		id = r.Id
	}
	return A.X{
		id,
		r.RoomName,        // 1
		r.BasePriceIDR,    // 2
		r.CurrentTenantId, // 3
		r.FirstUseAt,      // 4
		r.CreatedAt,       // 5
		r.CreatedBy,       // 6
		r.UpdatedAt,       // 7
		r.UpdatedBy,       // 8
		r.DeletedAt,       // 9
		r.DeletedBy,       // 10
		r.RestoredBy,      // 11
		r.BuildingId,      // 12
		r.LastUseAt,       // 13
		r.RoomSize,        // 14
		r.ImageUrl,        // 15
	}
}

// FromArray convert slice to receiver fields
func (r *Rooms) FromArray(a A.X) *Rooms { //nolint:dupl false positive
	r.Id = X.ToU(a[0])
	r.RoomName = X.ToS(a[1])
	r.BasePriceIDR = X.ToI(a[2])
	r.CurrentTenantId = X.ToU(a[3])
	r.FirstUseAt = X.ToS(a[4])
	r.CreatedAt = X.ToI(a[5])
	r.CreatedBy = X.ToU(a[6])
	r.UpdatedAt = X.ToI(a[7])
	r.UpdatedBy = X.ToU(a[8])
	r.DeletedAt = X.ToI(a[9])
	r.DeletedBy = X.ToU(a[10])
	r.RestoredBy = X.ToU(a[11])
	r.BuildingId = X.ToU(a[12])
	r.LastUseAt = X.ToS(a[13])
	r.RoomSize = X.ToS(a[14])
	r.ImageUrl = X.ToS(a[15])
	return r
}

// FromUncensoredArray convert slice to receiver fields
func (r *Rooms) FromUncensoredArray(a A.X) *Rooms { //nolint:dupl false positive
	r.Id = X.ToU(a[0])
	r.RoomName = X.ToS(a[1])
	r.BasePriceIDR = X.ToI(a[2])
	r.CurrentTenantId = X.ToU(a[3])
	r.FirstUseAt = X.ToS(a[4])
	r.CreatedAt = X.ToI(a[5])
	r.CreatedBy = X.ToU(a[6])
	r.UpdatedAt = X.ToI(a[7])
	r.UpdatedBy = X.ToU(a[8])
	r.DeletedAt = X.ToI(a[9])
	r.DeletedBy = X.ToU(a[10])
	r.RestoredBy = X.ToU(a[11])
	r.BuildingId = X.ToU(a[12])
	r.LastUseAt = X.ToS(a[13])
	r.RoomSize = X.ToS(a[14])
	r.ImageUrl = X.ToS(a[15])
	return r
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (r *Rooms) FindOffsetLimit(offset, limit uint32, idx string) []Rooms { //nolint:dupl false positive
	var rows []Rooms
	res, err := r.Adapter.Select(r.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Rooms.FindOffsetLimit failed: `+r.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Rooms{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (r *Rooms) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := r.Adapter.Select(r.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Rooms.FindOffsetLimit failed: `+r.SpaceName()) {
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
func (r *Rooms) Total() int64 { //nolint:dupl false positive
	rows := r.Adapter.CallBoxSpace(r.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// RoomsFieldTypeMap returns key value of field name and key
var RoomsFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:              Tt.Unsigned,
	`roomName`:        Tt.String,
	`basePriceIDR`:    Tt.Integer,
	`currentTenantId`: Tt.Unsigned,
	`firstUseAt`:      Tt.String,
	`createdAt`:       Tt.Integer,
	`createdBy`:       Tt.Unsigned,
	`updatedAt`:       Tt.Integer,
	`updatedBy`:       Tt.Unsigned,
	`deletedAt`:       Tt.Integer,
	`deletedBy`:       Tt.Unsigned,
	`restoredBy`:      Tt.Unsigned,
	`buildingId`:      Tt.Unsigned,
	`lastUseAt`:       Tt.String,
	`roomSize`:        Tt.String,
	`imageUrl`:        Tt.String,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Stocks DAO reader/query struct
type Stocks struct {
	Adapter      *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id           uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	StockName    string      `json:"stockName" form:"stockName" query:"stockName" long:"stockName" msg:"stockName"`
	StockAddedAt string      `json:"stockAddedAt" form:"stockAddedAt" query:"stockAddedAt" long:"stockAddedAt" msg:"stockAddedAt"`
	Quantity     int64       `json:"quantity" form:"quantity" query:"quantity" long:"quantity" msg:"quantity"`
	PriceIDR     int64       `json:"priceIDR" form:"priceIDR" query:"priceIDR" long:"priceIDR" msg:"priceIDR"`
	CreatedAt    int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy    uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt    int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy    uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt    int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy    uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy   uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
}

// NewStocks create new ORM reader/query object
func NewStocks(adapter *Tt.Adapter) *Stocks {
	return &Stocks{Adapter: adapter}
}

// SpaceName returns full package and table name
func (s *Stocks) SpaceName() string { //nolint:dupl false positive
	return string(mProperty.TableStocks) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (s *Stocks) SqlTableName() string { //nolint:dupl false positive
	return `"stocks"`
}

func (s *Stocks) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (s *Stocks) FindById() bool { //nolint:dupl false positive
	res, err := s.Adapter.Select(s.SpaceName(), s.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{s.Id})
	if L.IsError(err, `Stocks.FindById failed: `+s.SpaceName()) {
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
func (s *Stocks) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "stockName"
	, "stockAddedAt"
	, "quantity"
	, "priceIDR"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (s *Stocks) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "stockName"
	, "stockAddedAt"
	, "quantity"
	, "priceIDR"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// ToUpdateArray generate slice of update command
func (s *Stocks) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, s.Id},
		A.X{`=`, 1, s.StockName},
		A.X{`=`, 2, s.StockAddedAt},
		A.X{`=`, 3, s.Quantity},
		A.X{`=`, 4, s.PriceIDR},
		A.X{`=`, 5, s.CreatedAt},
		A.X{`=`, 6, s.CreatedBy},
		A.X{`=`, 7, s.UpdatedAt},
		A.X{`=`, 8, s.UpdatedBy},
		A.X{`=`, 9, s.DeletedAt},
		A.X{`=`, 10, s.DeletedBy},
		A.X{`=`, 11, s.RestoredBy},
	}
}

// IdxId return name of the index
func (s *Stocks) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (s *Stocks) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxStockName return name of the index
func (s *Stocks) IdxStockName() int { //nolint:dupl false positive
	return 1
}

// SqlStockName return name of the column being indexed
func (s *Stocks) SqlStockName() string { //nolint:dupl false positive
	return `"stockName"`
}

// IdxStockAddedAt return name of the index
func (s *Stocks) IdxStockAddedAt() int { //nolint:dupl false positive
	return 2
}

// SqlStockAddedAt return name of the column being indexed
func (s *Stocks) SqlStockAddedAt() string { //nolint:dupl false positive
	return `"stockAddedAt"`
}

// IdxQuantity return name of the index
func (s *Stocks) IdxQuantity() int { //nolint:dupl false positive
	return 3
}

// SqlQuantity return name of the column being indexed
func (s *Stocks) SqlQuantity() string { //nolint:dupl false positive
	return `"quantity"`
}

// IdxPriceIDR return name of the index
func (s *Stocks) IdxPriceIDR() int { //nolint:dupl false positive
	return 4
}

// SqlPriceIDR return name of the column being indexed
func (s *Stocks) SqlPriceIDR() string { //nolint:dupl false positive
	return `"priceIDR"`
}

// IdxCreatedAt return name of the index
func (s *Stocks) IdxCreatedAt() int { //nolint:dupl false positive
	return 5
}

// SqlCreatedAt return name of the column being indexed
func (s *Stocks) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (s *Stocks) IdxCreatedBy() int { //nolint:dupl false positive
	return 6
}

// SqlCreatedBy return name of the column being indexed
func (s *Stocks) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (s *Stocks) IdxUpdatedAt() int { //nolint:dupl false positive
	return 7
}

// SqlUpdatedAt return name of the column being indexed
func (s *Stocks) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (s *Stocks) IdxUpdatedBy() int { //nolint:dupl false positive
	return 8
}

// SqlUpdatedBy return name of the column being indexed
func (s *Stocks) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (s *Stocks) IdxDeletedAt() int { //nolint:dupl false positive
	return 9
}

// SqlDeletedAt return name of the column being indexed
func (s *Stocks) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (s *Stocks) IdxDeletedBy() int { //nolint:dupl false positive
	return 10
}

// SqlDeletedBy return name of the column being indexed
func (s *Stocks) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (s *Stocks) IdxRestoredBy() int { //nolint:dupl false positive
	return 11
}

// SqlRestoredBy return name of the column being indexed
func (s *Stocks) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// ToArray receiver fields to slice
func (s *Stocks) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if s.Id != 0 {
		id = s.Id
	}
	return A.X{
		id,
		s.StockName,    // 1
		s.StockAddedAt, // 2
		s.Quantity,     // 3
		s.PriceIDR,     // 4
		s.CreatedAt,    // 5
		s.CreatedBy,    // 6
		s.UpdatedAt,    // 7
		s.UpdatedBy,    // 8
		s.DeletedAt,    // 9
		s.DeletedBy,    // 10
		s.RestoredBy,   // 11
	}
}

// FromArray convert slice to receiver fields
func (s *Stocks) FromArray(a A.X) *Stocks { //nolint:dupl false positive
	s.Id = X.ToU(a[0])
	s.StockName = X.ToS(a[1])
	s.StockAddedAt = X.ToS(a[2])
	s.Quantity = X.ToI(a[3])
	s.PriceIDR = X.ToI(a[4])
	s.CreatedAt = X.ToI(a[5])
	s.CreatedBy = X.ToU(a[6])
	s.UpdatedAt = X.ToI(a[7])
	s.UpdatedBy = X.ToU(a[8])
	s.DeletedAt = X.ToI(a[9])
	s.DeletedBy = X.ToU(a[10])
	s.RestoredBy = X.ToU(a[11])
	return s
}

// FromUncensoredArray convert slice to receiver fields
func (s *Stocks) FromUncensoredArray(a A.X) *Stocks { //nolint:dupl false positive
	s.Id = X.ToU(a[0])
	s.StockName = X.ToS(a[1])
	s.StockAddedAt = X.ToS(a[2])
	s.Quantity = X.ToI(a[3])
	s.PriceIDR = X.ToI(a[4])
	s.CreatedAt = X.ToI(a[5])
	s.CreatedBy = X.ToU(a[6])
	s.UpdatedAt = X.ToI(a[7])
	s.UpdatedBy = X.ToU(a[8])
	s.DeletedAt = X.ToI(a[9])
	s.DeletedBy = X.ToU(a[10])
	s.RestoredBy = X.ToU(a[11])
	return s
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (s *Stocks) FindOffsetLimit(offset, limit uint32, idx string) []Stocks { //nolint:dupl false positive
	var rows []Stocks
	res, err := s.Adapter.Select(s.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Stocks.FindOffsetLimit failed: `+s.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Stocks{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (s *Stocks) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := s.Adapter.Select(s.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Stocks.FindOffsetLimit failed: `+s.SpaceName()) {
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
func (s *Stocks) Total() int64 { //nolint:dupl false positive
	rows := s.Adapter.CallBoxSpace(s.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// StocksFieldTypeMap returns key value of field name and key
var StocksFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:           Tt.Unsigned,
	`stockName`:    Tt.String,
	`stockAddedAt`: Tt.String,
	`quantity`:     Tt.Integer,
	`priceIDR`:     Tt.Integer,
	`createdAt`:    Tt.Integer,
	`createdBy`:    Tt.Unsigned,
	`updatedAt`:    Tt.Integer,
	`updatedBy`:    Tt.Unsigned,
	`deletedAt`:    Tt.Integer,
	`deletedBy`:    Tt.Unsigned,
	`restoredBy`:   Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// WifiDevices DAO reader/query struct
type WifiDevices struct {
	Adapter    *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id         uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	StartAt    string      `json:"startAt" form:"startAt" query:"startAt" long:"startAt" msg:"startAt"`
	EndAt      string      `json:"endAt" form:"endAt" query:"endAt" long:"endAt" msg:"endAt"`
	PaidAt     string      `json:"paidAt" form:"paidAt" query:"paidAt" long:"paidAt" msg:"paidAt"`
	PriceIDR   int64       `json:"priceIDR" form:"priceIDR" query:"priceIDR" long:"priceIDR" msg:"priceIDR"`
	TenantId   uint64      `json:"tenantId,string" form:"tenantId" query:"tenantId" long:"tenantId" msg:"tenantId"`
	MacAddress string      `json:"macAddress" form:"macAddress" query:"macAddress" long:"macAddress" msg:"macAddress"`
	RoomId     uint64      `json:"roomId,string" form:"roomId" query:"roomId" long:"roomId" msg:"roomId"`
	CreatedAt  int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy  uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt  int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy  uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt  int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy  uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
}

// NewWifiDevices create new ORM reader/query object
func NewWifiDevices(adapter *Tt.Adapter) *WifiDevices {
	return &WifiDevices{Adapter: adapter}
}

// SpaceName returns full package and table name
func (w *WifiDevices) SpaceName() string { //nolint:dupl false positive
	return string(mProperty.TableWifiDevices) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (w *WifiDevices) SqlTableName() string { //nolint:dupl false positive
	return `"wifiDevices"`
}

func (w *WifiDevices) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (w *WifiDevices) FindById() bool { //nolint:dupl false positive
	res, err := w.Adapter.Select(w.SpaceName(), w.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{w.Id})
	if L.IsError(err, `WifiDevices.FindById failed: `+w.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		w.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (w *WifiDevices) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "startAt"
	, "endAt"
	, "paidAt"
	, "priceIDR"
	, "tenantId"
	, "macAddress"
	, "roomId"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (w *WifiDevices) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "startAt"
	, "endAt"
	, "paidAt"
	, "priceIDR"
	, "tenantId"
	, "macAddress"
	, "roomId"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// ToUpdateArray generate slice of update command
func (w *WifiDevices) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, w.Id},
		A.X{`=`, 1, w.StartAt},
		A.X{`=`, 2, w.EndAt},
		A.X{`=`, 3, w.PaidAt},
		A.X{`=`, 4, w.PriceIDR},
		A.X{`=`, 5, w.TenantId},
		A.X{`=`, 6, w.MacAddress},
		A.X{`=`, 7, w.RoomId},
		A.X{`=`, 8, w.CreatedAt},
		A.X{`=`, 9, w.CreatedBy},
		A.X{`=`, 10, w.UpdatedAt},
		A.X{`=`, 11, w.UpdatedBy},
		A.X{`=`, 12, w.DeletedAt},
		A.X{`=`, 13, w.DeletedBy},
		A.X{`=`, 14, w.RestoredBy},
	}
}

// IdxId return name of the index
func (w *WifiDevices) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (w *WifiDevices) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxStartAt return name of the index
func (w *WifiDevices) IdxStartAt() int { //nolint:dupl false positive
	return 1
}

// SqlStartAt return name of the column being indexed
func (w *WifiDevices) SqlStartAt() string { //nolint:dupl false positive
	return `"startAt"`
}

// IdxEndAt return name of the index
func (w *WifiDevices) IdxEndAt() int { //nolint:dupl false positive
	return 2
}

// SqlEndAt return name of the column being indexed
func (w *WifiDevices) SqlEndAt() string { //nolint:dupl false positive
	return `"endAt"`
}

// IdxPaidAt return name of the index
func (w *WifiDevices) IdxPaidAt() int { //nolint:dupl false positive
	return 3
}

// SqlPaidAt return name of the column being indexed
func (w *WifiDevices) SqlPaidAt() string { //nolint:dupl false positive
	return `"paidAt"`
}

// IdxPriceIDR return name of the index
func (w *WifiDevices) IdxPriceIDR() int { //nolint:dupl false positive
	return 4
}

// SqlPriceIDR return name of the column being indexed
func (w *WifiDevices) SqlPriceIDR() string { //nolint:dupl false positive
	return `"priceIDR"`
}

// IdxTenantId return name of the index
func (w *WifiDevices) IdxTenantId() int { //nolint:dupl false positive
	return 5
}

// SqlTenantId return name of the column being indexed
func (w *WifiDevices) SqlTenantId() string { //nolint:dupl false positive
	return `"tenantId"`
}

// IdxMacAddress return name of the index
func (w *WifiDevices) IdxMacAddress() int { //nolint:dupl false positive
	return 6
}

// SqlMacAddress return name of the column being indexed
func (w *WifiDevices) SqlMacAddress() string { //nolint:dupl false positive
	return `"macAddress"`
}

// IdxRoomId return name of the index
func (w *WifiDevices) IdxRoomId() int { //nolint:dupl false positive
	return 7
}

// SqlRoomId return name of the column being indexed
func (w *WifiDevices) SqlRoomId() string { //nolint:dupl false positive
	return `"roomId"`
}

// IdxCreatedAt return name of the index
func (w *WifiDevices) IdxCreatedAt() int { //nolint:dupl false positive
	return 8
}

// SqlCreatedAt return name of the column being indexed
func (w *WifiDevices) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (w *WifiDevices) IdxCreatedBy() int { //nolint:dupl false positive
	return 9
}

// SqlCreatedBy return name of the column being indexed
func (w *WifiDevices) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (w *WifiDevices) IdxUpdatedAt() int { //nolint:dupl false positive
	return 10
}

// SqlUpdatedAt return name of the column being indexed
func (w *WifiDevices) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (w *WifiDevices) IdxUpdatedBy() int { //nolint:dupl false positive
	return 11
}

// SqlUpdatedBy return name of the column being indexed
func (w *WifiDevices) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (w *WifiDevices) IdxDeletedAt() int { //nolint:dupl false positive
	return 12
}

// SqlDeletedAt return name of the column being indexed
func (w *WifiDevices) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (w *WifiDevices) IdxDeletedBy() int { //nolint:dupl false positive
	return 13
}

// SqlDeletedBy return name of the column being indexed
func (w *WifiDevices) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (w *WifiDevices) IdxRestoredBy() int { //nolint:dupl false positive
	return 14
}

// SqlRestoredBy return name of the column being indexed
func (w *WifiDevices) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// ToArray receiver fields to slice
func (w *WifiDevices) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if w.Id != 0 {
		id = w.Id
	}
	return A.X{
		id,
		w.StartAt,    // 1
		w.EndAt,      // 2
		w.PaidAt,     // 3
		w.PriceIDR,   // 4
		w.TenantId,   // 5
		w.MacAddress, // 6
		w.RoomId,     // 7
		w.CreatedAt,  // 8
		w.CreatedBy,  // 9
		w.UpdatedAt,  // 10
		w.UpdatedBy,  // 11
		w.DeletedAt,  // 12
		w.DeletedBy,  // 13
		w.RestoredBy, // 14
	}
}

// FromArray convert slice to receiver fields
func (w *WifiDevices) FromArray(a A.X) *WifiDevices { //nolint:dupl false positive
	w.Id = X.ToU(a[0])
	w.StartAt = X.ToS(a[1])
	w.EndAt = X.ToS(a[2])
	w.PaidAt = X.ToS(a[3])
	w.PriceIDR = X.ToI(a[4])
	w.TenantId = X.ToU(a[5])
	w.MacAddress = X.ToS(a[6])
	w.RoomId = X.ToU(a[7])
	w.CreatedAt = X.ToI(a[8])
	w.CreatedBy = X.ToU(a[9])
	w.UpdatedAt = X.ToI(a[10])
	w.UpdatedBy = X.ToU(a[11])
	w.DeletedAt = X.ToI(a[12])
	w.DeletedBy = X.ToU(a[13])
	w.RestoredBy = X.ToU(a[14])
	return w
}

// FromUncensoredArray convert slice to receiver fields
func (w *WifiDevices) FromUncensoredArray(a A.X) *WifiDevices { //nolint:dupl false positive
	w.Id = X.ToU(a[0])
	w.StartAt = X.ToS(a[1])
	w.EndAt = X.ToS(a[2])
	w.PaidAt = X.ToS(a[3])
	w.PriceIDR = X.ToI(a[4])
	w.TenantId = X.ToU(a[5])
	w.MacAddress = X.ToS(a[6])
	w.RoomId = X.ToU(a[7])
	w.CreatedAt = X.ToI(a[8])
	w.CreatedBy = X.ToU(a[9])
	w.UpdatedAt = X.ToI(a[10])
	w.UpdatedBy = X.ToU(a[11])
	w.DeletedAt = X.ToI(a[12])
	w.DeletedBy = X.ToU(a[13])
	w.RestoredBy = X.ToU(a[14])
	return w
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (w *WifiDevices) FindOffsetLimit(offset, limit uint32, idx string) []WifiDevices { //nolint:dupl false positive
	var rows []WifiDevices
	res, err := w.Adapter.Select(w.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `WifiDevices.FindOffsetLimit failed: `+w.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := WifiDevices{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (w *WifiDevices) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := w.Adapter.Select(w.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `WifiDevices.FindOffsetLimit failed: `+w.SpaceName()) {
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
func (w *WifiDevices) Total() int64 { //nolint:dupl false positive
	rows := w.Adapter.CallBoxSpace(w.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// WifiDevicesFieldTypeMap returns key value of field name and key
var WifiDevicesFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:         Tt.Unsigned,
	`startAt`:    Tt.String,
	`endAt`:      Tt.String,
	`paidAt`:     Tt.String,
	`priceIDR`:   Tt.Integer,
	`tenantId`:   Tt.Unsigned,
	`macAddress`: Tt.String,
	`roomId`:     Tt.Unsigned,
	`createdAt`:  Tt.Integer,
	`createdBy`:  Tt.Unsigned,
	`updatedAt`:  Tt.Integer,
	`updatedBy`:  Tt.Unsigned,
	`deletedAt`:  Tt.Integer,
	`deletedBy`:  Tt.Unsigned,
	`restoredBy`: Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
