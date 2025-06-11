package domain

import (
	"kostjc/model/mProperty"
	"kostjc/model/mProperty/saProperty"
	"kostjc/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminLocationLogs.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminLocationLogs.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminLocationLogs.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminLocationLogs.go
//go:generate farify doublequote --file AdminLocationLogs.go

type (
	AdminLocationLogsIn struct {
		RequestCommon

		Pager    zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		WithMeta bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
	}

	AdminLocationLogsOut struct {
		ResponseCommon

		Pager zCrud.PagerOut            `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Logs  []saProperty.LocationLogs `json:"logs" form:"logs" query:"logs" long:"logs" msg:"logs"`
		Meta  *zCrud.Meta               `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
	}
)

const (
	AdminLocationLogsAction = `admin/locationLogs`
)

var (
	AdminLocationLogsMeta = zCrud.Meta{
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

func (d *Domain) AdminLocationLogs(in *AdminLocationLogsIn) (out AdminLocationLogsOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &AdminLocationLogsMeta
	}

	// if not set, always override by createdAt descending
	if len(in.Pager.Order) == 0 {
		in.Pager.Order = []string{`-createdAt`}
	}

	r := saProperty.NewLocationLogs(d.PropOlap)
	out.Logs = r.FindByPagination(&in.Pager, &out.Pager)

	return
}
