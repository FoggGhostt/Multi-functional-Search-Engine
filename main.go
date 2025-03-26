package main

import (
	"log"

	// "search-engine/pkg/indexer"

	"search-engine/pkg/API"
	"search-engine/pkg/config"
	"search-engine/pkg/parser"
)

func main() {
	cnfg, err := config.GetConfig()
	if err != nil {
		log.Fatalf("Cannot read config")
	}

	english_stop_words_path := cnfg.SearchEngineConfig.EnglishStopWordsPath

	russian_stop_words_path := cnfg.SearchEngineConfig.RussianStopWordsPath

	err = parser.StopWordsHandle.InitializeWordMap(english_stop_words_path)
	if err != nil {
		log.Fatalf("Cannot scan stop-words: %v %s", err, english_stop_words_path)
	}

	err = parser.StopWordsHandle.InitializeWordMap(russian_stop_words_path)
	if err != nil {
		log.Fatalf("Cannot scan stop-words: %v %s", err, russian_stop_words_path)
	}

	API.API_Init()
}
