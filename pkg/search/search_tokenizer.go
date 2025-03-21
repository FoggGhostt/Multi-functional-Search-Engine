package search

import (
	"search-engine/pkg/parser"
	"strings"
)

func TokenizeSearchRequest(req string) ([]string, error) {
	req_text := strings.ToLower(req)
	words := strings.FieldsFunc(req_text, parser.SplitFunc)
	normalized_req := make([]string, 0)
	for _, word := range words {
		if parser.StopWordsHandle.IsStopWord(word) || word == "" {
			continue
		}
		stemmed_word, err := parser.Stem(word)
		if err != nil {
			return nil, err
		}
		normalized_req = append(normalized_req, stemmed_word)
	}
	return normalized_req, nil
}
