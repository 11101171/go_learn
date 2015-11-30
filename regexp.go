package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : regexp.go
// Author      : ningzhong.zeng
// Revision    : 2015-11-28 10:47:34
// Description :
//****************************************************/

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Start Main func()")
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(r.MatchString("peach"))

	fmt.Println(r.FindString("peach plkakl"))
}
