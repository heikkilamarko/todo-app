package internal

import (
	"context"

	"github.com/mitchellh/mapstructure"
	"github.com/samber/lo"
)

type DBAuthZ struct {
	Repo Repository
}

func NewDBAuthZ(repo Repository) *DBAuthZ {
	return &DBAuthZ{repo}
}

func (az *DBAuthZ) Authorize(ctx context.Context, query *AuthZQuery) (*AuthZResult, error) {
	var payload struct {
		Sub               string `mapstructure:"sub"`
		PreferredUsername string `mapstructure:"preferred_username"`
		ResourceAccess    struct {
			TodoAPI struct {
				Roles []string `mapstructure:"roles"`
			} `mapstructure:"todo-api"`
		} `mapstructure:"resource_access"`
	}

	if err := mapstructure.Decode(query.Token, &payload); err != nil {
		return nil, err
	}

	permissions, err := az.Repo.GetPermissions(ctx, payload.ResourceAccess.TodoAPI.Roles)
	if err != nil {
		return nil, err
	}

	allow := lo.Contains(permissions, query.Permission)

	return &AuthZResult{
		Allow:       allow,
		Sub:         payload.Sub,
		Username:    payload.PreferredUsername,
		Permissions: permissions,
	}, nil
}
