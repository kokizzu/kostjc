package domain

import (
	"encoding/json"
	"testing"

	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/zCrud"

	"github.com/kokizzu/lexid"
	"github.com/stretchr/testify/require"
)

func TestStaffPaymentPromiseInputJSONUsesStringEncodedIds(t *testing.T) {
	var in StaffPaymentPromiseIn
	err := json.Unmarshal([]byte(`{
		"cmd": "upsert",
		"promise": {
			"tenantId": "123",
			"bookingId": "456",
			"promisedAt": "2026-07-25",
			"promisePayAt": "2026-07-29",
			"note": "payday"
		}
	}`), &in)
	require.NoError(t, err)
	require.Equal(t, uint64(123), in.Promise.TenantId)
	require.Equal(t, uint64(456), in.Promise.BookingId)
	require.Equal(t, `2026-07-25`, in.Promise.PromisedAt)
	require.Equal(t, `2026-07-29`, in.Promise.PromisePayAt)
	require.Equal(t, `payday`, in.Promise.Note)

	err = json.Unmarshal([]byte(`{
		"cmd": "upsert",
		"promise": {
			"tenantId": 123,
			"bookingId": 456,
			"promisedAt": "2026-07-25",
			"promisePayAt": "2026-07-29"
		}
	}`), &in)
	require.Error(t, err)
}

func TestStaffPaymentPromisePersistsPersonDatesAndNote(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	suffix := lexid.ID()
	tenant := createTestTenant(t, d, `Payment Promise Tenant`, `ktp-promise-`+suffix)
	room := createTestRoom(t, d, `payment-promise-room-`+suffix)
	booking := upsertTestBooking(t, d, rqProperty.Bookings{
		RoomId:        room.Id,
		TenantId:      tenant.Id,
		DateStart:     `2026-07-01`,
		DateEnd:       `2026-07-31`,
		BasePriceIDR:  1_000_000,
		TotalPriceIDR: 1_000_000,
	})

	out := d.StaffPaymentPromise(&StaffPaymentPromiseIn{
		RequestCommon: testStaffRequestCommon(StaffPaymentPromiseAction),
		Cmd:           zCrud.CmdUpsert,
		Promise: rqProperty.PaymentPromises{
			TenantId:     tenant.Id,
			BookingId:    booking.Id,
			PromisedAt:   `2026-07-25`,
			PromisePayAt: `2026-07-29`,
			Note:         `will transfer after payday`,
		},
	})
	require.False(t, out.HasError(), out.Error)
	require.NotZero(t, out.Promise.Id)
	require.Equal(t, tenant.Id, out.Promise.TenantId)
	require.Equal(t, booking.Id, out.Promise.BookingId)
	require.Equal(t, `2026-07-25`, out.Promise.PromisedAt)
	require.Equal(t, `2026-07-29`, out.Promise.PromisePayAt)
	require.Equal(t, `will transfer after payday`, out.Promise.Note)

	latest := rqProperty.NewPaymentPromises(d.PropOltp).FindLatestByTenantIds([]uint64{tenant.Id})
	require.Contains(t, latest, tenant.Id)
	require.Equal(t, out.Promise.Id, latest[tenant.Id].Id)
	require.Equal(t, `will transfer after payday`, latest[tenant.Id].Note)
}

func TestStaffPaymentPromiseRejectsBookingForDifferentPerson(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	suffix := lexid.ID()
	bookingTenant := createTestTenant(t, d, `Promise Booking Tenant`, `ktp-promise-booking-`+suffix)
	promiseTenant := createTestTenant(t, d, `Promise Other Tenant`, `ktp-promise-other-`+suffix)
	room := createTestRoom(t, d, `payment-promise-mismatch-room-`+suffix)
	booking := upsertTestBooking(t, d, rqProperty.Bookings{
		RoomId:        room.Id,
		TenantId:      bookingTenant.Id,
		DateStart:     `2026-07-01`,
		DateEnd:       `2026-07-31`,
		BasePriceIDR:  1_000_000,
		TotalPriceIDR: 1_000_000,
	})

	out := d.StaffPaymentPromise(&StaffPaymentPromiseIn{
		RequestCommon: testStaffRequestCommon(StaffPaymentPromiseAction),
		Cmd:           zCrud.CmdUpsert,
		Promise: rqProperty.PaymentPromises{
			TenantId:     promiseTenant.Id,
			BookingId:    booking.Id,
			PromisedAt:   `2026-07-25`,
			PromisePayAt: `2026-07-29`,
		},
	})
	require.True(t, out.HasError())
	require.Equal(t, ErrStaffPaymentPromiseBookingMismatch, out.Error)
}
