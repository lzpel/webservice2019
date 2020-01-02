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
func bucketname() string{
	return fmt.Sprintf("%s.appspot.com",os.Getenv("PROJECT_ID"))
}

func queryobject(prefix string) *storage.ObjectIterator{
	ctx := context.Background()
	if client, err := storage.NewClient(ctx) ; err == nil{
		return client.Bucket(bucketname()).Objects(ctx, &storage.Query{Prefix:prefix})
	}
	return nil
}

func getobject(name string) (*storage.ObjectHandle,context.Context){
	ctx := context.Background()
	if client, err := storage.NewClient(ctx) ; err == nil {
		return client.Bucket(bucketname()).Object(name), ctx
	}
	return nil, nil
}

func newobjectwriter(name string) *storage.Writer{
	obj, ctx :=getobject(name)
	return obj.NewWriter(ctx)
}
func newobjectreader(name string) *storage.Reader{
	obj, ctx :=getobject(name)
	r ,err:= obj.NewReader(ctx)
	if  err != nil {
		return nil
	}
	return r
}

func getuploadurl(name string, contenttype string) string{
	//name := time.Now().Format("20060102150405")
	url, err := storage.SignedURL(bucketname(), name, &storage.SignedURLOptions{
		GoogleAccessID: os.Getenv("CLIENT_EMAIL"),
		Method:         "PUT",
		Expires:        time.Now().Add(15 * time.Minute),
		ContentType:    contenttype,
		PrivateKey:		[]byte(os.Getenv("PRIVATE_KEY")),
	})
	if err != nil {
		logf("sign: failed to sign, err = %v\n", err)
		logf("failed to sign by internal server error")
	}
	return url
}

func serve_teststorage(credential string,filename string) {
	handle("/upload",
		func(w response, r request) {
			url:=getuploadurl(filename,"")
			req, err := http.NewRequest("PUT", url, bytes.NewReader([]byte("this file is large")))
			//req.Header.Add("Content-Type", "text/plain") cause 403 NG
			client := new(http.Client)
			resp, err := client.Do(req)
			log.Println(resp, err)
		})
	handle("/write",
		func(w response, r request) {
			ow:= newobjectwriter(filename)
			ow.Write([]byte(time.Now().Format("20060102150405")))
			ow.Close()
		})
	handle("/read",
		func(w response, r request) {
			or:=newobjectreader(filename)
			data, _:=ioutil.ReadAll(or)
			or.Close()
			w.Write(data)
		})
	serve(credential)
}