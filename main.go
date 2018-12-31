package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, %s!\n", p.ByName("name"))
}

func headers(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	h := r.Header
	fmt.Fprintln(w, h)
	fmt.Fprintf(w, "\n")
	fmt.Fprintln(w, h.Get("User-Agent"))
	fmt.Fprintf(w, "\n")
	fmt.Fprintln(w, h.Get("Accept"))
	fmt.Fprintf(w, "\n")
	fmt.Fprintln(w, h.Get("Accept-Encoding"))
	fmt.Fprintf(w, "\n")
	fmt.Fprintln(w, h.Get("Accept-Language"))
	fmt.Fprintf(w, "\n")
	fmt.Fprintln(w, h.Get("Cookie"))
	fmt.Fprintf(w, "\n")
	fmt.Fprintln(w, h.Get("Cache-Control"))
	fmt.Fprintf(w, "\n")
	fmt.Fprintln(w, h.Get("Upgrade-Insecure-Requests"))
	fmt.Fprintf(w, "\n")
}

func body(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	length := r.ContentLength

	body := make([]byte, length)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))

}

func processForm(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	fmt.Fprintln(w, r.Form)
}

func processPostForm(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	fmt.Fprintln(w, r.PostForm)
}

func processMultipartForm(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseMultipartForm(1024)
	fileHeader := r.MultipartForm.File["file"][0]
	file, err := fileHeader.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}

	fmt.Fprintln(w, r.MultipartForm)
}

func writeExample(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	str := `
	    <html>
        <head><title>Go Web Programming</title></head>
        <body><h1>Hello World</h1></body>
        </html>
	`
	w.Write([]byte(str))
}

func writeHeaderExample(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.WriteHeader(http.StatusNotImplemented)
	fmt.Fprintln(w, "No such service, try next door")
}

func headerExample(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Location", "http://google.com")
	w.WriteHeader(http.StatusTemporaryRedirect)
}

func main() {
	// Mux for handling routes
	// mux := http.NewServeMux()
	mux := httprouter.New()

	mux.GET("/hello/:name", hello)
	mux.GET("/headers", headers)
	mux.GET("/body", body)
	mux.POST("/process-form", processForm)
	mux.POST("/process-post-form", processPostForm)
	mux.POST("/process-multipart-form", processMultipartForm)
	mux.POST("/write", writeExample)
	mux.POST("/writeheader", writeHeaderExample)
	mux.GET("/redirect", headerExample)

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

	// http2.ConfigureServer(server, &http2.Server{})
	server.ListenAndServe()
	// server.ListenAndServeTLS("server.crt", "server.key")
}

func index(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Hello, World!")
}
