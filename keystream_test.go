package solitaire

// type KeyStreamSuite struct {
// 	suite.Suite
// }

// func (s *KeyStreamSuite) TestColorString() {
// 	testCases := []struct {
// 		color    color
// 		expected string
// 	}{
// 		{clubs, "C"},
// 		{diamonds, "D"},
// 		{hearts, "H"},
// 		{spades, "S"},
// 	}

// 	for _, tC := range testCases {
// 		s.Run(tC.expected, func() {
// 			s.Equal(tC.expected, tC.color.String())
// 		})
// 	}
// }
// func (s *KeyStreamSuite) TestColorValue() {
// 	testCases := []struct {
// 		color    color
// 		card     int
// 		expected int
// 	}{
// 		{clubs, 13, 0},
// 		{diamonds, 26, 0},
// 		{hearts, 39, 0},
// 		{spades, 52, 0},
// 	}

// 	for _, tC := range testCases {
// 		s.Run(tC.color.String(), func() {
// 			s.Equal(tC.expected, tC.color.Value(tC.card))
// 		})
// 	}
// }

// func (s *KeyStreamSuite) TestColorValueWithCards() {
// 	testCases := []struct {
// 		desc     string
// 		color    color
// 		card     int
// 		expected int
// 	}{
// 		{desc: "one of clubs", color: clubs, card: 1, expected: 1},
// 		{desc: "two of diamonds", color: diamonds, card: 2, expected: 15},
// 		{desc: "three of hearts", color: hearts, card: 3, expected: 3},
// 		{desc: "four of spades", color: spades, card: 4, expected: 17},
// 		{desc: "five of clubs", color: clubs, card: 5, expected: 5},
// 		{desc: "six of diamonds", color: diamonds, card: 6, expected: 19},
// 		{desc: "seven of hearts", color: hearts, card: 7, expected: 7},
// 		{desc: "eight of spades", color: spades, card: 8, expected: 21},
// 		{desc: "ace of clubs", color: clubs, card: 13, expected: 13},
// 		{desc: "ace of diamonds", color: diamonds, card: 13, expected: 26},
// 	}
// 	for _, tC := range testCases {
// 		s.Run(tC.desc, func() {
// 			s.Equal(tC.expected, tC.color.Value(tC.card))
// 		})
// 	}
// }

// //	func TestSetup(t *testing.T) {
// //		setup("foo")
// //	}
// // func TestKeyStreamSuite(t *testing.T) {
// // 	suite.Run(t, new(KeyStreamSuite))
// // }
