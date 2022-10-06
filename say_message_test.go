package main

import "testing"

func TestSayMessage(t *testing.T) {
	type args struct {
		msg string
	}
	tests := []struct {
		name string
		args args
		ans  string
	}{
		{
			name: "normal",
			args: args{msg: "say"},
			ans:  "Ahoy, world!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SayMessage(tt.args.msg); got != tt.ans {
				t.Errorf("SayMessage() = %v, ans %v", got, tt.args.msg)
			}
		})
	}
}
