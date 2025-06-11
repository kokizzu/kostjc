package domain

import (
	"kostjc/model/mProperty"
	"kostjc/model/mProperty/saProperty"
	"kostjc/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminFacilityLogs.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminFacilityLogs.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminFacilityLogs.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminFacilityLogs.go
//go:generate farify doublequote --file AdminFacilityLogs.go

type (
	AdminFacilityLogsIn struct {
		RequestCommon

		Pager    zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		WithMeta bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
	}

	AdminFacilityLogsOut struct {
		ResponseCommon

		Pager zCrud.PagerOut            `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Logs  []saProperty.FacilityLogs `json:"logs" form:"logs" query:"logs" long:"logs" msg:"logs"`
		Meta  *zCrud.Meta               `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
	}
)

const (
	AdminFacilityLogsAction = `admin/facilityLogs`
)

var (
	AdminFacilityLogsMeta = zCrud.Meta{
		Fields: []zCrud.Field{
			{
				Name:      mProperty.CreatedAt,
				Label:     `Created At`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeDateTime,
			},
			{
				Name:      mProperty.ActorId,
				Label:     `Actor`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeCombobox,
			},
			{
				Name:      mProperty.BeforeJSON,
				Label:     `Before (JSON)`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeJSON,
				InputType: zCrud.InputTypeTextArea,
			},
			{
				Name:      mProperty.AfterJSON,
				Label:     `After (JSON)`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeJSON,
				InputType: zCrud.InputTypeTextArea,
			},
		},
	}
)

func (d *Domain) AdminFacilityLogs(in *AdminFacilityLogsIn) (out AdminFacilityLogsOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &AdminFacilityLogsMeta
	}

	// if not set, always override by createdAt descending
	if len(in.Pager.Order) == 0 {
		in.Pager.Order = []string{`-createdAt`}
	}

	r := saProperty.NewFacilityLogs(d.PropOlap)
	out.Logs = r.FindByPagination(&in.Pager, &out.Pager)

	return
}
