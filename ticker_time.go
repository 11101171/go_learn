package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : ticker_time.go
// Author      : ningzhong.zeng
// Revision    : 2015-11-19 22:32:18
// Description :
//****************************************************/

import (
	"fmt"
	// "strconv"
	"time"
)

func Print(text string) {
	go func() {
		ticker := time.NewTicker(time.Second * 3)
		ch := make(chan bool, 1)
		for {
			select {
			case <-ticker.C:
				fmt.Printf("hello, I'am %s\n", text)
			case <-ch:
				return
			}
		}
		// for _ = range ticker.C {
		// fmt.Printf("Ticker at %v\n", time.Now())
		// fmt.Printf("hello, I'am %s, time %v\n", text, time.Now())
		// }
		close(ch)
		fmt.Println("close ", text)
	}()

	fmt.Println("Out close ", text)
	// time.Sleep(5 * time.Second)
}

func main() {
	// i := 1
	// for {
	// select {
	// default:
	// return
	// }
	// fmt.Println("for ", strconv.Itoa(i))
	// i++
	// }

	go func() {
		ticker := time.NewTicker(time.Second * 3)
		defer ticker.Stop()
		ch := make(chan bool)
		for {
			select {
			case <-ticker.C:
				fmt.Printf("hello, I'am ")
				// ch <- true
				// return
			case <-ch:
				return
			}
		}
		// for _ = range ticker.C {
		// fmt.Printf("Ticker at %v\n", time.Now())
		// fmt.Printf("hello, I'am %s, time %v\n", text, time.Now())
		// }
		close(ch)
		fmt.Println("close ")
	}()
	// for i := 0; i < 10; i++ {
	// Print(strconv.Itoa(i))
	// }

	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
