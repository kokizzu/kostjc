package domain

import (
	"kostjc/model/mProperty"
	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/mProperty/wcProperty"
	"kostjc/model/zCrud"
	"time"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserPayment.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserPayment.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserPayment.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserPayment.go
//go:generate farify doublequote --file UserPayment.go

type (
	UserPaymentIn struct {
		RequestCommon
		Cmd      string              `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool                `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn       `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Payment  rqProperty.Payments `json:"payment" form:"payment" query:"payment" long:"payment" msg:"payment"`
	}
	UserPaymentOut struct {
		ResponseCommon
		Pager    zCrud.PagerOut      `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta     *zCrud.Meta         `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Payment  rqProperty.Payments `json:"payment" form:"payment" query:"payment" long:"payment" msg:"payment"`
		Payments [][]any             `json:"payments" form:"payments" query:"payments" long:"payments" msg:"payments"`
	}
)

const (
	UserPaymentAction = `user/payment`

	ErrUserPaymentNotFound          = `payment not found`
	ErrUserPaymentBookingIdNotFound = `booking id not found for this payment`
	ErrUserPaymentSaveFailed        = `failed to save payment`
	ErrUserPaymentDeleteFailed      = `failed to delete payment`
	ErrUserPaymentRestoreFailed     = `failed to restore payment`
)

var UserPaymentMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mProperty.Id,
			Label:     `ID`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:      mProperty.BookingId,
			Label:     `Booking ID`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
			ReadOnly:  false,
		},
		{
			Name:      mProperty.PaymentAt,
			Label:     `Payment At`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  false,
		},
		{
			Name:      mProperty.PaidIDR,
			Label:     `Total Paid`,
			DataType:  zCrud.DataTypeCurrency,
			InputType: zCrud.InputTypeNumber,
			ReadOnly:  false,
			Mapping:   `IDR`,
		},
		{
			Name:      mProperty.PaymentMethod,
			Label:     `Payment Method`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
			ReadOnly:  false,
		},
		{
			Name:      mProperty.PaymentStatus,
			Label:     `Payment Status`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
			ReadOnly:  false,
		},
		{
			Name:      mProperty.Note,
			Label:     `Note`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeTextArea,
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

func (d *Domain) UserPayment(in *UserPaymentIn) (out UserPaymentOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.actor = sess.UserId
	out.refId = in.Payment.Id

	if in.WithMeta {
		out.Meta = &UserPaymentMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		pym := wcProperty.NewPaymentsMutator(d.PropOltp)
		pym.Id = in.Payment.Id
		if pym.Id > 0 {
			if !pym.FindById() {
				out.SetError(400, ErrUserPaymentNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if pym.DeletedAt == 0 {
					pym.SetDeletedAt(in.UnixNow())
				}
			}

			if in.Cmd == zCrud.CmdRestore {
				if pym.DeletedAt > 0 {
					pym.SetDeletedAt(0)
					pym.SetRestoredBy(sess.UserId)
				}
			}
		}

		if in.Payment.BookingId > 0 {
			bkd := rqProperty.NewBookings(d.PropOltp)
			bkd.Id = in.Payment.BookingId
			if !bkd.FindById() {
				out.SetError(400, ErrUserPaymentBookingIdNotFound)
				return
			}

			pym.SetBookingId(in.Payment.BookingId)
		}

		if mProperty.IsValidDate(in.Payment.PaymentAt, time.DateOnly) {
			pym.SetPaymentAt(in.Payment.PaymentAt)
		}

		if in.Payment.PaidIDR > 0 {
			pym.SetPaidIDR(in.Payment.PaidIDR)
		}

		if in.Payment.PaymentMethod != `` {
			pym.SetPaymentMethod(in.Payment.PaymentMethod)
		}

		if in.Payment.PaymentStatus != `` {
			pym.SetPaymentStatus(in.Payment.PaymentStatus)
		}

		if in.Payment.Note != `` {
			pym.SetNote(in.Payment.Note)
		}

		if pym.Id == 0 {
			pym.SetCreatedAt(in.UnixNow())
			pym.SetCreatedBy(sess.UserId)
		}

		pym.SetUpdatedAt(in.UnixNow())
		pym.SetUpdatedBy(sess.UserId)

		if !pym.DoUpsert() {
			out.SetError(500, ErrUserPaymentSaveFailed)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		pym := rqProperty.NewPayments(d.PropOltp)
		out.Payments = pym.FindByPagination(
			&UserPaymentMeta,
			&in.Pager,
			&out.Pager,
		)
	}

	return
}
