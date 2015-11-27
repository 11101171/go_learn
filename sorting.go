package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : sorting.go
// Author      : ningzhong.zeng
// Revision    : 2015-11-26 19:58:56
// Description :
//****************************************************/

import (
	"fmt"
	"sort"
)

// 自定义func排序
type ByLength []string

func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fmt.Println("Start Main func()")
	strs := []string{"a", "d", "l", "p"}
	sort.Strings(strs)
	fmt.Println("strs", strs)

	ints := []int{3, 6, 8, 5}
	sort.Ints(ints)
	fmt.Println("ints", ints)

	// 根据func排序
	fruits := []string{"ba", "pe", "ki"}
	sort.Sort(ByLength(fruits))
	fmt.Println("fruits", fruits)
}
