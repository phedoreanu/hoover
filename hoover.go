package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Patch defines a square in the room
// with X and Y coordinates
type Patch struct {
	X, Y int
}

// NewPatch returns a new Patch from string coordinates
func NewPatch(s string) Patch {
	c := strings.Split(s, " ")
	return Patch{X: stringToInt(c[0]), Y: stringToInt(c[1])}
}

// String handles Patch pretty print
func (p Patch) String() string {
	return fmt.Sprintf("%d %d", p.X, p.Y)
}

// Hoover is a imaginary robotic hoover
type Hoover struct {
	Location          Patch
	Room              map[Patch]bool
	Cleaned, Boundary int
}

// String handles Hoover pretty print
func (h *Hoover) String() string {
	return fmt.Sprintf("%s\n%d", h.Location, h.Cleaned)
}

func stringToInt(s string) (x int) {
	x, _ = strconv.Atoi(s)
	return
}

// Walk moves the robot through the room
func (h *Hoover) Walk(steps []byte) {
	for _, s := range steps {
		switch {
		case h.Location.Y < h.Boundary && s == 78: // N
			h.Location.Y++
		case h.Location.Y > 0 && s == 83: // S
			h.Location.Y--
		case h.Location.X < h.Boundary && s == 69: // E
			h.Location.X++
		case h.Location.X > 0 && s == 87: // W
			h.Location.X--
		}

		if h.Room[h.Location] {
			h.Cleaned++
			delete(h.Room, h.Location)
		}
	}
}

// Vacuum hoovers the room `filename`
func (h *Hoover) Vacuum(filename string) string {
	input, _ := ioutil.ReadFile(filename)
	bufSrc := bytes.NewReader(input)
	scanner := bufio.NewScanner(bufSrc)

	var steps []byte

	lineNo := 0
	for scanner.Scan() {
		if len(scanner.Bytes()) == 0 { // skip empty lines
			continue
		}

		b := scanner.Bytes()[0]
		switch {
		case 48 <= b && b <= 57: // numbers
			switch lineNo {
			case 0:
				h.Room = make(map[Patch]bool)
				h.Boundary = stringToInt(strings.Split(scanner.Text(), " ")[0])
			case 1:
				h.Location = NewPatch(scanner.Text())
			default:
				dirtyPatch := NewPatch(scanner.Text())
				h.Room[dirtyPatch] = true
			}
		case 65 <= b && b <= 90: // characters
			for _, c := range scanner.Bytes() {
				steps = append(steps, c)
			}
		}
		lineNo++
	}
	h.Walk(steps)

	return h.String()
}

func main() {
	res := new(Hoover).Vacuum("input.txt")

	fmt.Println(res)
}
