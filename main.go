package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

type WordPair [2]string

func main() {
	nums, path := parseFlags()

	words := readInputFile(path)

	possibles := buildWordMap(words)

	output := generateOutput(possibles, nums)

	printWrappedOutput(output)
}

func parseFlags() (int, string) {
	var nums int
	var path string

	flag.IntVar(&nums, "nums", 50, "Number of words a generated line must contain")
	flag.StringVar(&path, "path", "", "Path to the .txt file containing paragraphs to learn from")
	flag.Parse()

	if path == "" {
		log.Fatal("Path to the .txt file must be specified with -path flag")
	}

	if nums <= 0 {
		log.Fatal("Number of words must be greater than 0")
	}

	return nums, path
}

func readInputFile(path string) []string {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return strings.Fields(string(content))
}

func buildWordMap(words []string) map[WordPair][]string {
	possibles := make(map[WordPair][]string)
	word1, word2 := "", ""

	for _, word := range words {
		key := WordPair{word1, word2}
		possibles[key] = append(possibles[key], word)
		word1, word2 = word2, word
	}

	possibles[WordPair{word1, word2}] = append(possibles[WordPair{word1, word2}], "")
	possibles[WordPair{word2, ""}] = append(possibles[WordPair{word2, ""}], "")

	return possibles
}

func generateOutput(possibles map[WordPair][]string, nums int) []string {
	rand.Seed(time.Now().UnixNano())

	start := chooseStartingWords(possibles)
	word1, word2 := start[0], start[1]

	output := []string{word1, word2}

	for i := 0; i < nums; i++ {
		words, ok := possibles[WordPair{word1, word2}]
		if !ok || len(words) == 0 {
			break
		}
		word := words[rand.Intn(len(words))]
		if word == "" {
			continue
		}
		output = append(output, word)
		word1, word2 = word2, word
	}

	return output
}

func chooseStartingWords(possibles map[WordPair][]string) WordPair {
	var startKeys []WordPair
	for k := range possibles {
		if len(k[0]) > 0 && strings.ToUpper(k[0][:1]) == k[0][:1] {
			startKeys = append(startKeys, k)
		}
	}

	if len(startKeys) == 0 {
		log.Fatal("No valid starting word found")
	}

	return startKeys[rand.Intn(len(startKeys))]
}

func printWrappedOutput(output []string) {
	wrappedText := wrapText(strings.Join(output, " "), 90)
	fmt.Println(strings.Join(wrappedText, "\n"))
}

func wrapText(text string, width int) []string {
	var result []string
	words := strings.Fields(text)
	var line string

	for _, word := range words {
		if len(line)+len(word)+1 > width {
			result = append(result, line)
			line = word
		} else {
			if line != "" {
				line += " "
			}
			line += word
		}
	}

	if line != "" {
		result = append(result, line)
	}

	return result
}
