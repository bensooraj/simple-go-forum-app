package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"math/rand"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

// Post .
type Post struct {
	User    string
	Threads []string
}

// PostV1 .
type PostV1 struct {
	ID      int
	Content string
	Author  string
}

// PostByID .
var PostByID map[int]*PostV1

// PostsByAuthor .
var PostsByAuthor map[string][]*PostV1

func store(post PostV1) {
	PostByID[post.ID] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
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

func templateExampleOne(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, err := template.ParseFiles("templates/tmpl.html")
	if err != nil {
		fmt.Fprintln(w, "Couldn't parse the template html file.")
	} else {
		t.Execute(w, "I love you Hannah!")
	}
}

func templateExampleTwo(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rand.Seed(time.Now().Unix())
	t, err := template.ParseFiles("templates/conditional.html")
	if err != nil {
		fmt.Fprintln(w, "Couldn't parse the conditional html file.")
	} else {
		t.Execute(w, rand.Intn(10) > 5)
	}
}

func templateExampleThree(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	rand.Seed(time.Now().Unix())
	// daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	daysOfWeek := []string{}

	t, err := template.ParseFiles("templates/iterator.html")
	if err != nil {
		fmt.Fprintln(w, "Couldn't parse the conditional html file.")
	} else {
		t.Execute(w, daysOfWeek)
	}
}

func templateExampleFour(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, err := template.ParseFiles("templates/set_action.html")
	if err != nil {
		fmt.Fprintln(w, "Couldn't parse the conditional html file.")
	} else {
		t.Execute(w, "Hello, World!")
	}
}

func templateExampleFive(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, err := template.ParseFiles("templates/include_t1.html", "templates/include_t2.html")
	if err != nil {
		fmt.Fprintln(w, "Couldn't parse the include html file(s).")
	} else {
		fmt.Println(t)
		t.Execute(w, "Hello, World!")
	}
}

func formatDate(t time.Time) string {
	layout := "2019-01-01"
	return t.Format(layout)
}

func templateExampleFunc(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	funcMap := template.FuncMap{"fdate": formatDate}
	t := template.New("func_map.html").Funcs(funcMap)
	t, _ = t.ParseFiles("templates/func_map.html")
	t.Execute(w, time.Now())
}

func templateExampleSeven(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, err := template.ParseFiles("templates/context_1.html")
	if err != nil {
		fmt.Fprintln(w, "Couldn't parse the include html file(s).")
	} else {
		content := `I asked: <i>"What's up?"</i>`
		t.Execute(w, content)
	}
}

func templateExampleEight(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	t, err := template.ParseFiles("templates/xss_form.html")
	if err != nil {
		fmt.Fprintln(w, "Couldn't parse the XSS html form file.")
	} else {
		t.Execute(w, nil)
	}
}

func templateExampleXSSTest(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	w.Header().Set("X-XSS-Protection", "0")
	t, err := template.ParseFiles("templates/xss_test.html")
	if err != nil {
		fmt.Fprintln(w, "Couldn't parse the XSS test form file.")
	} else {
		// t.Execute(w, r.PostFormValue("comment"))
		t.Execute(w, template.HTML(r.PostFormValue("comment")))
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

	mux.GET("/template/example/one", templateExampleOne)
	mux.GET("/template/example/two", templateExampleTwo)
	mux.GET("/template/example/three", templateExampleThree)
	mux.GET("/template/example/four", templateExampleFour)
	mux.GET("/template/example/five", templateExampleFive)
	mux.GET("/template/example/six", templateExampleFunc)
	mux.GET("/template/example/seven", templateExampleSeven)

	// XSS Test
	mux.GET("/template/example/eight", templateExampleEight)
	mux.POST("/template/example/xss_test", templateExampleXSSTest)

	// App Memory Store test
	PostByID = make(map[int]*PostV1)
	PostsByAuthor = make(map[string][]*PostV1)

	post1 := PostV1{ID: 1, Content: "Hello World!", Author: "Sau Sheong"}
	post2 := PostV1{ID: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := PostV1{ID: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := PostV1{ID: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostByID[1])
	fmt.Println(PostByID[2])

	for _, authorPost := range PostsByAuthor["Sau Sheong"] {
		fmt.Println(authorPost)
	}

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
