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
		BuildingId uint64 `json:"buildingId" form:"buildingId" query:"buildingId" long:"buildingId" msg:"buildingId"`
	}
	StaffOccupancyReportOut struct {
		ResponseCommon
		User *rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		Segments        M.SB                       `json:"segments" form:"segments" query:"segments" long:"segments" msg:"segments"`
		BuildingChoices map[uint64]string          `json:"buildingChoices" form:"buildingChoices" query:"buildingChoices" long:"buildingChoices" msg:"buildingChoices"`
		RoomNames       []string                   `json:"roomNames" form:"roomNames" query:"roomNames" long:"roomNames" msg:"roomNames"`
		Bookings        []rqProperty.BookingDetail `json:"bookingsPerQuartal" form:"bookingsPerQuartal" query:"bookingsPerQuartal" long:"bookingsPerQuartal" msg:"bookingsPerQuartal"`
	}
)

const (
	StaffOccupancyReportAction = `staff/occupancyReport`

	ErrStaffOccupancyReportNotFound = `user not found`
)

func (d *Domain) StaffOccupancyReport(in *StaffOccupancyReportIn) (out StaffOccupancyReportOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustBelowStaff(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	building := rqProperty.NewBuildings(d.PropOltp)
	out.BuildingChoices = building.FindBuildingChoices()

	room := rqProperty.NewRooms(d.PropOltp)
	out.RoomNames = room.FindRoomNames(in.BuildingId)

	booking := rqProperty.NewBookings(d.PropOltp)
	out.Bookings = booking.FindBookingsPerQuartal(in.MonthStart, in.MonthEnd, in.BuildingId)

	return
}
