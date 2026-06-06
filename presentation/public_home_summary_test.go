package presentation

import (
	"testing"

	"kostjc/model/mProperty/rqProperty"
)

func TestBuildPublicLoginSummary(t *testing.T) {
	summary := buildPublicLoginSummary(
		[]rqProperty.UnpaidBookingTenant{
			{
				BookingId:  1,
				TenantName: "Alice",
				RoomName:   "A-101",
				TotalPaid:  1000,
				TotalPrice: 2000,
				DateStart:  "2026-03-01",
				DateEnd:    "2026-03-10",
			},
			{
				BookingId:  2,
				TenantName: "Bob",
				RoomName:   "B-202",
				TotalPaid:  0,
				TotalPrice: 3000,
				DateStart:  "2026-02-01",
				DateEnd:    "2026-02-20",
			},
			{
				BookingId:  3,
				TenantName: "Charlie",
				RoomName:   "C-303",
				TotalPaid:  0,
				TotalPrice: 3000,
				DateStart:  "2026-06-14",
				DateEnd:    "2026-06-30",
			},
		},
		[]rqProperty.AvailableRoom{
			{
				RoomName:       "A-101",
				AvailableAt:    "2026-04-01",
				IsAvailableNow: false,
				LastTenant:     "Alice",
				BasePriceIDR:   1000000,
				TotalPriceIDR:  1200000,
			},
			{
				RoomName:       "B-202",
				AvailableAt:    "2026-05-01",
				IsAvailableNow: false,
				LastTenant:     "Bob",
				BasePriceIDR:   2000000,
				TotalPriceIDR:  2200000,
			},
		},
		[]rqProperty.CurrentOccupant{
			{
				TenantName: "Zed",
				RoomName:   "A-999",
			},
			{
				TenantName: "Diana",
				RoomName:   "D-404",
			},
			{
				TenantName: "Diana",
				RoomName:   "D-405",
			},
			{
				TenantName: "Eve",
				RoomName:   "A-101",
			},
			{
				TenantName: "Frank",
				RoomName:   "F-606",
			},
		},
	)

	if len(summary.UnpaidTenants) != 2 {
		t.Fatalf("expected 2 unique unpaid tenants, got %d", len(summary.UnpaidTenants))
	}
	if summary.UnpaidTenants[0].TenantName != "Bob" || summary.UnpaidTenants[0].RoomName != "B-202" {
		t.Fatalf("unexpected first unpaid summary: %#v", summary.UnpaidTenants[0])
	}
	if summary.UnpaidTenants[0].LastDate != "2026-02-20" {
		t.Fatalf("unexpected first unpaid last date: %#v", summary.UnpaidTenants[0])
	}
	if summary.UnpaidTenants[1].TenantName != "Alice" || summary.UnpaidTenants[1].RoomName != "A-101" {
		t.Fatalf("unexpected second unpaid summary: %#v", summary.UnpaidTenants[1])
	}
	if summary.UnpaidTenants[1].LastDate != "2026-03-10" {
		t.Fatalf("unexpected second unpaid last date: %#v", summary.UnpaidTenants[1])
	}
	if len(summary.UnpaidTenants) != 2 {
		t.Fatalf("expected 2 unpaid tenants, got %d", len(summary.UnpaidTenants))
	}
	if len(summary.Rooms) != 3 {
		t.Fatalf("expected 3 room notices, got %d", len(summary.Rooms))
	}
	if summary.Rooms[0].RoomName != "A-101" || summary.Rooms[0].LastTenant != "Alice" || summary.Rooms[0].LastDate != "2026-04-01" {
		t.Fatalf("unexpected first room summary: %#v", summary.Rooms[0])
	}
	if summary.Rooms[1].RoomName != "B-202" || summary.Rooms[1].LastTenant != "Bob" || summary.Rooms[1].LastDate != "2026-05-01" {
		t.Fatalf("unexpected second room summary: %#v", summary.Rooms[1])
	}
	if summary.Rooms[2].RoomName != "C-303" || summary.Rooms[2].LastTenant != "Charlie" || summary.Rooms[2].LastDate != "2026-06-14" {
		t.Fatalf("unexpected future unpaid room summary: %#v", summary.Rooms[2])
	}
	if got := summary.UnpaidTenants[0].DaysAgo; got != unpaidDays("2026-02-01", "2026-02-20") {
		t.Fatalf("unexpected first unpaid days ago: got %d want %d", got, unpaidDays("2026-02-01", "2026-02-20"))
	}
	if got := summary.UnpaidTenants[1].DaysAgo; got != unpaidDays("2026-03-01", "2026-03-10") {
		t.Fatalf("unexpected second unpaid days ago: got %d want %d", got, unpaidDays("2026-03-01", "2026-03-10"))
	}
	if got := summary.Rooms[0].DaysAgo; got != signedDays("2026-04-01") {
		t.Fatalf("unexpected first room days ago: got %d want %d", got, signedDays("2026-04-01"))
	}
	if got := summary.Rooms[1].DaysAgo; got != signedDays("2026-05-01") {
		t.Fatalf("unexpected second room days ago: got %d want %d", got, signedDays("2026-05-01"))
	}
	if got := summary.Rooms[2].DaysAgo; got != signedDays("2026-06-14") {
		t.Fatalf("unexpected third room days ago: got %d want %d", got, signedDays("2026-06-14"))
	}
	if len(summary.Occupants) != 3 {
		t.Fatalf("expected 3 occupants, got %d", len(summary.Occupants))
	}
	if summary.Occupants[0].TenantName != "Zed" || len(summary.Occupants[0].RoomNames) != 1 || summary.Occupants[0].RoomNames[0] != "A-999" {
		t.Fatalf("unexpected first occupant summary: %#v", summary.Occupants[0])
	}
	if summary.Occupants[1].TenantName != "Diana" || len(summary.Occupants[1].RoomNames) != 2 || summary.Occupants[1].RoomNames[0] != "D-404" {
		t.Fatalf("unexpected second occupant summary: %#v", summary.Occupants[1])
	}
	if summary.Occupants[2].TenantName != "Frank" || len(summary.Occupants[2].RoomNames) != 1 || summary.Occupants[2].RoomNames[0] != "F-606" {
		t.Fatalf("unexpected third occupant summary: %#v", summary.Occupants[2])
	}
}
