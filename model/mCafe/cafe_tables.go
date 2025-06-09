package mCafe

import (
	"time"

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
	TableMenus Tt.TableName = `menus`

	Name         = `name`
	HppIDR       = `hppIDR`
	SalePriceIDR = `salePriceIDR`
	Detail       = `detail`
	ImageUrl     = `imageUrl`
)

const (
	TableSales Tt.TableName = `sales`

	Cashier       = `cashier`
	TenantId      = `tenantId`
	BuyerName     = `buyerName`
	MenuIds       = `menuIds`
	QrisIDR       = `qrisIDR`
	CashIDR       = `cashIDR`
	DebtIDR       = `debtIDR`
	TopupIDR      = `topupIDR`
	TotalPriceIDR = `totalPriceIDR`
	SalesDate     = `salesDate`
	PaidAt        = `paidAt`
	Note          = `note`
)

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TableMenus: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{Name, Tt.String},
			{HppIDR, Tt.Integer},
			{SalePriceIDR, Tt.Integer},
			{Detail, Tt.String},
			{ImageUrl, Tt.String},
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
		Unique1:         Name,
	},
	TableSales: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{Cashier, Tt.String},
			{TenantId, Tt.Unsigned},
			{BuyerName, Tt.String},
			{MenuIds, Tt.Array},
			{QrisIDR, Tt.Integer},
			{CashIDR, Tt.Integer},
			{DebtIDR, Tt.Integer},
			{TopupIDR, Tt.Integer},
			{TotalPriceIDR, Tt.Integer},
			{SalesDate, Tt.String},
			{PaidAt, Tt.String},
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
}
