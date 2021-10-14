package auth

import "testing"

func TestGetUserNameOK(t *testing.T) {
	token := map[string]interface{}{
		"name": "username",
	}

	userName := GetUserName(token)

	if userName != "username" {
		t.Errorf("expected 'username', got '%s'", userName)
	}
}

func TestGetUserNameEmptyToken(t *testing.T) {
	token := map[string]interface{}{}

	userName := GetUserName(token)

	if userName != "" {
		t.Errorf("expected '', got '%s'", userName)
	}
}

func TestGetUserNameNilToken(t *testing.T) {
	userName := GetUserName(nil)

	if userName != "" {
		t.Errorf("expected '', got '%s'", userName)
	}
}

func TestGetRolesOK(t *testing.T) {
	token := map[string]interface{}{
		"resource_access": map[string]interface{}{
			"todo-api": map[string]interface{}{
				"roles": []interface{}{"role1", "role2"},
			},
		},
	}

	roles := GetRoles(token)

	rolesCount := len(roles)

	if rolesCount != 2 {
		t.Errorf("expected 2 roles, got %d", rolesCount)
		return
	}

	role1 := roles[0]
	role2 := roles[1]

	if role1 != "role1" {
		t.Errorf("expected 'role1', got '%s'", role1)
	}

	if role2 != "role2" {
		t.Errorf("expected 'role2', got '%s'", role2)
	}
}

func TestGetRolesEmptyToken(t *testing.T) {
	token := map[string]interface{}{}

	roles := GetRoles(token)

	rolesCount := len(roles)

	if rolesCount != 0 {
		t.Errorf("expected 0 roles, got %d", rolesCount)
	}
}

func TestGetRolesNilToken(t *testing.T) {
	roles := GetRoles(nil)

	rolesCount := len(roles)

	if rolesCount != 0 {
		t.Errorf("expected 0 roles, got %d", rolesCount)
	}
}
