package domain

import "kostjc/model"

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file AdminSettingss.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type AdminSettingss.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type AdminSettingss.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type AdminSettingss.go
//go:generate farify doublequote --file AdminSettingss.go

type (
	AdminSettingsIn struct {
		RequestCommon
		Cmd string `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
	}
	AdminSettingsOut struct {
		ResponseCommon
		Message string `json:"message" form:"message" query:"message" long:"message" msg:"message"`
	}
)

const (
	AdminSettingsAction = `admin/settings`
	AdminSettingsCmdBackupDatabase = `backupDatabase`
)

func (d *Domain) AdminSettings(in *AdminSettingsIn) (out AdminSettingsOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	switch in.Cmd {
	case AdminSettingsCmdBackupDatabase:
		model.BackupDatabase(d.PropOltp)
		out.Message = `backup completed`
	}

	return
}
