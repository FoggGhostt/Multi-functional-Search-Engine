package main

import (
	"fmt"
	"log"
	"os"
	"search-engine/pkg/indexer"
	"search-engine/pkg/parser"
)

func main() {
	english_stop_words_path := os.Getenv("ENGLISH_STOP_WORDS_PATH")
	if english_stop_words_path == "" {
		english_stop_words_path = "../pkg/parser/utils/english_stop_words.txt"
	}

	russian_stop_words_path := os.Getenv("RUSSIAN_STOP_WORDS_PATH")
	if russian_stop_words_path == "" {
		russian_stop_words_path = "../pkg/parser/utils/russian_stop_words.txt"
	}

	fmt.Println(russian_stop_words_path)

	err := parser.StopWordsHandle.InitializeWordMap(english_stop_words_path)
	if err != nil {
		log.Fatalf("Cannot scan stop-words: %v %s", err, english_stop_words_path)
	}

	err = parser.StopWordsHandle.InitializeWordMap(russian_stop_words_path)
	if err != nil {
		log.Fatalf("Cannot scan stop-words: %v %s", err, russian_stop_words_path)
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
