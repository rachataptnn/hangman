package hangman

import "testing"

func TestInitDisplayWord(t *testing.T) {
	testCases := []struct {
		testName       string
		inputWord      string
		expectedOutput string
		expectedCount  int
	}{
		{"All letters", "hello", "_____", 5},
		{"With space", "hi there", "__ _____", 7},
		{"With number", "abc123", "___123", 3},
		{"Special chars", "go!", "__!", 2},
		{"Mixed case", "GoLang", "______", 6},
		{"Only spaces", "   ", "   ", 0},
		{"Empty string", "", "", 0},
	}

	for _, tc := range testCases {
		t.Run(tc.testName, func(t *testing.T) {
			h := &hangman{TargetWord: tc.inputWord}

			h.initDisplayWord()
			if h.DisplayWord != tc.expectedOutput {
				t.Errorf("For word '%s', expected DisplayWord '%s', got '%s'", tc.inputWord, tc.expectedOutput, h.DisplayWord)
			}
			if h.LetterToGuess != tc.expectedCount {
				t.Errorf("For word '%s', expected LetterToGuess '%d', got '%d'", tc.inputWord, tc.expectedCount, h.LetterToGuess)
			}
		})
	}
}
