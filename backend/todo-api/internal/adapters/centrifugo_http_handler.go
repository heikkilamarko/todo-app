package adapters

import (
	"errors"
	"net/http"

	"github.com/heikkilamarko/goutils"
	"github.com/rs/zerolog"
	"gopkg.in/square/go-jose.v2"
	"gopkg.in/square/go-jose.v2/jwt"
)

type tokenClaims struct {
	Subject string `json:"sub,omitempty"`
}

type tokenResponse struct {
	Token string `json:"token"`
}

type CentrifugoHTTPHandler struct {
	signingKey []byte
	logger     *zerolog.Logger
}

func NewCentrifugoHTTPHandler(signingKey string, logger *zerolog.Logger) *CentrifugoHTTPHandler {
	return &CentrifugoHTTPHandler{[]byte(signingKey), logger}
}

func (h *CentrifugoHTTPHandler) GetToken(w http.ResponseWriter, r *http.Request) {
	rt := goutils.TokenFromHeader(r)
	if rt == "" {
		h.logger.Error().Err(errors.New("token is empty")).Send()
		goutils.WriteUnauthorized(w, nil)
		return
	}

	t, err := jwt.ParseSigned(rt)
	if err != nil {
		h.logger.Error().Err(err).Send()
		goutils.WriteUnauthorized(w, nil)
		return
	}

	c := tokenClaims{}
	if err := t.UnsafeClaimsWithoutVerification(&c); err != nil {
		h.logger.Error().Err(err).Send()
		goutils.WriteUnauthorized(w, nil)
		return
	}

	s, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: h.signingKey}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		h.logger.Error().Err(err).Send()
		goutils.WriteInternalError(w, nil)
		return
	}

	token, err := jwt.Signed(s).Claims(c).CompactSerialize()
	if err != nil {
		h.logger.Error().Err(err).Send()
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteOK(w, &tokenResponse{token}, nil)
}
