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

func DeleteProductController(pm *model.ProductModel) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		productName := model.Product{}
		if err := productName.FromJSON(r.Body); err != nil {
			log.Printf("error occurred on getting clients argument : %v", err)

			er := &ErrorResponse{ErrorCode: "r-01", Message: "error on receiving clients argument"}
			rw.WriteHeader(http.StatusBadRequest)
			er.ToJSON(rw)
		}

		if err := service.DeleteSingleProduct(productName, pm); err != nil {
			log.Printf("error occurred on deleting the product : %v", err)

			er := &ErrorResponse{ErrorCode: "r-04", Message: "error on deleting the product"}
			rw.WriteHeader(http.StatusInternalServerError)
			er.ToJSON(rw)
		}

		sr := &SuccesResponse{Message: "Product deleted successfully"}
		rw.WriteHeader(http.StatusOK)
		sr.ToJSON(rw)
	}
}

func GetSingleProductController(pm *model.ProductModel) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		productName := model.Product{}
		if err := productName.FromJSON(r.Body); err != nil {
			log.Printf("error occurred on getting clients argument")

			er := &ErrorResponse{ErrorCode: "r-01", Message: "error on receiving clients arguments"}
			rw.WriteHeader(http.StatusBadRequest)
			er.ToJSON(rw)
		}

		product, err := service.GetSingleProduct(productName, pm)
		if err != nil {
			log.Printf("error occurred on getting single product")

			er := &ErrorResponse{ErrorCode: "r-05", Message: "error on getting single product"}
			rw.WriteHeader(http.StatusInternalServerError)
			er.ToJSON(rw)
		}

		rw.WriteHeader(http.StatusOK)
		product.ToJSON(rw)
	}
}

func UpdateProductController(pm *model.ProductModel) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type", "application/json")

		updateProduct := model.Product{}
		if err := updateProduct.FromJSON(r.Body); err != nil {
			log.Printf("error occurred on getting clients argument")

			er := &ErrorResponse{ErrorCode: "r-01", Message: "error on getting clients argument"}
			rw.WriteHeader(http.StatusBadRequest)
			er.ToJSON(rw)
		}

		if err := service.UpdateSingleProduct(updateProduct, pm); err != nil {
			log.Printf("error occurred on updating the product")

			er := &ErrorResponse{ErrorCode: "r-06", Message: "error on updating product"}
			rw.WriteHeader(http.StatusInternalServerError)
			er.ToJSON(rw)
		}

		sr := &SuccesResponse{Message: "Product Updated Successfully"}
		rw.WriteHeader(http.StatusOK)
		sr.ToJSON(rw)
	}
}
