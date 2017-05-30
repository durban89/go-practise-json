package main

import (
	"encoding/json"
	"fmt"
)

type Server struct {
	ServerName string
	ServerIP   string
}

type ServerSlice struct {
	Servers []Server
}

func DecodeJson() {
	var s ServerSlice
	str := `{
  "servers": [{
    "serverName": "Shanghai_VPN",
    "serverIP": "127.0.0.1"
  }, {
    "serverName": "Beijing_VPN",
    "serverIP": "127.0.0.2"
  }]
}`
	json.Unmarshal([]byte(str), &s)
	fmt.Println(s)
}

func EncodeJson() {
	var s ServerSlice
	s.Servers = append(s.Servers, Server{ServerName: "Shanghai_VPN", ServerIP: "127.0.0.1"})
	s.Servers = append(s.Servers, Server{ServerName: "Beijing_VPN", ServerIP: "127.0.0.2"})
	b, err := json.Marshal(s)
	if err != nil {
		fmt.Println("json err:", err)
	}

	fmt.Println(string(b))
}

func main() {
	DecodeJson()
	EncodeJson()
}
