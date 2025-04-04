/*
 *  Copyright 2025 Markus Mahlberg <138420+mwmahlberg@users.noreply.github.com>
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 *
 */

package solitaire

type Deck [54]Card

func (d *Deck) advance() {
	// Move the current position to the next card in the deck.
	a := d.findJokerA()
	d.move(a, 1)
	b := d.findJokerB()
	d.move(b, 2)
	d.tripleCut()
	d.defaultCountCut()
}

func (d *Deck) move(pos, by int) {
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
func (d *Deck) findJokerB() int {
	// Find the index of the black joker in the deck.
	return d.find(Card{rank: jokerB})
}

func (d *Deck) findJokerA() int {
	return d.find(Card{rank: jokerA})
}

func (d *Deck) findFirstJoker() int {
	// Find the index of the first joker in the deck.
	for i, c := range d {
		if c.rank == jokerB || c.rank == jokerA {
			return i
		}
	}
	return -1
}
func (d *Deck) findLastJoker() int {
	// Find the index of the last joker in the deck.
	for i := len(d) - 1; i >= 0; i-- {
		if d[i].rank == jokerA || d[i].rank == jokerB {
			return i
		}
	}
	return -1
}

func (d *Deck) tripleCut() {
	f := d.findFirstJoker()
	l := d.findLastJoker()

	// TODO: This feels _wrong_.
	// It should be manipulating the deck in place, not copying it.
	top := d[:f]
	bottom := d[l+1:]
	middle := d[f+1 : l]
	newOrder := append(append(append(append(bottom, d[f]), middle...), d[l]), top...)
	copy(d[:], newOrder)
}

func (d *Deck) defaultCountCut() {
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

func (d *Deck) insertCard(value Card, index int) {
	var n [54]Card
	copy(n[:], d[:index])
	copy(n[index:], []Card{value})
	copy(n[index+1:], d[index:])
	copy(d[:], n[:])
}

func (d *Deck) removeCard(index int) {
	copy(d[:], d[:index])
	copy(d[index:], d[index+1:])
}

func (d *Deck) moveCard(srcIndex int, dstIndex int) {
	value := d[srcIndex]
	// Remove the card from the source index
	d.removeCard(srcIndex)
	d.insertCard(value, dstIndex)
}

func (d *Deck) key() int {
	d.advance()
	// The key is determined by the value of the card at the top of the deck.
	val := d[d[0].Value()].Value()
	if val >= 53 {
		// Skip the jokers
		return d.key()
	}
	return val
}
