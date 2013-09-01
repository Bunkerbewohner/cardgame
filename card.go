package main

import gl "github.com/go-gl/gl"
import glh "github.com/go-gl/glh"
import rand "math/rand"

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

func drawArrows(arrows [8]bool) {
	for i := 0; i < 8; i++ {
		// determine arrow position on card
		x, y := 0.0, 0.0
		if i < 3 {
			x = float64(i) * (CardWidth / 2.0)
		} else if i == 3 {
			x = CardWidth
			y = CardHeight / 2.0
		} else if i < 6 {
			x = (6.0 - float64(i)) * (CardWidth / 2.0)
			y = CardHeight
		} else {
			y = CardHeight / 2.0
		}

		if arrows[i] {
			// draw the arrow
			gl.Color3f(1.0, 1.0, 0)
			glh.DrawQuadd(x, y, 2, 2)
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
