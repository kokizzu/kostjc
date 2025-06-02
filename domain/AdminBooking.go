package domain

import (
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty"
	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/mProperty/saProperty"
	"kostjc/model/mProperty/wcProperty"
	"kostjc/model/zCrud"
	"time"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/X"
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
		Pager    zCrud.PagerOut       `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta     *zCrud.Meta          `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Booking  *rqProperty.Bookings `json:"booking" form:"booking" query:"booking" long:"booking" msg:"booking"`
		Bookings [][]any              `json:"bookings" form:"bookings" query:"bookings" long:"bookings" msg:"bookings"`
	}
)

const (
	AdminBookingAction = `admin/booking`

	ErrAdminBookingNotFound       = `booking not found`
	ErrAdminBookingSaveFailed     = `failed to save booking`
	ErrAdminBookingDeleteFailed   = `failed to delete booking`
	ErrAdminBookingRestoreFailed  = `failed to restore booking`
	ErrAdminBookingTenantNotFound = `tenant not found`
	ErrAdminBookingRoomNotFound   = `room not found`
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
			Name:      mProperty.RoomId,
			Label:     `Room`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
			ReadOnly:  false,
		},
		{
			Name:      `totalPaidIDR`,
			Label:     `Total Paid`,
			DataType:  zCrud.DataTypeCurrency,
			InputType: zCrud.InputTypeNumber,
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
			Name:      mProperty.TenantId,
			Label:     `Tenant`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
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
			InputType: zCrud.InputTypeTextArea,
			ReadOnly:  false,
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
			Name:      mProperty.ExtraTenants,
			Label:     `Extra Tenants`,
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
		bk := rqProperty.NewBookings(d.PropOltp)
		bk.Id = in.Booking.Id
		if bk.Id > 0 {
			if !bk.FindById() {
				out.SetError(400, ErrAdminBookingNotFound)
				return
			}
		}
		out.Booking = bk
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		bk := wcProperty.NewBookingsMutator(d.PropOltp)
		bk.Id = in.Booking.Id

		isEdited := false
		beforeJson := []byte(``)
		if bk.Id > 0 {
			isEdited = true
			if !bk.FindById() {
				out.SetError(400, ErrAdminBookingNotFound)
				return
			}
			beforeJson, _ = json.Marshal(bk)

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

		if in.Booking.Id > 0 && in.Booking.FacilitiesObj != `` {
			bk.SetFacilitiesObj(in.Booking.FacilitiesObj)
		}

		bk.SetBasePriceIDR(in.Booking.BasePriceIDR)
		bk.SetTotalPriceIDR(in.Booking.TotalPriceIDR)

		if mProperty.IsValidDate(in.Booking.PaidAt, time.DateOnly) {
			bk.SetPaidAt(in.Booking.PaidAt)
		}

		if in.Booking.TenantId > 0 {
			tenant := rqAuth.NewTenants(d.AuthOltp)
			tenant.Id = in.Booking.TenantId
			if !tenant.FindById() {
				out.SetError(400, ErrAdminBookingTenantNotFound)
				return
			}

			bk.SetTenantId(in.Booking.TenantId)
		}

		for _, id := range in.Booking.ExtraTenants {
			// Skip extra tenant if same as main tenant
			if id == in.Booking.TenantId {
				continue
			}
			tenant := rqAuth.NewTenants(d.AuthOltp)
			tenant.Id = X.ToU(id)
			if !tenant.FindById() {
				out.SetError(400, ErrAdminBookingTenantNotFound)
				return
			}
		}

		if in.Cmd == zCrud.CmdUpsert {
			bk.SetExtraTenants(in.Booking.ExtraTenants)
		}

		if in.Booking.RoomId != 0 {
			room := wcProperty.NewRoomsMutator(d.PropOltp)
			room.Id = in.Booking.RoomId
			if !room.FindById() {
				out.SetError(400, ErrAdminBookingRoomNotFound)
				return
			}
			room.SetCurrentTenantId(in.Booking.TenantId)
			room.SetLastUseAt(in.Booking.DateEnd)
			room.SetUpdatedAt(in.UnixNow())
			room.SetUpdatedBy(sess.UserId)
			if !room.DoUpdateById() {
				out.SetError(500, ErrAdminBookingSaveFailed)
				return
			}
			bk.SetRoomId(in.Booking.RoomId)
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

		if isEdited {
			afterJson, _ := json.Marshal(bk)
			row := saProperty.BookingLogs{
				CreatedAt:  in.TimeNow(),
				ActorId:    sess.UserId,
				BeforeJson: string(beforeJson),
				AfterJson:  string(afterJson),
			}
			d.bookingLogs.Insert([]any{
				row.CreatedAt,
				row.ActorId,
				row.BeforeJson,
				row.AfterJson,
			})
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
