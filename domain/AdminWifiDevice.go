package domain

import (
	"kostjc/model/mProperty"
	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/mProperty/wcProperty"
	"kostjc/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminWifiDevice.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminWifiDevice.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminWifiDevice.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminWifiDevice.go
//go:generate farify doublequote --file AdminWifiDevice.go

type (
	AdminWifiDeviceIn struct {
		RequestCommon
		Cmd        string                 `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta   bool                   `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager      zCrud.PagerIn          `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		WifiDevice rqProperty.WifiDevices `json:"wifiDevice" form:"wifiDevice" query:"wifiDevice" long:"wifiDevice" msg:"wifiDevice"`
	}
	AdminWifiDeviceOut struct {
		ResponseCommon
		Pager       zCrud.PagerOut         `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta        *zCrud.Meta            `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		WifiDevice  rqProperty.WifiDevices `json:"wifiDevice" form:"wifiDevice" query:"wifiDevice" long:"wifiDevice" msg:"wifiDevice"`
		WifiDevices [][]any                `json:"wifiDevices" form:"wifiDevices" query:"wifiDevices" long:"wifiDevices" msg:"wifiDevices"`
	}
)

const (
	AdminWifiDeviceAction = `admin/wifiDevice`

	ErrAdminWifiDeviceNotFound      = `wifi device not found`
	ErrAdminWifiDeviceSaveFailed    = `failed to save wifi device`
	ErrAdminWifiDeviceDeleteFailed  = `failed to delete wifi device`
	ErrAdminWifiDeviceRestoreFailed = `failed to restore wifi device`
	ErrAdminWifiDeviceTypeInvalid   = `invalid wifi device type, must be Room/Building`
)

var AdminWifiDeviceMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mProperty.Id,
			Label:     `ID`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:      mProperty.StartAt,
			Label:     `Start At`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  false,
		},
		{
			Name:      mProperty.EndAt,
			Label:     `End At`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  false,
		},
		{
			Name:      mProperty.PaidAt,
			Label:     `Paid At`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  false,
		},
		{
			Name:        mProperty.PriceIDR,
			Label:       `Price`,
			DataType:    zCrud.DataTypeCurrency,
			InputType:   zCrud.InputTypeNumber,
			ReadOnly:    false,
			Mapping:     `IDR`,
			Description: `100`,
		},
		{
			Name:      mProperty.MacAddress,
			Label:     `Mac Address`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
			ReadOnly:  false,
		},
		{
			Name:      mProperty.RoomId,
			Label:     `Room`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
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

func (d *Domain) AdminWifiDevice(in *AdminWifiDeviceIn) (out AdminWifiDeviceOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.actor = sess.UserId
	out.refId = in.WifiDevice.Id

	if in.WithMeta {
		out.Meta = &AdminWifiDeviceMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		wd := wcProperty.NewWifiDevicesMutator(d.PropOltp)
		wd.Id = in.WifiDevice.Id

		if wd.Id > 0 {
			if !wd.FindById() {
				out.SetError(400, ErrAdminWifiDeviceNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if wd.DeletedAt == 0 {
					wd.SetDeletedAt(in.UnixNow())
				}
			}

			if in.Cmd == zCrud.CmdRestore {
				if wd.DeletedAt > 0 {
					wd.SetDeletedAt(0)
					wd.SetRestoredBy(sess.UserId)
				}
			}
		}

		// defer InsertPropertyLog(wd.Id, d.wifiDeviceLog, out.ResponseCommon, in.TimeNow(), sess.UserId, wd)()

		wd.SetStartAt(in.WifiDevice.StartAt)
		wd.SetEndAt(in.WifiDevice.EndAt)
		wd.SetPaidAt(in.WifiDevice.PaidAt)
		wd.SetPriceIDR(in.WifiDevice.PriceIDR)
		wd.SetTenantId(in.WifiDevice.TenantId)
		wd.SetRoomId(in.WifiDevice.RoomId)
		wd.SetMacAddress(in.WifiDevice.MacAddress)

		if wd.Id == 0 {
			wd.SetCreatedAt(in.UnixNow())
			wd.SetCreatedBy(sess.UserId)
		}

		wd.SetUpdatedAt(in.UnixNow())
		wd.SetUpdatedBy(sess.UserId)

		if !wd.DoUpsert() {
			out.SetError(500, ErrAdminWifiDeviceSaveFailed)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		wd := rqProperty.NewFacilities(d.PropOltp)
		out.WifiDevices = wd.FindByPagination(
			&AdminWifiDeviceMeta,
			&in.Pager,
			&out.Pager,
		)
	}

	return
}
