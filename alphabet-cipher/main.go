package main

import (
	"flag"
	"fmt"
	"os"
	"unicode"
)

const alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

func main() {
	var message = flag.String("message", "", "The message you want to encrypt")
	var secret = flag.String("secret", "", "The secret keyword to encrypt with")
	var doDecode = flag.Bool("decode", false, "Decode the message instead of encode")

	flag.Parse()

	if *message == "" || *secret == "" {
		flag.Usage()
		os.Exit(1)
	}

	t := makeTable()
	var res string
	if *doDecode {
		res = decode(t, *message, *secret)
	} else {
		res = encode(t, *message, *secret)
	}

	fmt.Printf("%s\n", res)
}

func makeTable() map[rune][]rune {
	t := make(map[rune][]rune)

	for i, c := range alphabet {
		first := toRunes(alphabet[i:len(alphabet)])
		second := toRunes(alphabet[0:i])
		t[c] = append(first, second...)
	}
	return t
}

func toRunes(chars string) []rune {
	runes := make([]rune, len(chars))
	for n, r := range chars {
		runes[n] = rune(r)
	}
	return runes
}

func encode(table map[rune][]rune, message, secret string) string {
	res := make([]rune, len(message))

	for i, char := range message {
		key := keyFromSecret(secret, i)
		index := findLetter(char, []rune(alphabet))
		combination := table[key]
		res[i] = unicode.ToLower(combination[index])
	}

	return string(res)
}

func keyFromSecret(secret string, index int) rune {
	length := len(secret)
	char := rune(secret[index%length])
	return unicode.ToUpper(char)
}

func findLetter(r rune, list []rune) int {
	r = unicode.ToUpper(r)
	for i, letter := range list {
		if letter == r {
			return i
		}
	}
	return 0
}

func decode(table map[rune][]rune, message, secret string) string {
	res := make([]rune, len(message))

	for i, char := range message {
		char = unicode.ToUpper(char)
		key := keyFromSecret(secret, i)
		combination := table[key]
		index := findLetter(char, combination)
		res[i] = unicode.ToLower(rune(alphabet[index]))
	}
	return string(res)
}
