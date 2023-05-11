package main

import (
	"fmt"

	"github.com/M2ABA2/five_card_game/deck"
	//"github.com/M2ABA2/five_card_game/rank"

	"github.com/chehsunliu/poker"
)

func main() {
	
	//Get shuffled cards from the dack package.
	cards := deck.New()
	c     := cards[0:5]
	fmt.Println("Your hand: ",c)


	deck := poker.NewDeck()
	hand := deck.Draw(7)
	fmt.Println(hand)

	rank := poker.Evaluate(hand)
	fmt.Println(rank)
	fmt.Println(poker.RankString(rank))
}
	//From the deck retuned, rank.
	// d := rank.RankHand(c)
	 //fmt.Println(d)

	 // Create a hand of cards
	/*hand := rank.Hand{
		rank.Card{c},
		
	}

	// Call the RankHand function
	result := rank.RankHand(hand)

	// Print the result
	fmt.Println(result)*/
//}
