package domain

import (
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mCafe"
	"kostjc/model/mCafe/rqCafe"
	"kostjc/model/mCafe/wcCafe"
	"kostjc/model/zCrud"
	"time"

	"github.com/kokizzu/gotro/X"
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
		TenantID uint64        `json:"tenantId" form:"tenantId" query:"tenantId" long:"tenantId" msg:"tenantId"`
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

	ErrAdminSaleNotFound         = `Sale not found`
	ErrAdminSaleTenantIdNotFound = `tenant id not found for this sale`
	ErrAdminSaleSaveFailed       = `failed to save sale`
	ErrAdminSaleInvalidMenu      = `invalid menu`
	ErrAdminSaleDeleteFailed     = `failed to delete sale`
	ErrAdminSaleRestoreFailed    = `failed to restore sale`
	ErrAdminSaleAlreadyExists    = `Sale already exists`
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
			Name:      mCafe.TenantId,
			Label:     `Tenant`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
			ReadOnly:  false,
		},
		{
			Name:      mCafe.BuyerName,
			Label:     `Buyer Name`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
			ReadOnly:  false,
		},
		{
			Name:      mCafe.MenuIds,
			Label:     `Menus`,
			DataType:  zCrud.DataTypeIntArr,
			InputType: zCrud.InputTypeMultiSelect,
			ReadOnly:  false,
		},
		{
			Name:      mCafe.PaymentMethod,
			Label:     `Payment Method`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeComboboxArr,
			ReadOnly:  false,
		},
		{
			Name:      mCafe.PaymentStatus,
			Label:     `Payment Status`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeComboboxArr,
			ReadOnly:  false,
		},
		{
			Name:      mCafe.TransferIDR,
			Label:     `Transfer IDR`,
			DataType:  zCrud.DataTypeCurrency,
			InputType: zCrud.InputTypeNumber,
			ReadOnly:  false,
			Mapping:   `IDR`,
		},
		{
			Name:      mCafe.QrisIDR,
			Label:     `QRIS IDR`,
			DataType:  zCrud.DataTypeCurrency,
			InputType: zCrud.InputTypeNumber,
			ReadOnly:  false,
			Mapping:   `IDR`,
		},
		{
			Name:      mCafe.CashIDR,
			Label:     `Cash IDR`,
			DataType:  zCrud.DataTypeCurrency,
			InputType: zCrud.InputTypeNumber,
			ReadOnly:  false,
			Mapping:   `IDR`,
		},
		{
			Name:      mCafe.DebtIDR,
			Label:     `Debt IDR`,
			DataType:  zCrud.DataTypeCurrency,
			InputType: zCrud.InputTypeNumber,
			ReadOnly:  false,
			Mapping:   `IDR`,
		},
		{
			Name:      mCafe.TopupIDR,
			Label:     `Topup IDR`,
			DataType:  zCrud.DataTypeCurrency,
			InputType: zCrud.InputTypeNumber,
			ReadOnly:  false,
			Mapping:   `IDR`,
		},
		{
			Name:      mCafe.Donation,
			Label:     `Donation`,
			DataType:  zCrud.DataTypeCurrency,
			InputType: zCrud.InputTypeNumber,
			ReadOnly:  false,
			Mapping:   `IDR`,
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
			Name:      mCafe.SalesDate,
			Label:     `Sales Date`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  false,
		},
		{
			Name:      mCafe.PaidAt,
			Label:     `Paid At`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
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
	case zCrud.CmdUpsert, zCrud.CmdUpdatePayment, zCrud.CmdDelete, zCrud.CmdRestore:
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
					sale.SetDeletedBy(sess.UserId)
				}
			}

			if in.Cmd == zCrud.CmdRestore {
				if sale.DeletedAt > 0 {
					sale.SetDeletedAt(0)
					sale.SetRestoredBy(sess.UserId)
				}
			}

			if in.Cmd == zCrud.CmdUpdatePayment {
				if in.Sale.PaymentMethod != "" {
					sale.SetPaymentMethod(in.Sale.PaymentMethod)
				}
				if in.Sale.PaymentStatus != "" {
					sale.SetPaymentStatus(in.Sale.PaymentStatus)
				}

				sale.SetTransferIDR(in.Sale.TransferIDR)
				sale.SetQrisIDR(in.Sale.QrisIDR)
				sale.SetCashIDR(in.Sale.CashIDR)
				sale.SetDebtIDR(in.Sale.DebtIDR)
				sale.SetTopupIDR(in.Sale.TopupIDR)
				sale.SetTotalPriceIDR(in.Sale.TotalPriceIDR)

				if mCafe.IsValidDate(in.Sale.PaidAt, time.DateOnly) {
					sale.SetPaidAt(in.Sale.PaidAt)
				}

				sale.SetUpdatedAt(in.UnixNow())
				sale.SetUpdatedBy(sess.UserId)

			}
		}

		defer InsertCafeLog(sale.Id, d.saleLogs, out.ResponseCommon, in.TimeNow(), sess.UserId, sale)()

		if in.Sale.TenantId > 0 {
			bkd := rqAuth.NewTenants(d.PropOltp)
			bkd.Id = in.Sale.TenantId
			if !bkd.FindById() {
				out.SetError(400, ErrAdminSaleTenantIdNotFound)
				return
			}

			sale.SetTenantId(in.Sale.TenantId)
		}

		for _, id := range in.Sale.MenuIds {
			men := rqCafe.NewMenus(d.PropOltp)
			men.Id = X.ToU(id)
			if !men.FindById() {
				out.SetError(400, ErrAdminSaleInvalidMenu)
				return
			}
		}

		if in.Cmd == zCrud.CmdUpsert {
			sale.SetMenuIds(in.Sale.MenuIds)
		}

		if in.Sale.PaymentMethod != `` {
			sale.SetPaymentMethod(in.Sale.PaymentMethod)
		}

		if in.Sale.PaymentStatus != `` {
			sale.SetPaymentStatus(in.Sale.PaymentStatus)
		}

		if in.Sale.Cashier != `` {
			sale.SetCashier(in.Sale.Cashier)
		}

		if in.Sale.BuyerName != `` {
			sale.SetBuyerName(in.Sale.BuyerName)
		}

		if mCafe.IsValidDate(in.Sale.SalesDate, time.DateOnly) {
			sale.SetSalesDate(in.Sale.SalesDate)
		}

		if mCafe.IsValidDate(in.Sale.PaidAt, time.DateOnly) {
			sale.SetPaidAt(in.Sale.PaidAt)
		}

		if in.Sale.Note != `` {
			sale.SetNote(in.Sale.Note)
		}

		sale.SetQrisIDR(in.Sale.QrisIDR)
		sale.SetCashIDR(in.Sale.CashIDR)
		sale.SetDebtIDR(in.Sale.DebtIDR)
		sale.SetTopupIDR(in.Sale.TopupIDR)
		sale.SetTotalPriceIDR(in.Sale.TotalPriceIDR)
		sale.SetDonation(in.Sale.Donation)
		sale.SetTransferIDR(in.Sale.TransferIDR)

		if sale.Id == 0 {
			sale.SetCreatedAt(in.UnixNow())
			sale.SetCreatedBy(sess.UserId)
		}

		sale.SetUpdatedAt(in.UnixNow())
		sale.SetUpdatedBy(sess.UserId)

		if !sale.DoUpsert() {
			out.SetError(500, ErrAdminSaleSaveFailed)
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
