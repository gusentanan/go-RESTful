package model

import (
	"database/sql"
	"encoding/json"
	"io"
	"log"
)

const (
	INSERT_PRODUCT_STMT       string = "INSERT INTO products (productName, price, shortDesc) VALUES (?,?,?)"
	GET_SPECIFIC_PRODUCT_STMT string = "SELECT * FROM products WHERE productName = ?"
	GET_ALL_PRODUCT_STMT      string = "SELECT * FROM products"
	DELETE_PRODUCT_STMT       string = "DELETE FROM products WHERE productName = ? "
	UPDATE_PRODUCT_STMT       string = "UPDATE products SET ProductName = ?, price = ?, shortDesc = ? WHERE ID = ?"
)

type Product struct {
	ID               int64  `json:"id"`
	ProductName      string `json:"productsName"`
	Price            int64  `json:"price"`
	ShortDescription string `json:"shortDesc"`
}

type Products []*Product

type ProductModel struct {
	DB *sql.DB
}

// to parse argument from non-json to json
func (product Products) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(product)
}

// to parse argument from json to non-json
func (product *Product) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(product)
}

// insert one product
func (pm *ProductModel) Insert(product Product) error {
	if _, err := pm.DB.Exec(INSERT_PRODUCT_STMT, product.ProductName, product.Price, product.ShortDescription); err != nil {
		log.Printf("error occured on inserting product : %s", err)
		return err
	}
	return nil
}

// get all products
func (pm *ProductModel) GetAll() (*Products, error) {
	res, err := pm.DB.Query(GET_ALL_PRODUCT_STMT)
	if err != nil {
		log.Printf("error occured on gx`etting all the products : %s", err)
	}

	defer res.Close()
	var products Products
	for res.Next() {
		product := &Product{}
		res.Scan(&product.ID, &product.ProductName, &product.Price, &product.ShortDescription)
		products = append(products, product)
	}

	return &products, nil
}

// get specific product
func (pm *ProductModel) GetSpec(productName string) (*Products, error) {
	res, err := pm.DB.Query(GET_SPECIFIC_PRODUCT_STMT, productName)
	if err != nil {
		log.Printf("error occurred on getting specific product : %s", err)
		return nil, err
	}

	defer res.Close()

	var singleProduct Products
	product := Product{}
	res.Scan(&product.ID, &product.ProductName, &product.Price, &product.ShortDescription)
	singleProduct = append(singleProduct, &product)

	return &singleProduct, nil
}

// delete single product
func (pm *ProductModel) Delete(productID int64) error {
	if _, err := pm.DB.Query(DELETE_PRODUCT_STMT, productID); err != nil {
		log.Printf("error occurred on deleting a specific product : %s", err)
		return err
	}
	// add some json response after successfull run
	return nil
}

// update single product
func (pm *ProductModel) Update(product Product) error {
	if _, err := pm.DB.Query(UPDATE_PRODUCT_STMT, product.ProductName, product.Price, product.ShortDescription); err != nil {
		log.Printf("error occurred on updating a specific product : %s", err)
		return err
	}
	return nil
}
