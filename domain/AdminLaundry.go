package domain

// import (
// 	"kostjc/model/mCafe"
// 	"kostjc/model/mCafe/rqCafe"
// 	"kostjc/model/mCafe/wcCafe"
// 	"kostjc/model/zCrud"
// 	"time"
// )

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminLaundry.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminLaundry.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminLaundry.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminLaundry.go
//go:generate farify doublequote --file AdminLaundry.go

// type (
// 	AdminLaundryIn struct {
// 		RequestCommon
// 		Cmd      string         `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
// 		WithMeta bool           `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
// 		Pager    zCrud.PagerIn  `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
// 		Laundry  rqCafe.Laundry `json:"laundry" form:"laundry" query:"laundry" long:"laundry" msg:"laundry"`
// 		TenantID uint64         `json:"tenantId" form:"tenantId" query:"tenantId" long:"tenantId" msg:"tenantId"`
// 	}
// 	AdminLaundryOut struct {
// 		ResponseCommon
// 		Pager     zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
// 		Meta      *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
// 		Laundry   rqCafe.Laundry `json:"laundry" form:"laundry" query:"laundry" long:"laundry" msg:"laundry"`
// 		Laundries [][]any        `json:"laundries" form:"laundries" query:"laundries" long:"laundries" msg:"laundries"`
// 	}
// )

// const (
// 	AdminLaundryAction = `admin/laundry`

// 	ErrAdminLaundryNotFound         = `Laundry not found`
// 	ErrAdminLaundryTenantIdNotFound = `tenant id not found for this Laundry`
// 	ErrAdminLaundrySaveFailed       = `failed to save Laundry`
// 	ErrAdminLaundryInvalidMenu      = `invalid menu`
// 	ErrAdminLaundryDeleteFailed     = `failed to delete Laundry`
// 	ErrAdminLaundryRestoreFailed    = `failed to restore Laundry`
// 	ErrAdminLaundryAlreadyExists    = `Laundry already exists`
// )

// var AdminLaundryMeta = zCrud.Meta{
// 	Fields: []zCrud.Field{
// 		{
// 			Name:      mCafe.Id,
// 			Label:     `ID`,
// 			DataType:  zCrud.DataTypeInt,
// 			InputType: zCrud.InputTypeHidden,
// 			ReadOnly:  true,
// 		},
// 		{
// 			Name:      mCafe.LaundryCustomer,
// 			Label:     `Customer`,
// 			DataType:  zCrud.DataTypeString,
// 			InputType: zCrud.InputTypeText,
// 			ReadOnly:  false,
// 		},
// 		{
// 			Name:      mCafe.LaundryItems,
// 			Label:     `Items`,
// 			DataType:  zCrud.DataTypeString,
// 			InputType: zCrud.InputTypeText,
// 			ReadOnly:  false,
// 		},
// 		{
// 			Name:      mCafe.LaundryStatus,
// 			Label:     `Status`,
// 			DataType:  zCrud.DataTypeString,
// 			InputType: zCrud.InputTypeComboboxArr,
// 			ReadOnly:  false,
// 		},
// 		{
// 			Name:      mCafe.LaundryAt,
// 			Label:     `Laundry At`,
// 			DataType:  zCrud.DataTypeString,
// 			InputType: zCrud.InputTypeDateTime,
// 			ReadOnly:  false,
// 		},
// 		{
// 			Name:      mCafe.CreatedAt,
// 			Label:     `Created At`,
// 			DataType:  zCrud.DataTypeInt,
// 			InputType: zCrud.InputTypeDateTime,
// 			ReadOnly:  true,
// 		},
// 		{
// 			Name:      mCafe.UpdatedAt,
// 			Label:     `Updated At`,
// 			DataType:  zCrud.DataTypeInt,
// 			InputType: zCrud.InputTypeDateTime,
// 			ReadOnly:  true,
// 		},
// 		{
// 			Name:      mCafe.DeletedAt,
// 			Label:     `Deleted At`,
// 			DataType:  zCrud.DataTypeInt,
// 			InputType: zCrud.InputTypeDateTime,
// 			ReadOnly:  true,
// 		},
// 	},
// }

// func (d *Domain) AdminBorrowedUtensils(in *AdminBorrowedUtensilsIn) (out AdminBorrowedUtensilsOut) {
// 	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
// 	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
// 	if sess == nil {
// 		return
// 	}

// 	out.actor = sess.UserId
// 	out.refId = in.BorrowedUtensil.Id

// 	if in.WithMeta {
// 		out.Meta = &AdminBorrowedUtensilsMeta
// 	}

// 	switch in.Cmd {
// 	case zCrud.CmdForm:
// 	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
// 		borrowedUtensil := wcCafe.NewBorrowedUtensilsMutator(d.PropOltp)
// 		borrowedUtensil.Id = in.BorrowedUtensil.Id
// 		if borrowedUtensil.Id > 0 {
// 			if !borrowedUtensil.FindById() {
// 				out.SetError(400, ErrAdminBorrowedUtensilsNotFound)
// 				return
// 			}

// 			if in.Cmd == zCrud.CmdDelete {
// 				if borrowedUtensil.DeletedAt == 0 {
// 					borrowedUtensil.SetDeletedAt(in.UnixNow())
// 					borrowedUtensil.SetDeletedBy(sess.UserId)
// 				}
// 			}

// 			if in.Cmd == zCrud.CmdRestore {
// 				if borrowedUtensil.DeletedAt > 0 {
// 					borrowedUtensil.SetDeletedAt(0)
// 					borrowedUtensil.SetRestoredBy(sess.UserId)
// 				}
// 			}
// 		}

// 		defer InsertCafeLog(borrowedUtensil.Id, d.borrowedUtensilLogs, out.ResponseCommon, in.TimeNow(), sess.UserId, borrowedUtensil)()

// 		if in.BorrowedUtensil.Customer != `` {
// 			borrowedUtensil.SetCustomer(in.BorrowedUtensil.Customer)
// 		}

// 		if in.BorrowedUtensil.Items != `` {
// 			borrowedUtensil.SetItems(in.BorrowedUtensil.Items)
// 		}

// 		if in.BorrowedUtensil.Qty != 0 {
// 			borrowedUtensil.SetQty(in.BorrowedUtensil.Qty)
// 		}

// 		if in.BorrowedUtensil.Status != `` {
// 			borrowedUtensil.SetStatus(in.BorrowedUtensil.Status)
// 		}

// 		if mCafe.IsValidDate(in.BorrowedUtensil.BorrowedAt, time.DateOnly) {
// 			borrowedUtensil.SetBorrowedAt(in.BorrowedUtensil.BorrowedAt)
// 		}

// 		if borrowedUtensil.Id == 0 {
// 			borrowedUtensil.SetCreatedAt(in.UnixNow())
// 			borrowedUtensil.SetCreatedBy(sess.UserId)
// 		}

// 		borrowedUtensil.SetUpdatedAt(in.UnixNow())
// 		borrowedUtensil.SetUpdatedBy(sess.UserId)

// 		if !borrowedUtensil.DoUpsert() {
// 			out.SetError(500, ErrAdminBorrowedUtensilsSaveFailed)
// 			return
// 		}

// 		if in.Pager.Page == 0 {
// 			break
// 		}

// 		fallthrough
// 	case zCrud.CmdList:
// 		borrowedUtensil := rqCafe.NewBorrowedUtensils(d.PropOltp)
// 		out.BorrowedUtensils = borrowedUtensil.FindByPagination(
// 			&AdminBorrowedUtensilsMeta,
// 			&in.Pager,
// 			&out.Pager,
// 		)
// 	}

// 	return
// }
