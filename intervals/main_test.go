package main

import (
	"reflect"
	"testing"
)

func Test_parseSequences(t *testing.T) {
	type args struct {
		seq string
	}
	tests := []struct {
		name    string
		args    args
		want    []interval
		wantErr bool
	}{
		{
			name: "simple sequence",
			args: args{
				seq: "5-10",
			},
			want:    []interval{interval{From: 5, To: 10}},
			wantErr: false,
		},
		{
			name: "two sequences",
			args: args{
				seq: "5-10,20-30",
			},
			want:    []interval{interval{From: 5, To: 10}, interval{From: 20, To: 30}},
			wantErr: false,
		},
		{
			name: "single number",
			args: args{
				seq: "5",
			},
			want:    []interval{interval{From: 5, To: 5}},
			wantErr: false,
		},
		{
			name: "missing start of interval",
			args: args{
				seq: "-10",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "missing end of interval",
			args: args{
				seq: "2-",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "start of sequence is lower than end",
			args: args{
				seq: "5-3",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseSequences(tt.args.seq)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseSequences() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeSequences(t *testing.T) {
	type args struct {
		seq []interval
	}
	tests := []struct {
		name    string
		args    args
		want    []interval
		wantErr bool
	}{
		{
			name: "returns same when input is a single interval",
			args: args{
				seq: []interval{
					interval{From: 10, To: 20},
				},
			},
			want: []interval{
				interval{From: 10, To: 20},
			},
			wantErr: false,
		},
		{
			name: "merges two non-overlapping intervals",
			args: args{
				seq: []interval{
					interval{From: 10, To: 20},
					interval{From: 30, To: 40},
				},
			},
			want: []interval{
				interval{From: 10, To: 20},
				interval{From: 30, To: 40},
			},
			wantErr: false,
		},
		{
			name: "merges two overlapping intervals into one",
			args: args{
				seq: []interval{
					interval{From: 10, To: 20},
					interval{From: 15, To: 30},
				},
			},
			want:    []interval{interval{From: 10, To: 30}},
			wantErr: false,
		},
		{
			name: "merges two overlapping intervals and leaves one not overlapping",
			args: args{
				seq: []interval{
					interval{From: 10, To: 20},
					interval{From: 30, To: 40},
					interval{From: 12, To: 21},
				},
			},
			want: []interval{
				interval{From: 10, To: 21},
				interval{From: 30, To: 40},
			},
			wantErr: false,
		},
		{
			name: "merges two sequential intervals",
			args: args{
				seq: []interval{
					interval{From: 10, To: 20},
					interval{From: 21, To: 30},
				},
			},
			want:    []interval{interval{From: 10, To: 30}},
			wantErr: false,
		},
		{
			name: "merges three intervals where one overlaps the others",
			args: args{
				seq: []interval{
					interval{From: 10, To: 20},
					interval{From: 25, To: 30},
					interval{From: 15, To: 28},
				},
			},
			want:    []interval{interval{From: 10, To: 30}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mergeSequences(tt.args.seq)
			if (err != nil) != tt.wantErr {
				t.Errorf("mergeSequences() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_overlapOrConsequtive(t *testing.T) {
	type args struct {
		first  interval
		second interval
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "no overlap",
			args: args{
				first:  interval{From: 10, To: 20},
				second: interval{From: 30, To: 40},
			},
			want: false,
		},
		{
			name: "first is before with overlap",
			args: args{
				first:  interval{From: 10, To: 20},
				second: interval{From: 15, To: 25},
			},
			want: true,
		},
		{
			name: "first is after with overlap",
			args: args{
				first:  interval{From: 10, To: 20},
				second: interval{From: 5, To: 15},
			},
			want: true,
		},
		{
			name: "first spans second completely",
			args: args{
				first:  interval{From: 10, To: 20},
				second: interval{From: 12, To: 18},
			},
			want: true,
		},
		{
			name: "intervals are consequtive",
			args: args{
				first:  interval{From: 10, To: 20},
				second: interval{From: 21, To: 30},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := overlapOrConsequtive(tt.args.first, tt.args.second); got != tt.want {
				t.Errorf("overlap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subtractSequences(t *testing.T) {
	type args struct {
		includes []interval
		excludes []interval
	}
	tests := []struct {
		name    string
		args    args
		want    []interval
		wantErr bool
	}{
		{
			name: "subtracting non-overlapping returns same",
			args: args{
				includes: []interval{
					interval{From: 10, To: 20},
				},
				excludes: []interval{
					interval{From: 30, To: 40},
				},
			},
			want:    []interval{interval{From: 10, To: 20}},
			wantErr: false,
		},
		{
			name: "subtracts one exclude from one include",
			args: args{
				includes: []interval{
					interval{From: 10, To: 20},
				},
				excludes: []interval{
					interval{From: 15, To: 30},
				},
			},
			want:    []interval{interval{From: 10, To: 14}},
			wantErr: false,
		},
		{
			name: "subtracts exclude overlapping multiple includes",
			args: args{
				includes: []interval{
					interval{From: 10, To: 20},
					interval{From: 25, To: 40},
				},
				excludes: []interval{
					interval{From: 15, To: 30},
				},
			},
			want: []interval{
				interval{From: 10, To: 14},
				interval{From: 31, To: 40},
			},
			wantErr: false,
		},
		{
			name: "regression",
			args: args{
				includes: []interval{
					interval{5, 10},
					interval{20, 30},
				},
				excludes: []interval{
					interval{7, 11},
					interval{29, 40},
				},
			},
			want: []interval{
				interval{5, 6},
				interval{20, 28},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := subtractSequences(tt.args.includes, tt.args.excludes)
			if (err != nil) != tt.wantErr {
				t.Errorf("subtractSequences() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subtractSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_subtract(t *testing.T) {
	type args struct {
		a interval
		b interval
	}
	tests := []struct {
		name string
		args args
		want []interval
	}{
		{
			name: "a is before b",
			args: args{
				a: interval{From: 10, To: 20},
				b: interval{From: 15, To: 25},
			},
			want: []interval{interval{10, 14}},
		},
		{
			name: "a is after b",
			args: args{
				a: interval{From: 10, To: 20},
				b: interval{From: 5, To: 15},
			},
			want: []interval{interval{16, 20}},
		},
		{
			name: "b splits a in two",
			args: args{
				a: interval{From: 10, To: 30},
				b: interval{From: 15, To: 20},
			},
			want: []interval{
				interval{10, 14},
				interval{21, 30},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := subtract(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subtract() = %v, want %v", got, tt.want)
			}
		})
	}
}
