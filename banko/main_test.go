package main

import (
	"testing"
)

func Test_row_String(t *testing.T) {
	tests := []struct {
		name string
		r    row
		want string
	}{
		{
			name: "prints a row of consecutive numbers",
			r:    row{1, 10, 20, 30, 40},
			want: "| 1|10|20|30|40|  |  |  |  |",
		},
		{
			name: "prints a row with gaps",
			r:    row{5, 26, 48, 61, 88},
			want: "| 5|  |26|  |48|  |61|  |88|",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.r.String(); got != tt.want {
				t.Errorf("row.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makePlates(t *testing.T) {
	g := makeGenerator()
	b := g.MakeRandomBoard()

	if len(b) != 6 {
		t.Errorf("board does not have six plates: %v", b)
	}

	numbers := make([]bool, 91)
	for i := 0; i < 91; i++ {
		numbers[i] = false
	}

	for _, p := range b {
		if len(p) != 3 {
			t.Errorf("plate does not have three rows: %v", p)
		}
		for _, r := range p {
			if len(r) != 5 {
				t.Errorf("row does not have five elements: %v", r)
			}
			for _, n := range r {
				if numbers[n] == true {
					t.Errorf("duplicate number %d", n)
				}
				numbers[n] = true
			}
		}
	}
}
