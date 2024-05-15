package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	b, err := ioutil.ReadFile("/tmp/json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		err := recover()
		if err != nil {
			fmt.Print(string(b))
		}
	}()
	for _, content := range strings.Split(string(b), "\n") {
		if len(content) == 0 {
			continue
		}
		start := strings.Index(content, "{")
		end := strings.LastIndex(content, "}")
		fmt.Println(content[start : end+1])
	}
}
