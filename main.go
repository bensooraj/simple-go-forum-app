package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/http2"

	"github.com/julienschmidt/httprouter"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s!\n", p.ByName("name"))
}

func main() {
	// Mux for handling routes
	// mux := http.NewServeMux()
	mux := httprouter.New()

	mux.GET("/hello/:name", hello)

	server := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: mux,
		// TLSConfig:    *tls.Config,
		// ReadTimeout:    time.Duration,
		// ReadHeaderTimeout:    time.Duration,
		// WriteTimeout:    time.Duration,
		// IdleTimeout:    time.Duration,
		// MaxHeaderBytes:    int,
		// TLSNextProto:    map[string]func(*Server, *tls.Conn, Handler),
		// ConnState:    func(net.Conn, ConnState),
		// ErrorLog:    *log.Logger,
		// disableKeepAlives:    int32,
		// inShutdown:        int32,
		// nextProtoOnce        sync.Once,
		// nextProtoErr        error,
		// mu        sync.Mutex,
		// listeners    map[*net.Listener]struct{},
		// activeConn    map[*conn]struct{},
		// doneChan    chan struct{},
		// onShutdown    []func(),
	}

	http2.ConfigureServer(server, &http2.Server{})
	// server.ListenAndServe()
	server.ListenAndServeTLS("server.crt", "server.key")
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, World!")
}
