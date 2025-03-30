package solitaire

type Deck struct {
	order    [54]Card
	position int
	Value    *Card
}

func (d *Deck) MoveCurrent(by int) {
	// Move the card at the current position by the specified number of positions
	// in the deck. If by is negative, move the card to the left; if by is positive, move it to the right.
	// If the card moves past the end of the deck, wrap around to the beginning.
	offset := by
	if d.position+by >= len(d.order) {
		offset = by + 1
	}
	if by > 0 {
		copy(d.order[:], moveInt(d.order[:], d.position, (d.position+offset)%len(d.order)))
	} else if by < 0 {
		copy(d.order[:], moveInt(d.order[:], d.position, (d.position+len(d.order)+offset)%len(d.order)))
	}
}
func (d *Deck) Position() int {
	// Return the current position in the deck.
	return d.position
}
func (d *Deck) SetPosition(pos int) {
	// Set the current position in the deck to the specified value.
	if pos < 0 || pos >= len(d.order) {
		panic("Position out of bounds")
	}
	d.position = pos
	d.Value = &d.order[d.position]
}

func (d *Deck) Next() {
	// Move to the next card in the deck.
	d.position = (d.position + 1) % len(d.order)
	d.Value = &d.order[d.position]
}

func (d *Deck) Previous() {
	// Move to the previous card in the deck.
	d.position = (d.position - 1 + len(d.order)) % len(d.order)
	d.Value = &d.order[d.position]
}

func (d *Deck) find(card Card) int {
	// Find the index of the specified card in the deck.
	for i, c := range d.order {
		if c == card {
			return i
		}
	}
	return -1
}
func (d *Deck) FindBlackJoker() int {
	// Find the index of the black joker in the deck.
	return d.find(Card{color: jokers, card: jokerB})
}

func (d *Deck) FindRedJoker() int {
	// Find the index of the red joker in the deck.
	return d.find(Card{color: jokers, card: jokerA})
}

func (d *Deck) FindFirstJoker() int {
	// Find the index of the first joker in the deck.
	for i, c := range d.order {
		if c.color == jokers {
			return i
		}
	}
	return -1
}
func (d *Deck) FindLastJoker() int {
	// Find the index of the last joker in the deck.
	for i := len(d.order) - 1; i >= 0; i-- {
		if d.order[i].color == jokers {
			return i
		}
	}
	return -1
}

func (d *Deck) TripleCut() {
	f := d.FindFirstJoker()
	l := d.FindLastJoker()

	top := d.order[:f]
	bottom := d.order[l+1:]
	middle := d.order[f+1 : l]
	newOrder := append(append(append(append(bottom, d.order[f]), middle...), d.order[l]), top...)
	copy(d.order[:], newOrder)
}

func (d *Deck) CountCut() {
	// Count the number of cards in the deck and cut the deck at that position.
	// The number of cards is determined by the value of the card at the bottom of the deck.
	bottomCard := d.order[len(d.order)-1]
	cutPosition := bottomCard.Value() % len(d.order)
	d.countCut(cutPosition)
}

func (d *Deck) countCut(cut int) {
	// Count the number of cards in the deck and cut the deck at that position.
	// The number of cards is determined by the value of the card at the bottom of the deck.
	if cut == 0 {
		return
	}

	top := d.order[:cut]
	bottom := d.order[cut : len(d.order)-1]
	// The last card is not included in the cut
	lastCard := d.order[len(d.order)-1]
	newOrder := append(append(bottom, top...), lastCard)
	copy(d.order[:], newOrder)
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func insertInt(array []Card, value Card, index int) []Card {
	return append(array[:index], append([]Card{value}, array[index:]...)...)
}

func removeInt(array []Card, index int) []Card {
	return append(array[:index], array[index+1:]...)
}

func moveInt(array []Card, srcIndex int, dstIndex int) []Card {
	value := array[srcIndex]
	return insertInt(removeInt(array, srcIndex), value, dstIndex)
}
