package main

import (
	"hangman/hangman"
)

func main() {
	wordsDirectory := "./words/"
	hangman := hangman.New(wordsDirectory)
	hangman.PrepareWordCategories()
	hangman.StartGame()
}
