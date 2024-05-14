package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		return
	}
	b, err := io.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		return
	}
	queryStr := os.Args[1]
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
		raw := json.RawMessage(content)
		value, err := findFirstValue(queryStr, raw)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println(value)
	}
}
func findFirstValue(queryKey string, raw json.RawMessage) (string, error) {
	var asMap map[string]*json.RawMessage
	if err := json.Unmarshal(raw, &asMap); err == nil {
		for key, val := range asMap {
			if strings.EqualFold(key, queryKey) {
				return string(*val), nil // 直接返回原始 JSON 字符串
			}

			value, err := findFirstValue(queryKey, *val)
			if err != nil {
				return "", err
			}
			if value != "" {
				return value, nil
			}
		}
		return "", nil
	}

	var asSlice []*json.RawMessage
	if err := json.Unmarshal(raw, &asSlice); err == nil {
		for _, val := range asSlice {
			value, err := findFirstValue(queryKey, *val)
			if err != nil {
				return "", err
			}
			if value != "" {
				return value, nil
			}
		}
	}

	return "", nil
}
