package hangman

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

type Hangman interface {
	PrepareWordCategories() error
	StartGame()
}

func New(wordsDirectory string) Hangman {
	return &hangman{
		wordsDirectory: wordsDirectory,
	}
}

type hangman struct {
	wordsDirectory   string
	WordCategories   []WordCategory
	SelectedCategory int
	TargetWord       word
	GuessedLetters   []string
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
	entries, err := os.ReadDir(h.wordsDirectory)
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
	fileContent, err := os.OpenFile(h.wordsDirectory+fileName, os.O_RDONLY, 0)
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
	word := category.Words[rand.Intn(len(category.Words))]

	fmt.Println("Hint: ", word.Hint)
	fmt.Println("")

	isWin, score := guess(word)
	if isWin {
		fmt.Println("Congratulations! You found the word!")
		fmt.Println("your score: ", score)
	}

	var playAgain string
	fmt.Println("Wanna play again? Y/n")
	fmt.Scanln(&playAgain)
	if strings.ToLower(playAgain) == "y" {
		h.StartGame()
	}

	fmt.Println("Goodbye!")
}

func selectCategory(wcs []WordCategory) int {
	fmt.Println("Select Category:")
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
		}
	}

	return categoryNumber
}

func guess(word word) (bool, int) {
	incorrectGuessLimit := 10
	score := 0

	process := ""
	n := len(word.Word)
	for _, v := range word.Word {
		if v == ' ' {
			process += " "
		} else {
			process += "_"
		}

	}
	showProcess(process, n, incorrectGuessLimit)

	lowerWord := strings.ToLower(word.Word)
	correctCount := 0
	for incorrectGuessLimit > 0 && correctCount != n {
		// var letter string
		// fmt.Print("Enter a letter: ")
		// fmt.Scanln(&letter)

		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter a string: ")
		letter, err := reader.ReadString('\n') // Reads input until a newline character
		if err != nil {
			fmt.Println(err)
			return false, score
		}

		letter = letter[:len(letter)-1]
		fmt.Println("You entered:", letter)

		letter = strings.ToLower(letter)

		if len(letter) > 1 {

			fmt.Println("")
			fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			fmt.Println("!! YOU GUESS THE WHOLE WORD !!")
			fmt.Println("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
			fmt.Println("")

			if letter == lowerWord {
				score += 20000
				showProcess(lowerWord, score, incorrectGuessLimit)
				return true, score
			} else {
				fmt.Printf("letter <%s>", letter)
				fmt.Printf("word <%s>", lowerWord)

				fmt.Println("Incorrect Word, score is decreased by 20")
				incorrectGuessLimit -= 1
				score -= 20
				showProcess(process, score, incorrectGuessLimit)
				// drawHangman(10 - incorrectGuessLimit)
			}
			continue
		}

		if strings.Contains(lowerWord, letter) {
			fmt.Println("")
			fmt.Println("Correct Letter!")
			fmt.Println("")

			correctCount += 1
			score += 10
			process = updateProcess(lowerWord, process, letter, n)
			showProcess(process, score, incorrectGuessLimit)

		} else {
			fmt.Println("")
			fmt.Println("Incorrect Letter~")
			fmt.Println("")

			incorrectGuessLimit--
			showProcess(process, score, incorrectGuessLimit)
			// drawHangman(10 - incorrectGuessLimit)

			if incorrectGuessLimit == 0 {
				fmt.Println("Game Over!")
				fmt.Println("Word: ", lowerWord)
				return false, score
			}
		}
	}

	return true, score
}

func showProcess(process string, score, incorrectGuessLimit int) {
	spacedProcess := strings.Join(strings.Split(process, ""), " ")
	fmt.Printf("%s	score:%d remaining incorrect guess:%d\n", spacedProcess, score, incorrectGuessLimit)
}

func updateProcess(word, process, letter string, n int) string {
	for i := 0; i < n; i++ {
		if word[i] == letter[0] && process[i] == '_' {
			process = process[:i] + string(letter) + process[i+1:]
			// break
		}
	}

	return process
}

// func drawHangman(left int) {
// 	switch left {
// 	case 0:
// 		fmt.Println(`
// +---+
// |   |
//     |
//     |
//     |
//     |
// ======
// `)
// 	case 1:
// 		fmt.Println(`
// +---+
// |   |
// O   |
//     |
//     |
//     |
// ======
// `)
// 	case 2:
// 		fmt.Println(`
// +---+
// |   |
// O   |
// |   |
//     |
//     |
// ======
// `)
// 	case 3:
// 		fmt.Println(`
// +---+
// |   |
// O   |
// /|   |
//     |
//     |
// ======
// `)
// 	case 4:
// 		fmt.Println(`
// +---+
// |   |
// O   |
// /|\  |
//     |
//     |
// ======
// `)
// 	case 5:
// 		fmt.Println(`
// +---+
// |   |
// O   |
// /|\  |
// /    |
//     |
// ======
// `)
// 	case 6:
// 		fmt.Println(`
// +---+
// |   |
// O   |
// /|\  |
// / \  |
//     |
// ======
// `)
// 	}
// }
