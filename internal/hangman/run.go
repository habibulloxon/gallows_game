package hangman

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

func Run() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("[1] Start a new game")
		fmt.Println("[2] Start a new game with custom words")
		fmt.Println("[3] Exit")
		fmt.Print("Enter your choice: ")

		choice, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			word := SelectRandomWord(nil)
			PlayGame(word)
		case "2":
			fmt.Print("Enter the path for .txt file: ")
			path, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading file path:", err)
				continue
			}
			path = strings.TrimSpace(path)
			words, err := ReadWordsFromFile(path)
			if err != nil {
				fmt.Println("Error processing file:", err)
				continue
			}
			if len(words) == 0 {
				fmt.Println("No valid words found in file!")
				continue
			}
			word := SelectRandomWord(words)
			PlayGame(word)
		case "3":
			fmt.Println("Exiting the game...")
			return
		default:
			fmt.Println("[500] Invalid option, please try again.")
		}
	}
}

func PlayGame(word string) {
	game := NewHangmanGame(word)
	const maxInvalid = 3
	re := regexp.MustCompile(`^[a-zA-Z]$`)
	scanner := bufio.NewScanner(os.Stdin)
	invalidAttempts := 0

	for {
		game.DisplayState()

		if game.IsWon() {
			fmt.Println("Congratulations! You've guessed the word!")
			break
		}
		if game.IsLost() {
			fmt.Printf("Game over! The word was: %s\n", game.Word)
			break
		}

		fmt.Print("Enter a letter: ")
		if !scanner.Scan() {
			fmt.Println("Error reading input.")
			break
		}
		input := strings.TrimSpace(scanner.Text())

		if !re.MatchString(input) {
			fmt.Println("Error: Please enter exactly one letter from the English alphabet.")
			invalidAttempts++
			if invalidAttempts >= maxInvalid {
				fmt.Println("Too many invalid attempts. Exiting game.")
				break
			}
			continue
		}

		letter := rune(input[0])
		if unicode.IsUpper(letter) {
			fmt.Println("Error: Please enter a lowercase letter")
			invalidAttempts++
			if invalidAttempts >= maxInvalid {
				fmt.Println("Too many invalid attempts. Exiting game.")
				break
			}
			continue	
		}

		if !unicode.IsLetter(letter) {
			fmt.Println("Error: Please enter a valid letter (no numbers or symbols).")
			invalidAttempts++
			if invalidAttempts >= maxInvalid {
				fmt.Println("Too many invalid attempts. Exiting game.")
				break
			}
			continue
		}
		invalidAttempts = 0

		if game.GuessedLetters[letter] {
			fmt.Println("You've already guessed that letter!")
			continue
		}

		game.ProcessGuess(letter)
	}
}
