package domain

import (
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty"
	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/mProperty/wcProperty"
	"kostjc/model/zCrud"
	"strings"
	"time"

	"github.com/kokizzu/gotro/S"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminRoom.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminRoom.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminRoom.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminRoom.go
//go:generate farify doublequote --file AdminRoom.go

type (
	AdminRoomIn struct {
		RequestCommon
		Cmd      string           `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool             `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn    `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Room     rqProperty.Rooms `json:"room" form:"room" query:"room" long:"room" msg:"room"`
	}
	AdminRoomOut struct {
		ResponseCommon
		Pager zCrud.PagerOut   `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta  *zCrud.Meta      `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Room  rqProperty.Rooms `json:"room" form:"room" query:"room" long:"room" msg:"room"`
		Rooms [][]any          `json:"rooms" form:"rooms" query:"rooms" long:"rooms" msg:"rooms"`
	}
)

const (
	AdminRoomAction = `admin/room`

	ErrAdminRoomNotFound         = `room not found`
	ErrAdminRoomSaveFailed       = `failed to save room`
	ErrAdminRoomDeleteFailed     = `failed to delete room`
	ErrAdminRoomRestoreFailed    = `failed to restore room`
	ErrAdminRoomTenantNotFound   = `tenant not found`
	ErrAdminRoomBuildingNotFound = `building not found`
	ErrAdminRoomInvalidSize      = `invalid room size, must be both Number x Number, e.g: 2x2`
)

var AdminRoomMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mProperty.Id,
			Label:     `ID`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:        mProperty.RoomName,
			Label:       `Room Name`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			ReadOnly:    false,
			Description: `D13`,
		},
		{
			Name:        mProperty.RoomSize,
			Label:       `Room Size`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			ReadOnly:    false,
			Description: `2x2.5`,
		},
		{
			Name:      mProperty.BasePriceIDR,
			Label:     `Base Price`,
			DataType:  zCrud.DataTypeCurrency,
			InputType: zCrud.InputTypeNumber,
			ReadOnly:  false,
			Mapping:   `IDR`,
		},
		{
			Name:      mProperty.CurrentTenantId,
			Label:     `Current Tenant`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
			ReadOnly:  false,
		},
		{
			Name:      mProperty.BuildingId,
			Label:     `Building`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
			ReadOnly:  false,
		},
		{
			Name:      mProperty.FirstUseAt,
			Label:     `First Use At`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  false,
		},
		{
			Name:      mProperty.LastUseAt,
			Label:     `Last Use At`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  false,
		},
		{
			Name:        mProperty.ImageUrl,
			Label:       `Image URL`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			ReadOnly:    false,
			Description: `img-D13.jpg`,
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

func (d *Domain) AdminRoom(in *AdminRoomIn) (out AdminRoomOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.actor = sess.UserId
	out.refId = in.Room.Id

	if in.WithMeta {
		out.Meta = &AdminRoomMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		room := wcProperty.NewRoomsMutator(d.PropOltp)
		room.Id = in.Room.Id
		if room.Id > 0 {
			if !room.FindById() {
				out.SetError(400, ErrAdminRoomNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if room.DeletedAt == 0 {
					room.SetDeletedAt(in.UnixNow())
				}
			}

			if in.Cmd == zCrud.CmdRestore {
				if room.DeletedAt > 0 {
					room.SetDeletedAt(0)
					room.SetRestoredBy(sess.UserId)
				}
			}
		}

		if in.Room.RoomName != `` {
			room.SetRoomName(in.Room.RoomName)
		}

		if in.Room.BasePriceIDR != 0 {
			room.SetBasePriceIDR(in.Room.BasePriceIDR)
		}

		if in.Room.CurrentTenantId != 0 {
			tenant := rqAuth.NewTenants(d.AuthOltp)
			tenant.Id = in.Room.CurrentTenantId
			if !tenant.FindById() {
				out.SetError(400, ErrAdminRoomTenantNotFound)
				return
			}
			room.SetCurrentTenantId(in.Room.CurrentTenantId)
		}

		if in.Room.BuildingId != 0 {
			building := rqProperty.NewBuildings(d.PropOltp)
			building.Id = in.Room.BuildingId
			if !building.FindById() {
				out.SetError(400, ErrAdminRoomBuildingNotFound)
				return
			}
			room.SetBuildingId(in.Room.BuildingId)
		}

		if mProperty.IsValidDate(in.Room.FirstUseAt, time.DateOnly) {
			room.SetFirstUseAt(in.Room.FirstUseAt)
		}

		if mProperty.IsValidDate(in.Room.LastUseAt, time.DateOnly) {
			room.SetLastUseAt(in.Room.LastUseAt)
		}

		if in.Room.ImageUrl != `` {
			room.SetImageUrl(in.Room.ImageUrl)
		}

		if in.Room.RoomSize != `` {
			if !isValidRoomSize(in.Room.RoomSize) {
				out.SetError(400, ErrAdminRoomInvalidSize)
				return
			}

			room.SetRoomSize(in.Room.RoomSize)
		}

		if room.Id == 0 {
			room.SetCreatedAt(in.UnixNow())
			room.SetCreatedBy(sess.UserId)
		}

		room.SetUpdatedAt(in.UnixNow())
		room.SetUpdatedBy(sess.UserId)

		if !room.DoUpsert() {
			out.SetError(500, ErrAdminRoomSaveFailed)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		room := rqProperty.NewRooms(d.PropOltp)
		out.Rooms = room.FindByPagination(
			&AdminRoomMeta,
			&in.Pager,
			&out.Pager,
		)
	}

	return
}

func isValidRoomSize(rawsize string) bool {
	if !strings.Contains(rawsize, `x`) {
		return false
	}

	sizeSlice := S.Split(rawsize, `x`)

	if len(sizeSlice) != 2 {
		return false
	}

	var sizeX, sizeY float64

	sizeX = S.ToF(sizeSlice[0])
	if sizeX <= 0 {
		sizeX = float64(S.ToU(sizeSlice[0]))
	}

	if sizeX <= 0 {
		return false
	}

	sizeY = S.ToF(sizeSlice[1])
	if sizeY <= 0 {
		sizeY = float64(S.ToU(sizeSlice[1]))
	}

	if sizeY <= 0 {
		return false
	}

	return true
}
