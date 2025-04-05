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

// A deck of cards used in the Solitaire algorithm.
// The deck consists of 54 cards: 52 standard playing cards and 2 jokers.
type deck [54]card

func (d *deck) advance() {
	// Move the current position to the next card in the deck.
	a := d.findJokerA()
	d.move(a, 1)
	b := d.findJokerB()
	d.move(b, 2)
	d.tripleCut()
	d.defaultCountCut()
}

func (d *deck) move(pos, by int) {
	// Move the card at the specified position by the specified number of positions
	// in the deck.
	offset := by
	if pos+by >= len(d) {
		offset = by + 1
	}
	d.moveCard(pos, (pos+offset)%len(d))
}

func (d *deck) find(card card) int {
	// Find the index of the specified card in the deck.
	for i, c := range d {
		if c == card {
			return i
		}
	}
	return -1
}
func (d *deck) findJokerA() int {
	return d.find(card{rank: jokerA})
}
func (d *deck) findJokerB() int {
	// Find the index of the fast joker in the deck.
	return d.find(card{rank: jokerB})
}

func (d *deck) findFirstJoker() int {
	// Find the index of the first joker in the deck.
	for i, c := range d {
		if c.IsJoker() {
			return i
		}
	}
	return -1
}
func (d *deck) findLastJoker() int {
	// Find the index of the last joker in the deck.
	for i := len(d) - 1; i >= 0; i-- {
		if d[i].IsJoker() {
			return i
		}
	}
	return -1
}

func (d *deck) tripleCut() {
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

func (d *deck) defaultCountCut() {
	// Count the number of cards in the deck and cut the deck at that position.
	// The number of cards is determined by the value of the card at the bottom of the deck.
	bottomCard := d[len(d)-1]
	cutPosition := bottomCard.Value() % len(d)
	d.countCut(cutPosition)
}

func (d *deck) countCut(cut int) {
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

func (d *deck) insertCard(value card, index int) {
	var n [54]card
	copy(n[:], d[:index])
	copy(n[index:], []card{value})
	copy(n[index+1:], d[index:])
	copy(d[:], n[:])
}

func (d *deck) removeCard(index int) {
	copy(d[:], d[:index])
	copy(d[index:], d[index+1:])
}

func (d *deck) moveCard(srcIndex int, dstIndex int) {
	value := d[srcIndex]
	// Remove the card from the source index
	d.removeCard(srcIndex)
	d.insertCard(value, dstIndex)
}

func (d *deck) key() int {
	d.advance()
	// The key is determined by the value of the card at the top of the deck.
	val := d[d[0].Value()].Value()
	if val >= 53 {
		// Skip the jokers
		return d.key()
	}
	return val
}
