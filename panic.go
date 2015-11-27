package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : panic.go
// Author      : ningzhong.zeng
// Revision    : 2015-11-26 20:09:42
// Description :
//****************************************************/

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Start Main func()")
	panic("a problem")

	_, err := os.Create("/sss/sss")
	if err != nil {
		panic(err)
	}
}
