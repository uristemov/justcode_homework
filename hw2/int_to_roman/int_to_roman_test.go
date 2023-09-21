package main

import "testing"

func Test_intToRoman(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test 1",
			args{
				12,
			},
			"XII",
		},
		{
			"Test 2",
			args{
				3657,
			},
			"MMMDCLVII",
		},
		{
			"Test 3",
			args{
				1937,
			},
			"MCMXXXVII",
		},
		{
			"Test 4",
			args{
				1994,
			},
			"MCMXCIV",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intToRoman(tt.args.num); got != tt.want {
				t.Errorf("intToRoman() = %v, want %v", got, tt.want)
			}
		})
	}
}
