package controller

import (
	"database/sql"

	"github.com/gorilla/mux"
	"github.com/gusentanan/go-RESTful/internal/products/model"
)

func ProductRouter(r *mux.Router, db *sql.DB) {
	pm := &model.ProductModel{DB: db}
	r.HandleFunc("/api/product/insert", InsertProductController(pm)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/product/getAll", GetAllProductController(pm)).Methods("GET", "OPTIONS")
	r.HandleFunc("/api/product/delete", DeleteProductController(pm)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/product/getSingle", GetSingleProductController(pm)).Methods("POST", "OPTIONS")
	r.HandleFunc("/api/product/update", UpdateProductController(pm)).Methods("POST", "OPTIONS")
}
