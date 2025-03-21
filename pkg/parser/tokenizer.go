package parser

import (
	"fmt"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"

	"github.com/kljensen/snowball"
)

var mtx sync.Mutex

func Stem(word string) (string, error) {
	r, _ := utf8.DecodeRuneInString(word)
	if unicode.Is(unicode.Latin, r) {
		return snowball.Stem(word, "english", false)
	} else {
		return snowball.Stem(word, "russian", false)
	}
}

func SplitFunc(r rune) bool {
	return unicode.IsSpace(r) || unicode.IsPunct(r)
}

func Tokenize(text string, sync_map *sync.Map) error {
	text = strings.ToLower(text)
	words := strings.FieldsFunc(text, SplitFunc)
	for _, word := range words {
		if StopWordsHandle.IsStopWord(word) || word == "" {
			continue
		}
		stemmed_word, err := Stem(word)
		if err != nil {
			return err
		}
		value, is_inside := sync_map.Load(stemmed_word)
		if is_inside {
			intValue, isCorrectType := value.(int64)
			if !isCorrectType {
				return fmt.Errorf("incorrect type in sync_map")
			}
			for !sync_map.CompareAndSwap(stemmed_word, intValue, intValue+1) {
				value, _ = sync_map.Load(stemmed_word)
				intValue, _ = value.(int64)
			}
		} else {
			mtx.Lock()
			value, is_inside = sync_map.Load(stemmed_word)
			intValue := 0
			if is_inside {
				intVal, isCorrectType := value.(int64)
				if !isCorrectType {
					return fmt.Errorf("incorrect type in sync_map")
				}
				intValue = int(intVal)
			}
			sync_map.Store(stemmed_word, int64(intValue+1))
			mtx.Unlock()
		}
	}
	return nil
}
