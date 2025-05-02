package domain

import (
	"kostjc/model/mProperty"
	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/mProperty/wcProperty"
	"kostjc/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminLocation.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminLocation.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminLocation.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminLocation.go
//go:generate farify doublequote --file AdminLocation.go

type (
	AdminLocationIn struct {
		RequestCommon
		Cmd      string               `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool                 `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn        `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Location rqProperty.Locations `json:"location" form:"location" query:"location" long:"location" msg:"location"`
	}
	AdminLocationOut struct {
		ResponseCommon
		Pager     zCrud.PagerOut       `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta      *zCrud.Meta          `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Location  rqProperty.Locations `json:"location" form:"location" query:"location" long:"location" msg:"location"`
		Locations [][]any              `json:"locations" form:"locations" query:"locations" long:"locations" msg:"locations"`
	}
)

const (
	AdminLocationAction = `admin/location`

	ErrAdminLocationNotFound      = `location not found`
	ErrAdminLocationSaveFailed    = `failed to save location`
	ErrAdminLocationDeleteFailed  = `failed to delete location`
	ErrAdminLocationRestoreFailed = `failed to restore location`
)

var AdminLocationMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mProperty.Id,
			Label:     `ID`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:        mProperty.Name,
			Label:       `Name`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			ReadOnly:    false,
			Description: `Cafe MVP`,
		},
		{
			Name:        mProperty.Address,
			Label:       `Address`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			ReadOnly:    false,
			Description: `Jl. Raya, No. 1, Kec. Kuta, Kabupaten Badung, Bali 80361`,
		},
		{
			Name:        mProperty.GmapLocation,
			Label:       `Gmap Location`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			ReadOnly:    false,
			Description: `https://maps.google.com/`,
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

func (d *Domain) AdminLocation(in *AdminLocationIn) (out AdminLocationOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.actor = sess.UserId
	out.refId = in.Location.Id

	if in.WithMeta {
		out.Meta = &AdminLocationMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		loc := wcProperty.NewLocationsMutator(d.PropOltp)
		loc.Id = in.Location.Id
		if loc.Id > 0 {
			if !loc.FindById() {
				out.SetError(400, ErrAdminLocationNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if loc.DeletedAt == 0 {
					loc.SetDeletedAt(in.UnixNow())
				}
			}

			if in.Cmd == zCrud.CmdRestore {
				if loc.DeletedAt > 0 {
					loc.SetDeletedAt(0)
					loc.SetRestoredBy(sess.UserId)
				}
			}
		}

		if in.Location.Name != `` {
			loc.SetName(in.Location.Name)
		}

		if in.Location.Address != `` {
			loc.SetAddress(in.Location.Address)
		}

		if in.Location.GmapLocation != `` {
			loc.SetGmapLocation(in.Location.GmapLocation)
		}

		if loc.Id == 0 {
			loc.SetCreatedAt(in.UnixNow())
			loc.SetCreatedBy(sess.UserId)
		}

		loc.SetUpdatedAt(in.UnixNow())
		loc.SetUpdatedBy(sess.UserId)

		if !loc.DoUpsert() {
			out.SetError(500, ErrAdminLocationSaveFailed)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		loc := rqProperty.NewLocations(d.PropOltp)
		out.Locations = loc.FindByPagination(
			&AdminLocationMeta,
			&in.Pager,
			&out.Pager,
		)
	}

	return
}
