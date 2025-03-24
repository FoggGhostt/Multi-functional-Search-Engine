package main

import (
	"log"
	"os"

	// "search-engine/pkg/indexer"
	"search-engine/pkg/API"
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

	err := parser.StopWordsHandle.InitializeWordMap(english_stop_words_path)
	if err != nil {
		log.Fatalf("Cannot scan stop-words: %v %s", err, english_stop_words_path)
	}

	err = parser.StopWordsHandle.InitializeWordMap(russian_stop_words_path)
	if err != nil {
		log.Fatalf("Cannot scan stop-words: %v %s", err, russian_stop_words_path)
	}

	API.API_Init()
}
