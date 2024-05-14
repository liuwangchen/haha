package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	b, err := io.ReadAll(os.Stdin)
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
		start := strings.Index(content, "{")
		end := strings.LastIndex(content, "}")
		fmt.Println(content[start : end+1])
	}
}