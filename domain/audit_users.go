package domain

import (
	"kostjc/model/mAuth/rqAuth"
	"kostjc/model/mProperty"
	"kostjc/model/zCrud"

	"github.com/kokizzu/gotro/X"
)

func (d *Domain) auditUserChoices(rows [][]any, meta *zCrud.Meta) map[uint64]string {
	ids := make([]uint64, 0)
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
				ids = append(ids, id)
			}
		}
	}

	user := rqAuth.NewUsers(d.AuthOltp)
	return user.FindUserChoicesByIds(ids)
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
