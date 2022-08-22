package main

import (
	"fmt"
	"os"
	"testing"

	jsoniter "github.com/json-iterator/go"
)

// 176	   6799224 ns/op	  603208 B/op	   30010 allocs/op

func BenchmarkJson(b *testing.B) {
	listPerson := []Person{}
	for i := 0; i < b.N; i++ {
		data, _ := os.ReadFile("data.json")

		jsoniter.Unmarshal(data, &listPerson)
		// jsoniter.Unmarshal(data, &listPerson)
	}
	fmt.Println("len list person: ", len(listPerson))
}
