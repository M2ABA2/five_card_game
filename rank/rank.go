package rank

/*import (
	"fmt"

	"github.com/M2ABA2/five_card_game/deck"
)*/

type Card struct {
    Suit string
    Value int
}

type Hand [5]Card

func RankHand(hand Hand) string {
    sortedHand := sortHand(hand)
    if isRoyalFlush(sortedHand) {
        return "Royal Flush"
    }
    if isStraightFlush(sortedHand) {
        return "Straight Flush"
    }
    if isFourOfAKind(sortedHand) {
        return "Four of a Kind"
    }
    if isFullHouse(sortedHand) {
        return "Full House"
    }
    if isFlush(sortedHand) {
        return "Flush"
    }
    if isStraight(sortedHand) {
        return "Straight"
    }
    if isThreeOfAKind(sortedHand) {
        return "Three of a Kind"
    }
    if isTwoPair(sortedHand) {
        return "Two Pair"
    }
    if isOnePair(sortedHand) {
        return "One Pair"
    }
    return "High Card"
}

func sortHand(hand Hand) Hand {
    for i := 0; i < len(hand)-1; i++ {
        for j := i + 1; j < len(hand); j++ {
            if hand[i].Value > hand[j].Value {
                hand[i], hand[j] = hand[j], hand[i]
            }
        }
    }
    return hand
}

func isRoyalFlush(hand Hand) bool {
    if hand[0].Value == 10 && hand[1].Value == 11 && hand[2].Value == 12 && hand[3].Value == 13 && hand[4].Value == 14 && isFlush(hand) {
        return true
    }
    return false
}

func isStraightFlush(hand Hand) bool {
    if isFlush(hand) && isStraight(hand) {
        return true
    }
    return false
}

func isFourOfAKind(hand Hand) bool {
    if hand[0].Value == hand[1].Value && hand[1].Value == hand[2].Value && hand[2].Value == hand[3].Value {
        return true
    }
    if hand[1].Value == hand[2].Value && hand[2].Value == hand[3].Value && hand[3].Value == hand[4].Value {
        return true
    }
    return false
}

func isFullHouse(hand Hand) bool {
    if hand[0].Value == hand[1].Value && hand[1].Value == hand[2].Value && hand[3].Value == hand[4].Value {
        return true
    }
    if hand[0].Value == hand[1].Value && hand[2].Value == hand[3].Value && hand[3].Value == hand[4].Value {
        return true
    }
    return false
}

func isFlush(hand Hand) bool {
    if hand[0].Suit == hand[1].Suit && hand[1].Suit == hand[2].Suit && hand[2].Suit == hand[3].Suit && hand[3].Suit == hand[4].Suit {
        return true
    }
    return false
}

func isStraight(hand Hand) bool {
    if hand[0].Value == hand[1].Value-1 && hand[1].Value == hand[2].Value-1 && hand[2].Value == hand[3].Value-1 && hand[3].Value == hand[4].Value-1 {
        return true
    }
    // check for ace low straight
    if hand[0].Value == 2 && hand[1].Value == 3 && hand[2].Value == 4 && hand[3].Value == 5 && hand[4].Value == 14 {
        return true
    }
    return false
}

func isThreeOfAKind(hand Hand) bool {
    for i := 0; i < len(hand)-2; i++ {
        if hand[i].Value == hand[i+1].Value && hand[i+1].Value == hand[i+2].Value {
            return true
        }
    }
    return false
}

func isTwoPair(hand Hand) bool {
    pairs := 0
    for i := 0; i < len(hand)-1; i++ {
        if hand[i].Value == hand[i+1].Value {
            pairs++
            i++ // skip over the pair
        }
    }
    return pairs == 2
}

func isOnePair(hand Hand) bool {
    for i := 0; i < len(hand)-1; i++ {
        if hand[i].Value == hand[i+1].Value {
            return true
        }
    }
    return false
}
