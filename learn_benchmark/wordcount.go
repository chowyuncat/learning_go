package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

func isLetter(r rune) bool {
	if (r < 'a' || r > 'z') && (r < 'A' || r > 'Z') {
		return false
	}

	return true
}

type WordCountMap map[string]int

// performs too many  reallocations with string concat
func WordCountNaive(s string) WordCountMap {
	wc := make(WordCountMap)
	word := ""
	for _, c := range s {
		if unicode.IsSpace(c) { // strings.Fields implementation calls unicode.IsSpace
			if len(word) > 0 {
				wc[word] += 1
			}
			word = ""
		} else {
			word += string(c)
		}
	}

	if len(word) > 0 {
		wc[word] += 1
	}

	return wc
}

// @TODO: allocate only once per string
func WordCountAllocEachStringOnlyOnce(s string) WordCountMap {
	return make(WordCountMap)
}

// @TODO: use slices (views into underlying array) rather than strings
func WordCountEfficientSlices(s string) WordCountMap {
	return make(WordCountMap)
}

func WordCountWithStringFields(s string) WordCountMap {
	wc := make(WordCountMap)
	for _, word := range strings.Fields(s) {
		wc[word]++
	}
	return wc
}

func main() {

	bytes, err := ioutil.ReadFile(os.Args[2])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	text := string(bytes)

	var fp func(s string) WordCountMap

	switch os.Args[1] {
	case "naive":
		fp = WordCountNaive
	default:
		fp = WordCountWithStringFields
	}

	for i := 0; i < 1000; i++ {
		fp(text)
	}

	// fmt.Println(len(wc))
}
