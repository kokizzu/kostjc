package domain

import (
	"kostjc/model/mAuth"
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mAuth/wcAuth"
	"kostjc/model/zCrud"

	"github.com/kokizzu/gotro/S"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminUsersManagement.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminUsersManagement.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminUsersManagement.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminUsersManagement.go
//go:generate farify doublequote --file AdminUsersManagement.go

type (
	AdminUsersManagementIn struct {
		RequestCommon
		Cmd      string        `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		User     rqAuth.Users  `json:"user" form:"user" query:"user" long:"user" msg:"user"`
	}
	AdminUsersManagementOut struct {
		ResponseCommon
		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta  *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		User  rqAuth.Users   `json:"user" form:"user" query:"user" long:"user" msg:"user"`
		Users [][]any        `json:"users" form:"users" query:"users" long:"users" msg:"users"`
	}
)

const (
	AdminUsersManagementAction = `admin/usersManagement`

	ErrAdminUsersManagementNotFound      = `user not found`
	ErrAdminUsersManagementSaveFailed    = `failed to save user`
	ErrAdminUsersManagementDeleteFailed  = `failed to delete user`
	ErrAdminUsersManagementRestoreFailed = `failed to restore user`
)

var AdminUsersManagementMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mAuth.Id,
			Label:     `ID`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:        mAuth.Email,
			Label:       `Email`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			ReadOnly:    false,
			Description: `komangzW3l4@example.com`,
		},
		{
			Name:        mAuth.UserName,
			Label:       `Username`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			ReadOnly:    false,
			Description: `komangz`,
		},
		{
			Name:        mAuth.FullName,
			Label:       `Full Name`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeText,
			ReadOnly:    false,
			Description: `I Komang Putra Sanjaya`,
		},
		{
			Name:        mAuth.Role,
			Label:       `Role`,
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeComboboxArr,
			ReadOnly:    false,
			Description: `Admin/User`,
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

func (d *Domain) AdminUsersManagement(in *AdminUsersManagementIn) (out AdminUsersManagementOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	out.actor = sess.UserId
	out.refId = in.User.Id

	if in.WithMeta {
		out.Meta = &AdminUsersManagementMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		usr := wcAuth.NewUsersMutator(d.AuthOltp)
		usr.Id = in.User.Id
		if usr.Id > 0 {
			if !usr.FindById() {
				out.SetError(400, ErrAdminUsersManagementNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if usr.DeletedAt == 0 {
					usr.SetDeletedAt(in.UnixNow())
				}
			}

			if in.Cmd == zCrud.CmdRestore {
				if usr.DeletedAt > 0 {
					usr.SetDeletedAt(0)
				}
			}
		}

		if in.User.UserName != `` {
			usr.SetUserName(S.ValidateIdent(in.User.UserName))
		}

		if in.User.Email != `` {
			usr.SetEmail(S.ValidateEmail(in.User.Email))
		}

		if in.User.FullName != `` {
			usr.SetFullName(in.User.FullName)
		}

		if mAuth.IsValidRole(in.User.Role) {
			usr.SetRole(in.User.Role)
		}

		if in.User.Id > 0 && in.User.Password != `` {
			if len(in.User.Password) >= minPassLength {
				usr.SetSecretCode(``)
				usr.SetSecretCodeAt(0)
				usr.SetEncryptedPassword(in.User.Password, in.UnixNow())
			}
		}

		if usr.Id == 0 {
			usr.SetCreatedAt(in.UnixNow())
			usr.SetCreatedBy(sess.UserId)
		}

		usr.SetUpdatedAt(in.UnixNow())
		usr.SetUpdatedBy(sess.UserId)

		if !usr.DoUpsert() {
			out.SetError(500, ErrAdminUsersManagementSaveFailed)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		usr := rqAuth.NewUsers(d.AuthOltp)
		out.Users = usr.FindByPagination(
			&AdminUsersManagementMeta,
			&in.Pager,
			&out.Pager,
		)
	}

	return
}
