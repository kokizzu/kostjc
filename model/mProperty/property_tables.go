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
	CurrentTenantId = `currentTenantId`
	FirstUseAt      = `firstUseAt`
	LastUseAt       = `lastUseAt`
	BuildingId      = `buildingId`
)

const (
	TableBookings Tt.TableName = `bookings`

	DateStart     = `dateStart`
	DateEnd       = `dateEnd`
	BasePriceIDR  = `basePriceIDR`
	Facilities    = `facilities`
	TotalPriceIDR = `totalPriceIDR`
	PaidAt        = `paidAt`
	TenantId      = `tenantId`
)

const (
	TablePayments Tt.TableName = `payments`

	BookingId     = `bookingId`
	PaymentAt     = `paymentAt`
	PaidIDR       = `paidIDR`
	PaymentMethod = `paymentMethod`
	PaymentStatus = `paymentStatus`
	Note          = `note`
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
			{BuildingId, Tt.Unsigned},
			{LastUseAt, Tt.Integer},
		},
		AutoIncrementId: true,
		Engine:          Tt.Memtx,
	},
	TableBookings: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{DateStart, Tt.String},
			{DateEnd, Tt.String},
			{BasePriceIDR, Tt.Integer},
			{Facilities, Tt.String},
			{TotalPriceIDR, Tt.Integer},
			{PaidAt, Tt.String},
			{TenantId, Tt.Unsigned},
			{CreatedAt, Tt.String},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.String},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.String},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
		},
		AutoIncrementId: true,
		Engine:          Tt.Memtx,
	},
	TablePayments: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{BookingId, Tt.Unsigned},
			{PaymentAt, Tt.String},
			{PaidIDR, Tt.Integer},
			{PaymentMethod, Tt.String},
			{PaymentStatus, Tt.String},
			{Note, Tt.String},
			{CreatedAt, Tt.String},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.String},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.String},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
		},
		AutoIncrementId: true,
		Engine:          Tt.Memtx,
	},
}
