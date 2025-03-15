package hangman

import (
	"fmt"
	"strings"
)

var gallowsStages = []string{
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

type Hangman struct {
	Word           string
	GuessedLetters map[rune]bool
	ErrorCount     int
	MaxErrors      int
}

func NewHangmanGame(word string) *Hangman {
	return &Hangman{
		Word:           word,
		GuessedLetters: make(map[rune]bool),
		ErrorCount:     0,
		MaxErrors:      len(gallowsStages) - 1,
	}
}

func (g *Hangman) DisplayWord() string {
	var displayed strings.Builder
	for _, c := range g.Word {
		if g.GuessedLetters[c] {
			displayed.WriteString(string(c) + " ")
		} else {
			displayed.WriteString("_ ")
		}
	}
	return displayed.String()
}

func (g *Hangman) ProcessGuess(letter rune) {
	if strings.ContainsRune(g.Word, letter) {
		g.GuessedLetters[letter] = true
	} else {
		g.ErrorCount++
	}
}

func (g *Hangman) IsWon() bool {
	for _, c := range g.Word {
		if !g.GuessedLetters[c] {
			return false
		}
	}
	return true
}

func (g *Hangman) IsLost() bool {
	return g.ErrorCount >= g.MaxErrors
}

func (g *Hangman) DisplayState() {
	fmt.Printf("\nError Count: %d\n", g.ErrorCount)
	fmt.Println(gallowsStages[g.ErrorCount])
	fmt.Printf("Word: %s\n", g.DisplayWord())
}
