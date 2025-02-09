package parser

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"

	"github.com/kljensen/snowball"
)

func stem(word string) (string, error) {
	r, _ := utf8.DecodeRuneInString(word)
	if unicode.Is(unicode.Latin, r) {
		return snowball.Stem(word, "english", false)
	} else {
		return snowball.Stem(word, "russian", false)
	}
}

func splitFunc(r rune) bool {
	return unicode.IsSpace(r) || unicode.IsPunct(r)
}

func Tokenize(text string, sync_map *sync.Map) error {
	text = strings.ToLower(text)
	words := strings.FieldsFunc(text, splitFunc)
	for _, word := range words {
		if StopWordsHandle.IsStopWord(word) || word == "" {
			continue
		}
		value, is_inside := sync_map.Load(word)
		stemmed_word, err := stem(word)
		if err != nil {
			return err
		}
		if is_inside {
			intValue, isCorrectType := value.(int64)
			if !isCorrectType {
				return fmt.Errorf("incorrect type in sync_map")
			}
			sync_map.Store(stemmed_word, intValue+1)
		} else {
			sync_map.Store(stemmed_word, int64(1))
		}
	}
	return nil
}
