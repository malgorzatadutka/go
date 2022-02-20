package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func CountSign(path string, sign string) {
	content, err := ioutil.ReadFile(path)

	if err != nil {
		panic(err)
	}

	fmt.Println(strings.Count(string(content), sign))

}

func CountWords(path string, num_frequent int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	words_counter := make(map[string]int)
	for {
		line, _ := reader.ReadString('\n')
		if line == "" {
			break
		}
		lines := strings.Fields(line)
		for _, word := range lines {
			word = strings.ToUpper(word)
			words_counter[word]++
		}
	}
	// Sort map with word and number of occurence
	words := make([]string, 0, len(words_counter))
	for word := range words_counter {
		words = append(words, word)
	}
	sort.Slice(words, func(i, j int) bool {
		return words_counter[words[i]] > words_counter[words[j]]
	})

	for _, word := range words[:num_frequent] {
		fmt.Printf("%v: %v\n", word, words_counter[word])
	}
}

func main() {
	CountSign("example.txt", "a")
	CountWords("example.txt", 5)
}
