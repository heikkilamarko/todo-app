package envplugin

import (
	"context"
	"fmt"
	"net/http"
)

type Config struct {
	Env string `json:"env,omitempty"`
}

func CreateConfig() *Config {
	return &Config{}
}

type PluginHandler struct {
	next http.Handler
	env  string
	name string
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	if config.Env == "" {
		return nil, fmt.Errorf("env cannot be empty")
	}

	return &PluginHandler{
		next: next,
		env:  config.Env,
		name: name,
	}, nil
}

func (h *PluginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("X-Env", h.env)
	h.next.ServeHTTP(w, r)
}
