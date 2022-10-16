package main

import (
	"context"
	"log"
	"micro-services/handler"
	"micro-services/pkg/data"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"

	"github.com/nicholasjackson/env"
)

var bindAddress = env.String("BIND_ADDRESS", false, ":9090", "Bind address for the server")

func main()  {
	
	env.Parse()

	l := log.New(os.Stdout, "product-api", log.LstdFlags)
	v := data.NewValidation()

	//create the handler
    ph := handler.NewProducts(l, v)
	

	// create a new serve mux and register the handlers
    sm := mux.NewRouter()

    // handlers for API
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products",ph.ListAll)
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle)

	// Put Router
	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products", ph.Update)
	putRouter.Use(ph.MiddlewareValidation)

	// Post Router
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.Create)
	postRouter.Use(ph.MiddlewareValidation)
    
	// Delete Router
	deleteRouter := sm.Methods(http.MethodPost).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}", ph.DeleteProduct)

	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh  := middleware.Redoc(opts, nil)


	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	//create a new server
    s := &http.Server{
		Addr: ":9090",  				//
		Handler: sm,
		IdleTimeout: 120*time.Second,
		ReadTimeout: 1 *time.Second,
		WriteTimeout: 1 *time.Second,
	}
	//start the server
	go func ()  {

		l.Println("Starting server on port 9090")
		err := s.ListenAndServe()

		if err != nil {
			l.Printf("Error starting server : %s \n", err)
			l.Fatal(err)
		}
	}()

	// trap or intrup and gracefully shutdown the server
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan,os.Kill)

	// Block until a signal is received
    sig := <- sigChan
	l.Println("Received terminate, graceful shutdown", sig)
	
	// gracefully shutdown the server, waiting max 30 seconds for current opreations to complete
	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}