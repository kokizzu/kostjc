package domain

import (
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mCafe/rqCafe"
	"kostjc/model/zCrud"

	"github.com/kokizzu/gotro/M"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file StaffSales.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type StaffSales.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type StaffSales.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type StaffSales.go
//go:generate farify doublequote --file StaffSales.go

type (
	StaffSalesIn struct {
		RequestCommon
		Cmd      string        `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Sale     rqCafe.Sales  `json:"sale" form:"sale" query:"sale" long:"sale" msg:"sale"`
		TenantID uint64        `json:"tenantId" form:"tenantId" query:"tenantId" long:"tenantId" msg:"tenantId"`

		StaffData string `json:"staffData" form:"staffData" query:"staffData" long:"staffData" msg:"staffData"`
	}
	StaffSalesOut struct {
		ResponseCommon
		User     *rqAuth.Users  `json:"user" form:"user" query:"user" long:"user" msg:"user"`
		Pager    zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta     *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Sale     rqCafe.Sales   `json:"sale" form:"sale" query:"sale" long:"sale" msg:"sale"`
		Sales    [][]any        `json:"sales" form:"sales" query:"sales" long:"sales" msg:"sales"`
		Segments M.SB           `json:"segments" form:"segments" query:"segments" long:"segments" msg:"segments"`
	}
)

const (
	StaffSalesAction = `staff/sales`

	ErrStaffSalesNotFound = `user not found`
)

func (d *Domain) StaffSales(in *StaffSalesIn) (out StaffSalesOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &AdminSaleMeta
	}

	switch in.Cmd {
	case zCrud.CmdList:
		sale := rqCafe.NewSales(d.PropOltp)
		out.Sales = sale.FindByPagination(
			&AdminSaleMeta,
			&in.Pager,
			&out.Pager,
		)
	}

	return
}
