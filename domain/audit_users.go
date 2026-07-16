package domain

import (
	"sort"

	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty"
	"kostjc/model/zCrud"

	"github.com/kokizzu/gotro/X"
)

func auditUserMetaFields() []zCrud.Field {
	return []zCrud.Field{
		{
			Name:      mProperty.CreatedBy,
			Label:     `Created By`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
			ReadOnly:  true,
		},
		{
			Name:      mProperty.UpdatedBy,
			Label:     `Updated By`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
			ReadOnly:  true,
		},
		{
			Name:      mProperty.DeletedBy,
			Label:     `Deleted By`,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
			ReadOnly:  true,
		},
	}
}

func (d *Domain) auditUserChoices(rows [][]any, meta *zCrud.Meta) map[uint64]string {
	ids := auditUserIds(rows, meta)

	user := rqAuth.NewUsers(d.AuthOltp)
	return user.FindUserChoicesByIds(ids)
}

func auditUserIds(rows [][]any, meta *zCrud.Meta) []uint64 {
	seen := map[uint64]struct{}{}
	for _, fieldName := range []string{
		mProperty.CreatedBy,
		mProperty.UpdatedBy,
		mProperty.DeletedBy,
	} {
		idx := metaFieldIndex(meta, fieldName)
		if idx < 0 {
			continue
		}
		for _, row := range rows {
			if idx >= len(row) {
				continue
			}
			id := X.ToU(row[idx])
			if id > 0 {
				seen[id] = struct{}{}
			}
		}
	}

	ids := make([]uint64, 0, len(seen))
	for id := range seen {
		ids = append(ids, id)
	}
	sort.Slice(ids, func(i, j int) bool {
		return ids[i] < ids[j]
	})

	return ids
}

func metaFieldIndex(meta *zCrud.Meta, fieldName string) int {
	if meta == nil {
		return -1
	}
	for idx, field := range meta.Fields {
		if field.Name == fieldName {
			return idx
		}
	}
	return -1
}
