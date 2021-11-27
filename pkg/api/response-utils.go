package api

import (
	"encoding/json"
	"fmt"
	"github.com/rs/zerolog/log"
	"net/http"
)

const (
	errTypeError = "ERROR"
)

func HandleOptions(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE, PATCH")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if request.Method == "OPTIONS" {
		response.WriteHeader(http.StatusOK)
		return
	}
}

func WriteErrModel(w http.ResponseWriter, errModel *ErrorModel, httpStatus int) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	jsonStr, err := json.Marshal(errModel)
	if err != nil {
		log.Error().Err(err).Msg("Failed marshaling JSON response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(httpStatus)
	fmt.Fprint(w, string(jsonStr))
}

func WriteModel(response http.ResponseWriter, model interface{}, httpStatus int) {
	jsonStr, err := json.Marshal(model)
	response.Header().Set("Access-Control-Allow-Origin", "*")
	response.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE, PATCH")
	response.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		log.Error().Err(err).Msg("Failed marshaling JSON response")
		return
	}
	response.WriteHeader(httpStatus)
	fmt.Fprint(response, string(jsonStr))
}


func NewErrorResponse(msg string) *ErrorModel {
	return &ErrorModel{
		Message: msg,
		Type:    errTypeError,
	}
}
