package solitaire

type Deck [54]Card

func (d *Deck) Advance() {
	// Move the current position to the next card in the deck.
	a := d.FindJokerA()
	d.Move(a, 1)
	b := d.FindJokerB()
	d.Move(b, 2)
	d.TripleCut()
	d.CountCut()
}

func (d *Deck) Move(pos, by int) {
	// Move the card at the specified position by the specified number of positions
	// in the deck.
	offset := by
	if pos+by >= len(d) {
		offset = by + 1
	}
	d.moveCard(pos, (pos+offset)%len(d))
}

func (d *Deck) find(card Card) int {
	// Find the index of the specified card in the deck.
	for i, c := range d {
		if c == card {
			return i
		}
	}
	return -1
}
func (d *Deck) FindJokerB() int {
	// Find the index of the black joker in the deck.
	return d.find(Card{color: Jokers, card: JokerB})
}

func (d *Deck) FindJokerA() int {
	return d.find(Card{color: Jokers, card: JokerA})
}

func (d *Deck) FindFirstJoker() int {
	// Find the index of the first joker in the deck.
	for i, c := range d {
		if c.color == Jokers {
			return i
		}
	}
	return -1
}
func (d *Deck) FindLastJoker() int {
	// Find the index of the last joker in the deck.
	for i := len(d) - 1; i >= 0; i-- {
		if d[i].color == Jokers {
			return i
		}
	}
	return -1
}

func (d *Deck) TripleCut() {
	f := d.FindFirstJoker()
	l := d.FindLastJoker()

	// TODO: This feels _wrong_.
	// It should be manipulating the deck in place, not copying it.
	top := d[:f]
	bottom := d[l+1:]
	middle := d[f+1 : l]
	newOrder := append(append(append(append(bottom, d[f]), middle...), d[l]), top...)
	copy(d[:], newOrder)
}

func (d *Deck) CountCut() {
	// Count the number of cards in the deck and cut the deck at that position.
	// The number of cards is determined by the value of the card at the bottom of the deck.
	bottomCard := d[len(d)-1]
	cutPosition := bottomCard.Value() % len(d)
	d.countCut(cutPosition)
}

func (d *Deck) countCut(cut int) {
	// Count the number of cards in the deck and cut the deck at that position.
	// The number of cards is determined by the value of the card at the bottom of the deck.
	if cut == 0 {
		return
	}

	top := d[:cut]
	bottom := d[cut : len(d)-1]
	// The last card is not included in the cut
	lastCard := d[len(d)-1]
	newOrder := append(append(bottom, top...), lastCard)
	copy(d[:], newOrder)
}

func (d *Deck) InsertCard(value Card, index int) {
	var n [54]Card
	copy(n[:], d[:index])
	copy(n[index:], []Card{value})
	copy(n[index+1:], d[index:])
	copy(d[:], n[:])
}

func (d *Deck) RemoveCard(index int) {
	copy(d[:], d[:index])
	copy(d[index:], d[index+1:])
}

func (d *Deck) moveCard(srcIndex int, dstIndex int) {
	value := d[srcIndex]
	// Remove the card from the source index
	d.RemoveCard(srcIndex)
	d.InsertCard(value, dstIndex)
}
