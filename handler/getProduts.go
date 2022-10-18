package handler

import (
	"micro-services/pkg/data"
	"net/http"

	"github.com/go-playground/validator"
)
type Validation struct {
	validate *validator.Validate
}
// swagger:route GET /products products listProducts
// Returns a list of products
// responses:
//     200: productsResponse

// ListAll handels Get requests and returns all current products
func (p *Products) ListAll(w http.ResponseWriter, r *http.Request){
        p.l.Println("[Debug] get all products")

        w.Header().Add("Content-Type", "application/json")
        prods := data.GetProducts()

        err := data.ToJSON(prods, w)

        if err != nil {
                // we should never be here but log the error just incase
                p.l.Println("[Error] serializing product", err)
        }
        
}

// swagger:route GET /products/{id} products listSingleProduct
// Returns a single product
// responses:
//     200: productsResponse
//     404: errorResponse

// ListSingle handels Get requests and return single product from database
func (p *Products) ListSingle(w http.ResponseWriter, r *http.Request) {
        id := getProductID(r)

        p.l.Println("[DEBUG] get record id",id)

        prod, err := data.GetProductsByID(id)

        switch err {
        case nil:
        case data.ErrorProductNotFound:
                p.l.Println("[Error] fetching product", err)

                w.WriteHeader(http.StatusNotFound)
                data.ToJSON(&GenericError{Message: err.Error()}, w)

                return

        default:
                p.l.Println("[Error] fetching product", err)

                w.WriteHeader(http.StatusNotFound)
                data.ToJSON(&GenericError{Message: err.Error()}, w)

                return
        }

        err = data.ToJSON(prod, w)

        if err != nil {
                // we should never be here but log the error just incase
                p.l.Println("[Error] serializing product", err)
        }
}