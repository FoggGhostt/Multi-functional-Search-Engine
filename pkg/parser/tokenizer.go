package parser

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
)

func splitFunc(r rune) bool {
	return unicode.IsSpace(r) || unicode.IsPunct(r)
}

func Tokenize(text string, sync_map *sync.Map) error {
	text = strings.ToLower(text)
	words := strings.FieldsFunc(text, splitFunc)
	for _, word := range words {
		if StopWordsHandle.IsStopWord(word) {
			continue
		}
		value, is_inside := sync_map.Load(word)
		if is_inside {
			intValue, isCorrectType := value.(int64)
			if !isCorrectType {
				return fmt.Errorf("incorrect type in sync_map")
			}
			sync_map.Store(word, intValue+1)
		} else {
			sync_map.Store(word, int64(1))
		}
	}
	return nil
}
