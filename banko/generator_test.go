package main

import (
	"math/rand"
	"testing"
)

func Test_generator_generateRow(t *testing.T) {
	type fields struct {
		randomizer *rand.Rand
	}
	type args struct {
		allRows [][]int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "generates a row",
			fields: fields{
				randomizer: rand.New(rand.NewSource(1)),
			},
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := generator{
				randomizer: tt.fields.randomizer,
				groups: [][]int{
					{1, 2, 7, 3, 6, 4, 8, 9, 5},
					{10, 19, 18, 14, 11, 15, 16, 13, 12, 17},
					{22, 25, 26, 29, 28, 23, 21, 24, 27, 20},
					{37, 30, 33, 39, 36, 35, 34, 32, 31, 38},
					{48, 44, 41, 40, 42, 47, 45, 49, 43, 46},
					{56, 50, 53, 57, 59, 58, 52, 54, 51, 55},
					{67, 65, 64, 69, 66, 60, 68, 62, 63, 61},
					{74, 71, 70, 78, 73, 72, 75, 76, 79, 77},
					{88, 80, 89, 90, 83, 86, 82, 85, 87, 84, 81},
				},
			}
			got := g.generateRow()
			if len(got) != 5 {
				t.Errorf("generator.generateRow() = %v, does not have five elements", got)
			}
			for _, x := range got {
				if x == 0 {
					t.Errorf("number is invalid: %v\n", got)
				}
			}
			t.Logf("generated row:%v\n", got)
		})
	}
}

func Test_generator_indexesBasedOnLongestGroups(t *testing.T) {
	type fields struct {
		randomizer *rand.Rand
		groups     [][]int
	}
	tests := []struct {
		name   string
		fields fields
		want   func([]int) bool
	}{
		{
			name: "generates indexes",
			fields: fields{
				randomizer: rand.New(rand.NewSource(1)),
				groups: [][]int{
					{1, 2, 7, 3, 6, 4, 8, 9, 5},
					{10, 19, 18, 14, 11, 15, 16, 13, 12, 17},
					{22, 25, 26, 29, 28, 23, 21, 24, 27, 20},
					{37, 30, 33, 39, 36, 35, 34, 32, 31, 38},
					{48, 44, 41, 40, 42, 47, 45, 49, 43, 46},
					{56, 50, 53, 57, 59, 58, 52, 54, 51, 55},
					{67, 65, 64, 69, 66, 60, 68, 62, 63, 61},
					{74, 71, 70, 78, 73, 72, 75, 76, 79, 77},
					{88, 80, 89, 90, 83, 86, 82, 85, 87, 84, 81},
				},
			},
			want: func(res []int) bool {
				return res[0] == 8
			},
		},
		{
			name: "generates indexes",
			fields: fields{
				randomizer: rand.New(rand.NewSource(1)),
				groups: [][]int{
					{1, 2, 5},
					{13, 12, 17},
					{29, 28, 23, 21, 24, 27, 20},
					{35, 34, 32, 31, 38},
					{48, 44, 41, 40, 42, 47, 45, 49, 43, 46},
					{56, 50, 55},
					{},
					{75, 76, 79, 77},
					{81},
				},
			},
			want: func(res []int) bool {
				return res[0] == 4
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := generator{
				randomizer: tt.fields.randomizer,
				groups:     tt.fields.groups,
			}
			if got := g.indexesBasedOnLongestGroups(); !tt.want(got) {
				t.Errorf("generator.indexesBasedOnLongestGroups() = %v", got)
			}
		})
	}
}
