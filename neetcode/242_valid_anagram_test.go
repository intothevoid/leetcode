package neetcode

import "testing"

func TestIsAnagram(t *testing.T) {
	type args struct {
		s string
		t string
	}
	tests := []struct {
		name string
		args args
	}{
		{name: "test1", args: args{s: "anagram", t: "gramana"}},
		{name: "test2", args: args{s: "weak", t: "kewa"}},
		{name: "test3", args: args{s: "stone", t: "onesi"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("%s - %t", tt.name, IsAnagram(tt.args.s, tt.args.t))
		})
	}
}
