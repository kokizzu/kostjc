package domain

import (
	"kostjc/model/mProperty"
	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/mProperty/wcProperty"
	"kostjc/model/zCrud"
	"time"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminStock.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminStock.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminStock.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminStock.go
//go:generate farify doublequote --file AdminStock.go

type (
	AdminStockIn struct {
		RequestCommon
		Cmd      string            `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool              `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn     `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Stock    rqProperty.Stocks `json:"stock" form:"stock" query:"stock" long:"stock" msg:"stock"`
	}
	AdminStockOut struct {
		ResponseCommon
		Pager  zCrud.PagerOut    `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta   *zCrud.Meta       `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Stock  rqProperty.Stocks `json:"stock" form:"stock" query:"stock" long:"stock" msg:"stock"`
		Stocks [][]any           `json:"stocks" form:"stocks" query:"stocks" long:"stocks" msg:"stocks"`
	}
)

const (
	AdminStockAction = `admin/stock`

	ErrAdminStockNotFound      = `stock not found`
	ErrAdminStockSaveFailed    = `failed to save stock`
	ErrAdminStockDeleteFailed  = `failed to delete stock`
	ErrAdminStockRestoreFailed = `failed to restore stock`
)

var AdminStockMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mProperty.Id,
			Label:     `ID`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:        mProperty.StockName,
			Label:       `Stock Name`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			Description: `Coffee`,
			ReadOnly:    false,
		},
		{
			Name:      mProperty.StockAddedAt,
			Label:     `Added At`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  false,
		},
		{
			Name:        mProperty.Quantity,
			Label:       `Quantity`,
			DataType:    zCrud.DataTypeInt,
			InputType:   zCrud.InputTypeNumber,
			Description: `10`,
			ReadOnly:    false,
		},
		{
			Name:        mProperty.PriceIDR,
			Label:       `Total Price`,
			DataType:    zCrud.DataTypeCurrency,
			InputType:   zCrud.InputTypeNumber,
			Description: `14000`,
			ReadOnly:    false,
			Mapping:     `IDR`,
		},
		{
			Name:      mProperty.CreatedAt,
			Label:     `Created At`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
		{
			Name:      mProperty.UpdatedAt,
			Label:     `Updated At`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
		{
			Name:      mProperty.DeletedAt,
			Label:     `Deleted At`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
	},
}

func (d *Domain) AdminStock(in *AdminStockIn) (out AdminStockOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.actor = sess.UserId
	out.refId = in.Stock.Id

	if in.WithMeta {
		out.Meta = &AdminStockMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		stck := wcProperty.NewStocksMutator(d.PropOltp)
		stck.Id = in.Stock.Id
		if stck.Id > 0 {
			if !stck.FindById() {
				out.SetError(400, ErrAdminStockNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if stck.DeletedAt == 0 {
					stck.SetDeletedAt(in.UnixNow())
				}
			}

			if in.Cmd == zCrud.CmdRestore {
				if stck.DeletedAt > 0 {
					stck.SetDeletedAt(0)
					stck.SetRestoredBy(sess.UserId)
				}
			}
		}

		if in.Stock.StockName != `` {
			stck.SetStockName(in.Stock.StockName)
		}

		if mProperty.IsValidDate(in.Stock.StockAddedAt, time.DateOnly) {
			stck.SetStockAddedAt(in.Stock.StockAddedAt)
		}

		if in.Stock.Quantity > 0 {
			stck.SetQuantity(in.Stock.Quantity)
		}

		if in.Stock.PriceIDR > 0 {
			stck.SetPriceIDR(in.Stock.PriceIDR)
		}

		if stck.Id == 0 {
			stck.SetCreatedAt(in.UnixNow())
			stck.SetCreatedBy(sess.UserId)
		}

		stck.SetUpdatedAt(in.UnixNow())
		stck.SetUpdatedBy(sess.UserId)

		if !stck.DoUpsert() {
			out.SetError(500, ErrAdminStockSaveFailed)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		stck := rqProperty.NewStocks(d.PropOltp)
		out.Stocks = stck.FindByPagination(
			&AdminStockMeta,
			&in.Pager,
			&out.Pager,
		)
	}

	return
}
