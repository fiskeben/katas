package main

import (
	"testing"
)

func equals(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}

	for i, c := range a {
		if c != b[i] {
			return false
		}
	}
	return true
}

func Test_makeTable(t *testing.T) {
	table := makeTable()
	if len(table) != 26 {
		t.Error("table does not contain 26 items")
	}

	if !equals(table['A'], []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")) {
		t.Errorf("A has wrong value '%v', expected '%v'", table['A'], "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}

	if !equals(table['T'], []rune("TUVWXYZABCDEFGHIJKLMNOPQRS")) {
		t.Errorf("T has wrong value '%v', expected '%v'", table['T'], "TUVWXYZABCDEFGHIJKLMNOPQRS")
	}
}

func Test_encode(t *testing.T) {
	table := makeTable()
	type args struct {
		table   map[rune][]rune
		message string
		secret  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "can encode given a secret keyword",
			args: args{
				table:   table,
				message: "meetmeontuesdayeveningatseven",
				secret:  "vigilance",
			},
			want: "hmkbxebpxpmyllyrxiiqtoltfgzzv",
		},
		{
			name: "can encode another given a secret keyword",
			args: args{
				table:   table,
				message: "meetmebythetree",
				secret:  "scones",
			},
			want: "egsgqwtahuiljgs",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := encode(tt.args.table, tt.args.message, tt.args.secret); got != tt.want {
				t.Errorf("encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decode(t *testing.T) {
	table := makeTable()
	type args struct {
		table   map[rune][]rune
		message string
		secret  string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "can decode a message given a secret keyword",
			args: args{
				table:   table,
				message: "hmkbxebpxpmyllyrxiiqtoltfgzzv",
				secret:  "vigilance",
			},
			want: "meetmeontuesdayeveningatseven",
		},
		{
			name: "can decode another message given a secret keyword",
			args: args{
				table:   table,
				message: "egsgqwtahuiljgs",
				secret:  "scones",
			},
			want: "meetmebythetree",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decode(tt.args.table, tt.args.message, tt.args.secret); got != tt.want {
				t.Errorf("encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
