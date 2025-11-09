package domain

import (
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty/rqProperty"

	"github.com/kokizzu/gotro/M"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file StaffOccupancyHeatmap.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type StaffOccupancyHeatmap.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type StaffOccupancyHeatmap.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type StaffOccupancyHeatmap.go
//go:generate farify doublequote --file StaffOccupancyHeatmap.go

type (
	StaffOccupancyHeatmapIn struct {
		RequestCommon
		YearMonth string `json:"yearMonth" form:"yearMonth" query:"yearMonth" long:"yearMonth" msg:"yearMonth"`
	}
	StaffOccupancyHeatmapOut struct {
		ResponseCommon
		User *rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		Segments  M.SB                               `json:"segments" form:"segments" query:"segments" long:"segments" msg:"segments"`
		RoomNames []string                           `json:"roomNames" form:"roomNames" query:"roomNames" long:"roomNames" msg:"roomNames"`
		Bookings  []rqProperty.BookingDetailPerMonth `json:"bookingsPerQuartal" form:"bookingsPerQuartal" query:"bookingsPerQuartal" long:"bookingsPerQuartal" msg:"bookingsPerQuartal"`
	}
)

const (
	StaffOccupancyHeatmapAction = `staff/occupancyHeatmap`

	ErrStaffOccupancyHeatmapNotFound = `user not found`
)

func (d *Domain) StaffOccupancyHeatmap(in *StaffOccupancyHeatmapIn) (out StaffOccupancyHeatmapOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	room := rqProperty.NewRooms(d.PropOltp)
	out.RoomNames = room.FindRoomNames()

	booking := rqProperty.NewBookings(d.PropOltp)
	out.Bookings = booking.FindBookingsPerMonth(in.YearMonth)

	return
}
