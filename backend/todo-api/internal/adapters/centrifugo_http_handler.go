package adapters

import (
	"errors"
	"net/http"
	"todo-api/internal/adapters/auth"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
	"github.com/heikkilamarko/goutils"
	"github.com/rs/zerolog"
)

type tokenClaims struct {
	Subject string `json:"sub"`
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
	if err := auth.AuthorizeRead(r); err != nil {
		h.logError(err)
		goutils.WriteUnauthorized(w, nil)
		return
	}

	sub := auth.GetSubject(r.Context())
	if sub == "" {
		h.logError(errors.New("sub claim is empty"))
		goutils.WriteUnauthorized(w, nil)
		return
	}

	sig, err := jose.NewSigner(jose.SigningKey{Algorithm: jose.HS256, Key: h.signingKey}, (&jose.SignerOptions{}).WithType("JWT"))
	if err != nil {
		h.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	token, err := jwt.Signed(sig).Claims(tokenClaims{sub}).CompactSerialize()
	if err != nil {
		h.logError(err)
		goutils.WriteInternalError(w, nil)
		return
	}

	goutils.WriteOK(w, &tokenResponse{token}, nil)
}

func (h *CentrifugoHTTPHandler) logError(err error) {
	h.logger.Error().Err(err).Send()
}
