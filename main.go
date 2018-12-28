package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Mux for handling routes
	mux := http.NewServeMux()

	fileDir := http.Dir("./public/")
	fileServer := http.FileServer(fileDir)

	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	mux.HandleFunc("/", index)
	// mux.Handle("/err", err)

	// mux.Handle("/login", login)
	// mux.Handle("/logout", logout)
	// mux.Handle("/signup", signup)
	// mux.Handle("/signup_account", signup_account)
	// mux.Handle("/authenticate", authenticate)

	// mux.Handle("/thread/new", newThread)
	// mux.Handle("/thread/create", createThread)
	// mux.Handle("/thread/post", postThread)
	// mux.Handle("/thread/read", readThread)

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

	server.ListenAndServe()
}

func index(w http.ResponseWriter, r *http.Request) {
	// files := []string{
	// 	"template/layout.html",
	// 	"template/navbar.html",
	// 	"template/index.html",
	// }

	// templates := template.Must(template.ParseFiles(files...))

	// threads, err := data.Threads()
	// if err == nil {
	// 	templates.ExecuteTemplate(w, "layout", threads)
	// }
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, World!")
}
