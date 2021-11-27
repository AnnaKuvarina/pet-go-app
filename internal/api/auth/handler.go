package auth

import (
	"encoding/json"
	. "github.com/AnnaKuvarina/pet-go-app/internal/postgre/credentials"
	"github.com/AnnaKuvarina/pet-go-app/internal/services/auth"
	"github.com/AnnaKuvarina/pet-go-app/pkg/api"
	"github.com/AnnaKuvarina/pet-go-app/pkg/utils"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
)

type Handler struct {
	Store       *UserCredsStore
	AuthService *auth.Service
}

func (h *Handler) Signup(resp http.ResponseWriter, req *http.Request) {
	user := req.Context().Value(UserRequestData{}).(UserRequestData)

	storedUser, err := h.Store.FindUserByEmail(user.Email)

	if storedUser != nil {
		log.Error().Msgf("User with email: %s is already exist", user.Email)
		api.WriteErrModel(resp, api.NewErrorResponse(""), http.StatusBadRequest)
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error().Err(err).Msg("unable to hash password")
		api.WriteErrModel(
			resp,
			api.NewErrorResponse("internal error"),
			http.StatusInternalServerError,
		)
		return
	}
	userId := uuid.New().String()
	storeUser := &UserCredItem{
		ID:    userId,
		Email:     user.Email,
		Password:  string(hashedPass),
		Username:  user.UserName,
		TokenHash: utils.GenerateRandomString(15),
	}

	err = h.Store.Create(storeUser)

	if err != nil {
		log.Error().Err(err).Msg("unable to insert user to db")
		api.WriteErrModel(resp, api.NewErrorResponse(""), http.StatusBadRequest)
		return
	}

	log.Debug().Msgf("user created successfully")
	api.WriteModel(resp, &SignupResponse{UserID: userId}, http.StatusCreated)
}

func (h *Handler) Login(resp http.ResponseWriter, req *http.Request) {
	// Read body
	body, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()

	loginRequestData := LoginRequestData{}
	if err = json.Unmarshal(body, &loginRequestData); err != nil {
		api.WriteModel(resp, api.NewErrorResponse("validation error"), http.StatusBadRequest)
		return
	}

	user, err := h.Store.FindUserByEmail(loginRequestData.Email)
	if err != nil {
		log.Error().Err(err).Msg("error fetching the user")
		api.WriteErrModel(resp, api.NewErrorResponse("failed to login"), http.StatusBadRequest)
		return
	}

	if valid := h.AuthService.Authenticate(&loginRequestData, user); !valid {
		api.WriteErrModel(resp, api.NewErrorResponse("failed to login"), http.StatusBadRequest)
		return
	}

	accessToken, err := h.AuthService.GenerateAccessToken(user)
	if err != nil {
		log.Error().Err(err).Msg("unable to generate access token")
		api.WriteErrModel(resp, api.NewErrorResponse("internal error"), http.StatusInternalServerError)
		return
	}

	api.WriteModel(resp, &LoginResponseData{AccessToken: accessToken, Username: user.Username}, http.StatusOK)
}
