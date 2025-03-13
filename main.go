package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

var gallows = []string{
	`+---+
|   |
    |
    |
    |
    |
=========`,
	`+---+
|   |
O   |
    |
    |
    |
=========`,
	`+---+
|   |
O   |
|   |
    |
    |
=========`,
	`+---+
|   |
O   |
/|   |
    |
    |
=========`,
	`+---+
|   |
O   |
/|\  |
    |
    |
=========`,
	`+---+
|   |
O   |
/|\  |
/    |
    |
=========`,
	`+---+
|   |
O   |
/|\  |
/ \  |
    |
=========`,
}

type HangmanGame struct {
	word           string
	guessedLetters map[rune]bool
	errorCount     int
	maxErrors      int
}

func NewHangmanGame(word string) *HangmanGame {
	return &HangmanGame{
		word:           word,
		guessedLetters: make(map[rune]bool),
		errorCount:     0,
		maxErrors:      len(gallows) - 1,
	}
}

func (g *HangmanGame) displayWord() string {
	var displayed strings.Builder
	for _, c := range g.word {
		if g.guessedLetters[c] {
			displayed.WriteString(string(c) + " ")
		} else {
			displayed.WriteString("_ ")
		}
	}
	return displayed.String()
}

func (g *HangmanGame) ProcessGuess(letter rune) {
	if strings.ContainsRune(g.word, letter) {
		g.guessedLetters[letter] = true
	} else {
		g.errorCount++
	}
}

func (g *HangmanGame) isWon() bool {
	for _, c := range g.word {
		if !g.guessedLetters[c] {
			return false
		}
	}
	return true
}

func (g *HangmanGame) isLost() bool {
	return g.errorCount >= g.maxErrors
}

func (g *HangmanGame) DisplayState() {
	fmt.Printf("\nError Count: %d\n", g.errorCount)
	fmt.Println(gallows[g.errorCount])
	fmt.Printf("Word: %s\n", g.displayWord())
}

func (g *HangmanGame) Play() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		g.DisplayState()

		if g.isWon() {
			fmt.Println("Congratulations! You've guessed the word!")
			break
		}
		if g.isLost() {
			fmt.Printf("Game over! The word was: %s\n", g.word)
			break
		}

		fmt.Print("Enter a letter: ")
		scanner.Scan()
		input := scanner.Text()
		input = strings.TrimSpace(input)
		if len(input) == 0 {
			continue
		}
		letter := rune(strings.ToLower(input)[0])

		if g.guessedLetters[letter] {
			fmt.Println("You've already guessed that letter!")
			continue
		}

		g.ProcessGuess(letter)
	}
}

func selectRandomWord() string {
	words := []string{
		"gosha",
		"sonyxm4",
		"aux",
		"kiwi",
		"hr",
	}
	rand.Seed(time.Now().UnixNano())
	return words[rand.Intn(len(words))]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("[1] Start a new game")
		fmt.Println("[2] Exit")
		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		if choice == "2" {
			fmt.Println("Exiting the game...")
			break
		} else if choice == "1" {
			word := selectRandomWord()
			game := NewHangmanGame(word)
			game.Play()
		} else {
			fmt.Println("[500] Invalid option, please try again.")
		}
	}
}
