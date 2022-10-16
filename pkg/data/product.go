package data

import (
	"fmt"
	"time"
)

// ErrProductNotFound is an error raised when a product can not be found in the database
var ErrorProductNotFound = fmt.Errorf("Product not found")

// Product define the structure for an API product
// swagger:model
type Product struct {
	// the id for this user
    //
    // required: true
    // min: 1
	ID           int       `json:"id"` // Unique identifier for the product

    // the name for this user
    // required: true
    // min length: 3
	Name 		 string    `json:"name" validate:"required"`

    // the dedcription for this product
    //
    // required: true
    // max length: 1000
	Description  string    `json:"description" validate:"required"`

	// the name for this user
    // required: true
    // min: 0.1
	Price 		 float64   `json:"price" validate:"gt=0"`

	// the name for this user
    // required: true
    // pattern: [a-z]+-[a-z]+-[a-z]+
	SKU   		 string    `json:"sku" validate:"required,sku"`
	CreatedOn    string	   `json:"-"`
	UpdatedOn    string	   `json:"-"`
	DeletedOn	 string	   `json:"-"`
}

// Products defines a slice of Product
type Products []*Product

// GetProducts returns all products from the database
func GetProducts()Products {
	return productList
}

// GetProductByID retuens a single product which matches the id from the
// database.
// if a product is not found this function returns a ProductnotFoundError error
func GetProductsByID(id int)(*Product, error){
	i := findIndexByProductID(id)
	if id == -1 {
		return nil, ErrorProductNotFound
	}

	return productList[i], nil
}

// addProduct adds a new product to the database
func AddProduct (p Product){
	// get the next ID in sequence
	maxID := productList[len(productList) -1].ID 
   p.ID = maxID + 1

   productList = append(productList, &p)

}
// DeleteProduct deletes a product from the database
func DeleteProduct(id int) error {
	i := findIndexByProductID(id)

	if i == -1 {
		return ErrorProductNotFound
	}
	productList = append(productList[:i],productList[i+1] )
	
	return nil
}

// to generate next ID
func getNextID() int {
 lp := productList[len(productList) -1]
 return lp.ID + 1
}

// UpdateProduct replace a product in the database with the given
// item.
// if a product with the given id does not exist in the database
// this function returns a ProductNotFoundError error
func UpdateProducts(p Product) error {
	i := findIndexByProductID(p.ID)

	if i != -1 {
	 return ErrorProductNotFound
	}
	
	//update the product in the DB
	productList[i] = &p
    
	return nil
}

// findIndex finds the index of a product in the database
// return -1 when no product can be found
func findIndexByProductID(id int) int {
	for i, p := range productList {
		if p.ID == id {
			return i
		}
	}

	return -1
}

var productList = []*Product {
	&Product{
		ID:           1,
		Name:         "Tea",
		Description:  "DoothPathi",
		Price:        5.9,
		SKU:          "Durn12",
		CreatedOn:    time.November.String(),
		UpdatedOn:    time.November.String(),
	},
	&Product{
		ID:           2,
		Name:         "Green Tea",
		Description:  "Qahwa",
		Price:        3.9,
		SKU:          "rn12",
		CreatedOn:    time.November.String(),
		UpdatedOn:    time.November.String(),
	},
}