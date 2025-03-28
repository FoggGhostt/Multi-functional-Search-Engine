package indexer

import (
	"context"
	"log"
	"os"
	"search-engine/pkg/mongodb"
	"search-engine/pkg/parser"
	"testing"
)

func TestIndexFiles(t *testing.T) {
	os.Setenv("CONFIG_PATH", "../../config.yaml")

	english_stop_words_path := "../parser/utils/english_stop_words.txt"

	russian_stop_words_path := "../parser/utils/russian_stop_words.txt"

	err := parser.StopWordsHandle.InitializeWordMap(english_stop_words_path)
	if err != nil {
		log.Fatalf("Cannot scan stop-words: %v %s", err, english_stop_words_path)
	}

	err = parser.StopWordsHandle.InitializeWordMap(russian_stop_words_path)
	if err != nil {
		log.Fatalf("Cannot scan stop-words: %v %s", err, russian_stop_words_path)
	}
	tests := []struct {
		name     string
		filepath string
	}{
		{"file 1", "../../search_test/1.txt"},
		{"file 2", "../../search_test/2.txt"},
		{"file 3", "../../search_test/3.txt"},
		{"file 4", "../../search_test/4.txt"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			syncMapPtr, err := parser.ParseFile(test.filepath)
			if err != nil {
				t.Error("Error of Parser function")
			}
			err = IndexFiles([]string{test.filepath})
			if err != nil {
				t.Errorf("Error of testing function: %s", err.Error())
			}
			db, err := mongodb.GetDB()
			if err != nil {
				t.Errorf("Error of mongodb function: %s", err.Error())
			}
			docInfo, err := db.GetFileIndex(context.Background(), test.filepath)
			if err != nil {
				t.Errorf("Error of mongodb function: %s", err.Error())
			}
			for _, tokenInfo := range docInfo.Tokens {
				fileCount, ok := syncMapPtr.Load(tokenInfo.Token)
				if !ok {
					t.Errorf("Token: %s has not be found in DB", tokenInfo.Token)
				}
				intFileCount, isCorrectType := fileCount.(int64)
				if !isCorrectType {
					t.Error("Incorrect type in syncMap")
				}
				if intFileCount != tokenInfo.OccureCount {
					t.Errorf("Incorrect Token count, expected: %d, recieved: %d", intFileCount, tokenInfo.OccureCount)
				}
			}
		})
	}
}
