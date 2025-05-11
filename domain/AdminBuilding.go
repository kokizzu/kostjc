package domain

import (
	"kostjc/model/mProperty"
	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/mProperty/wcProperty"
	"kostjc/model/zCrud"

	"github.com/kokizzu/gotro/X"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminBuilding.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminBuilding.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminBuilding.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminBuilding.go
//go:generate farify doublequote --file AdminBuilding.go

type (
	AdminBuildingIn struct {
		RequestCommon
		Cmd      string               `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool                 `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn        `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Building rqProperty.Buildings `json:"building" form:"building" query:"building" long:"building" msg:"building"`
	}
	AdminBuildingOut struct {
		ResponseCommon
		Pager     zCrud.PagerOut       `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta      *zCrud.Meta          `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Building  rqProperty.Buildings `json:"building" form:"building" query:"building" long:"building" msg:"building"`
		Buildings [][]any              `json:"buildings" form:"buildings" query:"buildings" long:"buildings" msg:"buildings"`
	}
)

const (
	AdminBuildingAction = `admin/building`

	ErrAdminBuildingNotFound          = `building not found`
	ErrAdminBuildingSaveFailed        = `failed to save building`
	ErrAdminBuildingDeleteFailed      = `failed to delete building`
	ErrAdminBuildingRestoreFailed     = `failed to restore building`
	ErrAdminBuildingLocationNotFound  = `location not found`
	ErrAdminBuildingInvalidFacilities = `invalid facilities`
)

var AdminBuildingMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mProperty.Id,
			Label:     `ID`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:      mProperty.BuildingName,
			Label:     `Building Name`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
			ReadOnly:  false,
		},
		{
			Name:      mProperty.LocationId,
			Label:     `Location`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
			ReadOnly:  false,
		},
		{
			Name:      mProperty.Facilities,
			Label:     `Facilities`,
			DataType:  zCrud.DataTypeIntArr,
			InputType: zCrud.InputTypeMultiSelect,
			ReadOnly:  false,
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

func (d *Domain) AdminBuilding(in *AdminBuildingIn) (out AdminBuildingOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.actor = sess.UserId
	out.refId = in.Building.Id

	if in.WithMeta {
		out.Meta = &AdminBuildingMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		bld := wcProperty.NewBuildingsMutator(d.PropOltp)
		bld.Id = in.Building.Id
		if bld.Id > 0 {
			if !bld.FindById() {
				out.SetError(400, ErrAdminBuildingNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if bld.DeletedAt == 0 {
					bld.SetDeletedAt(in.UnixNow())
				}
			}

			if in.Cmd == zCrud.CmdRestore {
				if bld.DeletedAt > 0 {
					bld.SetDeletedAt(0)
					bld.SetRestoredBy(sess.UserId)
				}
			}
		}

		if in.Building.BuildingName != "" {
			bld.SetBuildingName(in.Building.BuildingName)
		}

		if in.Building.LocationId > 0 {
			loc := rqProperty.NewLocations(d.PropOltp)
			loc.Id = in.Building.LocationId
			if !loc.FindById() {
				out.SetError(400, ErrAdminBuildingLocationNotFound)
				return
			}
			bld.SetLocationId(in.Building.LocationId)
		}

		for _, v := range in.Building.Facilities {
			facId := X.ToU(v)
			fac := rqProperty.NewFacilities(d.PropOltp)
			fac.Id = facId
			if !fac.FindById() {
				out.SetError(400, ErrAdminBuildingInvalidFacilities)
				return
			}

			if fac.FacilityType != mProperty.FacilityTypeBuilding {
				out.SetError(400, ErrAdminBuildingInvalidFacilities)
				return
			}
		}

		if in.Building.Id > 0 {
			bld.SetFacilities(in.Building.Facilities)
		}

		if bld.Id == 0 {
			bld.SetCreatedAt(in.UnixNow())
			bld.SetCreatedBy(sess.UserId)
		}

		bld.SetUpdatedAt(in.UnixNow())
		bld.SetUpdatedBy(sess.UserId)

		if !bld.DoUpsert() {
			out.SetError(500, ErrAdminBuildingSaveFailed)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		bld := rqProperty.NewBuildings(d.PropOltp)
		out.Buildings = bld.FindByPagination(
			&AdminBuildingMeta,
			&in.Pager,
			&out.Pager,
		)
	}

	return
}
