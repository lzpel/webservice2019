package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

func main() {
	//順番に注意
	//http://localhost:8080/music?id=5714772351778816
	handle("/public", func(w response, r request){
		servefile(w,r,r.URL.Path[1:])
	})
	handle("/upload", func(w response, r request) {
		for _, value := range multipartfile(r)["fileinput"] {
			if file, e := value.Open(); e == nil {
				m := base{Name: value.Filename, Area: "music"}
				put(&m)
				tempdir, err := ioutil.TempDir("", "beroringa")
				if err != nil {
					logfatal(err)
				}
				defer os.RemoveAll(tempdir)
				fu, err := os.Create(filepath.Join(tempdir, "i"))
				if err != nil {
					logfatal(err)
				}
				defer fu.Close()
				io.Copy(fu, file)
				cmd := exec.Command("ffmpeg", "-i", filepath.Join(tempdir, "i"), "-f", "ogg", filepath.Join(tempdir, "o"))
				var out bytes.Buffer
				var stderr bytes.Buffer
				cmd.Stdout = &out
				cmd.Stderr = &stderr
				err = cmd.Run()
				if err != nil {
					logs(err.Error() + ": " + stderr.String())
					return
				}
				logs("Result: " + out.String())
				fi, _ := os.Open(filepath.Join(tempdir, "o"))
				defer fi.Close()
				ow := newobjectwriter(fmt.Sprintf("%d", m.Key.ID))
				defer ow.Close()
				io.Copy(ow, fi)
			}
		}
		redirect(w, r, "/")
	})
	handle("/", func(w response, r request) {
		logs(r.URL.Path)
		if r.URL.Path == "/" {
			var data []base
			q := querydata("base").Filter("Area=", "music").Order("-TimeBorn").Limit(5)
			getall(q, &data)
			writetemplate(
				w,
				"index.html",
				dict{
					"music": data,
				},
				template.FuncMap{
					"signedurl": func(id int64) string {
						url:=signedurl("GET", fmt.Sprintf("%d",id), "")
						logs(id,url)
						return url
					},
				},
			)
		}
	})
	listen("crediantial.json")
}
