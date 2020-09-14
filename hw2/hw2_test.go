package hw2

import "testing"

func TestUnpack(t *testing.T) {
	type args struct {
		packedStr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{"1", args{"a4bc2d5e"}, "aaaabccddddde", false},
		{"2", args{"abcd"}, "abcd", false},
		{"3", args{"45"}, "", true},
		{"4", args{""}, "", false},
		/*
			{"5", args{`qwe\4\5`}, "qwe45", false},
			{"6", args{`qwe\45`}, "qwe44444", false},
			{"7", args{`qwe\\5`}, `qwe\\\\`, false},
		*/
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Unpack(tt.args.packedStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unpack() error = %v, wantErr %v, name %s", err, tt.wantErr, tt.name)
				return
			}
			if got != tt.want {
				t.Errorf("\nUnpack() = %s, want %s, name %s\n\n", got, tt.want, tt.name)
			}
		})
	}
}
