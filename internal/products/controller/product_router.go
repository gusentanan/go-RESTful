package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/gusentanan/go-RESTful/internal/products/model"
)

func ProductRouter(r *mux.Router, db *sql.DB) {
	pm := &model.ProductModel{DB: db}
	r.HandleFunc("/api/product", InsertProductController(pm)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/product", GetAllProductController(pm)).Methods("GET", "OPTIONS")
}
