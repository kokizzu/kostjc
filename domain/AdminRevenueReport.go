package domain

import (
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty/rqProperty"

	"github.com/kokizzu/gotro/M"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminRevenueReport.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminRevenueReport.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminRevenueReport.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminRevenueReport.go
//go:generate farify doublequote --file AdminRevenueReport.go

type (
	AdminRevenueReportIn struct {
		RequestCommon
		YearMonth string `json:"yearMonth" form:"yearMonth" query:"yearMonth" long:"yearMonth" msg:"yearMonth"`
	}
	AdminRevenueReportOut struct {
		ResponseCommon
		User *rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		Segments            M.SB                            `json:"segments" form:"segments" query:"segments" long:"segments" msg:"segments"`
		RevenueReports      []rqProperty.RevenueReport      `json:"revenueReports" form:"revenueReports" query:"revenueReports" long:"revenueReports" msg:"revenueReports"`
		ChartRevenueReports []rqProperty.ChartRevenueReport `json:"chartRevenueReports" form:"chartRevenueReports" query:"chartRevenueReports" long:"chartRevenueReports" msg:"chartRevenueReports"`
	}
)

const (
	AdminRevenueReportAction = `admin/revenueReport`

	ErrAdminRevenueReportNotFound = `user not found`
)

func (d *Domain) AdminRevenueReport(in *AdminRevenueReportIn) (out AdminRevenueReportOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	booking := rqProperty.NewBookings(d.PropOltp)
	out.RevenueReports = booking.FindRevenueReports(in.YearMonth)

	py := rqProperty.NewPayments(d.PropOltp)
	out.ChartRevenueReports = py.GetChartRevenueReports(in.YearMonth)

	return
}
