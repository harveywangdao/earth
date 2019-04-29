package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"test/reflectest/pac1"
)

type AAA struct {
	A int
	a int
}

func (a AAA) name() {
	fmt.Println("AAA name", a.A, a.a)
}

func (a AAA) Name() {
	fmt.Println("AAA Name", a.A, a.a)
}

func main() {
	h := AAA{}
	k := pac1.BBB{}

	ht := reflect.TypeOf(h)
	kt := reflect.TypeOf(k)

	fmt.Println(ht)
	fmt.Println(kt)

	fmt.Println(ht.Name())
	fmt.Println(kt.Name())

	hv := reflect.ValueOf(h)
	kv := reflect.ValueOf(k)

	fmt.Println(hv)
	fmt.Println(kv)

	fmt.Println(hv.Type())
	fmt.Println(kv.Type())

	fmt.Println(hv.Kind())
	fmt.Println(kv.Kind())

	fmt.Println(hv.Interface())
	fmt.Println(kv.Interface())

	hFN := hv.NumField()
	kFN := kv.NumField()

	fmt.Println(hFN)
	fmt.Println(kFN)

	/*	for i := 0; i < hFN; i++ {
			fmt.Println(ht.Field(i), ht.Field(i).Type, hv.Field(i).Interface())
		}

		for i := 0; i < kFN; i++ {
			fmt.Println(kt.Field(i).Name, kt.Field(i).Type, kv.Field(i).Interface())
		}*/

	hNM := ht.NumMethod()
	kNM := kt.NumMethod()

	fmt.Println(hNM)
	fmt.Println(kNM)

	for i := 0; i < hNM; i++ {
		fmt.Println(ht.Method(i).Name, ht.Method(i).Type)
		hv.Method(i).Call(nil)
	}

	for i := 0; i < kNM; i++ {
		fmt.Println(kt.Method(i).Name, kt.Method(i).Type)
		kv.Method(i).Call(nil)
	}

	he := reflect.ValueOf(&h).Elem()
	ke := reflect.ValueOf(&k).Elem()

	for i := 0; i < he.NumField(); i++ {
		f := he.Field(i)
		fmt.Println(he.Type().Field(i).Name, f.Type(), f.Interface())
	}

	for i := 0; i < ke.NumField(); i++ {
		f := ke.Field(i)
		fmt.Println(ke.Type().Field(i).Name, f.Type(), f.Interface())
	}

	json.Marshal(v)

	fmt.Println("End")
}
