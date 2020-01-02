package main

import (
	"fmt"
	"github.com/nlopes/slack"
	"io"
)
const(
	token="xoxp-3069876617-352171753584-823349114694-8643cc74117066609bd1468e2383cbd3"
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
		for _, value := range multipartfile(r)["fileinput"] {
			if file, e := value.Open(); e == nil {
				m := base{Name: value.Filename,Area: "photo"}
				put(&m)
				postmessage(fmt.Sprintf("画像が投稿されました https://%s/image?id=%d", r.Host,m.Key.ID))
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
func postmessage(message string){
	api := slack.New(token)
	channel,clientid,error:=api.PostMessage("CCP3876JK",slack.MsgOptionText(message,false))
	logs(message,channel,clientid,error)
}