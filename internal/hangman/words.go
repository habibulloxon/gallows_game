package hangman

import (
	"bufio"
	"math/rand"
	"os"
	"strings"
	"time"
)

func SelectRandomWord(arr []string) string {
	var words []string
	if arr != nil {
		words = arr
	} else {
		words = []string{
			"gosha",
			"sonyxm4",
			"aux",
			"kiwi",
			"hr",
		}
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return words[r.Intn(len(words))]
}

func ReadWordsFromFile(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			words = append(words, line)
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return words, nil
}
