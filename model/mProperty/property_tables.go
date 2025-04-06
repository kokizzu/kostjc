package mProperty

import "github.com/kokizzu/gotro/D/Tt"

const (
	Id         = `id`
	CreatedAt  = `createdAt`
	CreatedBy  = `createdBy`
	UpdatedAt  = `updatedAt`
	UpdatedBy  = `updatedBy`
	DeletedAt  = `deletedAt`
	DeletedBy  = `deletedBy`
	RestoredBy = `restoredBy`
)

const (
	TableLocations Tt.TableName = `locations`

	Name         = `name`
	Address      = `address`
	GmapLocation = `gmapLocation`
)

const (
	TableFacilities Tt.TableName = `facilities`

	FacilityName   = `facilityName`
	extraChargeIDR = `extraChargeIDR`
)

const (
	TableBuildings Tt.TableName = `buildings`

	BuildingName  = `buildingName`
	LocationId    = `locationId`
	FacilitiesObj = `facilitiesObj`
)

const (
	TableRooms Tt.TableName = `rooms`

	RoomName        = `roomName`
	BasePriceIDR    = `basePriceIDR`
	CurrentTenantId = `currentTenantId`
	FirstUseAt      = `firstUseAt`
)

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TableLocations: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{Name, Tt.String},
			{Address, Tt.String},
			{GmapLocation, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
		},
		AutoIncrementId: true,
		Engine:          Tt.Memtx,
	},
	TableFacilities: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{FacilityName, Tt.String},
			{extraChargeIDR, Tt.Integer},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
		},
		AutoIncrementId: true,
		Engine:          Tt.Memtx,
	},
	TableBuildings: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{BuildingName, Tt.String},
			{LocationId, Tt.Unsigned},
			{FacilitiesObj, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
		},
		AutoIncrementId: true,
		Engine:          Tt.Memtx,
	},
	TableRooms: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{RoomName, Tt.String},
			{BasePriceIDR, Tt.Integer},
			{CurrentTenantId, Tt.Unsigned},
			{FirstUseAt, Tt.Integer},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
		},
		AutoIncrementId: true,
		Engine:          Tt.Memtx,
	},
}
