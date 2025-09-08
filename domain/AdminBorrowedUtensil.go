package domain

import (
	"kostjc/model/mCafe"
	"kostjc/model/mCafe/rqCafe"
	"kostjc/model/mCafe/wcCafe"
	"kostjc/model/zCrud"
	"time"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminBorrowedUtensils.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminBorrowedUtensils.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminBorrowedUtensils.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminBorrowedUtensils.go
//go:generate farify doublequote --file AdminBorrowedUtensils.go

type (
	AdminBorrowedUtensilsIn struct {
		RequestCommon
		Cmd             string                  `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta        bool                    `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager           zCrud.PagerIn           `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		BorrowedUtensil rqCafe.BorrowedUtensils `json:"borrowedUtensil" form:"borrowedUtensil" query:"borrowedUtensil" long:"borrowedUtensil" msg:"borrowedUtensil"`
		TenantID        uint64                  `json:"tenantId" form:"tenantId" query:"tenantId" long:"tenantId" msg:"tenantId"`
	}
	AdminBorrowedUtensilsOut struct {
		ResponseCommon
		Pager            zCrud.PagerOut          `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta             *zCrud.Meta             `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		BorrowedUtensil  rqCafe.BorrowedUtensils `json:"borrowedUtensil" form:"borrowedUtensil" query:"borrowedUtensil" long:"borrowedUtensil" msg:"borrowedUtensil"`
		BorrowedUtensils [][]any                 `json:"borrowedUtensils" form:"borrowedUtensils" query:"borrowedUtensils" long:"borrowedUtensils" msg:"borrowedUtensils"`
	}
)

const (
	AdminBorrowedUtensilsAction = `admin/borrowedUtensils`

	ErrAdminBorrowedUtensilsNotFound         = `Borrowed Utensils not found`
	ErrAdminBorrowedUtensilsTenantIdNotFound = `tenant id not found for this Borrowed Utensils`
	ErrAdminBorrowedUtensilsSaveFailed       = `failed to save Borrowed Utensils`
	ErrAdminBorrowedUtensilsInvalidMenu      = `invalid menu`
	ErrAdminBorrowedUtensilsDeleteFailed     = `failed to delete Borrowed Utensils`
	ErrAdminBorrowedUtensilsRestoreFailed    = `failed to restore Borrowed Utensils`
	ErrAdminBorrowedUtensilsAlreadyExists    = `Borrowed Utensils already exists`
)

var AdminBorrowedUtensilsMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mCafe.Id,
			Label:     `ID`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:      mCafe.BorrowedCustomer,
			Label:     `Customer`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
			ReadOnly:  false,
		},
		{
			Name:      mCafe.BorrowedItems,
			Label:     `Items`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
			ReadOnly:  false,
		},
		{
			Name:      mCafe.BorrowedQty,
			Label:     `Qty`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeNumber,
			ReadOnly:  false,
		},
		{
			Name:      mCafe.BorrowedStatus,
			Label:     `Status`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeComboboxArr,
			ReadOnly:  false,
		},
		{
			Name:      mCafe.BorrowedAt,
			Label:     `Borrowed At`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  false,
		},
		{
			Name:      mCafe.CreatedAt,
			Label:     `Created At`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
		{
			Name:      mCafe.UpdatedAt,
			Label:     `Updated At`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
		{
			Name:      mCafe.DeletedAt,
			Label:     `Deleted At`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
	},
}

func (d *Domain) AdminBorrowedUtensils(in *AdminBorrowedUtensilsIn) (out AdminBorrowedUtensilsOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.actor = sess.UserId
	out.refId = in.BorrowedUtensil.Id

	if in.WithMeta {
		out.Meta = &AdminBorrowedUtensilsMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		borrowedUtensil := wcCafe.NewBorrowedUtensilsMutator(d.PropOltp)
		borrowedUtensil.Id = in.BorrowedUtensil.Id
		if borrowedUtensil.Id > 0 {
			if !borrowedUtensil.FindById() {
				out.SetError(400, ErrAdminBorrowedUtensilsNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if borrowedUtensil.DeletedAt == 0 {
					borrowedUtensil.SetDeletedAt(in.UnixNow())
					borrowedUtensil.SetDeletedBy(sess.UserId)
				}
			}

			if in.Cmd == zCrud.CmdRestore {
				if borrowedUtensil.DeletedAt > 0 {
					borrowedUtensil.SetDeletedAt(0)
					borrowedUtensil.SetRestoredBy(sess.UserId)
				}
			}
		}

		defer InsertCafeLog(borrowedUtensil.Id, d.borrowedUtensilLogs, out.ResponseCommon, in.TimeNow(), sess.UserId, borrowedUtensil)()

		if in.BorrowedUtensil.Customer != `` {
			borrowedUtensil.SetCustomer(in.BorrowedUtensil.Customer)
		}

		if in.BorrowedUtensil.Items != `` {
			borrowedUtensil.SetItems(in.BorrowedUtensil.Items)
		}

		if in.BorrowedUtensil.Qty != 0 {
			borrowedUtensil.SetQty(in.BorrowedUtensil.Qty)
		}

		if in.BorrowedUtensil.Status != `` {
			borrowedUtensil.SetStatus(in.BorrowedUtensil.Status)
		}

		if mCafe.IsValidDate(in.BorrowedUtensil.BorrowedAt, time.DateOnly) {
			borrowedUtensil.SetBorrowedAt(in.BorrowedUtensil.BorrowedAt)
		}

		if borrowedUtensil.Id == 0 {
			borrowedUtensil.SetCreatedAt(in.UnixNow())
			borrowedUtensil.SetCreatedBy(sess.UserId)
		}

		borrowedUtensil.SetUpdatedAt(in.UnixNow())
		borrowedUtensil.SetUpdatedBy(sess.UserId)

		if !borrowedUtensil.DoUpsert() {
			out.SetError(500, ErrAdminBorrowedUtensilsSaveFailed)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		borrowedUtensil := rqCafe.NewBorrowedUtensils(d.PropOltp)
		out.BorrowedUtensils = borrowedUtensil.FindByPagination(
			&AdminBorrowedUtensilsMeta,
			&in.Pager,
			&out.Pager,
		)
	}

	return
}
