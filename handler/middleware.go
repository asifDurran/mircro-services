package handler

import (
	"context"
	"micro-services/pkg/data"
	"net/http"
)

// MiddlewareValidation validate product in the request and calls next if ok
func (p *Products) MiddlewareValidation(next http.Handler) http.Handler{
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		prod := &data.Products{}


		err := data.FromJSON(prod, r.Body)
		if err != nil {
			p.l.Println("deserializing product", err)

			w.WriteHeader(http.StatusBadRequest)
			data.ToJSON(&GenericError{Message: err.Error()}, w)
			return
		}

		// validate the product
		errs := p.v.Validate(prod)

		if len(errs) != 0 {
				p.l.Println("validating product", errs)

				// return the validation messages as an array

				w.WriteHeader(http.StatusUnprocessableEntity)
				data.ToJSON(&ValidationError{Messages: errs.Errors()}, w)

				return
		}


		// add the product to the context

		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)

		r = r.WithContext(ctx)


		// call the next handler, which can be another middleware in the chian, or the final handler

		next.ServeHTTP(w, r)

	})
}