package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : same_key_diff_value.go
// Author      : ningzhong.zeng
// Revision    : 2015-11-19 22:22:30
// Description :
//****************************************************/

import (
	"fmt"
)

type User struct {
	location string
	age      int
}

func (user User) String() string {
	return fmt.Sprintf("(%s, %d)", user.location, user.age)
}

func main() {
	fmt.Println("Start Main func()")
	users := make(map[*User]string, 0)
	users[&User{location: "HangZhou", age: 10}] = "zhe'jiang"
	users[&User{location: "WuHang", age: 20}] = "hu'bei"
	users[&User{location: "NanNing", age: 30}] = "guang'xi"
	users[&User{location: "GuangZhouo", age: 40}] = "guang'dong"
	fmt.Println("map=>", users)
	
	maps := map[(HangZhou, 10):zheJiang (WuHang, 20):hubei]
	fmt.Println("maps=>", maps)
}
