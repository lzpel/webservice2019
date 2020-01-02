package main

import (
	"cloud.google.com/go/datastore"
	"context"
	"log"
	"math/rand"
	"os"
	"reflect"
	"time"
)

type Key=datastore.Key
func GetClient() (*datastore.Client, context.Context) {
	ctx := context.Background()
	cli, e := datastore.NewClient(ctx, os.Getenv("PROJECT_ID"))
	if e == nil {
		return cli, ctx
	}
	log.Println(e)
	return nil, nil
}
func KeyKind() string {
	x := os.Getenv("DATASTORE_KIND")
	if x != "" {
		return x
	}
	return "base"
}
func KeyMap(i int) int {
	if 'A' <= i && i <= 'Z' {
		return i - 'A' + 00
	} else if '0' <= i && i <= '9' {
		return i - '0' + 26
	} else if 00 <= i && i < 26 {
		return i - 00 + 'A'
	} else if 26 <= i && i < 36 {
		return i - 26 + '0'
	}
	return 0
}
func KeyGen(i rune) *Key {
	return datastore.IDKey(KeyKind(), int64(rand.Uint64()<<5)|int64(KeyMap(int(i))), nil)
}
func KeyEnc(key *Key) string {
	i, f := key.ID, func(x int64, y uint) rune {
		return rune(KeyMap(int(x>>y) & 31))
	}
	return string([]rune{f(i, 00), f(i, 05), f(i, 10), f(i, 15), f(i, 20), f(i, 25), f(i, 30), f(i, 35), f(i, 40), f(i, 45), f(i, 50), f(i, 55), f(i, 60),})
}
func KeyDec(x string) *Key {
	if len(x) == 64/5+1 {
		f := func(x string, y uint) int64 {
			return int64(KeyMap(int(x[y/5]))) << y
		}
		return datastore.IDKey(KeyKind(), f(x, 00)|f(x, 05)|f(x, 10)|f(x, 15)|f(x, 20)|f(x, 25)|f(x, 30)|f(x, 35)|f(x, 40)|f(x, 45)|f(x, 50)|f(x, 55)|f(x, 60), nil)
	}
	return nil
}

type Base struct {
	Key       *Key `datastore:"__key__"`
	KeyHead   string
	TimeNew   time.Time
	TimeRenew time.Time
}

type IOHook interface {
	Save()
	Load()
}

func (m *Base) Save() {
	m.TimeRenew = time.Now()
	if m.TimeNew.IsZero() {
		m.TimeNew = m.TimeRenew
	}
	m.KeyHead = KeyEnc(m.Key)[:1]
}
func (m *Base) Load() {
}

type Model struct {
	Base
	KeyUser    *Key
	KeyHost    *Key
	Name       string
	Code       string
	Word       string
	Tags       []string
	IndexView  int
	IndexLike  int
	IndexReply int
	IndexRate  float32
	Text       string `datastore:",noindex"`
	TextLarge  string `datastore:",noindex"`
	Byte       []byte `datastore:",noindex"`
	ByteLarge  []byte `datastore:",noindex"`
}

func Put(ptr IOHook, k *Key) {
	cli, ctx := GetClient()
	if cli != nil {
		ptr.Save()
		if _, e := cli.Put(ctx, k, ptr); e != nil {
			log.Println(e)
		}
	}
}

func NewQuery() *datastore.Query {
	return datastore.NewQuery(KeyKind())
}
func GetAll(u *datastore.Query, PtrToArray interface{}) {
	cli, ctx := GetClient()
	if cli != nil {
		if _, e := cli.GetAll(ctx, u, PtrToArray); e != nil {
			log.Println(e)
		}
	}
}
func GetEntity(k *Key, Ptr interface{}) {
	cli, ctx := GetClient()
	if k != nil && cli != nil {
		if e := cli.Get(ctx, k, Ptr); e != nil {
			log.Println(e)
		}
	}
}

func GetEntities(k []*Key, DstArray interface{}) {
	cli, ctx := GetClient()
	if k != nil && cli != nil {
		if e := cli.GetMulti(ctx, k, DstArray); e != nil {
			log.Println(e)
		}
	}
}

func GetEntitiesHelper(SrcArray interface{}, DstArray interface{},conv func(i interface{})*Key) {
	va:=reflect.Indirect(reflect.ValueOf(SrcArray))
	k:=make([]*Key,va.Len())
	for i:=0;i<va.Len();i++{
		k[i]=conv(va.Index(i).Interface())
	}
	GetEntities(k,DstArray)
}

func Delete(k *Key) {
	cli, ctx := GetClient()
	if k != nil && cli != nil {
		if e := cli.Delete(ctx, k); e != nil {
			log.Println(e)
		}
	}
}
func Count(u *datastore.Query) int {
	cli, ctx := GetClient()
	if cli != nil {
		if n, e := cli.Count(ctx, u); e == nil {
			return n
		} else {
			log.Println(e)
		}
	}
	return -1
}
