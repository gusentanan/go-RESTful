package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gusentanan/go-RESTful/internal/products/model"
	"github.com/gusentanan/go-RESTful/internal/products/service"
)

type ErrorResponse struct {
	ErrorCode string `json:"status"`
	Message   string `json:"message"`
}

type SuccesResponse struct {
	Message string `json:"message"`
}

func (sr *SuccesResponse) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(sr)
}

func (er *ErrorResponse) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(er)
}

func InsertProductController(pm *model.ProductModel) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		product := model.Product{}
		if err := product.FromJSON(r.Body); err != nil {
			log.Printf("error occurred on getting client argument : %v", err)

			er := &ErrorResponse{ErrorCode: "r-01", Message: "error on receiving clients argument"}
			rw.WriteHeader(http.StatusBadRequest)
			er.ToJSON(rw)
			return
		}

		if err := service.InsertSingleProduct(product, pm); err != nil {
			log.Printf("error occurred on inserting product : %v", err)

			er := &ErrorResponse{ErrorCode: "r-02", Message: "error on inserting product"}
			rw.WriteHeader(http.StatusInternalServerError)
			er.ToJSON(rw)
			return
		}

		sr := &SuccesResponse{Message: "Product Created Successfully"}
		rw.WriteHeader(http.StatusCreated)
		sr.ToJSON(rw)
	}
}

func GetAllProductController(pm *model.ProductModel) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		products, err := service.GetAllProduct(pm)
		if err != nil {
			er := &ErrorResponse{ErrorCode: "r-03", Message: "error on getting all product"}
			rw.WriteHeader(http.StatusInternalServerError)
			er.ToJSON(rw)
		}

		rw.WriteHeader(http.StatusOK)
		products.ToJSON(rw)
	}
}
