package handler

import (
	"log"
	"net/http"
)

type GoodBye struct {
	l *log.Logger

}

func NewGoodbye(l *log.Logger) *GoodBye {
	return &GoodBye{l}
}

func (g *GoodBye) ServeHTTP(w http.ResponseWriter, r *http.Request){
	    w.Write([]byte("bytes but I do not know what it is? \n"))
		g.l.Println("GoodBye Hello world!")
}