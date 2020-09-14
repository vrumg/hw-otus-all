package top10

import (
	"reflect"
	"testing"
)

func TestTop10(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"first",
			args{"1 2 2 3 3 3 4 4 4 4 5 5 5 5 5 6 6 6 6 6 6 7 7 7 7 7 7 7 8 8 8 8 8 8 8 8 9 9 9 9 9 9 9 9 9 0 0 0 0 0 0 0 0 0 0"},
			[]string{"0", "9", "8", "7", "6", "5", "4", "3", "2", "1"},
		},
		{
			"second",
			args{"1 2 2 3 3 3 4 4 4 4 5 5 5 5 5 6 6 6 6 6 6 7 7 7 7 7 7 7 8 8 8 8 8 8 8 8 9 9 9 9 9 9 9 9 9 0 0 0 0 0 0 0 0 0 0 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,"},
			[]string{"1,", "0", "9", "8", "7", "6", "5", "4", "3", "2"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Top10(tt.args.text); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Top10() = %v, want %v", got, tt.want)
			}
		})
	}
}
