package main

import (
	"fmt"
)

func main() {
	// ayyay
	fmt.Printf("ss")
	list := []string{"a", "b", "c", "d", "e", "f"}
	for k, v := range list {
		k1 := k
		fmt.Println(k1)
		fmt.Println(v)
		// fmt.Println("k=>" + k + ",v=>" + v)
	}

	// slice
	var array [100]int   // Create array, index from 0 to 99 4
	slice := array[0:99] // Create slice, index from 0 to 98

	var slice2 = make([]int, 6)
	s0 := []int{0, 0}
	s1 := append(s0, 2)
	n1 := copy(s, a[0:])

	// map
	monthdays := map[string]int{
		"Jan": 31, "Feb": 28, "Mar": 31,
		"Apr": 30, "May": 31, "Jun": 30,
		"Jul": 31, "Aug": 31, "Sep": 30,
		"Oct": 31, "Nov": 30, "Dec": 31, //← 逗号是必须的
	}
	monthdays["Undecim"] = 30 // ←  添加一个月
	monthdays["Feb"] = 29     // ←  重写

	year := 0
	for _, days := range monthdays { // ← 键没有使用，因此用 _, days
		year += day // ← 试试 //
	}

	var v1 int              // int v1 = 0;
	var v2 *int             // Integer v2 = null;
	var v3 string           // String v3 = "";
	var v4 [10]int          //  int[] v4 = new int[10]; // v4 is a value in Go.
	var v5 []int            // int[] v5 = null;
	var v6 *struct{ a int } //  C v6 = null; // Given: class C { int a; }
	var v7 map[string]int   // HashMap<String,Integer> v7 = null;
	var v8 func(a int) int  //  F v8 = null; // interface F { int f(int a); }
}
