/* {{.datetime}} */
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
	"text/template"
	"time"
)

type response = http.ResponseWriter
type request = *http.Request
type dict = map[string]interface{}

func logf(format string, v ...interface{}) {
	log.Printf(format, v...)
}
func logs(v ...interface{}) {
	log.Println(v...)
}
func logfatal(v ...interface{}){
	log.Fatalln(v...)
}
func handle(pattern string, handler func(response, request)) {
	http.HandleFunc(pattern, handler)
}
func servefile(w response,r request,path string){
	http.ServeFile(w, r, path)
}

func listen(credentialpath string) {
	if credentialpath != ""{
		var fc map[string]string
		bytes, _ := ioutil.ReadFile(credentialpath)
		json.Unmarshal(bytes, &fc)
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS",credentialpath)
		os.Setenv("PROJECT_ID",fc["project_id"])
		os.Setenv("CLIENT_EMAIL",fc["client_email"])
		os.Setenv("PRIVATE_KEY",fc["private_key"])
		logf("Setting credentials %s", fc["project_id"])
	}
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		logf("Defaulting to port %s", port)
	}
	logf("Listening on port %s", port)
	logf("Open http://localhost:%s in the browser", port)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}

func samplehandler(w response, r request) {
	//logs(r.URL.Path)
	//logs(r.FormValue("key"))
	writetemplate(w, "server.go", dict{"datetime": time.Now().Format("2006-01-02 15:04:05"),},nil)
}

func writetemplate(w io.Writer, filename string, data interface{},funcs template.FuncMap) {
	t, e1:=template.New(filename).Funcs(funcs).ParseFiles(filename)
	if e1!=nil{
		logfatal(e1)
	}
	e2:=t.Execute(w, data)
	if e1!=nil{
		logfatal(e2)
	}
}

func redirect(w response, r request,url string){
	http.Redirect(w,r,"/",301)
}

func multipartfile(r request) map[string][]*multipart.FileHeader{
	err := r.ParseMultipartForm(200000)
	if err != nil {
		logs(err)
		return nil
	}
	return r.MultipartForm.File
}