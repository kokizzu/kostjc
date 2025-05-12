package domain

import (
	"kostjc/model/mProperty"
	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/mProperty/wcProperty"
	"kostjc/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminFacility.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminFacility.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminFacility.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminFacility.go
//go:generate farify doublequote --file AdminFacility.go

type (
	AdminFacilityIn struct {
		RequestCommon
		Cmd      string                `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool                  `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn         `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Facility rqProperty.Facilities `json:"facility" form:"facility" query:"facility" long:"facility" msg:"facility"`
	}
	AdminFacilityOut struct {
		ResponseCommon
		Pager      zCrud.PagerOut        `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta       *zCrud.Meta           `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Facility   rqProperty.Facilities `json:"facility" form:"facility" query:"facility" long:"facility" msg:"facility"`
		Facilities [][]any               `json:"facilities" form:"facilities" query:"facilities" long:"facilities" msg:"facilities"`
	}
)

const (
	AdminFacilityAction = `admin/facility`

	ErrAdminFacilityNotFound      = `facility not found`
	ErrAdminFacilitySaveFailed    = `failed to save facility`
	ErrAdminFacilityDeleteFailed  = `failed to delete facility`
	ErrAdminFacilityRestoreFailed = `failed to restore facility`
	ErrAdminFacilityTypeInvalid   = `invalid facility type, must be Room/Building`
)

var AdminFacilityMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mProperty.Id,
			Label:     `ID`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:        mProperty.FacilityName,
			Label:       `Facility Name`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			ReadOnly:    false,
			Description: `Bedcover`,
		},
		{
			Name:        mProperty.FacilityType,
			Label:       `Facility Type`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeComboboxArr,
			ReadOnly:    false,
			Description: `Room/Building/Site`,
		},
		{
			Name:        mProperty.ExtraChargeIDR,
			Label:       `Extra Charge`,
			DataType:    zCrud.DataTypeCurrency,
			InputType:   zCrud.InputTypeNumber,
			ReadOnly:    false,
			Mapping:     `IDR`,
			Description: `40000`,
		},
		{
			Name:        mProperty.DescriptionEN,
			Label:       `Description (english)`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeTextArea,
			ReadOnly:    false,
			Description: `Facility for kids`,
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

func (d *Domain) AdminFacility(in *AdminFacilityIn) (out AdminFacilityOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.actor = sess.UserId
	out.refId = in.Facility.Id

	if in.WithMeta {
		out.Meta = &AdminFacilityMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		fac := wcProperty.NewFacilitiesMutator(d.PropOltp)
		fac.Id = in.Facility.Id
		if fac.Id > 0 {
			if !fac.FindById() {
				out.SetError(400, ErrAdminFacilityNotFound)
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

		if in.Facility.FacilityType != `` {
			if !mProperty.IsValidFacilityType(in.Facility.FacilityType) {
				out.SetError(400, ErrAdminFacilityTypeInvalid)
				return
			}
			fac.SetFacilityType(in.Facility.FacilityType)
		}

		if in.Facility.ExtraChargeIDR != 0 {
			fac.SetExtraChargeIDR(in.Facility.ExtraChargeIDR)
		}

		if in.Facility.DescriptionEN != `` {
			fac.SetDescriptionEN(in.Facility.DescriptionEN)
		}

		if fac.Id == 0 {
			fac.SetCreatedAt(in.UnixNow())
			fac.SetCreatedBy(sess.UserId)
		}

		fac.SetUpdatedAt(in.UnixNow())
		fac.SetUpdatedBy(sess.UserId)

		if !fac.DoUpsert() {
			out.SetError(500, ErrAdminFacilitySaveFailed)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		fac := rqProperty.NewFacilities(d.PropOltp)
		out.Facilities = fac.FindByPagination(
			&AdminFacilityMeta,
			&in.Pager,
			&out.Pager,
		)
	}

	return
}
