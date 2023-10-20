package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Player struct {
	name         string
	holding      int
	holdingRange HoldingRange
}

type HoldingRange struct {
	start int
	end   int
}

func (p Player) play() int {
	score := 0
	for score <= p.holding {
		diceValue := rand.Int()%6 + 1
		if diceValue == 1 {
			return 0
		}
		score += diceValue
	}
	return score
}

func game(p1, p2 Player) (int, int) {
	p1Score, p2Score := 0, 0

	for p1Score < 100 && p2Score < 100 {
		p1Score += p1.play()
		if p1Score >= 100 {
			break
		}
		p2Score += p2.play()
	}
	return p1Score, p2Score
}

func (h HoldingRange) isStartAndEndSame() bool{
	return h.start == h.end
}

func match(p1, p2 Player, games int) {
	singleStrategyVsMultipleStrategy := p1.holdingRange.isStartAndEndSame() || p2.holdingRange.isStartAndEndSame()

	for i := p1.holdingRange.start; i <= p1.holdingRange.end; i++ {
		p1Wins, p2Wins := 0, 0
		totalGames := 0
		for j := p2.holdingRange.start; j <= p2.holdingRange.end; j++ {
			if i == j {
				continue
			}
			totalGames += 1
			g := 0
			

			for g < games {
				p1.holding = i
				p2.holding = j
				p1Score, p2Score := game(p1, p2)
				if p1Score >= p2Score {
					p1Wins++
				} else {
					p2Wins++
				}
				g++
			}

			if(singleStrategyVsMultipleStrategy) {
				fmt.Printf("Holding at %d vs Holding at %d: wins: %d/%d (%0.1f%%), losses: %d/%d (%0.1f%%)\n", i, j,
					p1Wins, games, float32(p1Wins)*100/float32(games),
					p2Wins, games, float32(p2Wins)*100/float32(games))
				p1Wins, p2Wins = 0, 0	
			}
		}
		if(!singleStrategyVsMultipleStrategy) {
			totalGames := totalGames*games
			fmt.Printf("Wins, Losses staying a k = %d: %d/%d (%0.1f%%), losses: %d/%d (%0.1f%%)\n", i,
					p1Wins, totalGames, float32(p1Wins)*100/float32(totalGames),
					p2Wins, totalGames, float32(p2Wins)*100/float32(totalGames))
		}
	}
}

func createHoldingRange(strategy string) HoldingRange {
	strategySplit := strings.Split(strategy, "-")
	holdingRange := HoldingRange{}
	first, err := strconv.Atoi(strategySplit[0])
	if err != nil {
		panic(fmt.Sprintf("Strategy %s is not valid", strategy))
	}
	holdingRange.start = first
	if len(strategySplit) == 2 {
		second, err := strconv.Atoi(strategySplit[1])
		if err != nil {
			panic(fmt.Sprintf("Strategy %s is not valid", strategy))
		}
		holdingRange.end = second
	} else {
		holdingRange.end = first
	}
	return holdingRange
}

func main() {
	args := os.Args[1:]
	
	p1:=Player{name: "p1", holdingRange: createHoldingRange(args[0])}
	p2:=Player{name: "p2", holdingRange: createHoldingRange(args[1])}
	match(p1, p2, 10)
}
