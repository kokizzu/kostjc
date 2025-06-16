package domain

import (
	"kostjc/model/mCafe"
	"kostjc/model/mCafe/saCafe"
	"kostjc/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminSaleLogs.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminSaleLogs.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminSaleLogs.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminSaleLogs.go
//go:generate farify doublequote --file AdminSaleLogs.go

type (
	AdminSaleLogsIn struct {
		RequestCommon

		Pager    zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		WithMeta bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
	}

	AdminSaleLogsOut struct {
		ResponseCommon

		Pager zCrud.PagerOut    `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Logs  []saCafe.SaleLogs `json:"logs" form:"logs" query:"logs" long:"logs" msg:"logs"`
		Meta  *zCrud.Meta       `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
	}
)

const (
	AdminSaleLogsAction = `admin/saleLogs`
)

var (
	AdminSaleLogsMeta = zCrud.Meta{
		Fields: []zCrud.Field{
			{
				Name:      mCafe.CreatedAt,
				Label:     `Datetime`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeDateTime,
			},
			{
				Name:      mCafe.ActorId,
				Label:     `Actor`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeCombobox,
			},
			{
				Name:      mCafe.BeforeJSON,
				Label:     `Before (JSON)`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeJSON,
				InputType: zCrud.InputTypeTextArea,
			},
			{
				Name:      mCafe.AfterJSON,
				Label:     `After (JSON)`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeJSON,
				InputType: zCrud.InputTypeTextArea,
			},
		},
	}
)

func (d *Domain) AdminSaleLogs(in *AdminSaleLogsIn) (out AdminSaleLogsOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &AdminSaleLogsMeta
	}

	// if not set, always override by createdAt descending
	if len(in.Pager.Order) == 0 {
		in.Pager.Order = []string{`-createdAt`}
	}

	r := saCafe.NewSaleLogs(d.PropOlap)
	out.Logs = r.FindByPagination(&in.Pager, &out.Pager)

	return
}
