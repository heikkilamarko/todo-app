package authz

import future.keywords.contains
import future.keywords.in

role_permissions := {
	"todo-user": ["todo.read", "todo.write"],
	"todo-viewer": ["todo.read"],
}

default allow = false

allow {
	p := permissions[_]
	p == input.permission
}

sub := input.token.sub

username := input.token.preferred_username

permissions contains p {
	some r in input.token.resource_access["todo-api"].roles
	some p in role_permissions[r]
}
