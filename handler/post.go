package handler

import (
	"micro-services/pkg/data"
	"net/http"
)

// swagger:route POST /products products createProdcut
// Create a new product
// responses:
//     200: productsResponse
// 422: errorValidation
// 501: errorResponse

// Create handles POST request to add a new product
func (p *Products) Create(w http.ResponseWriter, r *http.Request){

  // fetch the product from the context
  prod := r.Context().Value(KeyProduct{}).(data.Product)

  p.l.Printf("Inserting product %#v \n", prod)

  data.AddProduct(prod)
}