package main

import gl "github.com/go-gl/gl"
import rand "math/rand"

type Slot struct {
	card    PlayCard
	x, y    int
	blocked bool
}

type Playfield struct {
	playerCards   [5]DeckCard
	opponentCards [5]DeckCard
	slots         [16]Slot
}

func NewPlayfield() *Playfield {
	playfield := new(Playfield)

	for i := 0; i < 5; i++ {
		playfield.playerCards[i] = *NewRandomDeckCard()
		playfield.playerCards[i].Owner = PlayerID

		playfield.opponentCards[i] = *NewRandomDeckCard()
		playfield.opponentCards[i].Owner = OpponentID
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			index := i*4 + j
			playfield.slots[index].x = i
			playfield.slots[index].y = j
			playfield.slots[index].card = *new(PlayCard)
			playfield.slots[index].card.Card = NewRandomDeckCard()
			playfield.slots[index].card.Owner = 1 + rand.Int()%2
			playfield.slots[index].card.X = i
			playfield.slots[index].card.Y = j
		}
	}

	return playfield
}

func drawPool(cards [5]DeckCard, x float32) {
	gl.PushMatrix()
	gl.Translatef(x, 590-CardHeight, 0)

	for i, card := range cards {
		if i > 0 {
			gl.Translatef(0, -(CardHeight + 5), 0)
		}

		card.Draw()
	}
	gl.PopMatrix()
}

func (p *Playfield) Draw() {
	drawPool(p.playerCards, 10)
	drawPool(p.opponentCards, 790-CardWidth)

	// draw the slots
	gl.PushMatrix()
	gl.Translatef(CardWidth+60, 580-CardHeight, 0)
	for i, slot := range p.slots {
		if i > 0 {
			gl.Translatef(CardWidth+2, 0, 0)
			if i%4 == 0 {
				gl.Translatef(-((CardWidth + 2) * 4), -(CardHeight + 2), 0)
			}
		}
		slot.card.Draw()
	}
	gl.PopMatrix()
}
