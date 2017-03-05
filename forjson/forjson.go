package main

import "fmt"

/*
import (
	"encoding/json"
	"fmt"
)

type newinfo struct {
	Clothes string `json:"clothes"`
}

type info struct {
	Name  string   `json:"name"`
	Age   int      `json:"age"`
	Class []string `json:"class"`
	Cloth newinfo  `json:"cloth"`
}

func (m *info) Show() {
	fmt.Println("Show info:")
	fmt.Println("\tName\t:", m.Name)
	fmt.Println("\tAge\t:", m.Age)
	fmt.Println("\tClass\t:", m.Class)
	fmt.Println("\tCloth\t:", m.Cloth.Clothes)
}

func main() {
	m := &info{
		Name:  "mush",
		Age:   25,
		Class: []string{"english", "math"},
		Cloth: newinfo{Clothes: "red"},
	}

	m.Show()

	data, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(data))

	savebyte := new(info)
	err1 := json.Unmarshal(data, savebyte)
	if err != nil {

		panic(err1)
	}

	savebyte.Show()
}
*/