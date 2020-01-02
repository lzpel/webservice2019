package main

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)
func BucketName() string{
	return fmt.Sprintf("%s.appspot.com",os.Getenv("PROJECT_ID"))
}

func queryobject(prefix string) *storage.ObjectIterator{
	ctx := context.Background()
	if client, err := storage.NewClient(ctx) ; err == nil{
		return client.Bucket(BucketName()).Objects(ctx, &storage.Query{Prefix: prefix})
	}
	return nil
}

func GetObject(name string) (*storage.ObjectHandle,context.Context){
	ctx := context.Background()
	if client, err := storage.NewClient(ctx) ; err == nil {
		return client.Bucket(BucketName()).Object(name), ctx
	}
	return nil, nil
}

func NewObjectWriter(name string) *storage.Writer{
	obj, ctx := GetObject(name)
	return obj.NewWriter(ctx)
}
func NewObjectReader(name string) *storage.Reader{
	obj, ctx := GetObject(name)
	r ,_:= obj.NewReader(ctx)
	return r
}
func NewObjectRangeReader(name string,offset,length int64) *storage.Reader{
	obj, ctx := GetObject(name)
	r ,_:= obj.NewRangeReader(ctx, offset,length)
	return r
}

func SignedURL(Method, Name, ContentType string) string{
	url, e := storage.SignedURL(BucketName(), Name, &storage.SignedURLOptions{
		GoogleAccessID: os.Getenv("CLIENT_EMAIL"),
		Method:         Method,
		Expires:        time.Now().Add(15 * time.Minute),
		//ContentType:    ContentType,
		PrivateKey:		[]byte(os.Getenv("PRIVATE_KEY")),
	})
	if e != nil {
		log.Printf("sign: failed to sign, err = %v\n", e)
	}
	return url
}

func TestStorage(credential string,filename string) {
	Handle("/upload",
		func(w Response, r Request) {
			url:=SignedURL("POST",filename,"")
			req, err := http.NewRequest("PUT", url, bytes.NewReader([]byte("this file is large")))
			//req.Header.Add("Content-Type", "text/plain") cause 403 NG
			client := new(http.Client)
			resp, err := client.Do(req)
			log.Println(resp, err)
		})
	Handle("/write",
		func(w Response, r Request) {
			ow:= NewObjectWriter(filename)
			ow.Write([]byte(time.Now().Format("20060102150405")))
			ow.Close()
		})
	Handle("/read",
		func(w Response, r Request) {
			or:= NewObjectReader(filename)
			data, _:=ioutil.ReadAll(or)
			or.Close()
			w.Write(data)
		})
	Environment(credential,"")
	Listen()
}