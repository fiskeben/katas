package main

import "fmt"

type row []int

func (r row) String() string {
	s := "|"
	i := 0
	for factor := 1; factor < 10; factor++ {
		if i < len(r) && r[i] < factor*10 {
			n := fmt.Sprintf("%d", r[i])

			if r[i] < 10 {
				n = fmt.Sprintf(" %s", n)
			}
			s = fmt.Sprintf("%s%s|", s, n)
			i++
		} else {
			s = fmt.Sprintf("%s  |", s)
		}
	}
	return s
}

type card []row

func (p card) String() string {
	s := ".__.__.__.__.__.__.__.__.__.\n"
	for _, r := range p {
		s = fmt.Sprintf("%s%s\n", s, r.String())
	}
	s = fmt.Sprintf("%s.__.__.__.__.__.__.__.__.__.\n", s)
	return s
}

type board []card

func (b board) String() string {
	s := ""
	for _, p := range b {
		s = fmt.Sprintf("%s%s\n", s, p.String())
	}
	return s
}
