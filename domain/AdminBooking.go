package domain

import (
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty"
	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/mProperty/wcProperty"
	"kostjc/model/zCrud"
	"time"

	"github.com/goccy/go-json"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminBooking.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminBooking.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminBooking.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminBooking.go
//go:generate farify doublequote --file AdminBooking.go

type (
	AdminBookingIn struct {
		RequestCommon
		Cmd        string              `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta   bool                `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager      zCrud.PagerIn       `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Facilities []uint64            `json:"facilities" form:"facilities" query:"facilities" long:"facilities" msg:"facilities"`
		Booking    rqProperty.Bookings `json:"booking" form:"booking" query:"booking" long:"booking" msg:"booking"`
	}
	AdminBookingOut struct {
		ResponseCommon
		Pager    zCrud.PagerOut      `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta     *zCrud.Meta         `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Booking  rqProperty.Bookings `json:"booking" form:"booking" query:"booking" long:"booking" msg:"booking"`
		Bookings [][]any             `json:"bookings" form:"bookings" query:"bookings" long:"bookings" msg:"bookings"`
	}
)

const (
	AdminBookingAction = `admin/booking`

	ErrAdminBookingNotFound      = `booking not found`
	ErrAdminBookingSaveFailed    = `failed to save booking`
	ErrAdminBookingDeleteFailed  = `failed to delete booking`
	ErrAdminBookingRestoreFailed = `failed to restore booking`
)

var AdminBookingMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mProperty.Id,
			Label:     `ID`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:      mProperty.DateStart,
			Label:     `Date Start`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  false,
		},
		{
			Name:      mProperty.DateEnd,
			Label:     `Date End`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  false,
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
			Name:      mProperty.FacilitiesObj,
			Label:     `Facilities`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeCombobox,
			ReadOnly:  true,
		},
		{
			Name:      mProperty.TotalPriceIDR,
			Label:     `Total Price`,
			DataType:  zCrud.DataTypeCurrency,
			InputType: zCrud.InputTypeNumber,
			ReadOnly:  false,
			Mapping:   `IDR`,
		},
		{
			Name:      mProperty.PaidAt,
			Label:     `Paid At`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  false,
		},
		{
			Name:      mProperty.TenantId,
			Label:     `Tenant`,
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

func (d *Domain) AdminBooking(in *AdminBookingIn) (out AdminBookingOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.actor = sess.UserId
	out.refId = in.Booking.Id

	if in.WithMeta {
		out.Meta = &AdminBookingMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		bk := wcProperty.NewBookingsMutator(d.PropOltp)
		bk.Id = in.Booking.Id
		if bk.Id > 0 {
			if !bk.FindById() {
				out.SetError(400, ErrAdminBookingNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if bk.DeletedAt == 0 {
					bk.SetDeletedAt(in.UnixNow())
					bk.SetDeletedBy(sess.UserId)
				}
			}

			if in.Cmd == zCrud.CmdRestore {
				if bk.DeletedAt > 0 {
					bk.SetDeletedAt(0)
					bk.SetRestoredBy(sess.UserId)
				}
			}
		}

		if mProperty.IsValidDate(in.Booking.DateStart, time.DateOnly) {
			bk.SetDateStart(in.Booking.DateStart)
		}

		if mProperty.IsValidDate(in.Booking.DateEnd, time.DateOnly) {
			bk.SetDateEnd(in.Booking.DateEnd)
		}

		if len(in.Facilities) > 0 {
			facilities := []mProperty.FacilityObj{}
			for _, id := range in.Facilities {
				fac := rqProperty.NewFacilities(d.PropOltp)
				fac.Id = id
				if fac.FindById() {
					facilities = append(facilities, mProperty.FacilityObj{
						FacilityName:   fac.FacilityName,
						ExtraChargeIDR: fac.ExtraChargeIDR,
					})
				}
			}

			facilitiesByt, err := json.Marshal(facilities)
			if err != nil {
				out.SetError(500, ErrAdminBookingSaveFailed)
				return
			} else {
				bk.SetFacilitiesObj(string(facilitiesByt))
			}
		}

		if in.Booking.BasePriceIDR > 0 {
			bk.SetBasePriceIDR(in.Booking.BasePriceIDR)
		}

		if in.Booking.TotalPriceIDR > 0 {
			bk.SetTotalPriceIDR(in.Booking.TotalPriceIDR)
		}

		if mProperty.IsValidDate(in.Booking.PaidAt, time.DateOnly) {
			bk.SetPaidAt(in.Booking.PaidAt)
		}

		if in.Booking.TenantId > 0 {
			tenant := rqAuth.NewTenants(d.AuthOltp)
			tenant.Id = in.Booking.TenantId

			if tenant.FindById() {
				bk.SetTenantId(in.Booking.TenantId)
			}
		}

		if bk.Id == 0 {
			bk.SetCreatedAt(in.UnixNow())
			bk.SetCreatedBy(sess.UserId)
		}

		bk.SetUpdatedAt(in.UnixNow())
		bk.SetUpdatedBy(sess.UserId)

		if !bk.DoUpsert() {
			out.SetError(500, ErrAdminBookingSaveFailed)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		bk := rqProperty.NewBookings(d.PropOltp)
		out.Bookings = bk.FindByPagination(
			&AdminBookingMeta,
			&in.Pager,
			&out.Pager,
		)
	}

	return
}
