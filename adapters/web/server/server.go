package server

import (
	"log"
	"net/http"
	"os"
	"time"

	application "github.com/emanuelvss13/go-hexagonal/app"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
)

type WebServer struct {
	Service application.ProductInterface
}

func NewWebServer() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}

	err := server.ListenAndServe()

	if err != nil {
		log.Fatal(err)
	}

}