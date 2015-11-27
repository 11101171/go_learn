package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : const_test.go
// Author      : ningzhong.zeng
// Revision    : 2015-11-19 15:55:42
// Description :
//****************************************************/

import (
	"fmt"
	"go_import"
)

type State int

const (
	Running State = iota
	Stopped
	Rebooting
	Terminated
)

func (this State) String() string {
	switch this {
	case Running:
		return "Running"
	case Stopped:
		return "Stopped"
	default:
		return "Unknow"
	}
}

func main() {
	fmt.Println("Start Main func()")
	go_import.Say()

	fmt.Println("what's fuck -> ", go_import.What)

	//
	state := Stopped
	fmt.Println("state", state)
}
