package api

import (
	"context"
	"github.com/AnnaKuvarina/pet-go-app/internal/services/auth"
	"github.com/AnnaKuvarina/pet-go-app/pkg/utils"
	"github.com/rs/zerolog/log"
	"net/http"
	"strings"
)

func TrimSuffixMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// remove the trailing slash from our URL Path if it's not root
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimSuffix(r.URL.Path, "/")
		}

		// Call the next handler, which can be another middleware in the chain, or the final handler.
		next.ServeHTTP(w, r)
	})
}

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

func MiddlewareValidateAccessToken(next http.Handler, authService *auth.Service) http.Handler {
	return http.HandlerFunc(func(resp http.ResponseWriter, req *http.Request) {
		log.Debug().Msg("validate access token")

		token, err := utils.ExtractToken(req)
		if err != nil {
			WriteErrModel(resp, NewErrorResponse("token invalid"), http.StatusBadRequest)
			return
		}
		log.Debug().Msgf("got access token %s", token)

		userID, err := authService.ValidateAccessToken(token)
		if err != nil {
			log.Error().Err(err).Msg("token validation failed")
			WriteErrModel(resp, NewErrorResponse("token validation failed"), http.StatusBadRequest)
			return
		}
		log.Debug().Msg("access token validated")

		ctx := context.WithValue(req.Context(), UserIDKey{}, userID)
		req = req.WithContext(ctx)

		next.ServeHTTP(resp, req)
	})
}
