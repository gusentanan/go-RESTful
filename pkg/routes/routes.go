package routes

import (
	"context"

	"github.com/gorilla/mux"
	"github.com/gusentanan/go-RESTful/pkg/database"
)

func Router() *mux.Router {

	ctx := context.Background()
	router := mux.NewRouter()
	go database.InitializeMainDatabase(ctx)

	return router
}
