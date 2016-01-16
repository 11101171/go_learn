package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : url.go
// Author      : ningzhong.zeng
// Revision    : 2016-01-14 17:22:48
// Description :
//****************************************************/

import (
	"fmt"
	"net/url"
)

func main() {
	v := url.Values{}
	v.Set("name", "Ava")
	v.Add("friend", "Jess")
	v.Add("friend", "Sarah")
	v.Add("friend", "Zoe")
	// v.Encode() == "name=Ava&friend=Jess&friend=Sarah&friend=Zoe"
	fmt.Println(v.Get("name"))
	fmt.Println(v.Get("friend"))
	fmt.Println(v["friend"])
}
