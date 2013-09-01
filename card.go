package main

import gl "github.com/go-gl/gl"
import glh "github.com/go-gl/glh"

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
	NE Direction = iota
	N
	NW
	E
	SE
	S
	SW
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

// A deck card put into play
type PlayCard struct {
	Card *DeckCard

	// position on the playfield
	X, Y int

	// current owner of this card
	Owner int
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
}
