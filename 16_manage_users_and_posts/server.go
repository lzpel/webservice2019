package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"reflect"
	"text/template"
	"time"
)

type Response = http.ResponseWriter
type Request = *http.Request
type Dict = map[string]interface{}

func Handle(pattern string, handler func(w Response, r Request)) {
	http.HandleFunc(pattern, handler)
}

func HandleSample() {
	// {{ .datetime }} {{ .datetime | length }} {{ length .datetime}}
	Handle(`/sample`, func(w Response, r Request) {
		Writef(w, Dict{
			"datetime": time.Now().Format("2006-01-02 15:04:05"),
			"length": func(s string) int {
				return len(s)
			},
		}, "server.go")
	})
}

func HandleFile() {
	Handle(`/file/`, func(w Response, r Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
	})
}

func Environment(CredentialPath string, DatastoreKind string) {
	var e error
	if CredentialPath != "" {
		var fc map[string]string
		bytes, _ := ioutil.ReadFile(CredentialPath)
		e = json.Unmarshal(bytes, &fc)
		if e == nil {
			e = os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", CredentialPath)
			e = os.Setenv("PROJECT_ID", fc["project_id"])
			e = os.Setenv("CLIENT_EMAIL", fc["client_email"])
			e = os.Setenv("PRIVATE_KEY", fc["private_key"])
		}
		if e != nil {
			log.Println(e)
		} else {
			log.Printf("Setting credentials %s", fc["project_id"])
		}
	}
	e = os.Setenv("DATASTORE_KIND", DatastoreKind)
}
func Listen() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Printf("http://localhost:%s", port)
	e := http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
	if e != nil {
		log.Println(e)
	}
}

func Writef(w io.Writer, value Dict, filename ...string, ) {
	funcs := template.FuncMap{}
	for k, v := range value {
		if reflect.ValueOf(v).Kind() == reflect.Func {
			funcs[k] = v
			delete(value, k)
		}
	}
	t, e := template.New(filename[0]).Funcs(funcs).ParseFiles(filename...)
	if e == nil {
		e = t.Execute(w, value)
	}
	if e != nil {
		log.Println(e)
	}
}

func Redirect(w Response, r Request, url string) {
	http.Redirect(w, r, url, 301)
}

func GetMultipartFileHeaders(r Request) map[string][]*multipart.FileHeader {
	e := r.ParseMultipartForm(200000)
	if e == nil {
		return r.MultipartForm.File
	}
	log.Println(e)
	return nil
}

func CookieSet(w Response, k, v string, days int) {
	http.SetCookie(w, &http.Cookie{
		Path:   "/",
		Name:   k,
		Value:  v,
		MaxAge: 86400 * days,
	})
}

func CookieGet(r Request, k string) string {
	cookie, err := r.Cookie(k)
	if err == nil {
		return cookie.Value
	}
	return ""
}