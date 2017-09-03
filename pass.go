// Package xkcdpass generates secure random passwords that are (relatively) easy to memorize.
// It uses a dictionary provided specifically for passphrase generation by the EFF (https://www.eff.org/deeplinks/2016/07/new-wordlists-random-passphrases) and has an option for checking strength of the generated passwords using the zxcvbn library (https://github.com/nbutton23/zxcvbn-go).
// Inspired by https://xkcd.com/936/
package xkcdpass

import (
	"bytes"
	"crypto/rand"
	"math/big"
	"strings"

	"github.com/nbutton23/zxcvbn-go"
)

//go:generate go-bindata -nomemcopy static/...

var words []string
var max *big.Int

const defaultLength = 4
const defaultStrength = 4

func init() {
	data, err := Asset("static/default")
	if err != nil {
		panic("Failed to read word list: " + err.Error())
	}

	dict := string(data)
	words = strings.Split(dict, "\n")
	for index, w := range words {
		words[index] = strings.Title(strings.TrimSpace(w))
	}

	max = big.NewInt(int64(len(words)) - 1)
}

// GenerateDefault generates a random 4-word password.
func GenerateDefault() string {
	return makePassword(defaultLength)
}

// GenerateWithLength generates a random password with the specified number of words.
func GenerateWithLength(numWords int) string {
	return makePassword(numWords)
}

// GenerateDefaultChecked generates a random 4-word password with a zxcvbn strength score of 4.
func GenerateDefaultChecked() string {
	return makePasswordWithStrength(defaultLength, defaultStrength)
}

// GenerateWithLengthAndStrength generates a random password with the specified number of words and a score equal to or higher to the specified strength.
func GenerateWithLengthAndStrength(length int, strength int) string {
	return makePasswordWithStrength(length, strength)
}

func makePassword(numWords int) string {
	var b bytes.Buffer
	for i := 0; i < numWords; i++ {
		n, err := rand.Int(rand.Reader, max)
		index := int(n.Int64())

		if err != nil {
			panic("Failed to generate random number: " + err.Error())
		}

		w := strings.TrimSpace(words[index])
		b.WriteString(strings.Title(w))
	}

	return b.String()
}

func makePasswordWithStrength(numWords int, strength int) string {
	currenctStrength := 0
	var result string
	for currenctStrength < strength {
		result = makePassword(numWords)
		entropy := zxcvbn.PasswordStrength(result, nil)
		currenctStrength = entropy.Score
	}

	return result
}
