package domain

import (
	"kostjc/model/mProperty/rqProperty"
	"kostjc/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminSettingFixInconsistencies.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminSettingFixInconsistencies.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminSettingFixInconsistencies.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminSettingFixInconsistencies.go
//go:generate farify doublequote --file AdminSettingFixInconsistencies.go

type (
	AdminSettingFixInconsistenciesIn struct {
		RequestCommon
		Cmd string `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
	}
	AdminSettingFixInconsistenciesOut struct {
		ResponseCommon
		RoomIncosistencies []rqProperty.RoomBookingInconsistency `json:"roomIncosistencies" form:"roomIncosistencies" query:"roomIncosistencies" long:"roomIncosistencies" msg:"roomIncosistencies"`
	}
)

const (
	AdminSettingFixInconsistenciesAction = `admin/settingFixInconsistencies`
)

func (d *Domain) AdminSettingFixInconsistencies(in *AdminSettingFixInconsistenciesIn) (out AdminSettingFixInconsistenciesOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	switch in.Cmd {
	case zCrud.CmdUpsert:
		fallthrough
	case zCrud.CmdList:
		room := rqProperty.NewRooms(d.PropOltp)
		out.RoomIncosistencies = room.FindRoomBookingInconsistencies()
	}

	return
}
