package internal

import (
	"context"
	_ "embed"
	"errors"
	"strings"

	"github.com/open-policy-agent/opa/rego"
	"github.com/samber/lo"
)

//go:embed authz/authz.rego
var authzRego string

type OPAAuthZ struct {
	evalQuery rego.PreparedEvalQuery
}

func NewOPAAuthZ(ctx context.Context) (*OPAAuthZ, error) {
	q, err := rego.New(
		rego.Query(
			strings.Join([]string{
				"allow = data.authz.allow",
				"sub = data.authz.sub",
				"username = data.authz.username",
				"permissions = data.authz.permissions",
			}, ";")),
		rego.Module("authz.rego", authzRego),
	).PrepareForEval(ctx)

	if err != nil {
		return nil, err
	}

	return &OPAAuthZ{q}, nil
}

func (az *OPAAuthZ) Authorize(ctx context.Context, query *AuthZQuery) (*AuthZResult, error) {
	result, err := az.evalQuery.Eval(ctx, rego.EvalInput(query))
	if err != nil {
		return nil, err
	}

	if len(result) == 0 {
		return nil, errors.New("undefined authz result")
	}

	bindings := result[0].Bindings

	allow, ok := bindings["allow"].(bool)
	if !ok {
		return nil, errors.New("unexpected authz result type: 'allow' must be bool")
	}

	sub, ok := bindings["sub"].(string)
	if !ok {
		return nil, errors.New("unexpected authz result type: 'sub' must be string")
	}

	username, ok := bindings["username"].(string)
	if !ok {
		return nil, errors.New("unexpected authz result type: 'username' must be string")
	}

	p, ok := bindings["permissions"].([]any)
	if !ok {
		return nil, errors.New("unexpected authz result type: 'permissions' must be []string")
	}

	permissions, ok := lo.FromAnySlice[string](p)
	if !ok {
		return nil, errors.New("unexpected authz result type: 'permissions' must be []string")
	}

	return &AuthZResult{allow, sub, username, permissions}, nil
}
