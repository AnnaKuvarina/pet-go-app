package catalog

import (
	"github.com/AnnaKuvarina/pet-go-app/internal/mongo"
	"net/http"
)

type Handler struct {
	CatalogStore *mongo.AppMongoStore
}

func (h *Handler) GetCatalogByCategory(resp http.ResponseWriter, req *http.Request) {

}

func (h *Handler) GetProducDetails(resp http.ResponseWriter, req *http.Request) {

}
