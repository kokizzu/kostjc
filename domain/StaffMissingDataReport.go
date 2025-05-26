package domain

import (
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty/rqProperty"

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
	}
	StaffMissingDataReportOut struct {
		ResponseCommon
		User *rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		Segments    M.SB                               `json:"segments" form:"segments" query:"segments" long:"segments" msg:"segments"`
		MissingData []rqProperty.RoomMissingTenantData `json:"missingData" form:"missingData" query:"missingData" long:"missingData" msg:"missingData"`
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

	room := rqProperty.NewRooms(d.PropOltp)
	out.MissingData = room.FindMissingTenantsData()

	return
}
