package deck

import (
	"fmt"
	"math/rand"
	"time"
)

type Suit int
type Deck [52]Card

type Card struct {
	suit  Suit
	value int
}

const (
	Spades Suit = iota
	Hearts
	Diamonds
	Clubs
)

//Cards returned as their respective unitcodes.
func (s Suit) String() string {
	switch s {
	case Spades:
		return  "♠"//"Spades"
	case Hearts:
		return "♥"//"Hearts"
	case Diamonds:
		return "♦"//"Diamonds"
	case Clubs:
		return "♣"//"Clubs"
	default:
		//Do nothing
		panic("Invalid Card")
	}
}

//Returns all the cards
func (c Card) String() string {
	
	return fmt.Sprintf("%d %s", c.value, c.suit)
}

func NewCard(s Suit, v int) Card {
	return Card{
		suit:  s,
		value: v,
	}
}


//Card Implementation

func New() Deck {
	var (
		nSuits = 4
		nCards = 13
		d      = [52]Card{}
	)

	x := 0
	for i := 0; i < nSuits; i++ {
		for j := 0; j < nCards; j++ {
			d[x] = NewCard(Suit(i), j+1)
			x++
		}
	}
	//Returns all the cards
	//return d
    
	//Random number generator
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Shuffling ... Shuffling ... Shufling ...")

	//Deck Shuffled
	for i := range d {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}

    // Create a new deck with only the first 5 shuffled cards
	for i := 0; i < 5; i++ {
		//fmt.Println(d[i])
		
	}
	
	var result Deck
	copy(result[:], d[:5])

	return result
}

