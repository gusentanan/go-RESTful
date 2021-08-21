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
	UPDATE_PRODUCT_STMT       string = "UPDATE products SET productName = ?, price = ?, shortDesc = ? WHERE id = ?"
)

type Product struct {
	ID               int64  `json:"id"`
	ProductName      string `json:"productName"`
	Price            int64  `json:"price"`
	ShortDescription string `json:"shortDesc"`
}

type Products []*Product

type ProductModel struct {
	DB *sql.DB
}

// bundle data with respect to json formats
func (product Products) ToJSON(w io.Writer) error {
	return json.NewEncoder(w).Encode(product)
}

// revert encode data back to original form
func (product *Product) FromJSON(r io.Reader) error {
	return json.NewDecoder(r).Decode(product)
}

// insert single product to MySQL database
func (pm *ProductModel) Insert(product Product) error {
	if _, err := pm.DB.Exec(INSERT_PRODUCT_STMT, product.ProductName, product.Price, product.ShortDescription); err != nil {
		log.Printf("error occured on inserting product : %s", err)
		return err
	}
	return nil
}

// get all products from MySQL database
func (pm *ProductModel) GetAll() (*Products, error) {
	res, err := pm.DB.Query(GET_ALL_PRODUCT_STMT)
	if err != nil {
		log.Printf("error occured on getting all the products : %s", err)
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

// get specific product from MySQL database
func (pm *ProductModel) GetSpec(p Product) (*Products, error) {
	res, err := pm.DB.Query(GET_SPECIFIC_PRODUCT_STMT, p.ProductName)
	if err != nil {
		log.Printf("error occurred on getting specific product : %s", err)
		return nil, err
	}

	defer res.Close()

	var singleProduct Products
	for res.Next() {
		product := &Product{}
		res.Scan(&product.ID, &product.ProductName, &product.Price, &product.ShortDescription)
		singleProduct = append(singleProduct, product)
	}
	return &singleProduct, nil
}

// delete single product from MySQL database
func (pm *ProductModel) Delete(product Product) error {
	if _, err := pm.DB.Query(DELETE_PRODUCT_STMT, product.ProductName); err != nil {
		log.Printf("error occurred on deleting a specific product : %s", err)
		return err
	}
	return nil
}

// update product from MySQL database
func (pm *ProductModel) Update(product Product) error {
	if _, err := pm.DB.Query(UPDATE_PRODUCT_STMT, product.ProductName, product.Price, product.ShortDescription, product.ID); err != nil {
		log.Printf("error occurred on updating a specific product : %s", err)
		return err
	}
	return nil
}
