package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	//"os"
)

type Result struct {
	//Status int32 `json:"status"`
	//Msg string `json:"msg"`
	Data []NodeInfo `json:"data"`
}

type NodeInfo struct {
	Ip   string `json:"ip"`
	Name string `json:"name"`
}

func (s *NodeInfo) Show() {
	//fmt.Println("\tShow\t:")
	fmt.Printf("Ip:%s", s.Ip)
	fmt.Printf("Name:%s\n", s.Name)
}

//func (r *Result) Show() {
//	fmt.Println("\tMsg:\t",r.Msg)
//}

func Get() []byte {

	c := new(http.Client)
	//req, err := http.NewRequest("GET", , nil)
	req, err := http.NewRequest("GET","xxx",nil)
	if err != nil {
		panic(err)
	}

	resp, err1 := c.Do(req)
	if err1 != nil {
		panic(err1)
	}

	defer resp.Body.Close()

	response, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		panic(err2)
	}

	return response
}

func main() {
	// URL
	// req
	// c.Do(req)
	// resp
	// ReadAll() -> []byte -> json
	// json.Unmarshal()
	// xx.xx
	data := Get()

	error2 := ioutil.WriteFile("/Users/Desktop/ipinfo", data, 0644)
	if error2 != nil {
		panic(error2)
	}

	r := new(Result)
	error1 := json.Unmarshal(data, r)
	if error1 != nil {
		panic(error1)
	}

	fmt.Println(len(r.Data))

	ipbyte := make([]byte, 0)
	for _, d := range r.Data {
		if strings.Contains(d.Ip, "/") {
			d.Show()
			//	fmt.Printf(d.Name)
			//fmt.Println(r.Msg)

			ipbyte = append(ipbyte, []byte(d.Ip + d.Name + "\n")...)

		}

	}

	err := ioutil.WriteFile("/Users/Desktop/ipinfo", ipbyte, 0644)
	if err != nil {
		panic(err)
	}

	//	r.Data[3].Show()

	//	for _, d := range r.Data {
	//		d.Show()
	//	}

	//savebyte := new(NodeInfo)

	//savebyte.Sh
	//
	//
	// ow()
}
