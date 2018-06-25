package pkg

import (
	"fmt"
	"strings"
)

const (
	statusError   = "ERR"
	statusIllegal = "ILL"
)

var digitmap = map[rune]digit{
	'1': digit{[]rune{' ', ' ', ' '}, []rune{' ', ' ', '|'}, []rune{' ', ' ', '|'}},
	'2': digit{[]rune{' ', '_', ' '}, []rune{' ', '_', '|'}, []rune{'|', '_', ' '}},
	'3': digit{[]rune{' ', '_', ' '}, []rune{' ', '_', '|'}, []rune{' ', '_', '|'}},
	'4': digit{[]rune{' ', ' ', ' '}, []rune{'|', '_', '|'}, []rune{' ', ' ', '|'}},
	'5': digit{[]rune{' ', '_', ' '}, []rune{'|', '_', ' '}, []rune{' ', '_', '|'}},
	'6': digit{[]rune{' ', '_', ' '}, []rune{'|', '_', ' '}, []rune{'|', '_', '|'}},
	'7': digit{[]rune{' ', '_', ' '}, []rune{' ', ' ', '|'}, []rune{' ', ' ', '|'}},
	'8': digit{[]rune{' ', '_', ' '}, []rune{'|', '_', '|'}, []rune{'|', '_', '|'}},
	'9': digit{[]rune{' ', '_', ' '}, []rune{'|', '_', '|'}, []rune{' ', '_', '|'}},
	'0': digit{[]rune{' ', '_', ' '}, []rune{'|', ' ', '|'}, []rune{'|', '_', '|'}},
}

// MakeReport accepts a string containing the data
// and returns a report of it.
func MakeReport(data string) string {
	numbers := parse(data)
	out := make([]string, len(numbers))

	for i, n := range numbers {
		status := ""
		if strings.ContainsAny(string(n), "?") {
			status = statusIllegal
		} else if !checksum(n) {
			status = statusError
		}
		out[i] = strings.TrimSpace(fmt.Sprintf("%s %s", n, status))
	}
	return strings.Join(out, "\n")
}

func parse(data string) []accountnumber {
	rawNumbers := split(data)
	accountnumbers := make([]accountnumber, len(rawNumbers))

	for i, rawNumber := range rawNumbers {
		n := parseRawNumber(rawNumber)
		accountnumbers[i] = numberToValue(n)
	}

	return accountnumbers
}

func split(data string) []rawAccountNumber {
	lines := strings.Split(data, "\n")
	res := []rawAccountNumber{}

	currentAccountNumber := make([]string, 3)
	c := 0
	for _, l := range lines {
		if c == 3 {
			res = append(res, currentAccountNumber)
			currentAccountNumber = make([]string, 3)
			c = 0
			continue
		}
		currentAccountNumber[c] = l
		c++
	}
	return res
}

func parseRawNumber(acn rawAccountNumber) number {
	numberOfDigits := len(acn[0]) / 3
	n := make([]digit, numberOfDigits)
	digitIndex := 0
	for i := 0; i < numberOfDigits; i++ {
		digitIndex = i * 3
		d := digit{}
		for j := 0; j < 3; j++ {
			r1 := rune(acn[j][digitIndex])
			r2 := rune(acn[j][digitIndex+1])
			r3 := rune(acn[j][digitIndex+2])
			d = append(d, []rune{r1, r2, r3})
		}
		n[i] = d
	}
	return n
}

func numberToValue(n number) accountnumber {
	res := make([]rune, len(n))
	for i, d := range n {
		res[i] = deduce(d)
	}
	return accountnumber(res)
}

func deduce(a digit) rune {
digitloop:
	for r, candidate := range digitmap {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if a[i][j] != candidate[i][j] {
					continue digitloop
				}
			}
		}
		return r
	}
	return '?'
}

func checksum(n accountnumber) bool {
	sum := 0
	c := 1
	for x := len(n) - 1; x >= 0; x-- {
		v := int(n[x]-'0') * c
		sum += v
		c++
	}
	return sum%11 == 0
}
