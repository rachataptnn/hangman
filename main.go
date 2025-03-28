package main

import (
	"hangman/hangman"
)

func main() {
	wordsSrcDir := "./words/"
	hangman := hangman.New(wordsSrcDir)
	hangman.PrepareWordCategories()
	hangman.StartGame()
}
