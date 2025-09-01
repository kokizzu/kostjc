package domain

import (
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty/rqProperty"

	"github.com/kokizzu/gotro/M"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file StaffPricePerDayReport.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type StaffPricePerDayReport.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type StaffPricePerDayReport.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type StaffPricePerDayReport.go
//go:generate farify doublequote --file StaffPricePerDayReport.go

type (
	StaffPricePerDayReportIn struct {
		RequestCommon
		YearMonth string `json:"yearMonth" form:"yearMonth" query:"yearMonth" long:"yearMonth" msg:"yearMonth"`
	}
	StaffPricePerDayReportOut struct {
		ResponseCommon
		User        *rqAuth.Users                  `json:"user" form:"user" query:"user" long:"user" msg:"user"`
		PricePerDay []rqProperty.PricePerDayReport `json:"pricePerDay" form:"pricePerDay" query:"pricePerDay" long:"pricePerDay" msg:"pricePerDay"`

		Segments M.SB `json:"segments" form:"segments" query:"segments" long:"segments" msg:"segments"`
	}
)

const (
	StaffPricePerDayReportAction = `staff/pricePerDayReport`

	ErrStaffPricePerDayReportNotFound = `user not found`
)

func (d *Domain) StaffPricePerDayReport(in *StaffPricePerDayReportIn) (out StaffPricePerDayReportOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	booking := rqProperty.NewBookings(d.PropOltp)
	out.PricePerDay = booking.FindPricePerDayReport(in.YearMonth)

	return
}
