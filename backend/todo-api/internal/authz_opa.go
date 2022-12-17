package internal

import (
	"context"
	_ "embed"
	"errors"
	"strings"

	"github.com/open-policy-agent/opa/rego"
	"github.com/samber/lo"
)

//go:embed rego/authz.rego
var authzRego string

type OPAAuthZ struct {
	evalQuery rego.PreparedEvalQuery
}

func NewOPAAuthZ(ctx context.Context) (*OPAAuthZ, error) {
	q, err := rego.New(
		rego.Query(
			strings.Join([]string{
				"allow=data.authz.allow",
				"sub=data.authz.sub",
				"username=data.authz.username",
				"permissions=data.authz.permissions",
			}, " ")),
		rego.Module("authz.rego", authzRego),
	).PrepareForEval(ctx)

	if err != nil {
		return nil, err
	}

	return &OPAAuthZ{q}, nil
}

func (az *OPAAuthZ) Authorize(ctx context.Context, query *AuthZQuery) (*AuthZResult, error) {
	r, err := az.evalQuery.Eval(ctx, rego.EvalInput(query))
	if err != nil {
		return nil, err
	}

	allow, ok := r[0].Bindings["allow"].(bool)
	if !ok {
		return nil, errors.New("invalid authz result: 'allow' must be bool")
	}

	sub, ok := r[0].Bindings["sub"].(string)
	if !ok {
		return nil, errors.New("invalid authz result: 'sub' must be string")
	}

	username, ok := r[0].Bindings["username"].(string)
	if !ok {
		return nil, errors.New("invalid authz result: 'username' must be string")
	}

	p, ok := r[0].Bindings["permissions"].([]any)
	if !ok {
		return nil, errors.New("invalid authz result: 'permissions' must be []string")
	}

	permissions, ok := lo.FromAnySlice[string](p)
	if !ok {
		return nil, errors.New("invalid authz result: 'permissions' must be []string")
	}

	return &AuthZResult{allow, sub, username, permissions}, nil
}
