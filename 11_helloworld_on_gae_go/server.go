/* {{.datetime}} */
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
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
	log.Print(v...)
}
func handle(pattern string, handler func(response, request)) {
	http.HandleFunc(pattern, handler)
}
func serve(credentialpath string) {
	if credentialpath != ""{
		var fc map[string]string
		bytes, _ := ioutil.ReadFile(credentialpath)
		json.Unmarshal(bytes, &fc)
		os.Setenv("GOOGLE_APPLICATION_CREDENTIALS",credentialpath)
		os.Setenv("PROJECT_ID",fc["project_id"])
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
	writetemplate(w, "server.go", dict{"datetime": time.Now().Format("2006-01-02 15:04:05"),})
}

func writetemplate(w io.Writer, filename string, data interface{}) {
	template.Must(template.ParseFiles(filename)).Execute(w, data)
}
