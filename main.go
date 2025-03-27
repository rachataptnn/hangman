package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func main() {
	wcs, err := prepareWordCategories()
	if err != nil {
		fmt.Println(err)
		return
	}

	hangman(wcs)
}

type WordCategory struct {
	Name  string
	Words []word
}

type word struct {
	Word string
	Hint string
}

func prepareWordCategories() ([]WordCategory, error) {
	entries, err := os.ReadDir("./words")
	if err != nil {
		return nil, err
	}

	wc := make([]WordCategory, 0)
	for _, entry := range entries {
		if !entry.IsDir() {
			wordCategory, err := prepareWordCategory(entry.Name())
			if err != nil {
				return nil, err
			}

			wc = append(wc, wordCategory)
		}

	}

	return wc, nil
}

func prepareWordCategory(fileName string) (WordCategory, error) {
	fileContent, err := os.OpenFile("./words/"+fileName, os.O_RDONLY, 0)
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

//

func hangman(wcs []WordCategory) {
	categoryNumber := selectCategory(wcs)

	category := wcs[categoryNumber-1]
	word := category.Words[rand.Intn(len(category.Words))]

	fmt.Println("Hint: ", word.Hint)
	fmt.Println("")

	isWin, score := guess(word)
	if isWin {
		fmt.Println("Congratulations! You found the word!")
		fmt.Println("your score: ", score)
	}

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

func guess(word word) (isPass bool, score int) {
	incorrectGuessLimit := 10

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

	correctCount := 0
	for incorrectGuessLimit > 0 && correctCount != n {
		var letter string
		fmt.Print("Enter a letter: ")
		fmt.Scanln(&letter)

		if len(letter) != 1 {
			fmt.Println("")
			fmt.Println("Invalid Letter")
			fmt.Println("")
			continue
		}

		if strings.Contains(word.Word, letter) {
			fmt.Println("")
			fmt.Println("Correct Letter!")
			fmt.Println("")

			correctCount += 1
			score += 10
			process = updateProcess(word.Word, process, letter, n)
			showProcess(process, score, incorrectGuessLimit)

		} else {
			fmt.Println("")
			fmt.Println("Incorrect Letter~")
			fmt.Println("")

			incorrectGuessLimit--
			showProcess(process, score, incorrectGuessLimit)

			if incorrectGuessLimit == 0 {
				fmt.Println("Game Over!")
				fmt.Println("Word: ", word.Word)
				return false, score
			}
		}
	}

	return true, score
}

func showProcess(process string, score, incorrectGuessLimit int) {
	fmt.Printf("%s	score:%d remaining incorrect guess:%d\n", process, score, incorrectGuessLimit)
}

func updateProcess(word, process, letter string, n int) string {
	for i := 0; i < n; i++ {
		if word[i] == letter[0] && process[i] == '_' {
			process = process[:i] + string(letter) + process[i+1:]
			break
		}
	}

	return process
}
