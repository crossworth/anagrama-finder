package main

import (
	"io/ioutil"
	"log"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func main() {
	wordsBytes, err := ioutil.ReadFile("palavras.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make([]string, 0)

	wordsString := strings.ReplaceAll(string(wordsBytes), "\r\n", "\n")
	lines := strings.Split(wordsString, "\n")

	for _, line := range lines {
		word := unaccent(strings.ToLower(line))
		word = strings.ReplaceAll(word, "-", "")
		words = append(words, word)
	}

	input := "aergaonuoj"

	possibleWords := make([]string, 0)
	missByOneWords := make([]string, 0)

	lenInput := len(input)
	for _, word := range words {
		if len(word) != lenInput {
			continue
		}

		w := word

		for _, c := range input {
			w = strings.Replace(w, string(c), "", 1)
		}

		if len(w) == 1 {
			missByOneWords = append(missByOneWords, word)
		}

		if len(w) == 0 {
			possibleWords = append(possibleWords, word)
		}
	}

	log.Printf("Total de palavras verificadas: %d\n", len(words))

	if len(possibleWords) > 0 {
		log.Printf("Encontrado %d possíveis palavras\n", len(possibleWords))

		for _, word := range possibleWords {
			log.Printf("\t%s\n", word)
		}

		return
	}

	if len(missByOneWords) > 0 {
		log.Printf("Não foi encontrado encontrado nenhuma palavra exata, porém %d palavras com diferença de 1 letra\n", len(missByOneWords))

		for _, word := range missByOneWords {
			log.Printf("\t%s\n", word)
		}

		return
	}

	log.Printf("Não foi encontrado nenhum anagrama\n")
}

func unaccent(input string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ := transform.String(t, input)
	return s
}
