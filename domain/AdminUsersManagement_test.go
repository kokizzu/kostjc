package domain

import (
	"testing"

	"kostjc/model/mAuth/wcAuth"
)

func TestApplyAdminUserPassword(t *testing.T) {
	usr := wcAuth.NewUsersMutator(nil)

	if updated := applyAdminUserPassword(usr, "", 123); updated {
		t.Fatal("expected empty password to be ignored")
	}
	if usr.Password != "" {
		t.Fatalf("expected password to stay empty, got %q", usr.Password)
	}
	if usr.PasswordSetAt != 0 {
		t.Fatalf("expected passwordSetAt to stay 0, got %d", usr.PasswordSetAt)
	}

	if updated := applyAdminUserPassword(usr, "new-secret", 456); !updated {
		t.Fatal("expected non-empty password to be applied")
	}
	if usr.Password == "" {
		t.Fatal("expected password to be set")
	}
	if usr.PasswordSetAt != 456 {
		t.Fatalf("expected passwordSetAt=456, got %d", usr.PasswordSetAt)
	}
}
