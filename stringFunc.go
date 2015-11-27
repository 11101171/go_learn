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

func main() {
	fmt.Println("Start Main func()")
	// map 直接调用函数
	foos := map[string]func(){"foor": foor}
	funcs["foor"]()
	// 反射
	xfuncs := map[string]interface{}{"say": say}
	Call(xfuncs, "say", 123)

}
