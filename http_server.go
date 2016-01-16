package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : http_server.go
// Author      : ningzhong.zeng
// Revision    : 2015-12-1 10:17:41
// Description :
//****************************************************/

import (
	"fmt"
	"net/http"
)

// const http_root = dd

func main() {
	fmt.Println("Start Main func()")

	http.HandleFunc("/", rootHander)
	http.ListenAndServe(":9090", nil)
}

func rootHander(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "rootHandler: %s\n", r.URL.Path)
	fmt.Fprintf(w, "URL: %s\n", r.URL)
	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "RequestURI: %s\n", r.RequestURI)
	fmt.Fprintf(w, "Proto: %s\n", r.Proto)
	fmt.Fprintf(w, "HOST: %s\n", r.Host)
}
