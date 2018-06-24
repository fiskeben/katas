package main

import (
	"reflect"
	"testing"
)

func Test_parseRawNumber(t *testing.T) {
	type args struct {
		number rawAccountNumber
	}
	tests := []struct {
		name string
		args args
		want number
	}{
		{
			name: "parses 1",
			args: args{
				number: rawAccountNumber{"    ", "  |", "  |"},
			},
			want: number{digit{[]rune{' ', ' ', ' '}, []rune{' ', ' ', '|'}, []rune{' ', ' ', '|'}}},
		},
		{
			name: "parses 27",
			args: args{
				number: rawAccountNumber{" _  _ ", " _|  |", "|_   |"},
			},
			want: number{
				digit{[]rune{' ', '_', ' '}, []rune{' ', '_', '|'}, []rune{'|', '_', ' '}},
				digit{[]rune{' ', '_', ' '}, []rune{' ', ' ', '|'}, []rune{' ', ' ', '|'}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseRawNumber(tt.args.number); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseRawNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deduce(t *testing.T) {
	type args struct {
		a digit
	}
	tests := []struct {
		name string
		args args
		want rune
	}{
		{
			name: "can deduce 1",
			args: args{
				a: digit{[]rune{' ', ' ', ' '}, []rune{' ', ' ', '|'}, []rune{' ', ' ', '|'}},
			},
			want: '1',
		},
		{
			name: "can deduce 2",
			args: args{
				a: digit{[]rune{' ', '_', ' '}, []rune{' ', '_', '|'}, []rune{'|', '_', ' '}},
			},
			want: '2',
		},
		{
			name: "can deduce 3",
			args: args{
				a: digit{[]rune{' ', '_', ' '}, []rune{' ', '_', '|'}, []rune{' ', '_', '|'}},
			},
			want: '3',
		},
		{
			name: "can deduce 4",
			args: args{
				a: digit{[]rune{' ', ' ', ' '}, []rune{'|', '_', '|'}, []rune{' ', ' ', '|'}},
			},
			want: '4',
		},
		{
			name: "can deduce 5",
			args: args{
				a: digit{[]rune{' ', '_', ' '}, []rune{'|', '_', ' '}, []rune{' ', '_', '|'}},
			},
			want: '5',
		},
		{
			name: "can deduce 6",
			args: args{
				a: digit{[]rune{' ', '_', ' '}, []rune{'|', '_', ' '}, []rune{'|', '_', '|'}},
			},
			want: '6',
		},
		{
			name: "can deduce 7",
			args: args{
				a: digit{[]rune{' ', '_', ' '}, []rune{' ', ' ', '|'}, []rune{' ', ' ', '|'}},
			},
			want: '7',
		},
		{
			name: "can deduce 8",
			args: args{
				a: digit{[]rune{' ', '_', ' '}, []rune{'|', '_', '|'}, []rune{'|', '_', '|'}},
			},
			want: '8',
		},
		{
			name: "can deduce 9",
			args: args{
				a: digit{[]rune{' ', '_', ' '}, []rune{'|', '_', '|'}, []rune{' ', '_', '|'}},
			},
			want: '9',
		},
		{
			name: "can deduce 0",
			args: args{
				a: digit{[]rune{' ', '_', ' '}, []rune{'|', ' ', '|'}, []rune{'|', '_', '|'}},
			},
			want: '0',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deduce(tt.args.a); got != tt.want {
				t.Errorf("deduce() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parse(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want []accountnumber
	}{
		{
			name: "parses all zeros",
			args: args{
				data: ` _  _  _  _  _  _  _  _  _ 
| || || || || || || || || |
|_||_||_||_||_||_||_||_||_|
`,
			},
			want: []accountnumber{"000000000"},
		},
		{
			name: "parses 111111111",
			args: args{
				data: `                           
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |
`,
			},
			want: []accountnumber{"111111111"},
		},
		{
			name: "parses 222222222",
			args: args{
				data: ` _  _  _  _  _  _  _  _  _ 
 _| _| _| _| _| _| _| _| _|
|_ |_ |_ |_ |_ |_ |_ |_ |_ 
`,
			},
			want: []accountnumber{"222222222"},
		},
		{
			name: "parses 333333333",
			args: args{
				data: ` _  _  _  _  _  _  _  _  _ 
 _| _| _| _| _| _| _| _| _|
 _| _| _| _| _| _| _| _| _|
 `,
			},
			want: []accountnumber{"333333333"},
		},
		{
			name: "parses 444444444",
			args: args{
				data: `                           
|_||_||_||_||_||_||_||_||_|
  |  |  |  |  |  |  |  |  |
  `,
			},
			want: []accountnumber{"444444444"},
		},
		{
			name: "parses 555555555",
			args: args{
				data: ` _  _  _  _  _  _  _  _  _ 
|_ |_ |_ |_ |_ |_ |_ |_ |_ 
 _| _| _| _| _| _| _| _| _|
 `,
			},
			want: []accountnumber{"555555555"},
		},
		{
			name: "parses 666666666",
			args: args{
				data: ` _  _  _  _  _  _  _  _  _ 
|_ |_ |_ |_ |_ |_ |_ |_ |_ 
|_||_||_||_||_||_||_||_||_|
`,
			},
			want: []accountnumber{"666666666"},
		},
		{
			name: "parses 777777777",
			args: args{
				data: ` _  _  _  _  _  _  _  _  _ 
  |  |  |  |  |  |  |  |  |
  |  |  |  |  |  |  |  |  |
  `,
			},
			want: []accountnumber{"777777777"},
		},
		{
			name: "parses 888888888",
			args: args{
				data: ` _  _  _  _  _  _  _  _  _ 
|_||_||_||_||_||_||_||_||_|
|_||_||_||_||_||_||_||_||_|
`,
			},
			want: []accountnumber{"888888888"},
		},
		{
			name: "parses 999999999",
			args: args{
				data: ` _  _  _  _  _  _  _  _  _ 
|_||_||_||_||_||_||_||_||_|
 _| _| _| _| _| _| _| _| _|
 `,
			},
			want: []accountnumber{"999999999"},
		},
		{
			name: "parses 123456789",
			args: args{
				data: `    _  _     _  _  _  _  _ 
  | _| _||_||_ |_   ||_||_|
  ||_  _|  | _||_|  ||_| _|
  `,
			},
			want: []accountnumber{"123456789"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parse(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checksum(t *testing.T) {
	type args struct {
		n accountnumber
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "711111111 is valid",
			args: args{n: "711111111"},
			want: true,
		},
		{
			name: "123456789 is valid",
			args: args{n: "123456789"},
			want: true,
		},
		{
			name: "490867715 is valid",
			args: args{n: "490867715"},
			want: true,
		},
		{
			name: "888888888 is invalid",
			args: args{n: "888888888"},
			want: false,
		},
		{
			name: "490067715 is invalid",
			args: args{n: "490067715"},
			want: false,
		},
		{
			name: "012345678 is invalid",
			args: args{n: "012345678"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checksum(tt.args.n); got != tt.want {
				t.Errorf("checksum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeReport(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "generates OK report for 000000051",
			args: args{
				data: ` _  _  _  _  _  _  _  _     
| || || || || || || ||_   |
|_||_||_||_||_||_||_| _|  |
`,
			},
			want: "000000051",
		},
		{
			name: "49006771? is illegal",
			args: args{
				data: `    _  _  _  _  _  _     _ 
|_||_|| || ||_   |  |  | _ 
  | _||_||_||_|  |  |  | _|
`,
			},
			want: "49006771? ILL",
		},
		{
			name: "1234?678? is illegal",
			args: args{
				data: `    _  _     _  _  _  _  _ 
  | _| _||_| _ |_   ||_||_|
  ||_  _|  | _||_|  ||_| _ 
`,
			},
			want: "1234?678? ILL",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makeReport(tt.args.data); got != tt.want {
				t.Errorf("makeReport() = %v, want %v", got, tt.want)
			}
		})
	}
}
