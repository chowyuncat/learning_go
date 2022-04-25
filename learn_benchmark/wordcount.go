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

// performs too many reallocations with string concat
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
// [ a bc  de ]
//  0123456789
// [1:2] => 'a'
// [3:5] => 'bc'
// [6:6] => ''
// [7:9] => 'de'

// [a bc  de ]
//  0123456789
// [1:1] => ''
// [2:4] => 'bc'
// [5:5] => ''
// [6:8] => 'de'
// [9:9] => ''
func WordCountAllocEachStringOnlyOnce(s string) WordCountMap {
	wc := make(WordCountMap)
	last_space_idx := -1
	i := 0
	var c rune
	for i, c = range s {
		if unicode.IsSpace(c) { // strings.Fields implementation calls unicode.IsSpace
			if last_space_idx + 1 != i {
				word := s[last_space_idx + 1:i]
				// fmt.Printf("[%d:%d] => \"%s\"\n", last_space_idx + 1, i, word)
				wc[word]++
			}
			last_space_idx = i
		} else {
			// fmt.Printf("[%d] => '%c'\n", i, c)
		}
	}


	if last_space_idx != i + 1 {
		word := s[last_space_idx + 1:i + 1]
		wc[word] += 1
	}

	return wc
}

func WordCountWithStringsFields(s string) WordCountMap {
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
		os.Exit(1)
	}

	text := string(bytes)

	var fp func(s string) WordCountMap

	switch os.Args[1] {
	case "naive":
		fp = WordCountNaive
	case "fields":
		fp = WordCountWithStringsFields
	case "wordalloc":
		fp = WordCountAllocEachStringOnlyOnce
	default:
		fmt.Println("implementation", os.Args[1], "unknown")
		os.Exit(1)
	}

	for i := 0; i < 1; i++ {
		wc := fp(text)
		if i == 0 {
			fmt.Println(len(wc))
			fmt.Println(wc)
		}
	}

}
