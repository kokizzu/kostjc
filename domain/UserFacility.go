package domain

import (
	"kostjc/model/mProperty"
	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/mProperty/wcProperty"
	"kostjc/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserFacility.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserFacility.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserFacility.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserFacility.go
//go:generate farify doublequote --file UserFacility.go

type (
	UserFacilityIn struct {
		RequestCommon
		Cmd      string                `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool                  `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn         `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Facility rqProperty.Facilities `json:"facility" form:"facility" query:"facility" long:"facility" msg:"facility"`
	}
	UserFacilityOut struct {
		ResponseCommon
		Pager      zCrud.PagerOut        `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta       *zCrud.Meta           `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Facility   rqProperty.Facilities `json:"facility" form:"facility" query:"facility" long:"facility" msg:"facility"`
		Facilities [][]any               `json:"facilities" form:"facilities" query:"facilities" long:"facilities" msg:"facilities"`
	}
)

const (
	UserFacilityAction = `user/facility`

	ErrUserFacilityNotFound      = `facility not found`
	ErrUserFacilitySaveFailed    = `failed to save facility`
	ErrUserFacilityDeleteFailed  = `failed to delete facility`
	ErrUserFacilityRestoreFailed = `failed to restore facility`
)

var UserFacilityMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mProperty.Id,
			Label:     `ID`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:      mProperty.FacilityName,
			Label:     `Facility Name`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
			ReadOnly:  false,
		},
		{
			Name:      mProperty.ExtraChargeIDR,
			Label:     `Extra Charge`,
			DataType:  zCrud.DataTypeCurrency,
			InputType: zCrud.InputTypeNumber,
			ReadOnly:  false,
			Mapping:   `IDR`,
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

func (d *Domain) UserFacility(in *UserFacilityIn) (out UserFacilityOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.actor = sess.UserId
	out.refId = in.Facility.Id

	if in.WithMeta {
		out.Meta = &UserFacilityMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		fac := wcProperty.NewFacilitiesMutator(d.PropOltp)
		fac.Id = in.Facility.Id
		if fac.Id > 0 {
			if !fac.FindById() {
				out.SetError(400, ErrUserFacilityNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if fac.DeletedAt == 0 {
					fac.SetDeletedAt(in.UnixNow())
				}
			}

			if in.Cmd == zCrud.CmdRestore {
				if fac.DeletedAt > 0 {
					fac.SetDeletedAt(0)
					fac.SetRestoredBy(sess.UserId)
				}
			}
		}

		if in.Facility.FacilityName != `` {
			fac.SetFacilityName(in.Facility.FacilityName)
		}

		if in.Facility.ExtraChargeIDR != 0 {
			fac.SetExtraChargeIDR(in.Facility.ExtraChargeIDR)
		}

		if fac.Id == 0 {
			fac.SetCreatedAt(in.UnixNow())
			fac.SetCreatedBy(sess.UserId)
		}

		fac.SetUpdatedAt(in.UnixNow())
		fac.SetUpdatedBy(sess.UserId)

		if !fac.DoUpsert() {
			out.SetError(500, ErrUserFacilitySaveFailed)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		fac := rqProperty.NewFacilities(d.PropOltp)
		out.Facilities = fac.FindByPagination(
			&UserFacilityMeta,
			&in.Pager,
			&out.Pager,
		)
	}

	return
}
