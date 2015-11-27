package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : engine.go
// Author      : ningzhong.zeng
// Revision    : 2015-11-16 12:02:27
// Description :
//****************************************************/

import (
	"const_main"
	"fmt"
)

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
}

func (this *Car) Start() {
	fmt.Println("Car->Start()")
}

//

type show interface {
	Show()
}

type Person struct {
	Name string
	Age  int
	//show
}

func (this *Person) Show() {
	fmt.Printf("Person's name is %s,age is %d", this.Name, this.Age)
}

type Man struct {
	Person
}

func main() {
	fmt.Println("Start Main func()")
	car := new(Car)
	car.Start()

	man := Man{Person{Name: "xiaoming"}}
	fmt.Println("My Name is %s", man.Name)
	man.Show()

	fmt.Println(const_main.win)
}
