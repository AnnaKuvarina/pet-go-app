package auth

import (
	. "github.com/AnnaKuvarina/pet-go-app/internal/api/auth"
	. "github.com/AnnaKuvarina/pet-go-app/internal/pg-store"
	"github.com/dgrijalva/jwt-go"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"time"
)

const issuer = "test"

var (
	AccessTokenGenerationError   = errors.New("could not generate access token")
	UnexpectedMethodError        = errors.New("Unexpected signing method in auth token")
	AccessTokenVerificationError = errors.New("failed to verify access token")
	AuthenticationError          = errors.New("invalid token: authentication failed")
)

type IService interface {
	GenerateAccessToken(user *UserCredItem) (string, error)
}

type Service struct {
	IService
}

type AccessTokenClaims struct {
	jwt.StandardClaims
	UserID    string `json:"user_id"`
	TokenType string `json:"token_type"`
}

func NewAuthService() *Service {
	return &Service{}
}

func (s *Service) GenerateAccessToken(user *UserCredItem) (string, error) {
	userID := user.ID
	tokenType := "access"

	claims := AccessTokenClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * time.Duration(60)).Unix(),
			Issuer:    issuer,
		},
		UserID:    userID,
		TokenType: tokenType,
	}

	//TODO: set path
	signBytes, err := ioutil.ReadFile("../keys/privateKey")
	if err != nil {
		log.Err(err).Msg("enable to read private key")
		return "", AccessTokenGenerationError
	}

	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
	if err != nil {
		log.Err(err).Msg("unable to parse private key")
		return "", AccessTokenGenerationError
	}

	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	return token.SignedString(signKey)
}

func (s *Service) ValidateAccessToken(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &AccessTokenClaims{}, verifyKey)
	if err != nil {
		log.Err(err).Msg("unable to parse access token claims")
		return "", err
	}

	claims, ok := token.Claims.(*AccessTokenClaims)
	if !ok || !token.Valid || claims.UserID == "" || claims.TokenType != "access" {
		return "", AuthenticationError
	}

	return claims.UserID, nil
}

func verifyKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
		log.Error().Msg("Unexpected signing method in auth token")
		return nil, UnexpectedMethodError
	}

	//TODO: set path
	verifyBytes, err := ioutil.ReadFile("../keys/publicKey")
	if err != nil {
		log.Err(err).Msg("enable to read public key")
		return "", AccessTokenVerificationError
	}

	verifyKey, err := jwt.ParseRSAPrivateKeyFromPEM(verifyBytes)
	if err != nil {
		log.Err(err).Msg("unable to parse private key")
		return "", AccessTokenVerificationError
	}

	return verifyKey, nil
}

func (s *Service) Authenticate(reqUser *LoginRequestData, user *UserCredItem) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqUser.Password)); err != nil {
		log.Debug().Msg("password hashes are not same")
		return false
	}
	return true
}