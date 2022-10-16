package handler

import (
	"micro-services/pkg/data"
	"net/http"
)

// swagger:route DELETE /products/{id} products deleteProduct
// Update a products details
//
// responses:
//      201: noContent
//  404: errorResponse
//  501: errorResponse

// Delete handels DELETE requests ans remove items from the database
func (p *Products) DeleteProduct(w http.ResponseWriter, r *http.Request){
	id := getProductID(r)

	p.l.Println("Handel Delete Product", id)

	err := data.DeleteProduct(id)

	if err == data.ErrorProductNotFound{
		p.l.Println("deleting record id does not exist")

		w.WriteHeader(http.StatusNotFound)
		data.ToJSON(&GenericError{Message: err.Error()}, w)

		return 
	}

	if err != nil {
		p.l.Println("deleting record", err)

		w.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(&GenericError{Message: err.Error()}, w)

		return
	}

	w.WriteHeader(http.StatusNoContent)
}