package domain

import (
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty/rqProperty"

	"github.com/kokizzu/gotro/M"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file StaffRevenueReport.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type StaffRevenueReport.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type StaffRevenueReport.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type StaffRevenueReport.go
//go:generate farify doublequote --file StaffRevenueReport.go

type (
	StaffRevenueReportIn struct {
		RequestCommon
		YearMonth string `json:"yearMonth" form:"yearMonth" query:"yearMonth" long:"yearMonth" msg:"yearMonth"`
	}
	StaffRevenueReportOut struct {
		ResponseCommon
		User *rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		Segments       M.SB                       `json:"segments" form:"segments" query:"segments" long:"segments" msg:"segments"`
		RevenueReports []rqProperty.RevenueReport `json:"revenueReports" form:"revenueReports" query:"revenueReports" long:"revenueReports" msg:"revenueReports"`
	}
)

const (
	StaffRevenueReportAction = `staff/revenueReport`

	ErrStaffRevenueReportNotFound = `user not found`
)

func (d *Domain) StaffRevenueReport(in *StaffRevenueReportIn) (out StaffRevenueReportOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	booking := rqProperty.NewBookings(d.PropOltp)
	out.RevenueReports = booking.FindRevenueReports(in.YearMonth)

	return
}
