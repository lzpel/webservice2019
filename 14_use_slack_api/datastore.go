package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"os"
	"time"
)
type item interface{
	kind() string
	keyget() *datastore.Key
	keyset(*datastore.Key)
}

type base struct {
	//大文字=public::小文字=private
	//indexed content
	Key *datastore.Key `datastore:"__key__"`
	TimeBorn time.Time
	TimeUpdate time.Time
	Area string
	Name string
	Tags []string
	IndexView int
	IndexLike int
	IndexReply int
	IndexRate float32
	//not indexed content
	Text string `datastore:",noindex"`
	TextLarge string `datastore:",noindex"`
	Byte []byte `datastore:",noindex"`
	ByteLarge []byte `datastore:",noindex"`
}

func (u *base) kind() string{
	return "base"
}

func (u *base) keyget() *datastore.Key{
	return u.Key
}

func (u *base) keyset(k*datastore.Key) {
	u.TimeUpdate=time.Now()
	if(u.Key==nil){
		u.TimeBorn=u.TimeUpdate
	}
	u.Key=k
}

func put(ptr item){
	ctx := context.Background()
	cli, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if(err != nil){
		logf("new client error")
	}else{
		if(ptr.keyget()==nil){
			k:=datastore.IncompleteKey(ptr.kind(),nil)
			ptr.keyset(k)
		}
		k, err := cli.Put(ctx, ptr.keyget(), ptr)
		if err != nil {
			logf("put error")
		}
		ptr.keyset(k)
	}
}

func querydata(kind string) *datastore.Query{
	return datastore.NewQuery(kind)
}
func getall(u *datastore.Query, ptrtoarray interface{}){
	ctx := context.Background()
	cli, err := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if(err != nil){
		logf(err.Error())
	}else {
		if _, err := cli.GetAll(ctx, u, ptrtoarray); err != nil {
			logf(err.Error())
		}
	}
}

/*
- Signed integers (int, int8, int16, int32 and int64)
- bool
- string
- float32 and float64
- []byte (up to 1 megabyte in length)
- Any type whose underlying type is one of the above predeclared types
- *Key
- GeoPoint
- time.Time (stored with microsecond precision, retrieved as local time)
- Structs whose fields are all valid value types
- Pointers to structs whose fields are all valid value types
- Slices of any of the above
- Pointers to a signed integer, bool, string, float32, or float64

- omitempty:0を保存しない
- noindex:インデックスしない
- flatten:部分構造についてネスト構造を平滑化する

A and B are renamed to a and b.
A, C and J are not indexed.
D's tag is equivalent to having no tag at all (E).
I is ignored entirely by the datastore.
J has tag information for both the datastore and json packages.
type TaggedStruct struct {
	A int `datastore:"a,noindex"`
	B int `datastore:"b"`
	C int `datastore:",noindex"`
	D int `datastore:""`
	E int
	I int `datastore:"-"`
	J int `datastore:",noindex" json:"j"`
}
*/