package parser

import (
	"log"
	"sync"
	"testing"
)

func TestStem(t *testing.T) {
	tests := []struct {
		name     string
		words    []string
		expected []string
	}{
		{"Long English words", []string{"Pronunciation", "Conscientiousness", "Prioritisation"}, []string{"pronunci", "conscienti", "prioritis"}},
		{"Long Russian words", []string{"Продолжительность", "Неподвижность", "Домостроение", "Индустриализация", "Пуританство", "Необъятный"}, []string{"продолжительн", "неподвижн", "домостроен", "индустриализац", "пуританств", "необъятн"}},
		{"Short English words", []string{"House", "River", "Genious", "Long"}, []string{"hous", "river", "genious", "long"}},
		{"Short Russian words", []string{"Ручка", "Кровать", "Пенал", "Часы", "Помада"}, []string{"ручк", "крова", "пена", "час", "помад"}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			for i, word := range test.words {
				res, err := Stem(word)
				if err != nil {
					t.Error("Error of testing function")
				}
				if res != test.expected[i] {
					t.Errorf("Expected value: %s, recieved word: %s", test.expected[i], res)
				}
			}
		})
	}
}

func TestTokenize(t *testing.T) {
	english_stop_words_path := "./utils/english_stop_words.txt"

	russian_stop_words_path := "./utils/russian_stop_words.txt"

	err := StopWordsHandle.InitializeWordMap(english_stop_words_path)
	if err != nil {
		log.Fatalf("Cannot scan stop-words: %v %s", err, english_stop_words_path)
	}

	err = StopWordsHandle.InitializeWordMap(russian_stop_words_path)
	if err != nil {
		log.Fatalf("Cannot scan stop-words: %v %s", err, russian_stop_words_path)
	}
	tests := []struct {
		name     string
		text     string
		expected map[string]int64
	}{
		{
			"Text 1",
			"Шла Саша по шоссе и СоСаЛа Сушку сушка вкусную такую, не могла она остановиться саша",
			map[string]int64{
				"вкусн":   1,
				"останов": 1,
				"саш":     2,
				"соса":    1,
				"сушк":    2,
				"шла":     1,
				"шосс":    1,
			},
		},
		{
			"Text 2",
			"What are you doing tonight? I want to meet with yon tonIghT on Pokrovski belveder 11 in class with even number",
			map[string]int64{
				"belved":    1,
				"class":     1,
				"even":      1,
				"meet":      1,
				"number":    1,
				"pokrovski": 1,
				"tonight":   2,
				"want":      1,
				"yon":       1,
				"11":        1,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var tokens sync.Map
			err := Tokenize(test.text, &tokens)
			if err != nil {
				t.Errorf("Ошибка токенизации: %v", err)
			}

			var AreEqual bool = true
			result := make(map[string]int64)
			tokens.Range(func(key, value any) bool {
				token, ok1 := key.(string)
				count, ok2 := value.(int64)
				if !ok1 || !ok2 {
					t.Errorf("Incorrect data type in sync.Map: key %T, value %T", key, value)
					return false
				}
				if count != test.expected[token] {
					AreEqual = false
				}
				result[token] = count
				return true
			})

			if !AreEqual {
				t.Errorf("Ожидалось %v, получено %v", test.expected, result)
			}
		})
	}
}
