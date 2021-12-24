package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputs, boards := ReadInput("test.txt")
	solution := Play(boards, inputs)
	fmt.Println(solution)
}

type boardNum struct {
	number int
	drawn  bool
}

type board struct {
	numbers [5][5]boardNum
}

func Play(boards []board, inputs []int) int {
	for _, value := range inputs {
		for i := 0; i < len(boards); i++ {
			boards[i] = boards[i].PlayBoard(value)
			win := boards[i].IsWinning()
			if win {
				count := boards[i].Count()
				remainings := append(boards[:i], boards[i+1:]...)
				fmt.Println("remainings boards to check", len(remainings))

				if len(boards) == 1 {
					return count * value
				} else {
					return Play(remainings, inputs)
				}
			}
		}
	}
	return -1
}

func (b board) Count() int {
	var total int
	for _, col := range b.numbers {
		for _, val := range col {
			if !val.drawn {
				total += val.number
			}
		}
	}
	return total
}

func (b board) IsWinning() bool {
	for x := 0; x < 5; x++ {
		count := 0
		for y := 0; y < 5; y++ {
			if b.numbers[x][y].drawn {
				count = count + 1
			}
		}
		if count == 5 {
			fmt.Println("Win 1")
			return true
		}
	}

	for y := 0; y < 5; y++ {
		count := 0
		for x := 0; x < 5; x++ {
			if b.numbers[x][y].drawn {
				count = count + 1
			}
		}
		if count == 5 {
			fmt.Println("Win 2")
			return true
		}
	}

	return false
}

func (b board) PlayBoard(num int) board {
	for x := 0; x < 5; x++ {
		for y := 0; y < 5; y++ {
			b.numbers[x][y] = b.numbers[x][y].Draw(num)
		}
	}
	return b
}

func (bn boardNum) Draw(num int) boardNum {
	if num == bn.number {
		bn.drawn = true
		return bn
	}
	return bn
}

func ReadInput(path string) ([]int, []board) {
	file, err := os.Open(path)
	var out []string
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		out = append(out, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	var drawns []int

	for _, value := range strings.Split(out[0], ",") {
		i, _ := strconv.Atoi(value)
		drawns = append(drawns, i)
	}

	boardStrings := out[1:]
	boards := []board{}
	for i := 0; i < (len(boardStrings)); i += 6 {
		fmt.Println("new board", i, len(boardStrings))
		blines := [5][5]boardNum{}
		lines := boardStrings[i+1 : i+6]
		for col, val := range lines {
			sanitized := strings.ReplaceAll(strings.Trim(val, " "), "  ", " ")
			nums := strings.Split(sanitized, " ")
			for row, nval := range nums {
				number, _ := strconv.Atoi(nval)
				blines[col][row] = boardNum{
					number: number,
					drawn:  false,
				}
			}
		}
		boards = append(boards, board{
			numbers: blines,
		})
	}

	return drawns, boards
}
