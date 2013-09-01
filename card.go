package main

import gl "github.com/go-gl/gl"
import glh "github.com/go-gl/glh"
import rand "math/rand"
import "math"

type BattleClass int
type Direction int
type Strength byte // 0-15 (0..9,A,B,C,D,E)

const (
	PlayerID   = 1
	OpponentID = 2
)

const (
	P BattleClass = iota // physical
	M                    // magical
	X                    // flexible
	A                    // assault
)

const (
	SW Direction = iota
	S
	SE
	E
	NE
	N
	NW
	W
)

const (
	CardHeight = 140
	CardWidth  = 110
)

// A card in the deck that can be selected for games
type DeckCard struct {
	Owner int

	// stats
	Power           Strength
	Class           BattleClass
	PhysicalDefense Strength
	MagicalDefense  Strength

	// arrows
	Arrows [8]bool
}

func NewRandomDeckCard() *DeckCard {
	card := new(DeckCard)
	for i := 0; i < 8; i++ {
		card.Power = Strength(rand.Int() % 15)
		card.PhysicalDefense = Strength(rand.Int() % 15)
		card.MagicalDefense = Strength(rand.Int() % 15)
		card.Class = BattleClass(rand.Int() % 4)
		card.Arrows[i] = rand.Int()%2 == 0
	}
	return card
}

// A deck card put into play
type PlayCard struct {
	Card *DeckCard

	// position on the playfield
	X, Y int

	// current owner of this card
	Owner int
}

// Maps an arrow index (0=SW, 1=S, ..., 7=W) to
// an radian angle between PI/8 and 2PI
func arrowIndexToAngle(i Direction) float64 {
	return (2 * math.Pi) / 8.0 * float64(1+i)
}

// arrowCos computes the counter-clockwise rounded cosinus for arrow i
// arrows on the left edge yield -1, arrows in the middle 0, arrows on the right
// edge 1.
func arrowCos(i Direction) float64 {
	angle := arrowIndexToAngle(i)
	cos := math.Cos(angle)
	return -1 * Round(cos)
}

// arrowSin computes the counter-clockwise rounded sinus for arrow i
// arrows on the bottom edge yield -1, arrows in the middle 0, arrows on the
// top edge 1
func arrowSin(i Direction) float64 {
	angle := arrowIndexToAngle(i)
	sin := math.Sin(angle)
	return -1 * Round(sin)
}

func drawArrows(arrows [8]bool) {
	for i := Direction(0); i < 8; i++ {
		// determine arrow position on card
		x := 0.5 * (1 + arrowCos(i)) * CardWidth
		y := 0.5 * (1 + arrowSin(i)) * CardHeight

		if arrows[i] {
			// draw the arrow
			gl.Color3f(1.0, 1.0, 0)
			glh.DrawQuadd(x-2, y-2, 4, 4)
		}
	}
}

func (card *DeckCard) Draw() {
	if card.Owner == PlayerID {
		gl.Color3f(0.5, 0.8, 0.3)
	} else if card.Owner == OpponentID {
		gl.Color3f(0.8, 0.5, 0.3)
	} else {
		gl.Color3f(0.3, 0.5, 0.8)
	}
	glh.DrawQuadi(0, 0, CardWidth, CardHeight)
	drawArrows(card.Arrows)
}

func (card *PlayCard) Draw() {
	if card.Owner == PlayerID {
		gl.Color3f(0.5, 0.8, 0.3)
	} else if card.Owner == OpponentID {
		gl.Color3f(0.8, 0.5, 0.3)
	} else {
		gl.Color3f(0.8, 0.8, 0.8)
	}
	glh.DrawQuadi(0, 0, CardWidth, CardHeight)
	drawArrows(card.Card.Arrows)
}
