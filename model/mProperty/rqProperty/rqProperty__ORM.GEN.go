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
	Facilities    string      `json:"facilities" form:"facilities" query:"facilities" long:"facilities" msg:"facilities"`
	TotalPriceIDR int64       `json:"totalPriceIDR" form:"totalPriceIDR" query:"totalPriceIDR" long:"totalPriceIDR" msg:"totalPriceIDR"`
	PaidAt        string      `json:"paidAt" form:"paidAt" query:"paidAt" long:"paidAt" msg:"paidAt"`
	TenantId      uint64      `json:"tenantId,string" form:"tenantId" query:"tenantId" long:"tenantId" msg:"tenantId"`
	CreatedAt     string      `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy     uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt     string      `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy     uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt     string      `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy     uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy    uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
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
	, "facilities"
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
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (b *Bookings) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "dateStart"
	, "dateEnd"
	, "basePriceIDR"
	, "facilities"
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
	`
}

// ToUpdateArray generate slice of update command
func (b *Bookings) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, b.Id},
		A.X{`=`, 1, b.DateStart},
		A.X{`=`, 2, b.DateEnd},
		A.X{`=`, 3, b.BasePriceIDR},
		A.X{`=`, 4, b.Facilities},
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

// IdxFacilities return name of the index
func (b *Bookings) IdxFacilities() int { //nolint:dupl false positive
	return 4
}

// SqlFacilities return name of the column being indexed
func (b *Bookings) SqlFacilities() string { //nolint:dupl false positive
	return `"facilities"`
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
		b.Facilities,    // 4
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
	}
}

// FromArray convert slice to receiver fields
func (b *Bookings) FromArray(a A.X) *Bookings { //nolint:dupl false positive
	b.Id = X.ToU(a[0])
	b.DateStart = X.ToS(a[1])
	b.DateEnd = X.ToS(a[2])
	b.BasePriceIDR = X.ToI(a[3])
	b.Facilities = X.ToS(a[4])
	b.TotalPriceIDR = X.ToI(a[5])
	b.PaidAt = X.ToS(a[6])
	b.TenantId = X.ToU(a[7])
	b.CreatedAt = X.ToS(a[8])
	b.CreatedBy = X.ToU(a[9])
	b.UpdatedAt = X.ToS(a[10])
	b.UpdatedBy = X.ToU(a[11])
	b.DeletedAt = X.ToS(a[12])
	b.DeletedBy = X.ToU(a[13])
	b.RestoredBy = X.ToU(a[14])
	return b
}

// FromUncensoredArray convert slice to receiver fields
func (b *Bookings) FromUncensoredArray(a A.X) *Bookings { //nolint:dupl false positive
	b.Id = X.ToU(a[0])
	b.DateStart = X.ToS(a[1])
	b.DateEnd = X.ToS(a[2])
	b.BasePriceIDR = X.ToI(a[3])
	b.Facilities = X.ToS(a[4])
	b.TotalPriceIDR = X.ToI(a[5])
	b.PaidAt = X.ToS(a[6])
	b.TenantId = X.ToU(a[7])
	b.CreatedAt = X.ToS(a[8])
	b.CreatedBy = X.ToU(a[9])
	b.UpdatedAt = X.ToS(a[10])
	b.UpdatedBy = X.ToU(a[11])
	b.DeletedAt = X.ToS(a[12])
	b.DeletedBy = X.ToU(a[13])
	b.RestoredBy = X.ToU(a[14])
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
	`facilities`:    Tt.String,
	`totalPriceIDR`: Tt.Integer,
	`paidAt`:        Tt.String,
	`tenantId`:      Tt.Unsigned,
	`createdAt`:     Tt.String,
	`createdBy`:     Tt.Unsigned,
	`updatedAt`:     Tt.String,
	`updatedBy`:     Tt.Unsigned,
	`deletedAt`:     Tt.String,
	`deletedBy`:     Tt.Unsigned,
	`restoredBy`:    Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Buildings DAO reader/query struct
type Buildings struct {
	Adapter       *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id            uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	BuildingName  string      `json:"buildingName" form:"buildingName" query:"buildingName" long:"buildingName" msg:"buildingName"`
	LocationId    uint64      `json:"locationId,string" form:"locationId" query:"locationId" long:"locationId" msg:"locationId"`
	FacilitiesObj string      `json:"facilitiesObj" form:"facilitiesObj" query:"facilitiesObj" long:"facilitiesObj" msg:"facilitiesObj"`
	CreatedAt     int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy     uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt     int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy     uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt     int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy     uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy    uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
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

// SqlSelectAllFields generate Sql select fields
func (b *Buildings) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "buildingName"
	, "locationId"
	, "facilitiesObj"
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
	, "facilitiesObj"
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
		A.X{`=`, 3, b.FacilitiesObj},
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

// IdxFacilitiesObj return name of the index
func (b *Buildings) IdxFacilitiesObj() int { //nolint:dupl false positive
	return 3
}

// SqlFacilitiesObj return name of the column being indexed
func (b *Buildings) SqlFacilitiesObj() string { //nolint:dupl false positive
	return `"facilitiesObj"`
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
		b.BuildingName,  // 1
		b.LocationId,    // 2
		b.FacilitiesObj, // 3
		b.CreatedAt,     // 4
		b.CreatedBy,     // 5
		b.UpdatedAt,     // 6
		b.UpdatedBy,     // 7
		b.DeletedAt,     // 8
		b.DeletedBy,     // 9
		b.RestoredBy,    // 10
	}
}

// FromArray convert slice to receiver fields
func (b *Buildings) FromArray(a A.X) *Buildings { //nolint:dupl false positive
	b.Id = X.ToU(a[0])
	b.BuildingName = X.ToS(a[1])
	b.LocationId = X.ToU(a[2])
	b.FacilitiesObj = X.ToS(a[3])
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
	b.FacilitiesObj = X.ToS(a[3])
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
	`id`:            Tt.Unsigned,
	`buildingName`:  Tt.String,
	`locationId`:    Tt.Unsigned,
	`facilitiesObj`: Tt.String,
	`createdAt`:     Tt.Integer,
	`createdBy`:     Tt.Unsigned,
	`updatedAt`:     Tt.Integer,
	`updatedBy`:     Tt.Unsigned,
	`deletedAt`:     Tt.Integer,
	`deletedBy`:     Tt.Unsigned,
	`restoredBy`:    Tt.Unsigned,
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
	CreatedAt     string      `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy     uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt     string      `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy     uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt     string      `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
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
	p.CreatedAt = X.ToS(a[7])
	p.CreatedBy = X.ToU(a[8])
	p.UpdatedAt = X.ToS(a[9])
	p.UpdatedBy = X.ToU(a[10])
	p.DeletedAt = X.ToS(a[11])
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
	p.CreatedAt = X.ToS(a[7])
	p.CreatedBy = X.ToU(a[8])
	p.UpdatedAt = X.ToS(a[9])
	p.UpdatedBy = X.ToU(a[10])
	p.DeletedAt = X.ToS(a[11])
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
	`createdAt`:     Tt.String,
	`createdBy`:     Tt.Unsigned,
	`updatedAt`:     Tt.String,
	`updatedBy`:     Tt.Unsigned,
	`deletedAt`:     Tt.String,
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
	FirstUseAt      int64       `json:"firstUseAt" form:"firstUseAt" query:"firstUseAt" long:"firstUseAt" msg:"firstUseAt"`
	CreatedAt       int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy       uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt       int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy       uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt       int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy       uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy      uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
	BuildingId      uint64      `json:"buildingId,string" form:"buildingId" query:"buildingId" long:"buildingId" msg:"buildingId"`
	LastUseAt       int64       `json:"lastUseAt" form:"lastUseAt" query:"lastUseAt" long:"lastUseAt" msg:"lastUseAt"`
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
	}
}

// FromArray convert slice to receiver fields
func (r *Rooms) FromArray(a A.X) *Rooms { //nolint:dupl false positive
	r.Id = X.ToU(a[0])
	r.RoomName = X.ToS(a[1])
	r.BasePriceIDR = X.ToI(a[2])
	r.CurrentTenantId = X.ToU(a[3])
	r.FirstUseAt = X.ToI(a[4])
	r.CreatedAt = X.ToI(a[5])
	r.CreatedBy = X.ToU(a[6])
	r.UpdatedAt = X.ToI(a[7])
	r.UpdatedBy = X.ToU(a[8])
	r.DeletedAt = X.ToI(a[9])
	r.DeletedBy = X.ToU(a[10])
	r.RestoredBy = X.ToU(a[11])
	r.BuildingId = X.ToU(a[12])
	r.LastUseAt = X.ToI(a[13])
	return r
}

// FromUncensoredArray convert slice to receiver fields
func (r *Rooms) FromUncensoredArray(a A.X) *Rooms { //nolint:dupl false positive
	r.Id = X.ToU(a[0])
	r.RoomName = X.ToS(a[1])
	r.BasePriceIDR = X.ToI(a[2])
	r.CurrentTenantId = X.ToU(a[3])
	r.FirstUseAt = X.ToI(a[4])
	r.CreatedAt = X.ToI(a[5])
	r.CreatedBy = X.ToU(a[6])
	r.UpdatedAt = X.ToI(a[7])
	r.UpdatedBy = X.ToU(a[8])
	r.DeletedAt = X.ToI(a[9])
	r.DeletedBy = X.ToU(a[10])
	r.RestoredBy = X.ToU(a[11])
	r.BuildingId = X.ToU(a[12])
	r.LastUseAt = X.ToI(a[13])
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
	`firstUseAt`:      Tt.Integer,
	`createdAt`:       Tt.Integer,
	`createdBy`:       Tt.Unsigned,
	`updatedAt`:       Tt.Integer,
	`updatedBy`:       Tt.Unsigned,
	`deletedAt`:       Tt.Integer,
	`deletedBy`:       Tt.Unsigned,
	`restoredBy`:      Tt.Unsigned,
	`buildingId`:      Tt.Unsigned,
	`lastUseAt`:       Tt.Integer,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
