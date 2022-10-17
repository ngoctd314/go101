package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	os.MkdirAll("./logs/test", 0777)
	os.MkdirAll("./logs/test1", 0777)
	f, err := os.Create("./logs/test/test.txt")
	f.Close()

	os.Rename(f.Name(), "./logs/test1/teset.txt")
	fmt.Println(f.Name(), err, filepath.Base(f.Name()))
}
