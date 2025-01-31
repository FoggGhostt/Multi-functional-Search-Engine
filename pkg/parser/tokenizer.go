package parser

import (
	"strings"
	"sync"
	"unicode"
)

func splitFunc(r rune) bool {
	return unicode.IsSpace(r) || unicode.IsPunct(r)
}

func Tokenize(text string, sync_map *sync.Map) {
	text = strings.ToLower(text)
	words := strings.FieldsFunc(text, splitFunc)
	for _, word := range words {
		value, is_inside := sync_map.Load(word)
		intValue, _ := value.(int)
		if is_inside {
			sync_map.Store(word, intValue+1)
		} else {
			sync_map.Store(word, 1)
		}
	}
}
