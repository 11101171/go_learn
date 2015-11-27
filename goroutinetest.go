package main

import (
	"fmt"
	//"runtime"
	"time"
)

var c chan int

func ready(w string, sec int) {
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Println(w, "is ready")
	c <- 1
}

func main() {
	//runtime.GOMAXPROCS(2)

	c = make(chan int)

	go ready("Tea", 2)
	go ready("Coffee", 1)
	fmt.Println("i'am waiting, but too long")
	<-c
	<-c
	fmt.Println("end...")
	// var i int = 0
	// for {

	// 	select {
	// 	case <-c:
	// 		i++
	// 		if i > 0 {
	// 			break
	// 		}
	// 	}
	// }
}
