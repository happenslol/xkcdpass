package xkcdpass

import (
	"fmt"
	"testing"
	"unicode"

	"github.com/nbutton23/zxcvbn-go"
)

func TestGenerateDefault(t *testing.T) {
	pass := GenerateDefault()

	numUppercase := 0
	for _, c := range pass {
		if unicode.IsUpper(c) {
			numUppercase++
		}
	}

	if numUppercase != defaultLength {
		t.Error("Number of words was not 4")
	}
}

func TestGenerateDefaultChecked(t *testing.T) {
	for i := 0; i < 5; i++ {
		pass := GenerateDefaultChecked(i)
		entropy := zxcvbn.PasswordStrength(pass, nil)
		if entropy.Score < i {
			t.Errorf("Failed to generate password with a score of at least %v", i)
		}
	}
}
