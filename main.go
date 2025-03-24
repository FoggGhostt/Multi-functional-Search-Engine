package main

import (
	"fmt"
	"log"
	"os"
	"search-engine/pkg/indexer"
	"search-engine/pkg/parser"
	"search-engine/pkg/search"
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

	err := parser.StopWordsHandle.InitializeWordMap(english_stop_words_path)
	if err != nil {
		log.Fatalf("Cannot scan stop-words: %v %s", err, english_stop_words_path)
	}

	err = parser.StopWordsHandle.InitializeWordMap(russian_stop_words_path)
	if err != nil {
		log.Fatalf("Cannot scan stop-words: %v %s", err, russian_stop_words_path)
	}

	var filepath string

	for range 4 {
		fmt.Scan(&filepath)
		if err := indexer.IndexFiles([]string{filepath}); err == nil {
			fmt.Println("okey")
		} else {
			fmt.Println(err)
		}
	}

	tokens, err := search.Search("дом пенал кровать")
	if err != nil {
		fmt.Println(err)
	}

	if len(tokens) != 0 {
		fmt.Println(tokens[0])
	}
}
