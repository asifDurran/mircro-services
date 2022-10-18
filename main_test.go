package main

import (
	"fmt"
	"micro-services/sdk/client"
	"micro-services/sdk/client/products"
	"testing"
)



func TestOurClient(t *testing.T){
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("Products %#v", prod.GetPayload()[0])
	t.Fail()
	 
}