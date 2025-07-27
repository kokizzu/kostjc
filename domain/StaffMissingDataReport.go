package domain

import (
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/zCrud"

	"github.com/kokizzu/gotro/M"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file StaffMissingDataReport.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type StaffMissingDataReport.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type StaffMissingDataReport.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type StaffMissingDataReport.go
//go:generate farify doublequote --file StaffMissingDataReport.go

type (
	StaffMissingDataReportIn struct {
		RequestCommon
		Cmd      string `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		TenantId uint64 `json:"tenantId" form:"tenantId" query:"tenantId" long:"tenantId" msg:"tenantId"`
	}
	StaffMissingDataReportOut struct {
		ResponseCommon
		User *rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		Segments    M.SB                               `json:"segments" form:"segments" query:"segments" long:"segments" msg:"segments"`
		MissingData []rqProperty.RoomMissingTenantData `json:"missingData" form:"missingData" query:"missingData" long:"missingData" msg:"missingData"`
		Payments    []rqProperty.PaymentOfBooking      `json:"payments" form:"payments" query:"payments" long:"payments" msg:"payments"`
		Bookings    []rqProperty.TenantBookingDetail   `json:"bookings" form:"bookings" query:"bookings" long:"bookings" msg:"bookings"`
	}
)

const (
	StaffMissingDataReportAction = `staff/missingDataReport`
)

func (d *Domain) StaffMissingDataReport(in *StaffMissingDataReportIn) (out StaffMissingDataReportOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	switch in.Cmd {
	case zCrud.CmdForm:
		if in.TenantId == 0 {
			break
		}

		booking := rqProperty.NewBookings(d.PropOltp)
		bookings := booking.GetBookingsByTenantId(in.TenantId)
		out.Bookings = bookings

		payment := rqProperty.NewPayments(d.PropOltp)
		payments := payment.GetPaymentsByTenantId(in.TenantId)
		out.Payments = payments
	default:
		room := rqProperty.NewRooms(d.PropOltp)
		out.MissingData = room.FindMissingTenantsData()
	}

	return
}
