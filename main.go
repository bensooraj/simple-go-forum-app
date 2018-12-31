package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Post .
type Post struct {
	User    string
	Threads []string
}

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
	w.WriteHeader(301)
}

func jsonExample(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	post := Post{
		User: "Ben Sooraj",
		Threads: []string{
			"Hello, World!",
			"Test thread #1",
			"Test thread #2",
			"Test thread #3",
		},
	}
	jsonData, _ := json.Marshal(&post)
	w.Write(jsonData)
}

func setCookie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c1 := http.Cookie{
		Name:     "first_cookie",
		Value:    "Go Web Programming0",
		HttpOnly: true,
	}

	c2 := http.Cookie{
		Name:     "second_cookie",
		Value:    "Manning Publications Co0",
		HttpOnly: true,
	}

	// w.Header().Set("Set-Cookie", c1.String())
	// w.Header().Add("Set-Cookie", c2.String())
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getCookie(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// h := r.Header["Cookie"]
	// fmt.Fprintln(w, h)

	c1, err := r.Cookie("first_cookie")
	if err != nil {
		fmt.Fprintln(w, "Cannot get the first cookie")
	}
	fmt.Fprintln(w, c1)

	cs := r.Cookies()
	fmt.Fprintln(w, cs)
}

func setMessage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	msg := []byte("Hello, World!")
	c := http.Cookie{
		Name:  "flash",
		Value: base64.URLEncoding.EncodeToString(msg),
	}
	http.SetCookie(w, &c)
}

func showMessage(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	c, err := r.Cookie("flash")
	if err != nil {
		if err == http.ErrNoCookie {
			fmt.Fprintln(w, "Flash Message Not Found")
		}
	} else {

		rc := http.Cookie{
			Name:    "flash",
			MaxAge:  -1,
			Expires: time.Unix(1, 0),
		}
		http.SetCookie(w, &rc)

		flashMessage, _ := base64.URLEncoding.DecodeString(c.Value)
		fmt.Fprintln(w, string(flashMessage))
	}
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
	mux.GET("/json", jsonExample)

	mux.GET("/set_cookie", setCookie)
	mux.GET("/get_cookie", getCookie)

	mux.GET("/set_message", setMessage)
	mux.GET("/show_message", showMessage)

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
