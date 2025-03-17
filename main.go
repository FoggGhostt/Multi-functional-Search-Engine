package main

import (
	"fmt"
	"search-engine/pkg/indexer"
	"search-engine/pkg/parser"
)

func main() {
	err := parser.StopWordsHandle.InitializeWordMap("../pkg/parser/utils/english_stop_words.txt")
	if err != nil {
		panic("Cannot scan stop-words")
	}
	err = parser.StopWordsHandle.InitializeWordMap("../pkg/parser/utils/russian_stop_words.txt")
	if err != nil {
		panic("Cannot scan stop-words")
	}
	// parser.StopWordsHandle.WordsMap.Range(func(key, value any) bool {
	// 	fmt.Println(key)
	// 	return true
	// })

	var filepath string
	fmt.Scan(&filepath)

	// if res, err := parser.ParseFile(filepath); err == nil {
	// 	fmt.Println("okey")
	// 	res.Range(func(key, value any) bool {
	// 		fmt.Println(key, "-->", value)
	// 		return true
	// 	})
	// } else {
	// 	fmt.Println(err)
	// }
	if err := indexer.IndexFiles([]string{filepath}); err == nil {
		fmt.Println("okey")
	} else {
		fmt.Println(err)
	}
}
