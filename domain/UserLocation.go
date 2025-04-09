package domain

import (
	"kostjc/model/mProperty"
	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/mProperty/wcProperty"
	"kostjc/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserLocation.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserLocation.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserLocation.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserLocation.go
//go:generate farify doublequote --file UserLocation.go

type (
	UserLocationIn struct {
		RequestCommon
		Cmd      string               `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool                 `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn        `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Location rqProperty.Locations `json:"location" form:"location" query:"location" long:"location" msg:"location"`
	}
	UserLocationOut struct {
		ResponseCommon
		Pager     zCrud.PagerOut       `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta      *zCrud.Meta          `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Location  rqProperty.Locations `json:"location" form:"location" query:"location" long:"location" msg:"location"`
		Locations [][]any              `json:"locations" form:"locations" query:"locations" long:"locations" msg:"locations"`
	}
)

const (
	UserLocationAction = `user/location`

	ErrUserLocationNotFound      = `location not found`
	ErrUserLocationSaveFailed    = `failed to save location`
	ErrUserLocationDeleteFailed  = `failed to delete location`
	ErrUserLocationRestoreFailed = `failed to restore location`
)

var UserLocationMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mProperty.Id,
			Label:     `ID`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:      mProperty.Name,
			Label:     `Name`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
			ReadOnly:  false,
		},
		{
			Name:      mProperty.Address,
			Label:     `Address`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
			ReadOnly:  false,
		},
		{
			Name:      mProperty.GmapLocation,
			Label:     `Gmap Location`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
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
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
	},
}

func (d *Domain) UserLocation(in *UserLocationIn) (out UserLocationOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.actor = sess.UserId
	out.refId = in.Location.Id

	if in.WithMeta {
		out.Meta = &UserLocationMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		loc := wcProperty.NewLocationsMutator(d.PropOltp)
		loc.Id = in.Location.Id
		if loc.Id > 0 {
			if !loc.FindById() {
				out.SetError(400, ErrUserLocationNotFound)
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
			out.SetError(500, ErrUserLocationSaveFailed)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		loc := rqProperty.NewLocations(d.PropOltp)
		out.Locations = loc.FindByPagination(
			&UserLocationMeta,
			&in.Pager,
			&out.Pager,
		)
	}

	return
}
