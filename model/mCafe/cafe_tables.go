package mCafe

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
	TableMenus Tt.TableName = `menus`

	Name         = `name`
	HppIDR       = `hppIDR`
	SalePriceIDR = `salePriceIDR`
	Detail       = `detail`
	ImageUrl     = `imageUrl`
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
}
