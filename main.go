package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	prepareWordCategories()
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
