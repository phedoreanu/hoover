// Package hoover contains the source code of an imaginary robot Hoover
package hoover

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
	X, Y int64
}

// String handles Patch pretty print
func (p Patch) String() string {
	return fmt.Sprintf("%d %d", p.X, p.Y)
}

// NewPatch returns a new Patch from string coordinates
func NewPatch(s string) Patch {
	c := strings.Split(s, " ")
	return Patch{X: parseInt(c[0]), Y: parseInt(c[1])}
}

// Update calculates the next Patch to go
func (p *Patch) Update(n, b *Patch) {
	if p.X >= b.X && p.Y >= b.Y {
		return // skid in place
	}
	p.X += n.X
	p.Y += n.Y
}

// Hoover is an imaginary robotic hoover
type Hoover struct {
	Location, Boundary Patch
	Room               map[Patch]bool
	Dirty              int
}

// String returns the Hoover's last position
// and the number of Patches cleaned
func (h *Hoover) String() string {
	return fmt.Sprintf("%s\n%d\n", h.Location, h.Dirty-len(h.Room))
}

func parseInt(s string) int64 {
	i64, _ := strconv.ParseInt(s, 10, 64)
	return i64
}

// vacuum hoovers the room for given moves
func (h *Hoover) vacuum(moves string) {
	directions := map[byte]*Patch{
		78: {X: 0, Y: 1},  // N
		83: {X: 0, Y: -1}, // S
		69: {X: 1, Y: 0},  // E
		87: {X: -1, Y: 0}, // W
	}

	loc := &h.Location
	for _, d := range []byte(moves) {
		loc.Update(directions[d], &h.Boundary)
		delete(h.Room, h.Location) // clean dirty Patch
	}
}

func (h *Hoover) placeDirtyPatches(coordinates []string) {
	for _, c := range coordinates {
		dirtyPatch := NewPatch(c)
		h.Room[dirtyPatch] = true
		h.Dirty++
	}
}

// NewHoover returns a new `Hoover` instance
// from the given input filename
func NewHoover(filename string) (h *Hoover) {
	data, _ := ioutil.ReadFile(filename)
	if len(data) == 0 {
		return
	}
	scanner := bufio.NewScanner(bytes.NewReader(data))

	var cfg []string
	for scanner.Scan() {
		cfg = append(cfg, scanner.Text())
	}

	h = &Hoover{
		Boundary: NewPatch(cfg[0]),
		Location: NewPatch(cfg[1]),
		Room:     make(map[Patch]bool),
	}
	h.placeDirtyPatches(cfg[2 : len(cfg)-1])
	h.vacuum(cfg[len(cfg)-1])

	return
}
