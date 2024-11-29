package authz

import rego.v1

role_permissions := {
	"todo-user": ["todo.read", "todo.write"],
	"todo-viewer": ["todo.read"],
}

default allow := false

allow if {
	some p in permissions
	p == input.permission
}

sub := input.token.sub

username := input.token.preferred_username

permissions contains p if {
	some r in input.token.resource_access["todo-api"].roles
	some p in role_permissions[r]
}
