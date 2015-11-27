package main

//****************************************************/
//Copyright(c) 2015 Tencent, all rights reserved
// File        : server_list.go
// Author      : ningzhong.zeng
// Revision    : 2015-11-19 18:32:12
// Description :
//****************************************************/

import (
	"errors"
	"fmt"
	"strings"
)

type Servers []Server

type Server struct {
	Name string
}

func ListServer() Servers {
	return []Server{
		{Name: "app1"},
		{Name: "app2"},
		{Name: "app3"},
		{Name: "app4"},
	}
}

func (servers Servers) Filter(name string) (Server, error) {
	for _, server := range servers {
		if strings.Contains(name, server.Name) {
			return server, nil
		}
	}
	return Server{}, errors.New("Not found that server")
}

func main() {
	fmt.Println("Start Main func()")
	servers := ListServer()
	if server, err := servers.Filter("app1"); err == nil {
		fmt.Println("This server is working=>", server.Name)
	} else {
		fmt.Println("This server is error=>", server.Name)
	}
}
