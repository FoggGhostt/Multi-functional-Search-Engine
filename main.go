package main

import (
	"fmt"
	"search-engine/pkg/parser"
)

func main() {
	var filepath string
	fmt.Scan(&filepath)
	if res, err := parser.ParseFile(filepath); err == nil {
		res.Range(func(key, value any) bool {
			fmt.Println(key, "-->", value)
			return true
		})
	}
}
