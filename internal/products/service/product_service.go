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

func DeleteSingleProduct(productID int64, pm *model.ProductModel) error {
	err := pm.Delete(productID)
	return err
}

func GetSingleProduct(productName string, pm *model.ProductModel) (*model.Products, error) {
	product, err := pm.GetSpec(productName)
	return product, err
}

func UpdateSingleProduct(product model.Product, pm *model.ProductModel) error {
	err := pm.Update(product)
	return err
}
