package main

import (
	"log"
	"regexp"
)

func recheck(reg, str string) bool {
	return regexp.MustCompile(reg).Match([]byte(str))
}
func Expression(r Request) Dict {
	var m Model
	var mt []Model
	log.Println("e1")
	GetEntity(KeyDec(CookieGet(r, "AUTH")), &m)
	log.Println("e2")
	GetAll(NewQuery().Filter("KeyHead=", "I").Order("-TimeNew"), &mt)
	log.Println("e3")
	mu:=make([]Model,len(mt))
	GetEntitiesHelper(mt, mu, func(i interface{}) *Key {
		return i.(Model).KeyUser
	})
	log.Println("e4")
	return Dict{
		"AUTH": m,
		"ITEM": mt,
		"USER": mu,
	}
}
func main() {
	HandleFile()
	Handle(`/signup/`, func(w Response, r Request) {
		if r.Method == "GET" {
			Writef(w, Expression(r), "base.html", "signup.html")
		} else {
			r.ParseForm()
			m := Model{
				Name: r.Form.Get("name"),
				Code: r.Form.Get("code"),
				Word: r.Form.Get("word"),
			}
			if (recheck(`.+`, m.Name) && recheck(`^[a-zA-Z0-9.!#$%&'*+\/=?^_{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$`, m.Code) && recheck(`\w{6,}`, m.Word)) == false {
				Redirect(w, r, "/signup/?e=wrong")
			} else if Count(NewQuery().Filter("KeyHead=", "U").Filter("Code=", m.Code)) > 0 {
				Redirect(w, r, "/signup/?e=exist")
			} else {
				m.Key = KeyGen('U')
				Put(&m, m.Key)
				CookieSet(w, "AUTH", KeyEnc(m.Key), 100)
				Redirect(w,r,"/")
			}
		}
	})
	Handle(`/signin/`, func(w Response, r Request) {
		if r.Method == "GET" {
			Writef(w, Expression(r), "base.html", "signin.html")
		} else {
			r.ParseForm()
			m := []Model{}
			GetAll(NewQuery().Filter("KeyHead=", "U").Filter("Code=", r.Form.Get("code")).Filter("Word=", r.Form.Get("word")), &m)
			if len(m) > 0 {
				CookieSet(w, "AUTH", KeyEnc(m[0].Key), 100)
				Redirect(w,r,"/")
			} else {
				Redirect(w, r, "/signin?e=True")
			}
		}
	})
	Handle(`/signout/`, func(w Response, r Request) {
		CookieSet(w, "AUTH", "", -100)
		Redirect(w, r, "/")
	})
	Handle(`/delete/`, func(w Response, r Request) {
		if r.Method == "GET" {
			Writef(w, Expression(r), "base.html", "delete.html")
		} else {
			Delete(KeyDec(CookieGet(r, "AUTH")))
			Redirect(w, r, "/")
		}
	})
	Handle(`/itemadd/`, func(w Response, r Request) {
		exp:=Expression(r)
		if r.Method == "GET" {
			Writef(w, exp, "base.html", "itemadd.html")
		} else if exp["AUTH"].(Model).Key!=nil{
			r.ParseForm()
			m := Model{
				Text: r.Form.Get("text"),
				KeyUser:exp["AUTH"].(Model).Key,
			}
			m.Key = KeyGen('I')
			Put(&m, m.Key)
			Redirect(w, r, "/")
		}
	})
	Handle(`/`, func(w Response, r Request) {
		log.Println("1")
		if r.URL.Path == "/" {
			log.Println("2")
			x:=Expression(r)
			log.Println("3")
			Writef(w, x, "base.html", "index.html")
			log.Println("4")
		}
	})
	Environment("crediantial.json", "base")
	Listen()
}
