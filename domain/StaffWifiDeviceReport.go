package domain

import (
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty/rqProperty"

	"github.com/kokizzu/gotro/M"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file StaffWifiDeviceReport.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type StaffWifiDeviceReport.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type StaffWifiDeviceReport.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type StaffWifiDeviceReport.go
//go:generate farify doublequote --file StaffWifiDeviceReport.go

type (
	StaffWifiDeviceReportIn struct {
		RequestCommon
		YearMonth string `json:"yearMonth" form:"yearMonth" query:"yearMonth" long:"yearMonth" msg:"yearMonth"`
	}
	StaffWifiDeviceReportOut struct {
		ResponseCommon
		User *rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		Segments          M.SB                          `json:"segments" form:"segments" query:"segments" long:"segments" msg:"segments"`
		WifiDeviceReports []rqProperty.WifiDeviceReport `json:"wifiDeviceReports" form:"wifiDeviceReports" query:"wifiDeviceReports" long:"wifiDeviceReports" msg:"wifiDeviceReports"`
	}
)

const (
	StaffWifiDeviceReportAction = `staff/wifiDeviceReport`

	ErrStaffWifiDeviceReportNotFound = `user not found`
)

func (d *Domain) StaffWifiDeviceReport(in *StaffWifiDeviceReportIn) (out StaffWifiDeviceReportOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	wifiDevice := rqProperty.NewWifiDevices(d.PropOltp)
	out.WifiDeviceReports = wifiDevice.FindWifiDeviceReports(in.YearMonth)

	return
}
