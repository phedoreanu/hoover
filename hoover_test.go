package hoover_test

import (
	"fmt"
	"testing"

	"github.com/phedoreanu/hoover"
)

func Example() {
	fmt.Println(hoover.NewHoover("input.txt"))
	// Output:
	// 1 3
	// 1
}

func ExampleNewHoover() {
	fmt.Println(hoover.NewHoover("input_boundary.txt"))
	// Output:
	// 2 2
	// 2
}

func TestHoover(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"input_face.txt", "<nil>"},
		{"input_empty.txt", "<nil>"},
		{"input_room_clean.txt", "1 3\n0"},
	}

	for _, c := range cases {
		got := fmt.Sprint(hoover.NewHoover(c.in))
		if got != c.want {
			t.Errorf("Vacuum(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
