package parser

import (
	"testing"
)

func TestParseDigit(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		{"test 1",
			args{"1"},
			"",
			'1',
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ParseDigit(tt.args.src)
			if got != tt.want {
				t.Errorf("ParseDigit() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseDigit() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestParseNumber(t *testing.T) {
	type args struct {
		src string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		{
			"test1",
			args{"123456"},
			"",
			123456,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := ParseNumber(tt.args.src)
			if got != tt.want {
				t.Errorf("ParseNumber() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("ParseNumber() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
