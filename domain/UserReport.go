package domain

import (
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty/rqProperty"

	"github.com/kokizzu/gotro/M"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserReport.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserReport.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserReport.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserReport.go
//go:generate farify doublequote --file UserReport.go

type (
	UserReportIn struct {
		RequestCommon
		MonthStart string `json:"monthStart" form:"monthStart" query:"monthStart" long:"monthStart" msg:"monthStart"`
		MonthEnd   string `json:"monthEnd" form:"monthEnd" query:"monthEnd" long:"monthEnd" msg:"monthEnd"`
	}
	UserReportOut struct {
		ResponseCommon
		User *rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		Segments  M.SB                       `json:"segments" form:"segments" query:"segments" long:"segments" msg:"segments"`
		RoomNames []string                   `json:"roomNames" form:"roomNames" query:"roomNames" long:"roomNames" msg:"roomNames"`
		Bookings  []rqProperty.BookingDetail `json:"bookingsPerQuartal" form:"bookingsPerQuartal" query:"bookingsPerQuartal" long:"bookingsPerQuartal" msg:"bookingsPerQuartal"`
	}
)

const (
	UserReportAction = `user/occupancyReport`

	ErrUserReportNotFound = `user not found`
)

func (d *Domain) UserReport(in *UserReportIn) (out UserReportOut) {
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
