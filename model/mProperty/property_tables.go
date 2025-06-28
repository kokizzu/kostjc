package mProperty

import (
	"time"

	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
)

func IsValidDate(dateStr, format string) bool {
	_, err := time.Parse(format, dateStr)
	return err == nil
}

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
	FacilityTypeRoom     = `Room`
	FacilityTypeBuilding = `Building`
	FacilityTypeSite     = `Site`
)

func IsValidFacilityType(facilityType string) bool {
	switch facilityType {
	case FacilityTypeRoom, FacilityTypeBuilding, FacilityTypeSite:
		return true
	default:
		return false
	}
}

const (
	TableFacilities Tt.TableName = `facilities`

	FacilityName   = `facilityName`
	ExtraChargeIDR = `extraChargeIDR`
	FacilityType   = `facilityType`
	DescriptionEN  = `descriptionEN`
)

type FacilityObj struct {
	FacilityName   string `json:"facilityName"`
	ExtraChargeIDR int64  `json:"extraChargeIDR"`
}

const (
	TableBuildings Tt.TableName = `buildings`

	BuildingName = `buildingName`
	LocationId   = `locationId`
	Facilities   = `facilities`
)

const (
	TableRooms Tt.TableName = `rooms`

	RoomName        = `roomName`
	CurrentTenantId = `currentTenantId`
	FirstUseAt      = `firstUseAt`
	LastUseAt       = `lastUseAt`
	BuildingId      = `buildingId`
	RoomSize        = `roomSize`
	ImageUrl        = `imageUrl`
)

const (
	TableBookings Tt.TableName = `bookings`

	DateStart     = `dateStart`
	DateEnd       = `dateEnd`
	BasePriceIDR  = `basePriceIDR`
	TotalPriceIDR = `totalPriceIDR`
	PaidAt        = `paidAt`
	FacilitiesObj = `facilitiesObj`
	TenantId      = `tenantId`
	ExtraTenants  = `extraTenants`
	RoomId        = `roomId`

	TotalPaidIDR = `totalPaidIDR` // Read Only
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

const (
	TableStocks Tt.TableName = `stocks`

	StockName    = `stockName`
	StockAddedAt = `stockAddedAt`
	Quantity     = `quantity`
	PriceIDR     = `priceIDR`
)

const (
	TableWifiDevices Tt.TableName = `wifiDevices`

	StartAt    = `startAt` // UNIX timestamp
	EndAt      = `endAt`   // UNIX timestamp
	MacAddress = `macAddress`
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
			{ExtraChargeIDR, Tt.Integer},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
			{FacilityType, Tt.String},
			{DescriptionEN, Tt.String},
		},
		AutoIncrementId: true,
		Engine:          Tt.Memtx,
	},
	TableBuildings: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{BuildingName, Tt.String},
			{LocationId, Tt.Unsigned},
			{Facilities, Tt.Array},
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
		Unique1:         BuildingName,
	},
	TableRooms: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{RoomName, Tt.String},
			{BasePriceIDR, Tt.Integer},
			{CurrentTenantId, Tt.Unsigned},
			{FirstUseAt, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
			{BuildingId, Tt.Unsigned},
			{LastUseAt, Tt.String},
			{RoomSize, Tt.String},
			{ImageUrl, Tt.String},
		},
		AutoIncrementId: true,
		Engine:          Tt.Memtx,
		Unique1:         RoomName,
	},
	TableBookings: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{DateStart, Tt.String},
			{DateEnd, Tt.String},
			{BasePriceIDR, Tt.Integer},
			{FacilitiesObj, Tt.String},
			{TotalPriceIDR, Tt.Integer},
			{PaidAt, Tt.String},
			{TenantId, Tt.Unsigned},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
			{ExtraTenants, Tt.Array},
			{RoomId, Tt.Unsigned},
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
	TableStocks: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{StockName, Tt.String},
			{StockAddedAt, Tt.String},
			{Quantity, Tt.Integer},
			{PriceIDR, Tt.Integer},
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
	TableWifiDevices: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{StartAt, Tt.String},
			{EndAt, Tt.String},
			{PaidAt, Tt.String},
			{PriceIDR, Tt.Integer},
			{TenantId, Tt.Unsigned},
			{MacAddress, Tt.String},
			{RoomId, Tt.Unsigned},
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

const (
	TableLocationLogs   Ch.TableName = `locationLogs`
	TableFacilityLogs   Ch.TableName = `facilityLogs`
	TableBuildingLogs   Ch.TableName = `buildingLogs`
	TableRoomLogs       Ch.TableName = `roomLogs`
	TableBookingLogs    Ch.TableName = `bookingLogs`
	TablePaymentLogs    Ch.TableName = `paymentLogs`
	TableStockLogs      Ch.TableName = `stockLogs`
	TableWifiDeviceLogs Ch.TableName = `wifiDeviceLogs`
)

const (
	ActorId    = `actorId`
	BeforeJSON = `beforeJson`
	AfterJSON  = `afterJson`
)

var ClickhouseTables = map[Ch.TableName]*Ch.TableProp{
	TableLocationLogs: {
		Fields: []Ch.Field{
			{CreatedAt, Ch.DateTime},
			{ActorId, Ch.UInt64},
			{BeforeJSON, Ch.String},
			{AfterJSON, Ch.String},
		},
		Orders: []string{CreatedAt, ActorId},
	},
	TableFacilityLogs: {
		Fields: []Ch.Field{
			{CreatedAt, Ch.DateTime},
			{ActorId, Ch.UInt64},
			{BeforeJSON, Ch.String},
			{AfterJSON, Ch.String},
		},
		Orders: []string{CreatedAt, ActorId},
	},
	TableBuildingLogs: {
		Fields: []Ch.Field{
			{CreatedAt, Ch.DateTime},
			{ActorId, Ch.UInt64},
			{BeforeJSON, Ch.String},
			{AfterJSON, Ch.String},
		},
		Orders: []string{CreatedAt, ActorId},
	},
	TableRoomLogs: {
		Fields: []Ch.Field{
			{CreatedAt, Ch.DateTime},
			{ActorId, Ch.UInt64},
			{BeforeJSON, Ch.String},
			{AfterJSON, Ch.String},
		},
		Orders: []string{CreatedAt, ActorId},
	},
	TableBookingLogs: {
		Fields: []Ch.Field{
			{CreatedAt, Ch.DateTime},
			{ActorId, Ch.UInt64},
			{BeforeJSON, Ch.String},
			{AfterJSON, Ch.String},
		},
		Orders: []string{CreatedAt, ActorId},
	},
	TablePaymentLogs: {
		Fields: []Ch.Field{
			{CreatedAt, Ch.DateTime},
			{ActorId, Ch.UInt64},
			{BeforeJSON, Ch.String},
			{AfterJSON, Ch.String},
		},
		Orders: []string{CreatedAt, ActorId},
	},
	TableStockLogs: {
		Fields: []Ch.Field{
			{CreatedAt, Ch.DateTime},
			{ActorId, Ch.UInt64},
			{BeforeJSON, Ch.String},
			{AfterJSON, Ch.String},
		},
		Orders: []string{CreatedAt, ActorId},
	},
	TableWifiDeviceLogs: {
		Fields: []Ch.Field{
			{CreatedAt, Ch.DateTime},
			{ActorId, Ch.UInt64},
			{BeforeJSON, Ch.String},
			{AfterJSON, Ch.String},
		},
		Orders: []string{CreatedAt, ActorId},
	},
}
