package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : stringFunc.go
// Author      : ningzhong.zeng
// Revision    : 2015-11-16 16:52:50
// Description :
//****************************************************/

import (
	"errors"
	"fmt"
	"reflect"
)

func foor() {
	fmt.Println("Start->foor()")
}

func say(number int) {
	fmt.Printf("This text is %d", number)
}

func Call(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		err = errors.New("The number of params is not adapted.")
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

type T struct {
	A int
	B string
}
type S struct {
	F string `species:"gopher" color:"blue"`
}

func main() {
	fmt.Println("Start Main func()")

	// map 直接调用函数
	funcs := map[string]func(){"foor": foor}
	funcs["foor"]()
	// 反射
	xfuncs := map[string]interface{}{"say": say}
	Call(xfuncs, "say", 123)

	fmt.Println()
	fmt.Println()
	fmt.Println()
	var x int = 1
	fmt.Println("type: ", reflect.TypeOf(x))
	fmt.Println("value: ", reflect.ValueOf(x))
	fmt.Println("Kind:  ", reflect.ValueOf(x).Kind())
	// fmt.Println("Kind is Int? ", reflect.ValueOf(x).Kind() == reflect.int)

	t := T{12, "skidoo"}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d %s %s = %v\n", i, typeOfT.Field(i).Name, f.Type(), f.Interface())
	}

	sv := S{}
	st := reflect.TypeOf(sv)
	field := st.Field(0)
	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))

}
