package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

const (
	chances = 6
	Rules   = `
=====
Rules 
=====
1. You have 6 chances to guess the word
2. You can guess one letter at a time
3. You can guess the whole word at once by entering more than one letter`
)

type Hangman interface {
	PrepareWordCategories() error
	StartGame()
}

func New(wordsSrcDir string) Hangman {
	return &hangman{
		wordsSrcDir:      wordsSrcDir,
		Score:            0,
		CorrectGuesses:   0,
		IncorrectGuesses: 0,
	}
}

type hangman struct {
	wordsSrcDir      string
	WordCategories   []WordCategory
	SelectedCategory int
	TargetWord       string
	DisplayWord      string
	LetterToGuess    int
	CorrectGuesses   int
	IncorrectGuesses int
	IncorrectLetters []string
	Score            int
}

type WordCategory struct {
	Name  string
	Words []word
}

type word struct {
	Word string
	Hint string
}

func (h *hangman) PrepareWordCategories() error {
	entries, err := os.ReadDir(h.wordsSrcDir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			wordCategory, err := h.prepareWordCategory(entry.Name())
			if err != nil {
				return err
			}

			h.WordCategories = append(h.WordCategories, wordCategory)
		}
	}

	return nil
}

func (h *hangman) prepareWordCategory(fileName string) (WordCategory, error) {
	fileContent, err := os.OpenFile(h.wordsSrcDir+fileName, os.O_RDONLY, 0)
	if err != nil {
		return WordCategory{}, err
	}

	scanner := bufio.NewScanner(fileContent)
	defer fileContent.Close()

	var lineNumber int
	var category WordCategory
	for scanner.Scan() {
		lineContent := scanner.Text()
		if lineNumber == 0 {
			category.Name = lineContent
		} else {
			parts := strings.Split(lineContent, ",")
			if len(parts) == 2 {
				category.Words = append(category.Words, word{Word: parts[0], Hint: parts[1]})
			} else {
				fmt.Printf("File: %s found invalid line: %s\n", fileName, lineContent)
			}
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		return WordCategory{}, err
	}

	return category, nil
}

func (h *hangman) StartGame() {
	categoryNumber := selectCategory(h.WordCategories)
	category := h.WordCategories[categoryNumber-1]

	targetWord := category.Words[rand.Intn(len(category.Words))]
	h.TargetWord = strings.ToLower(targetWord.Word)

	fmt.Println(Rules)
	printWithPadding("GAME STARTED!! Hint: " + targetWord.Hint)

	h.initDisplayWord()
	h.showRoundSummary()

	isWin := h.guess()
	if isWin {
		fmt.Println("Congratulations! You found the word!")
		fmt.Println("your score: ", h.Score)
	} else {
		fmt.Println("Game Over! The word was: ", targetWord.Word)
		fmt.Println("your score: ", h.Score)
	}

	var playAgain string
	printWithPadding("Wanna play again? Y/n")
	fmt.Scanln(&playAgain)
	for strings.ToLower(playAgain) != "y" && strings.ToLower(playAgain) != "n" {
		fmt.Println("Invalid input. Please enter Y or n")
		fmt.Scanln(&playAgain)
	}
	if strings.ToLower(playAgain) == "y" {
		h = New(h.wordsSrcDir).(*hangman)
		h.PrepareWordCategories()
		h.StartGame()
	}

	fmt.Println("Goodbye!")
}

func (h *hangman) initDisplayWord() {
	for _, v := range h.TargetWord {
		if v == ' ' {
			h.DisplayWord += " "
		} else if (v >= 'a' && v <= 'z') || (v >= 'A' && v <= 'Z') {
			h.DisplayWord += "_"
			h.LetterToGuess++
		} else {
			h.DisplayWord += string(v)
		}
	}
}

func (h *hangman) updateDisplayWord(letter string) {
	for i := 0; i < len(h.TargetWord); i++ {
		if h.TargetWord[i] == letter[0] && h.DisplayWord[i] == '_' {
			h.DisplayWord = h.DisplayWord[:i] + string(letter) + h.DisplayWord[i+1:]
			h.CorrectGuesses++
		}
	}
}

func (h *hangman) showRoundSummary() {
	spacedDisplayWord := strings.Join(strings.Split(h.DisplayWord, ""), " ")
	fmt.Printf("\n%s	score: %d, incorrect guess: %d\n\n", spacedDisplayWord, h.Score, h.IncorrectGuesses)

	if len(h.IncorrectLetters) > 0 {
		fmt.Printf("Incorrect Letters: %s\n", strings.Join(h.IncorrectLetters, ", "))
	}
}

func (h *hangman) guess() bool {
	usedLetters := make(map[string]bool)

	for h.IncorrectGuesses < chances && h.CorrectGuesses != h.LetterToGuess {
		letter := readLetter()
		if len(letter) < 1 {
			fmt.Println("Please enter at least one letter")
			continue
		}

		letter = strings.ToLower(letter)
		if len(letter) > 1 {
			isWin := h.checkWholeWord(letter)
			if isWin {
				return true
			}

			continue
		}

		_, ok := usedLetters[letter]
		if ok {
			fmt.Printf("You already guessed this letter, try another one\n\n")
			continue
		}

		usedLetters[letter] = true
		if strings.Contains(h.TargetWord, letter) {
			fmt.Println("Correct!")

			h.Score += 10

			h.updateDisplayWord(letter)
			h.showRoundSummary()

		} else {
			fmt.Println("WRONG")

			h.Score -= 5
			h.IncorrectGuesses++
			h.IncorrectLetters = append(h.IncorrectLetters, letter)

			h.drawHangman()
			h.showRoundSummary()

			if h.IncorrectGuesses == chances {
				return false
			}
		}
	}

	return true
}

func (h *hangman) checkWholeWord(letter string) bool {
	printWithPadding("!! YOU GUESS THE WHOLE WORD !!")

	if letter == h.TargetWord {
		fmt.Println("wow! you found the word! (score is increased by 20000)")
		h.Score += 20000
		return true
	}

	fmt.Println("Nice try... but not this Word. (score is decreased by 20)")

	h.IncorrectGuesses++
	h.Score -= 20

	h.showRoundSummary()
	h.drawHangman()

	return false
}

func (h *hangman) drawHangman() {
	switch h.IncorrectGuesses {
	case 0:
		fmt.Println(` +---+
 |   |
	 |
	 |
	 |
	 |
======`)
	case 1:
		fmt.Println(` +---+
 |   |
 O   |
     |
     |
     |
======`)
	case 2:
		fmt.Println(` +---+
 |   |
 O   |
 |   |
     |
     |
======`)
	case 3:
		fmt.Println(` +---+
 |   |
 O   |
/|   |
     |
     |
======`)
	case 4:
		fmt.Println(` +---+
 |   |
 O   |
/|\  |
     |
     |
======`)
	case 5:
		fmt.Println(` +---+
 |   |
 O   |
/|\  |
/    |
     |
======`)
	case 6:
		fmt.Println(` +---+
 |   |
 O   |
/|\  |
/ \  |
     |
======`)
	}
}

func selectCategory(wcs []WordCategory) int {
	printWithPadding("Select Category:")
	for i, wc := range wcs {
		fmt.Printf("%d. %s\n", i+1, wc.Name)
	}
	fmt.Println("")

	var categoryNumber int
	for categoryNumber <= 0 || categoryNumber > len(wcs) {
		fmt.Print("Enter Category Number: ")
		fmt.Scanln(&categoryNumber)

		if categoryNumber < 0 || categoryNumber > len(wcs) {
			fmt.Println("Invalid Category Number")
		} else {
			fmt.Println("You choose: ", wcs[categoryNumber-1].Name)
		}
	}

	return categoryNumber
}

func readLetter() string {
	fmt.Printf("\nlet's guess: ")

	reader := bufio.NewReader(os.Stdin)
	letter, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("An error occurred while reading the input:" + err.Error())
		os.Exit(1)
	}

	if len(letter) < 1 {
		return ""
	}

	letterWithOutNewLine := strings.ToLower(letter[:len(letter)-1])
	return letterWithOutNewLine
}

func printWithPadding(msg string) {
	fmt.Printf("\n%s\n\n", msg)
}
