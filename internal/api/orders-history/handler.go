package ordershistory

import (
	"github.com/AnnaKuvarina/pet-go-app/internal/mongo"
	"github.com/AnnaKuvarina/pet-go-app/internal/pg-store"
	"net/http"
)

type Handler struct {
	HistoryStore *pg_store.OrdersHistoryStore
	CatalogStore *mongo.AppMongoStore
}

func (h *Handler) GetUserOrdersHistory(resp http.ResponseWriter, req *http.Request) {

}

func (h *Handler) NewOrder(resp http.ResponseWriter, req *http.Request) {

}
