package domain

import (
	"kostjc/model/mCafe"
	"kostjc/model/mCafe/rqCafe"
	"kostjc/model/mCafe/wcCafe"
	"kostjc/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminMenu.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminMenu.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminMenu.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminMenu.go
//go:generate farify doublequote --file AdminMenu.go

type (
	AdminMenuIn struct {
		RequestCommon
		Cmd      string        `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Menu     rqCafe.Menus  `json:"menu" form:"menu" query:"menu" long:"menu" msg:"menu"`
	}
	AdminMenuOut struct {
		ResponseCommon
		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta  *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Menu  rqCafe.Menus   `json:"menu" form:"menu" query:"menu" long:"menu" msg:"menu"`
		Menus [][]any        `json:"menus" form:"menus" query:"menus" long:"menus" msg:"menus"`
	}
)

const (
	AdminMenuAction = `admin/menu`

	ErrAdminMenuNotFound      = `Menu not found`
	ErrAdminMenuSaveFailed    = `failed to save menu`
	ErrAdminMenuDeleteFailed  = `failed to delete menu`
	ErrAdminMenuRestoreFailed = `failed to restore menu`
	ErrAdminMenuAlreadyExists = `Menu already exists`
)

var AdminMenuMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mCafe.Id,
			Label:     `ID`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:        mCafe.Name,
			Label:       `Name`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			ReadOnly:    false,
			Description: `Indomie Soto`,
		},
		{
			Name:      mCafe.HppIDR,
			Label:     `Hpp`,
			DataType:  zCrud.DataTypeCurrency,
			InputType: zCrud.InputTypeNumber,
			ReadOnly:  false,
			Mapping:   `IDR`,
		},
		{
			Name:      mCafe.SalePriceIDR,
			Label:     `Sale Price`,
			DataType:  zCrud.DataTypeCurrency,
			InputType: zCrud.InputTypeNumber,
			ReadOnly:  false,
			Mapping:   `IDR`,
		},
		{
			Name:        mCafe.Detail,
			Label:       `Detail`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeTextArea,
			ReadOnly:    false,
			Description: `Indomie rebus`,
		},
		{
			Name:        mCafe.ImageUrl,
			Label:       `Image URL`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			ReadOnly:    false,
			Description: `img-indomie-soto.jpg`,
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

func (d *Domain) AdminMenu(in *AdminMenuIn) (out AdminMenuOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.actor = sess.UserId
	out.refId = in.Menu.Id

	if in.WithMeta {
		out.Meta = &AdminMenuMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		menu := wcCafe.NewMenusMutator(d.PropOltp)
		menu.Id = in.Menu.Id
		if menu.Id > 0 {
			if !menu.FindById() {
				out.SetError(400, ErrAdminMenuNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if menu.DeletedAt == 0 {
					menu.SetDeletedAt(in.UnixNow())
				}
			}

			if in.Cmd == zCrud.CmdRestore {
				if menu.DeletedAt > 0 {
					menu.SetDeletedAt(0)
					menu.SetRestoredBy(sess.UserId)
				}
			}
		}

		if in.Menu.Name != `` {
			menu.SetName(in.Menu.Name)
		}

		if in.Menu.HppIDR != 0 {
			menu.SetHppIDR(in.Menu.HppIDR)
		}

		if in.Menu.SalePriceIDR != 0 {
			menu.SetSalePriceIDR(in.Menu.SalePriceIDR)
		}

		if in.Menu.ImageUrl != `` {
			menu.SetImageUrl(in.Menu.ImageUrl)
		}

		if in.Menu.Detail != `` {
			menu.SetDetail(in.Menu.Detail)
		}

		if menu.Id == 0 {
			if menu.FindByName() {
				out.SetError(400, ErrAdminMenuAlreadyExists)
				return
			}

			menu.SetCreatedAt(in.UnixNow())
			menu.SetCreatedBy(sess.UserId)
		}

		menu.SetUpdatedAt(in.UnixNow())
		menu.SetUpdatedBy(sess.UserId)

		if !menu.DoUpsert() {
			out.SetError(500, ErrAdminMenuSaveFailed)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		menu := rqCafe.NewMenus(d.PropOltp)
		out.Menus = menu.FindByPagination(
			&AdminMenuMeta,
			&in.Pager,
			&out.Pager,
		)
	}

	return
}
