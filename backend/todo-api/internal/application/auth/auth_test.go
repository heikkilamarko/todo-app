package auth

import "testing"

func TestGetUserNameOK(t *testing.T) {
	want := "username"

	token := map[string]interface{}{
		"name": want,
	}

	got := GetUserName(token)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGetUserNameEmptyToken(t *testing.T) {
	want := ""

	token := map[string]interface{}{}

	got := GetUserName(token)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestGetUserNameNilToken(t *testing.T) {
	want := ""

	got := GetUserName(nil)

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

	got := GetRoles(token)

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

	got := GetRoles(token)

	if len(got) != want {
		t.Errorf("got %d roles, want %d", len(got), want)
	}
}

func TestGetRolesNilToken(t *testing.T) {
	want := 0

	got := GetRoles(nil)

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

	want := true
	got := IsInRole(token, role)

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

	want := false
	got := IsInRole(token, "x")

	if got != want {
		t.Errorf("got %t, want %t", got, want)
		return
	}
}
