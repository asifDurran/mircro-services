// Package classification of Product API
//
// Documentation for Product API222
//
//	Schemes: http
//	BasePath: /
//	Version: 1.0.0
//
//	Consumes:
//	- application/json
//
//	Produces:
//	- application/json
//
// swagger:meta
package handler

import (
	"fmt"
	"log"
	"micro-services/pkg/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// swagger:response noContent
type productsNoContent struct {

}

// swagger:parameters deleteProduct
type productsIDParameterWrapper struct {
	// The id of the product to delete from the database
	// in: path
	// required: true
	ID int   `json:"id"`
}

// Products handler for getting and updating products
type Products struct {
 l  *log.Logger
 v *data.Validation
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages  []string `json: "messages"`
}

// NewProducts returns a new products handler with the given logger
func NewProducts(l*log.Logger, v *data.Validation) *Products {
	return &Products{l, v}
}

// ErrInvalidProductPath is an error message when the product path is not valid
var ErrInvalidProductPath = fmt.Errorf("Invalid Path, path should be /products/[id]")

// GenericError is a generic error message when the product path is not valid
type GenericError struct{
	Message string `json:"message"`
}

// KeyProduct is a key used for the product object in th context
type KeyProduct struct{}

// getProductID returns the product ID from the URL
// Panics if cannot convert the id into an integer
// this should never happen as the router ensures that 
// this is a valid number
func getProductID(r *http.Request) int{
	// parse the product id from the url
	vars := mux.Vars(r)
    
	// convert the id into an integer and return 
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		// should never happen
		panic(err)
	}
	
	return id
}