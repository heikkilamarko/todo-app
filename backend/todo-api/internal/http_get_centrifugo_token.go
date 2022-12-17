package internal

import (
	"errors"
	"net/http"

	"github.com/go-jose/go-jose/v3"
	"github.com/go-jose/go-jose/v3/jwt"
	"github.com/rs/zerolog"
)

type tokenClaims struct {
	Subject string `json:"sub"`
}

type tokenResponse struct {
	Token string `json:"token"`
}

type GetCentrifugoTokenHandler struct {
	Config *Config
	Logger *zerolog.Logger
}

func (h *GetCentrifugoTokenHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ar := GetAuthZResult(r.Context())

	sub := ar.Sub
	if sub == "" {
		h.Logger.Error().Err(errors.New("sub claim is empty")).Send()
		WriteResponse(w, http.StatusUnauthorized, nil)
		return
	}

	sig, err := jose.NewSigner(jose.SigningKey{
		Algorithm: jose.HS256,
		Key:       h.Config.CentrifugoTokenHMACSecretKey,
	}, (&jose.SignerOptions{}).WithType("JWT"))

	if err != nil {
		h.Logger.Error().Err(err).Send()
		WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}

	token, err := jwt.Signed(sig).Claims(tokenClaims{sub}).CompactSerialize()
	if err != nil {
		h.Logger.Error().Err(err).Send()
		WriteResponse(w, http.StatusInternalServerError, nil)
		return
	}

	WriteResponse(w, http.StatusOK, NewDataResponse(&tokenResponse{token}, nil))
}
