package domain

import (
	"kostjc/model/mAuth/rqAuth"

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

		StaffData string `json:"staffData" form:"staffData" query:"staffData" long:"staffData" msg:"staffData"`
	}
	StaffSalesOut struct {
		ResponseCommon
		User *rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`

		Segments M.SB `json:"segments" form:"segments" query:"segments" long:"segments" msg:"segments"`
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

	return
}
