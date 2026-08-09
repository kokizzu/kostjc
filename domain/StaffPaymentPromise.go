package domain

import (
	"time"

	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty"
	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/mProperty/wcProperty"
	"kostjc/model/zCrud"
)

type (
	StaffPaymentPromiseIn struct {
		RequestCommon
		Cmd     string                     `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		Promise rqProperty.PaymentPromises `json:"promise" form:"promise" query:"promise" long:"promise" msg:"promise"`
	}
	StaffPaymentPromiseOut struct {
		ResponseCommon
		Promise rqProperty.PaymentPromises `json:"promise" form:"promise" query:"promise" long:"promise" msg:"promise"`
	}
)

const (
	StaffPaymentPromiseAction = `staff/paymentPromise`

	ErrStaffPaymentPromiseTenantRequired  = `tenant is required for payment promise`
	ErrStaffPaymentPromiseTenantNotFound  = `tenant not found for payment promise`
	ErrStaffPaymentPromiseBookingNotFound = `booking not found for payment promise`
	ErrStaffPaymentPromiseBookingMismatch = `booking does not belong to payment promise tenant`
	ErrStaffPaymentPromisePromisedAt      = `valid promised date is required`
	ErrStaffPaymentPromisePayAt           = `valid promised payment date is required`
	ErrStaffPaymentPromiseSaveFailed      = `failed to save payment promise`
)

func (d *Domain) StaffPaymentPromise(in *StaffPaymentPromiseIn) (out StaffPaymentPromiseOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustBelowStaff(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.actor = sess.UserId
	out.refId = in.Promise.Id

	switch in.Cmd {
	case zCrud.CmdUpsert:
		if in.Promise.TenantId == 0 {
			out.SetError(400, ErrStaffPaymentPromiseTenantRequired)
			return
		}

		tenant := rqAuth.NewTenants(d.AuthOltp)
		tenant.Id = in.Promise.TenantId
		if !tenant.FindById() {
			out.SetError(400, ErrStaffPaymentPromiseTenantNotFound)
			return
		}

		if !mProperty.IsValidDate(in.Promise.PromisedAt, time.DateOnly) {
			out.SetError(400, ErrStaffPaymentPromisePromisedAt)
			return
		}

		if !mProperty.IsValidDate(in.Promise.PromisePayAt, time.DateOnly) {
			out.SetError(400, ErrStaffPaymentPromisePayAt)
			return
		}

		if in.Promise.BookingId > 0 {
			booking := rqProperty.NewBookings(d.PropOltp)
			booking.Id = in.Promise.BookingId
			if !booking.FindById() {
				out.SetError(400, ErrStaffPaymentPromiseBookingNotFound)
				return
			}
			if booking.TenantId != in.Promise.TenantId {
				out.SetError(400, ErrStaffPaymentPromiseBookingMismatch)
				return
			}
		}

		promise := wcProperty.NewPaymentPromisesMutator(d.PropOltp)
		promise.SetTenantId(in.Promise.TenantId)
		promise.SetBookingId(in.Promise.BookingId)
		promise.SetPromisedAt(in.Promise.PromisedAt)
		promise.SetPromisePayAt(in.Promise.PromisePayAt)
		promise.SetNote(in.Promise.Note)
		promise.SetCreatedAt(in.UnixNow())
		promise.SetCreatedBy(sess.UserId)
		promise.SetUpdatedAt(in.UnixNow())
		promise.SetUpdatedBy(sess.UserId)

		if !promise.DoUpsert() {
			out.SetError(500, ErrStaffPaymentPromiseSaveFailed)
			return
		}

		out.Promise = promise.PaymentPromises
		out.refId = promise.Id
	}

	return
}
