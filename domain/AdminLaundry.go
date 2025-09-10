package domain

import (
	"kostjc/model/mCafe"
	"kostjc/model/mCafe/rqCafe"
	"kostjc/model/mCafe/wcCafe"
	"kostjc/model/zCrud"
	"time"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminLaundry.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminLaundry.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminLaundry.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminLaundry.go
//go:generate farify doublequote --file AdminLaundry.go

type (
	AdminLaundryIn struct {
		RequestCommon
		Cmd      string         `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool           `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn  `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Laundry  rqCafe.Laundry `json:"laundry" form:"laundry" query:"laundry" long:"laundry" msg:"laundry"`
	}
	AdminLaundryOut struct {
		ResponseCommon
		Pager     zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta      *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Laundry   rqCafe.Laundry `json:"laundry" form:"laundry" query:"laundry" long:"laundry" msg:"laundry"`
		Laundries [][]any        `json:"laundries" form:"laundries" query:"laundries" long:"laundries" msg:"laundries"`
	}
)

const (
	AdminLaundryAction = `admin/laundry`

	ErrAdminLaundryNotFound      = `laundry not found`
	ErrAdminLaundrySaveFailed    = `failed to save laundry`
	ErrAdminLaundryDeleteFailed  = `failed to delete laundry`
	ErrAdminLaundryRestoreFailed = `failed to restore laundry`
)

var AdminLaundryMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mCafe.Id,
			Label:     `ID`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:      mCafe.LaundryCustomer,
			Label:     `Customer`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mCafe.LaundryItems,
			Label:     `Items`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mCafe.LaundryWeight,
			Label:     `Weight (kg)`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeNumber,
		},
		{
			Name:      mCafe.LaundryNote,
			Label:     `Note`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeTextArea,
		},
		{
			Name:      mCafe.LaundryAt,
			Label:     `Laudry At`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mCafe.LaundryPaidAt,
			Label:     `Paid At`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mCafe.LaundryPaidBy,
			Label:     `Paid By`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
		},
		{
			Name:      mCafe.LaundryWashAt,
			Label:     `Wash At`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mCafe.LaundryWashBy,
			Label:     `Wash By`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
		},
		{
			Name:      mCafe.LaundryDryAt,
			Label:     `Dry At`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mCafe.LaundryDryBy,
			Label:     `Dry By`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
		},
		{
			Name:      mCafe.LaundryFoldAt,
			Label:     `Fold At`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mCafe.LaundryFoldBy,
			Label:     `Fold By`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
		},
		{
			Name:      mCafe.LaundryNotifyAt,
			Label:     `Notify At`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mCafe.LaundryNotifyBy,
			Label:     `Notify By`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
		},
		{
			Name:      mCafe.LaundryGivenAt,
			Label:     `Given At`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mCafe.LaundryGivenBys,
			Label:     `Given By`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
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

func (d *Domain) AdminLaundry(in *AdminLaundryIn) (out AdminLaundryOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.actor = sess.UserId
	out.refId = in.Laundry.Id

	if in.WithMeta {
		out.Meta = &AdminLaundryMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		lnd := wcCafe.NewLaundryMutator(d.PropOltp)
		lnd.Id = in.Laundry.Id

		if lnd.Id > 0 {
			if !lnd.FindById() {
				out.SetError(400, ErrAdminLaundryNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if lnd.DeletedAt == 0 {
					lnd.SetDeletedAt(in.UnixNow())
				}
			}

			if in.Cmd == zCrud.CmdRestore {
				if lnd.DeletedAt > 0 {
					lnd.SetDeletedAt(0)
					lnd.SetRestoredBy(sess.UserId)
				}
			}
		}

		defer InsertPropertyLog(lnd.Id, d.laundryLogs, out.ResponseCommon, in.TimeNow(), sess.UserId, lnd)()

		lnd.SetCustomer(in.Laundry.Customer)
		lnd.SetItems(in.Laundry.Items)
		lnd.SetWeight(in.Laundry.Weight)
		lnd.SetNote(in.Laundry.Note)
		if mCafe.IsValidDate(in.Laundry.LaundryAt, time.DateOnly) {
			lnd.SetLaundryAt(in.Laundry.LaundryAt)
		}
		if mCafe.IsValidDate(in.Laundry.PaidAt, time.DateOnly) {
			lnd.SetPaidAt(in.Laundry.PaidAt)
		}
		lnd.SetPaidBy(in.Laundry.PaidBy)
		if mCafe.IsValidDate(in.Laundry.WashAt, time.DateOnly) {
			lnd.SetWashAt(in.Laundry.WashAt)
		}
		lnd.SetWashBy(in.Laundry.WashBy)
		if mCafe.IsValidDate(in.Laundry.DryAt, time.DateOnly) {
			lnd.SetDryAt(in.Laundry.DryAt)
		}
		lnd.SetDryBy(in.Laundry.DryBy)
		if mCafe.IsValidDate(in.Laundry.FoldAt, time.DateOnly) {
			lnd.SetFoldAt(in.Laundry.FoldAt)
		}
		lnd.SetFoldBy(in.Laundry.FoldBy)
		if mCafe.IsValidDate(in.Laundry.NotifyAt, time.DateOnly) {
			lnd.SetNotifyAt(in.Laundry.NotifyAt)
		}
		lnd.SetNotifyBy(in.Laundry.NotifyBy)
		if mCafe.IsValidDate(in.Laundry.GivenAt, time.DateOnly) {
			lnd.SetGivenAt(in.Laundry.GivenAt)
		}
		lnd.SetGivenBys(in.Laundry.GivenBys)

		if lnd.Id == 0 {
			lnd.SetCreatedAt(in.UnixNow())
			lnd.SetCreatedBy(sess.UserId)
		}

		lnd.SetUpdatedAt(in.UnixNow())
		lnd.SetUpdatedBy(sess.UserId)

		if !lnd.DoUpsert() {
			out.SetError(500, ErrAdminLaundrySaveFailed)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		lnd := rqCafe.NewLaundry(d.PropOltp)
		out.Laundries = lnd.FindByPagination(
			&AdminLaundryMeta,
			&in.Pager,
			&out.Pager,
		)
	}

	return
}
