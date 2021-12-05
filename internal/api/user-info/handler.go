package user_info

import (
	. "github.com/AnnaKuvarina/pet-go-app/internal/pg-store"
	"net/http"
)

type Handler struct {
	Store *UserInfoStore
}

func (h *Handler) GetUserData(resp http.ResponseWriter, req *http.Request) {

}
