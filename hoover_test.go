package hoover_test

import (
	"fmt"

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

func ExampleFakeInput() {
	fmt.Println(hoover.NewHoover("input_fake.txt"))
	// Output: <nil>
}

func ExampleEmptyInput() {
	fmt.Println(hoover.NewHoover("input_empty.txt"))
	// Output: <nil>
}

func ExampleCleanRoom() {
	fmt.Println(hoover.NewHoover("input_room_clean.txt"))
	// Output:
	// 1 3
	// 0
}
