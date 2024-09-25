package main

import (
	"AOC2023/utils"
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

//////////////
//HAND TYPE///
//////////////

type Hand struct {
	cardString       string
	cardMultiplicity map[int]int
	cards            []int
	bid              int
	handType         int
}

type HandSorter []Hand

func (hs HandSorter) Len() int {
	return len(hs)
}

func (hs HandSorter) Less(i, j int) bool {
	if hs[i].handType != hs[j].handType {
		return hs[i].handType < hs[j].handType
	}
	for k := 0; k < len(hs[i].cards); k++ {
		if hs[i].cards[k] != hs[j].cards[k] {
			return hs[i].cards[k] < hs[j].cards[k]
		}
	}
	return false
}

func (hs HandSorter) Swap(i, j int) {
	hs[i], hs[j] = hs[j], hs[i]
}

/////////////
//CARD MAPS//
/////////////

var CARD map[rune]int = map[rune]int{
	'2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8,
	'T': 9, 'J': 10, 'Q': 11, 'K': 12, 'A': 13}

var CARD2 map[rune]int = map[rune]int{
	'J': 0, '2': 1, '3': 2, '4': 3, '5': 4, '6': 5, '7': 6, '8': 7, '9': 8,
	'T': 9, 'Q': 10, 'K': 11, 'A': 12}

////////////
//PART ONE//
////////////

// 5 of a kind -> 6, 4 of a kind -> 5, Full house -> 4,
// 3 of a kind -> 3, 2 pairs -> 2, 1 pair -> 1,
// High-card -> 0
func getType(cards map[int]int) int {
	if len(cards) == 5 {
		return 0
	} else if len(cards) == 1 {
		return 6
	} else if len(cards) == 2 {
		for card := range cards {
			if cards[card] == 4 || cards[card] == 1 {
				return 5
			} else {
				return 4
			}
		}
	} else if len(cards) == 3 {
		for card := range cards {
			if cards[card] == 3 {
				return 3
			}
		}
		return 2
	} else if len(cards) == 4 {
		return 1
	}
	panic("Not supposed to happen.")
}

func getHandFromLine(line string) Hand {
	var components = strings.Split(line, " ")
	var h Hand = Hand{}
	h.cardMultiplicity = make(map[int]int)
	for _, r := range components[0] {
		h.cardMultiplicity[CARD[r]] += 1
		h.cards = append(h.cards, CARD[r])
	}
	h.bid, _ = strconv.Atoi(components[1])
	h.handType = getType(h.cardMultiplicity)
	h.cardString = components[0]
	return h
}

func getHands(input []string) []Hand {
	var hands = make([]Hand, 0)
	for _, line := range input {
		var h Hand = getHandFromLine(line)
		hands = append(hands, h)
	}
	return hands
}

func partOne(input string) int {
	var score = 0
	var sortedHands = getHands(utils.LineByLine(input))
	sort.Sort(HandSorter(sortedHands))
	for i, hand := range sortedHands {
		score += (i + 1) * hand.bid
	}
	return score
}

////////////
//PART TWO//
////////////

func getType2(cards map[int]int) int {
	var originalType = getType(cards)
	var numberJokers = cards[CARD2['J']]
	if numberJokers == 0 || numberJokers == 5 {
		return originalType
	} else if numberJokers == 4 {
		return 6
	} else if numberJokers == 3 {
		return originalType + 2
	} else if numberJokers == 2 {
		if originalType == 1 {
			return 3
		} else if originalType == 2 {
			return 5
		} else if originalType == 4 {
			return 6
		}
	} else if numberJokers == 1 {
		if originalType == 0 || originalType == 5 {
			return originalType + 1
		} else if originalType == 1 || originalType == 2 || originalType == 3 {
			return originalType + 2
		}
	}
	panic("Impossible case")
}

func getHandFromLine2(line string) Hand {
	var components = strings.Split(line, " ")
	var h Hand = Hand{}
	h.cardMultiplicity = make(map[int]int)
	for _, r := range components[0] {
		h.cardMultiplicity[CARD2[r]] += 1
		h.cards = append(h.cards, CARD2[r])
	}
	h.bid, _ = strconv.Atoi(components[1])
	h.handType = getType2(h.cardMultiplicity)
	h.cardString = components[0]
	return h
}

func getHands2(input []string) []Hand {
	var hands = make([]Hand, 0)

	for _, line := range input {
		var h Hand = getHandFromLine2(line)
		hands = append(hands, h)
	}
	return hands
}

func partTwo(input string) int {
	var score = 0
	var sortedHands = getHands2(utils.LineByLine(input))
	sort.Sort(HandSorter(sortedHands))
	for i, hand := range sortedHands {
		score += (i + 1) * hand.bid
	}
	return score
}

func main() {
	var start = time.Now()
	fmt.Println(partOne("07/input.txt"))
	fmt.Println(time.Since(start))
	start = time.Now()
	fmt.Println(partTwo("07/input.txt"))
	fmt.Println(time.Since(start))
}
