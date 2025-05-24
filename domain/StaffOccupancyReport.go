package domain

import (
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty/rqProperty"

	"github.com/kokizzu/gotro/M"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file StaffOccupancyReport.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type StaffOccupancyReport.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type StaffOccupancyReport.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type StaffOccupancyReport.go
//go:generate farify doublequote --file StaffOccupancyReport.go

type (
	StaffOccupancyReportIn struct {
		RequestCommon
		MonthStart string `json:"monthStart" form:"monthStart" query:"monthStart" long:"monthStart" msg:"monthStart"`
		MonthEnd   string `json:"monthEnd" form:"monthEnd" query:"monthEnd" long:"monthEnd" msg:"monthEnd"`
	}
	StaffOccupancyReportOut struct {
		ResponseCommon
		User *rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		Segments  M.SB                       `json:"segments" form:"segments" query:"segments" long:"segments" msg:"segments"`
		RoomNames []string                   `json:"roomNames" form:"roomNames" query:"roomNames" long:"roomNames" msg:"roomNames"`
		Bookings  []rqProperty.BookingDetail `json:"bookingsPerQuartal" form:"bookingsPerQuartal" query:"bookingsPerQuartal" long:"bookingsPerQuartal" msg:"bookingsPerQuartal"`
	}
)

const (
	StaffOccupancyReportAction = `staff/occupancyReport`

	ErrStaffOccupancyReportNotFound = `user not found`
)

func (d *Domain) StaffOccupancyReport(in *StaffOccupancyReportIn) (out StaffOccupancyReportOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	room := rqProperty.NewRooms(d.PropOltp)
	out.RoomNames = room.FindRoomNames()

	booking := rqProperty.NewBookings(d.PropOltp)
	out.Bookings = booking.FindBookingsPerQuartal(in.MonthStart, in.MonthEnd)

	return
}
