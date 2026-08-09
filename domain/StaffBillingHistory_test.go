package domain

import (
	"testing"

	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/zCrud"

	"github.com/kokizzu/lexid"
	"github.com/stretchr/testify/require"
)

func TestStaffBillingHistoryReturnsTenantBookingsAndPayments(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	suffix := lexid.ID()
	tenant := createTestTenant(t, d, `Billing History Tenant`, `ktp-billing-`+suffix)
	room := createTestRoom(t, d, `billing-history-room-`+suffix)
	booking := upsertTestBooking(t, d, rqProperty.Bookings{
		RoomId:        room.Id,
		TenantId:      tenant.Id,
		DateStart:     `2026-07-10`,
		DateEnd:       `2026-08-09`,
		BasePriceIDR:  1_000_000,
		TotalPriceIDR: 1_000_000,
	})

	paymentOut := d.AdminPayment(&AdminPaymentIn{
		RequestCommon: testAdminRequestCommon(AdminPaymentAction),
		Cmd:           zCrud.CmdUpsert,
		Payment: rqProperty.Payments{
			BookingId:     booking.Id,
			PaymentAt:     `2026-07-15`,
			PaidIDR:       600_000,
			PaymentMethod: `Transfer`,
			PaymentStatus: `Paid`,
			Note:          `billing history regression test`,
		},
	})
	require.False(t, paymentOut.HasError(), paymentOut.Error)

	historyOut := d.StaffMissingDataReport(&StaffMissingDataReportIn{
		RequestCommon: testStaffRequestCommon(StaffMissingDataReportAction),
		Cmd:           zCrud.CmdForm,
		TenantId:      tenant.Id,
	})
	require.False(t, historyOut.HasError(), historyOut.Error)
	require.Len(t, historyOut.Bookings, 1)
	require.Equal(t, booking.Id, historyOut.Bookings[0].BookingId)
	require.Equal(t, int64(1_000_000), historyOut.Bookings[0].TotalPriceIDR)
	require.Equal(t, int64(600_000), historyOut.Bookings[0].TotalPaidIDR)
	require.Len(t, historyOut.Payments, 1)
	require.Equal(t, booking.Id, historyOut.Payments[0].BookingId)
	require.Equal(t, `2026-07-15`, historyOut.Payments[0].PaymentAt)
	require.Equal(t, int64(600_000), historyOut.Payments[0].PaidIDR)
	require.Equal(t, `Transfer`, historyOut.Payments[0].PaymentMethod)
	require.Equal(t, `billing history regression test`, historyOut.Payments[0].Note)

	unpaidBookings := rqProperty.NewBookings(d.PropOltp).FindUnpaidBookingTenants()
	var matchingUnpaid *rqProperty.UnpaidBookingTenant
	for index := range unpaidBookings {
		if unpaidBookings[index].BookingId == booking.Id {
			matchingUnpaid = &unpaidBookings[index]
			break
		}
	}
	require.NotNil(t, matchingUnpaid)
	require.Equal(t, tenant.Id, matchingUnpaid.TenantId)
	require.Equal(t, int64(600_000), matchingUnpaid.TotalPaid)
}
