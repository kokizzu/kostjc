package domain

import (
	"kostjc/model/mAuth"
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mAuth/wcAuth"
	"kostjc/model/zCrud"

	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminTenants.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminTenants.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminTenants.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminTenants.go
//go:generate farify doublequote --file AdminTenants.go

type (
	AdminTenantsIn struct {
		RequestCommon
		Cmd      string         `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool           `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn  `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Tenant   rqAuth.Tenants `json:"tenant" form:"tenant" query:"tenant" long:"tenant" msg:"tenant"`
	}
	AdminTenantsOut struct {
		ResponseCommon
		Pager   zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta    *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Tenant  rqAuth.Tenants `json:"tenant" form:"tenant" query:"tenant" long:"tenant" msg:"tenant"`
		Tenants [][]any        `json:"tenants" form:"tenants" query:"tenants" long:"tenants" msg:"tenants"`
	}
)

const (
	AdminTenantsAction = `admin/tenants`

	ErrAdminTenantsNotFound          = `building not found`
	ErrAdminTenantsSaveFailed        = `failed to save building`
	ErrAdminTenantsDeleteFailed      = `failed to delete building`
	ErrAdminTenantsRestoreFailed     = `failed to restore building`
	ErrAdminTenantsLocationNotFound  = `location not found`
	ErrAdminTenantsInvalidFacilities = `invalid facilities`
)

var AdminTenantsMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mAuth.Id,
			Label:     `ID`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:        mAuth.TenantName,
			Label:       `Nama Penyewa`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			Description: `John Doe`,
			ReadOnly:    false,
		},
		{
			Name:        mAuth.KtpRegion,
			Label:       `Regional`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			Description: `PROVINSI BALI KABUPATEN BADUNG`,
			ReadOnly:    false,
		},
		{
			Name:        mAuth.KtpNumber,
			Label:       `Nomor KTP/NIK`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			Description: `1234567890123456`,
			ReadOnly:    false,
		},
		{
			Name:        mAuth.KtpName,
			Label:       `Nama Lengkap (sesuai KTP)`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			Description: `I Komang John Doe`,
			ReadOnly:    false,
		},
		{
			Name:        mAuth.KtpPlaceBirth,
			Label:       `Tempat Lahir`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			Description: `Denpasar`,
			ReadOnly:    false,
		},
		{
			Name:      mAuth.KtpDateBirth,
			Label:     `Tanggal Lahir (min. 17 tahun)`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  false,
		},
		{
			Name:        mAuth.KtpGender,
			Label:       `Jenis Kelamin`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeComboboxArr,
			ReadOnly:    false,
			Description: `Laki - Laki / Perempuan`,
		},
		{
			Name:        mAuth.KtpAddress,
			Label:       `Alamat`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			Description: `Jl. Jalan Raya No. 1`,
			ReadOnly:    false,
		},
		{
			Name:        mAuth.KtpRtRw,
			Label:       `RT/RW`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			ReadOnly:    false,
			Description: `001/004`,
		},
		{
			Name:      mAuth.KtpKelurahanDesa,
			Label:     `Kelurahan/Desa`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
			ReadOnly:  false,
		},
		{
			Name:      mAuth.KtpKecamatan,
			Label:     `Kecamatan`,
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
			ReadOnly:  false,
		},
		{
			Name:        mAuth.KtpReligion,
			Label:       `Agama`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			ReadOnly:    false,
			Description: `Hindu`,
		},
		{
			Name:        mAuth.KtpMaritalStatus,
			Label:       `Status Perkawinan`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeComboboxArr,
			Description: `Kawin / Belum Kawin`,
			ReadOnly:    false,
		},
		{
			Name:        mAuth.KtpCitizenship,
			Label:       `Kewarganegaraan`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeComboboxArr,
			Description: `WNI / WNA`,
			ReadOnly:    false,
		},
		{
			Name:        mAuth.TelegramUsername,
			Label:       `Username Telegram`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			Description: `@johndoe`,
			ReadOnly:    false,
		},
		{
			Name:        mAuth.WhatsappNumber,
			Label:       `Nomor Whatsapp`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			Description: `+6281234567890`,
			ReadOnly:    false,
		},
		{
			Name:      mAuth.CreatedAt,
			Label:     `Created At`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
		{
			Name:      mAuth.UpdatedAt,
			Label:     `Updated At`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
		{
			Name:      mAuth.DeletedAt,
			Label:     `Deleted At`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
	},
}

func (d *Domain) AdminTenants(in *AdminTenantsIn) (out AdminTenantsOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.actor = sess.UserId
	out.refId = in.Tenant.Id

	if in.WithMeta {
		out.Meta = &AdminTenantsMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
		tenant.Id = in.Tenant.Id
		if tenant.Id > 0 {
			if !tenant.FindById() {
				out.SetError(400, ErrAdminTenantsNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if tenant.DeletedAt == 0 {
					tenant.SetDeletedAt(in.UnixNow())
				}
			}

			if in.Cmd == zCrud.CmdRestore {
				if tenant.DeletedAt > 0 {
					tenant.SetDeletedAt(0)
					tenant.SetRestoredBy(sess.UserId)
				}
			}
		}

		if in.Tenant.KtpRegion != "" {
			in.Tenant.KtpRegion = S.ToUpper(in.Tenant.KtpRegion)
		}

		tenant.SetAll(in.Tenant, M.SB{}, M.SB{})

		if tenant.Id == 0 {
			tenant.SetCreatedAt(in.UnixNow())
			tenant.SetCreatedBy(sess.UserId)
		}

		tenant.SetUpdatedAt(in.UnixNow())
		tenant.SetUpdatedBy(sess.UserId)

		if !tenant.DoUpsert() {
			out.SetError(500, ErrAdminTenantsSaveFailed)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		tenant := rqAuth.NewTenants(d.AuthOltp)
		out.Tenants = tenant.FindByPagination(
			&AdminTenantsMeta,
			&in.Pager,
			&out.Pager,
		)
	}

	return
}
