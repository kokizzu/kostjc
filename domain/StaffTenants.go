package domain

import (
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mAuth/wcAuth"
	"kostjc/model/zCrud"

	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
)

type (
	StaffTenantsIn struct {
		RequestCommon
		Cmd      string         `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool           `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn  `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Tenant   rqAuth.Tenants `json:"tenant" form:"tenant" query:"tenant" long:"tenant" msg:"tenant"`
	}
	StaffTenantsOut struct {
		ResponseCommon
		Pager   zCrud.PagerOut  `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta    *zCrud.Meta     `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Tenant  *rqAuth.Tenants `json:"tenant" form:"tenant" query:"tenant" long:"tenant" msg:"tenant"`
		Tenants [][]any         `json:"tenants" form:"tenants" query:"tenants" long:"tenants" msg:"tenants"`
	}
)

const StaffTenantsAction = `staff/tenants`

func (d *Domain) handleTenants(
	in *RequestCommon,
	cmd string,
	withMeta bool,
	pager *zCrud.PagerIn,
	tenantIn *rqAuth.Tenants,
	out *ResponseCommon,
	sess *Session,
) (meta *zCrud.Meta, tenant *rqAuth.Tenants, tenants [][]any, pagerOut zCrud.PagerOut) {
	out.actor = sess.UserId
	out.refId = tenantIn.Id

	if withMeta {
		meta = &AdminTenantsMeta
	}

	switch cmd {
	case zCrud.CmdForm:
		tnt := rqAuth.NewTenants(d.AuthOltp)
		tnt.Id = tenantIn.Id
		if tnt.Id > 0 {
			if !tnt.FindById() {
				out.SetError(400, ErrAdminTenantsNotFound)
				return
			}
		}
		tenant = tnt
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore, zCrud.CmdToggleWaAdded, zCrud.CmdToggleTeleAdded:
		tenantMut := wcAuth.NewTenantsMutator(d.AuthOltp)
		tenantMut.Id = tenantIn.Id
		if tenantMut.Id > 0 {
			if !tenantMut.FindById() {
				out.SetError(400, ErrAdminTenantsNotFound)
				return
			}

			if cmd == zCrud.CmdDelete && tenantMut.DeletedAt == 0 {
				tenantMut.SetDeletedAt(in.UnixNow())
			}

			if cmd == zCrud.CmdRestore && tenantMut.DeletedAt > 0 {
				tenantMut.SetDeletedAt(0)
				tenantMut.SetRestoredBy(sess.UserId)
			}
			if cmd == zCrud.CmdToggleWaAdded {
				tenantMut.SetWaAddedAt(tenantIn.WaAddedAt)
				tenantMut.SetUpdatedAt(in.UnixNow())
				tenantMut.SetUpdatedBy(sess.UserId)

				if !tenantMut.DoUpsert() {
					out.SetError(500, ErrAdminTenantsSaveFailed)
					return
				}
			}
			if cmd == zCrud.CmdToggleTeleAdded {
				tenantMut.SetTeleAddedAt(tenantIn.TeleAddedAt)
				tenantMut.SetUpdatedAt(in.UnixNow())
				tenantMut.SetUpdatedBy(sess.UserId)

				if !tenantMut.DoUpsert() {
					out.SetError(500, ErrAdminTenantsSaveFailed)
					return
				}
			}
		}

		defer InsertPropertyLog(tenantMut.Id, d.tenantLogs, *out, in.TimeNow(), sess.UserId, tenantMut)()

		tenantIn.KtpPlaceBirth = S.ToUpper(tenantIn.KtpPlaceBirth)
		tenantIn.KtpRegion = S.ToUpper(tenantIn.KtpRegion)
		tenantIn.KtpOccupation = S.ToTitle(tenantIn.KtpOccupation)
		tenantIn.TenantName = S.ToTitle(tenantIn.TenantName)
		tenantIn.KtpName = S.ToTitle(tenantIn.KtpName)
		tenantIn.KtpGender = S.ToTitle(tenantIn.KtpGender)
		tenantIn.KtpCitizenship = S.ToUpper(tenantIn.KtpCitizenship)
		tenantIn.KtpAddress = S.ToUpper(tenantIn.KtpAddress)
		tenantIn.KtpKelurahanDesa = S.ToTitle(tenantIn.KtpKelurahanDesa)
		tenantIn.KtpKecamatan = S.ToTitle(tenantIn.KtpKecamatan)

		tenantMut.SetAll(*tenantIn, M.SB{}, M.SB{})

		if tenantMut.Id == 0 {
			if tenantMut.FindByKtpNumber() {
				out.SetError(400, ErrAdminTenantsAlreadyExists)
				return
			}
			tenantMut.SetCreatedAt(in.UnixNow())
			tenantMut.SetCreatedBy(sess.UserId)
		}

		tenantMut.SetUpdatedAt(in.UnixNow())
		tenantMut.SetUpdatedBy(sess.UserId)

		if !tenantMut.DoUpsert() {
			out.SetError(500, ErrAdminTenantsSaveFailed)
			return
		}

		if pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		tenantQry := rqAuth.NewTenants(d.AuthOltp)
		tenants = tenantQry.FindByPagination(
			&AdminTenantsMeta,
			pager,
			&pagerOut,
		)
	}

	return
}

func (d *Domain) StaffTenants(in *StaffTenantsIn) (out StaffTenantsOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustBelowStaff(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.Meta, out.Tenant, out.Tenants, out.Pager = d.handleTenants(
		&in.RequestCommon,
		in.Cmd,
		in.WithMeta,
		&in.Pager,
		&in.Tenant,
		&out.ResponseCommon,
		sess,
	)
	return
}
