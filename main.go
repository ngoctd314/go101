package main

import (
	"encoding/json"
	"os"
)

func main() {
	jsonn()
}

// Person ...
type Person struct {
	Name    string
	Age     int
	Address string
}

func jsonn() {
	p := Person{
		Name:    "TDN",
		Age:     22,
		Address: "Ha Noi",
	}
	listPerson := []Person{}
	for i := 0; i < 10000; i++ {
		listPerson = append(listPerson, p)
	}
	data, _ := json.Marshal(listPerson)
	os.WriteFile("data.json", data, 0600)
}
