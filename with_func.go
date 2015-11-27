package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : with_func.go
// Author      : ningzhong.zeng
// Revision    : 2015-11-19 20:04:45
// Description :
//****************************************************/

import (
	"fmt"
)

type Do struct {
	Name string
}

func (this Do) Lock() {
	fmt.Println(this.Name + " Lock ...")
}

func (this Do) UnLock() {
	fmt.Println(this.Name + " Unlock ~")
}

func Say() {
	do := Do{Name: "Say"}
	do.Lock()
	defer do.UnLock()
	fmt.Println("say doing")
}

func Fuck() {
	do := Do{Name: "Fuck"}
	do.Lock()
	defer do.UnLock()
	fmt.Println("fuck doing")
}

// 改造 －withContext
func withDoContext(fn func()) {
	do := Do{Name: "do"}
	do.Lock()
	defer do.UnLock()

	fn()
}

func Quit() {
	withDoContext(func() {
		fmt.Println("quit doing")
	})
}

func Reboot() {
	fmt.Println("reboot doing")
}

// 带有返回值的context

// 改造 －withContext
func withDoContextString(fn func() string) string {
	do := Do{Name: "do"}
	do.Lock()
	defer do.UnLock()

	return fn()
}

func main() {
	fmt.Println("Start Main func()")
	Say()
	Fuck()
	// 改造后
	Quit()
	withDoContext(Reboot)

	msg := withDoContextString(func() string { return "hahah" })
	fmt.Println("msg ", msg)
}
