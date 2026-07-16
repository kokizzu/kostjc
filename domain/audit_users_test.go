package domain

import (
	"reflect"
	"testing"

	"kostjc/model/mProperty"
	"kostjc/model/zCrud"
)

func TestAuditUserIds(t *testing.T) {
	meta := &zCrud.Meta{
		Fields: []zCrud.Field{
			{Name: mProperty.Id},
			{Name: mProperty.CreatedBy},
			{Name: `notAudit`},
			{Name: mProperty.UpdatedBy},
			{Name: mProperty.DeletedBy},
		},
	}
	rows := [][]any{
		{uint64(1), uint64(7), `x`, uint64(9), uint64(0)},
		{uint64(2), uint64(7), `x`, uint64(0), uint64(11)},
		{uint64(3), uint64(0), `x`, uint64(13)},
		{uint64(4)},
	}

	got := auditUserIds(rows, meta)
	want := []uint64{7, 9, 11, 13}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("auditUserIds() = %#v, want %#v", got, want)
	}
}

func TestAuditUserIdsMissingMeta(t *testing.T) {
	if got := auditUserIds([][]any{{uint64(1), uint64(7)}}, nil); len(got) != 0 {
		t.Fatalf("auditUserIds() with nil meta = %#v, want empty", got)
	}

	meta := &zCrud.Meta{Fields: []zCrud.Field{{Name: mProperty.Id}}}
	if got := auditUserIds([][]any{{uint64(1), uint64(7)}}, meta); len(got) != 0 {
		t.Fatalf("auditUserIds() with no audit fields = %#v, want empty", got)
	}
}

func TestAuditUserMetaFields(t *testing.T) {
	fields := auditUserMetaFields()
	wantNames := []string{mProperty.CreatedBy, mProperty.UpdatedBy, mProperty.DeletedBy}
	if len(fields) != len(wantNames) {
		t.Fatalf("len(auditUserMetaFields()) = %d, want %d", len(fields), len(wantNames))
	}

	for i, field := range fields {
		if field.Name != wantNames[i] {
			t.Fatalf("audit field %d name = %q, want %q", i, field.Name, wantNames[i])
		}
		if !field.ReadOnly {
			t.Fatalf("audit field %q should be read-only", field.Name)
		}
		if field.InputType != zCrud.InputTypeCombobox {
			t.Fatalf("audit field %q input type = %q, want %q", field.Name, field.InputType, zCrud.InputTypeCombobox)
		}
	}
}
