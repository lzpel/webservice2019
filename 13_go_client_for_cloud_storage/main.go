package main

import (
	"fmt"
	"io"
	"log"
)

func main() {
	handle("/", func(w response, r request) {
		var data []base
		q := querydata("base").Filter("Area=", "photo").Order("-TimeBorn").Limit(100)
		getall(q, &data)
		writetemplate(w, "index.html", dict{
			"photos": data,
		})
	})
	handle("/image", func(w response, r request) {
		id := r.URL.Query().Get("id")
		if or := newobjectreader(id); or != nil {
			io.Copy(w, or)
			or.Close()
		}
	})
	handle("/upload", func(w response, r request) {
		err := r.ParseMultipartForm(200000)
		if err != nil {
			log.Fatal(err)
			return
		}
		// loop through the files one by one
		for _, v := range r.MultipartForm.File["fileinput"] {
			if file, e := v.Open(); e == nil {
				m := base{
					Name: v.Filename,
					Area: "photo",
				}
				put(&m)
				ow := newobjectwriter(fmt.Sprintf("%d", m.Key.ID))
				io.Copy(ow,file)
				file.Close()
				ow.Close()
			}
		}
		redirect(w, r, "/")
	})
	serve("crediantial.json")
}