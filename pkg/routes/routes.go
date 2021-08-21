package routes

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/gusentanan/go-RESTful/internal/products/controller"
	"github.com/gusentanan/go-RESTful/pkg/database"
)

func Router() *mux.Router {

	ctx := context.Background()
	router := mux.NewRouter()
	db := database.InitializeMainDatabase(ctx)

	controller.ProductRouter(router, db)

	return router
}
