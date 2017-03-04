package web

import (
	"net/http"
	"fmt"
	"log"
	"time"
	"strings"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from bmt")
}

func echoHandler(w http.ResponseWriter, r *http.Request) {
	if (strings.HasPrefix(r.RequestURI, "/echo/")) {
		fmt.Fprintf(w, r.RequestURI[5:]);
	}
	if (strings.HasPrefix(r.RequestURI, "/echo") && strings.HasSuffix(r.RequestURI, "/echo")) {
		fmt.Fprintf(w, "nothing to display...")
	}
}

func handlerMaker(w http.ResponseWriter, r *http.Request) {
	if (strings.HasPrefix(r.RequestURI, "/echo")) {
		echoHandler(w, r)
	} else {
		rootHandler(w, r)
	}
}

func StartHttpServer() {

	http.HandleFunc("/echo*", echoHandler)
	http.HandleFunc("/", handlerMaker)



	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
