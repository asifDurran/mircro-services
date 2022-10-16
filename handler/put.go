package handler

import (
	"micro-services/pkg/data"
	"net/http"
)

// swagger:route PUT /products products updateProduct
// Create a new product
// responses:
//     200: productsResponse
// 404: errorResponse
// 433: errorValidation

// Update handles PUT request to update a product
func (p *Products) Update(w http.ResponseWriter, r *http.Request) {

	//fetch the product from the context 
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	p.l.Println("updating reconrd id", prod.ID)


	err := data.UpdateProducts(prod)

	if err == data.ErrorProductNotFound {
		p.l.Println("product not found", err)


		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: "Product not found in database"}, w)

		return
	}

	// write the no content success header
	w.WriteHeader(http.StatusNoContent)

}