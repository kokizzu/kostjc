package domain

import (
	"kostjc/model/mCafe"
	"kostjc/model/mCafe/rqCafe"
	"kostjc/model/mCafe/wcCafe"
	"kostjc/model/zCrud"
	"time"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminSales.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminSales.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminSales.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminSales.go
//go:generate farify doublequote --file AdminSales.go

type (
	AdminSaleIn struct {
		RequestCommon
		Cmd      string        `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Sale     rqCafe.Sales  `json:"sale" form:"sale" query:"sale" long:"sale" msg:"sale"`
	}
	AdminSaleOut struct {
		ResponseCommon
		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta  *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Sale  rqCafe.Sales   `json:"sale" form:"sale" query:"sale" long:"sale" msg:"sale"`
		Sales [][]any        `json:"sales" form:"sales" query:"sales" long:"sales" msg:"sales"`
	}
)

const (
	AdminSaleAction = `admin/sale`

	ErrAdminSaleNotFound      = `Sale not found`
	ErrAdminSaleSaveFailed    = `failed to save sale`
	ErrAdminSaleDeleteFailed  = `failed to delete sale`
	ErrAdminSaleRestoreFailed = `failed to restore sale`
	ErrAdminSaleAlreadyExists = `Sale already exists`
)

var AdminSaleMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mCafe.Id,
			Label:     `ID`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:      mCafe.Cashier,
			Label:     `Cashier`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
			ReadOnly:  false,
		},
		{
			Name:      mCafe.Who,
			Label:     `Who`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
			ReadOnly:  false,
		},
		{
			Name:      mCafe.SalesDate,
			Label:     `Sales Date`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  false,
			Mapping:   `IDR`,
		},
		{
			Name:      mCafe.PaidAt,
			Label:     `Paid At`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  false,
		},
		{
			Name:      mCafe.PaidWith,
			Label:     `Paid With`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeComboboxArr,
			ReadOnly:  false,
		},
		{
			Name:      mCafe.Note,
			Label:     `Note`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeTextArea,
			ReadOnly:  false,
		},
		{
			Name:      mCafe.TotalPriceIDR,
			Label:     `Total Price IDR`,
			DataType:  zCrud.DataTypeCurrency,
			InputType: zCrud.InputTypeNumber,
			ReadOnly:  false,
			Mapping:   `IDR`,
		},
		{
			Name:      mCafe.CreatedAt,
			Label:     `Created At`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
		{
			Name:      mCafe.UpdatedAt,
			Label:     `Updated At`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
		{
			Name:      mCafe.DeletedAt,
			Label:     `Deleted At`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
	},
}

func (d *Domain) AdminSale(in *AdminSaleIn) (out AdminSaleOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.actor = sess.UserId
	out.refId = in.Sale.Id

	if in.WithMeta {
		out.Meta = &AdminSaleMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		sale := wcCafe.NewSalesMutator(d.PropOltp)
		sale.Id = in.Sale.Id
		if sale.Id > 0 {
			if !sale.FindById() {
				out.SetError(400, ErrAdminSaleNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if sale.DeletedAt == 0 {
					sale.SetDeletedAt(in.UnixNow())
				}
			}

			if in.Cmd == zCrud.CmdRestore {
				if sale.DeletedAt > 0 {
					sale.SetDeletedAt(0)
					sale.SetRestoredBy(sess.UserId)
				}
			}
		}

		if in.Sale.Cashier != `` {
			sale.SetCashier(in.Sale.Cashier)
		}

		if in.Sale.Who != `` {
			sale.SetWho(in.Sale.Who)
		}

		if mCafe.IsValidDate(in.Sale.SalesDate, time.DateOnly) {
			sale.SetSalesDate(in.Sale.SalesDate)
		}

		if mCafe.IsValidDate(in.Sale.PaidAt, time.DateOnly) {
			sale.SetPaidAt(in.Sale.PaidAt)
		}

		if in.Sale.PaidWith != `` {
			sale.SetPaidWith(in.Sale.PaidWith)
		}

		if in.Sale.Note != `` {
			sale.SetNote(in.Sale.Note)
		}

		sale.SetTotalPriceIDR(in.Sale.TotalPriceIDR)

		if sale.Id == 0 {
			sale.SetCreatedAt(in.UnixNow())
			sale.SetCreatedBy(sess.UserId)
		}

		sale.SetUpdatedAt(in.UnixNow())
		sale.SetUpdatedBy(sess.UserId)

		if !sale.DoUpsert() {
			out.SetError(500, ErrAdminRoomSaveFailed)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		sale := rqCafe.NewSales(d.PropOltp)
		out.Sales = sale.FindByPagination(
			&AdminSaleMeta,
			&in.Pager,
			&out.Pager,
		)
	}

	return
}
