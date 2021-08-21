package service

import "github.com/gusentanan/go-RESTful/internal/products/model"

func InsertSingleProduct(product model.Product, pm *model.ProductModel) error {
	err := pm.Insert(product)
	return err
}

func GetAllProduct(pm *model.ProductModel) (*model.Products, error) {
	products, err := pm.GetAll()
	return products, err
}

func DeleteSingleProduct(product model.Product, pm *model.ProductModel) error {
	err := pm.Delete(product)
	return err
}

func GetSingleProduct(product model.Product, pm *model.ProductModel) (*model.Products, error) {
	singleProduct, err := pm.GetSpec(product)
	return singleProduct, err
}

func UpdateSingleProduct(product model.Product, pm *model.ProductModel) error {
	err := pm.Update(product)
	return err
}
