package auth

import (
	"context"
	"testing"
)

func TestGetUserNameOK(t *testing.T) {
	want := "username"

	token := map[string]interface{}{
		"name": want,
	}

	ctx := contextWithToken(token)

	got := GetUserName(ctx)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGetUserNameEmptyToken(t *testing.T) {
	want := ""

	token := map[string]interface{}{}

	ctx := contextWithToken(token)

	got := GetUserName(ctx)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGetUserNameNilToken(t *testing.T) {
	want := ""

	ctx := contextWithToken(nil)

	got := GetUserName(ctx)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGetRolesOK(t *testing.T) {
	want := []string{"role1", "role2"}

	token := map[string]interface{}{
		"resource_access": map[string]interface{}{
			"todo-api": map[string]interface{}{
				"roles": want,
			},
		},
	}

	ctx := contextWithToken(token)

	got := GetRoles(ctx)

	if len(got) != len(want) {
		t.Errorf("got %d roles, want %d", len(got), len(want))
		return
	}

	for i := range got {
		if got[i] != want[i] {
			t.Errorf("got %q, want %q", got[i], want[i])
		}
	}
}

func TestGetRolesEmptyToken(t *testing.T) {
	want := 0

	token := map[string]interface{}{}

	ctx := contextWithToken(token)

	got := GetRoles(ctx)

	if len(got) != want {
		t.Errorf("got %d roles, want %d", len(got), want)
	}
}

func TestGetRolesNilToken(t *testing.T) {
	want := 0

	ctx := contextWithToken(nil)

	got := GetRoles(ctx)

	if len(got) != want {
		t.Errorf("got %d roles, want %d", len(got), want)
	}
}

func TestIsInRoleTrue(t *testing.T) {
	role := "role1"

	token := map[string]interface{}{
		"resource_access": map[string]interface{}{
			"todo-api": map[string]interface{}{
				"roles": []string{role},
			},
		},
	}

	ctx := contextWithToken(token)

	want := true
	got := IsInRole(ctx, role)

	if got != want {
		t.Errorf("got %t, want %t", got, want)
		return
	}
}

func TestIsInRoleFalse(t *testing.T) {
	role := "role1"

	token := map[string]interface{}{
		"resource_access": map[string]interface{}{
			"todo-api": map[string]interface{}{
				"roles": []string{role},
			},
		},
	}

	ctx := contextWithToken(token)

	want := false
	got := IsInRole(ctx, "x")

	if got != want {
		t.Errorf("got %t, want %t", got, want)
		return
	}
}

func contextWithToken(token map[string]interface{}) context.Context {
	return context.WithValue(context.Background(), ContextKeyAccessToken, token)
}
