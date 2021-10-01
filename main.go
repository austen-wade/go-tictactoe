package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

const (
	min    int = 1
	size   int = 3
	spaces int = size * size
)

var (
	boardSlice []rune         = []rune{}
	boardMap   map[rune][]int = make(map[rune][]int)
	gameLoop   bool           = true
	winSets    [][]int        = [][]int{}
)

func main() {
	for i := spaces; i > 0; i-- {
		boardSlice = append(boardSlice, ' ')
	}
	rand.Seed(time.Now().UnixNano())
	generateWinSets()
	start()
	display()
}

func start() {
	if playerGoesFirst() {
		time.Sleep(time.Second)
		for gameLoop {
			if playerGo() {
				break
			}
			if computerGo() {
				break
			}
		}
	} else {
		time.Sleep(time.Second)
		for gameLoop {
			if computerGo() {
				break
			}
			if playerGo() {
				break
			}
		}
	}
}

func playerGoesFirst() bool {
	fmt.Println("Let's flip a coin to see who goes first.")
	time.Sleep(time.Second)
	for _, i := range []int{1, 2, 3} {
		fmt.Printf("%d...\n", i)
		time.Sleep(time.Second / 2)
	}

	if 1 == rand.Intn(2) {
		fmt.Println("Heads! Player goes first.")
		return true
	}

	fmt.Println("Tails! Computer goes first.")
	return false
}

func playerGo() bool {
	display()
	sameMove := true
	for sameMove {
		var userMove string
		fmt.Println("Select which square you would like to set [1-9]")
		fmt.Scanln(&userMove)

		if input, err := strconv.Atoi(userMove); err == nil && input >= min && input <= spaces {
			input -= 1
			if boardSlice[input] == ' ' {
				boardSlice[input] = 'x'
				boardMap['x'] = append(boardMap['x'], input)
				sameMove = false

				if isWin('x') {
					fmt.Println("Player wins!")
					return true
				} else if isFull() {
					fmt.Println("Board is full!")
					return true
				}
			} else {
				fmt.Println("That square is already taken")
			}
		} else {
			fmt.Println("Please input a valid number from [1-9]")
		}
	}
	return false
}

func computerGo() bool {
	display()
	fmt.Println("The computer is thinking ...")
	time.Sleep(time.Second * 1)
	emptyIdxes := []int{}
	for i, val := range boardSlice {
		if val == ' ' {
			emptyIdxes = append(emptyIdxes, i)
		}
	}

	boardPos := emptyIdxes[rand.Intn(len(emptyIdxes))]
	boardSlice[boardPos] = 'o'
	boardMap['o'] = append(boardMap['o'], boardPos)

	if isWin('o') {
		fmt.Println("Computer wins!")
		return true
	} else if isFull() {
		fmt.Println("Board is full!")
		return true
	}
	return false
}

func generateWinSets() {
	// diagonal
	winSets = append(winSets, createWinCond(0, spaces-1, size+1))

	// anti diagonal
	winSets = append(winSets, createWinCond(size-1, (size*2)+1, size-1))

	for i := 0; i <= size-1; i++ {
		// v
		winSets = append(winSets, createWinCond(i, spaces-1, size))

		// h
		winSets = append(winSets, createWinCond(i*size, (i*size)+size, 1))
	}
}

func isWin(piece rune) bool {
	for _, set := range winSets {
		if len(sliceIntersection(boardMap[piece], set)) == size {
			return true
		}
	}

	return false
}

func createWinCond(start, end, step int) (cond []int) {
	for i, j := start, 0; j < size; i, j = i+step, j+1 {
		cond = append(cond, i)
	}
	return
}

func isFull() bool {
	for _, i := range boardSlice {
		if i == ' ' {
			return false
		}
	}
	return true
}

func display() {
	fmt.Println()

	count := 0
	for _, i := range boardSlice {
		if count == size || count == size*2 {
			fmt.Println()
		}
		fmt.Printf("[%c]", i)
		count++
	}
	fmt.Printf("\n\n")
}

func sliceIntersection(s1, s2 []int) (inter []int) {
	hash := make(map[int]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] {
			inter = append(inter, e)
		}
	}
	return
}
