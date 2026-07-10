package domain

import (
	"fmt"
	"testing"

	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/zCrud"

	"github.com/kokizzu/lexid"
	"github.com/stretchr/testify/require"
)

func TestAdminBookingEditingOldBookingDoesNotOverwriteCurrentTenant(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	suffix := lexid.ID()
	olderTenant := createTestTenant(t, d, `Old Booking Tenant`, `ktp-old-`+suffix)
	newerTenant := createTestTenant(t, d, `Newer Booking Tenant`, `ktp-new-`+suffix)
	room := createTestRoom(t, d, `booking-regression-room-`+suffix)

	olderBooking := upsertTestBooking(t, d, rqProperty.Bookings{
		RoomId:        room.Id,
		TenantId:      olderTenant.Id,
		DateStart:     `2026-01-01`,
		DateEnd:       `2026-01-31`,
		BasePriceIDR:  1_000_000,
		TotalPriceIDR: 1_000_000,
	})

	upsertTestBooking(t, d, rqProperty.Bookings{
		RoomId:        room.Id,
		TenantId:      newerTenant.Id,
		DateStart:     `2026-02-01`,
		DateEnd:       `2026-02-28`,
		BasePriceIDR:  1_000_000,
		TotalPriceIDR: 1_000_000,
	})

	roomAfterNewerBooking := rqProperty.NewRooms(testTt)
	roomAfterNewerBooking.Id = room.Id
	require.True(t, roomAfterNewerBooking.FindById())
	require.Equal(t, newerTenant.Id, roomAfterNewerBooking.CurrentTenantId)

	upsertTestBooking(t, d, rqProperty.Bookings{
		Id:            olderBooking.Id,
		RoomId:        room.Id,
		TenantId:      olderTenant.Id,
		DateStart:     `2026-01-01`,
		DateEnd:       `2026-01-30`,
		BasePriceIDR:  1_100_000,
		TotalPriceIDR: 1_100_000,
	})

	roomAfterOldBookingEdit := rqProperty.NewRooms(testTt)
	roomAfterOldBookingEdit.Id = room.Id
	require.True(t, roomAfterOldBookingEdit.FindById())
	require.Equal(t, newerTenant.Id, roomAfterOldBookingEdit.CurrentTenantId)
	require.Equal(t, `2026-02-28`, roomAfterOldBookingEdit.LastUseAt)
}

func createTestTenant(t *testing.T, d *Domain, tenantName, ktpNumber string) rqAuth.Tenants {
	t.Helper()

	out := d.AdminTenants(&AdminTenantsIn{
		RequestCommon: testAdminRequestCommon(AdminTenantsAction),
		Cmd:           zCrud.CmdUpsert,
		Tenant: rqAuth.Tenants{
			TenantName:       fmt.Sprintf(`%s %s`, tenantName, ktpNumber),
			KtpRegion:        `BALI`,
			KtpNumber:        ktpNumber,
			KtpName:          tenantName,
			KtpPlaceBirth:    `DENPASAR`,
			KtpDateBirth:     `1990-01-01`,
			KtpGender:        `Laki - Laki`,
			KtpAddress:       `JALAN TEST`,
			KtpRtRw:          `001/001`,
			KtpKelurahanDesa: `Dalung`,
			KtpKecamatan:     `Kuta Utara`,
			KtpReligion:      `Islam`,
			KtpMaritalStatus: `Belum Kawin`,
			KtpCitizenship:   `WNI`,
			KtpOccupation:    `Programmer`,
			WhatsappNumber:   `+628000000000`,
		},
	})
	require.False(t, out.HasError(), out.Error)

	tenant := rqAuth.NewTenants(testTt)
	tenant.KtpNumber = ktpNumber
	require.True(t, tenant.FindByKtpNumber())
	require.NotZero(t, tenant.Id)

	return *tenant
}

func createTestRoom(t *testing.T, d *Domain, roomName string) rqProperty.Rooms {
	t.Helper()

	out := d.AdminRoom(&AdminRoomIn{
		RequestCommon: testAdminRequestCommon(AdminRoomAction),
		Cmd:           zCrud.CmdUpsert,
		Room: rqProperty.Rooms{
			RoomName:     roomName,
			RoomSize:     `2x2`,
			BasePriceIDR: 1_000_000,
		},
	})
	require.False(t, out.HasError(), out.Error)

	room := rqProperty.NewRooms(testTt)
	room.RoomName = roomName
	require.True(t, room.FindByRoomName())
	require.NotZero(t, room.Id)

	return *room
}

func upsertTestBooking(t *testing.T, d *Domain, booking rqProperty.Bookings) rqProperty.Bookings {
	t.Helper()

	if booking.ExtraTenants == nil {
		booking.ExtraTenants = []any{}
	}

	out := d.AdminBooking(&AdminBookingIn{
		RequestCommon: testAdminRequestCommon(AdminBookingAction),
		Cmd:           zCrud.CmdUpsert,
		Booking:       booking,
	})
	require.False(t, out.HasError(), out.Error)

	savedBooking := rqProperty.NewBookings(testTt)
	if booking.Id > 0 {
		savedBooking.Id = booking.Id
		require.True(t, savedBooking.FindById())
		return *savedBooking
	}

	query := fmt.Sprintf(`
SELECT %s
FROM "bookings"
WHERE "roomId" = %d
	AND "tenantId" = %d
	AND "dateStart" = '%s'
	AND "dateEnd" = '%s'
ORDER BY "id" DESC
LIMIT 1`,
		savedBooking.SqlSelectAllFields(),
		booking.RoomId,
		booking.TenantId,
		booking.DateStart,
		booking.DateEnd,
	)

	found := false
	savedBooking.Adapter.QuerySql(query, func(row []any) {
		savedBooking.FromArray(row)
		found = true
	})
	require.True(t, found)
	require.NotZero(t, savedBooking.Id)

	return *savedBooking
}
