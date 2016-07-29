package main

import "testing"

func TestHoover(t *testing.T) {

	cases := []struct {
		in, want string
	}{
		{"input.txt", "1 3\n1"},
		{"input_boundary.txt", "2 2\n2"},
		{"input_room_clean.txt", "1 3\n0"},
		{"input_empty.txt", "0 0\n0"},
	}

	for _, c := range cases {
		got := new(Hoover).Vacuum(c.in)
		if got != c.want {
			t.Errorf("Vacuum(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
