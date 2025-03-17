package parser

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type StopWordsHandleStruct struct {
	WordsMap sync.Map
}

func (wordsHandle *StopWordsHandleStruct) InitializeWordMap(filePath string) error {
	does_exists, err := searchFile(filePath)
	if err != nil {
		return fmt.Errorf("no such file")
	}
	if !does_exists {
		return fmt.Errorf("file %s didn't found", filePath)
	}
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("cant open the file, %s", filePath)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		wordsHandle.WordsMap.Store(line, true)
	}
	return nil
}

func (wordsHandle *StopWordsHandleStruct) IsStopWord(word string) bool {
	_, is_inside := wordsHandle.WordsMap.Load(word)
	if is_inside {
		return true
	} else {
		return false
	}
}

var StopWordsHandle StopWordsHandleStruct
