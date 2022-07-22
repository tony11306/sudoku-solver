package main

import (
	"fmt"
)

type Graph [][]byte
const SODOKU_SIZE int = 3 // which means its n*n blocks, which is n^2*n^2 array size
var symbols []byte
const EMPTY_SYMBOL byte = 0

// make an empty graph
func MakeGraph() Graph {
	g := make([][]byte, SODOKU_SIZE*SODOKU_SIZE)
	
	for i := 0; i < len(g); i++ {
		g[i] = make([]byte, SODOKU_SIZE*SODOKU_SIZE)
	}

	return g
}

// check is number valid to put in g[gi][gj]
// and assuming that g is valid before passing in
// and 0 <= gi, gj < SODOKU_SIZE * SODOKU_SIZE
func IsDoable(g Graph, gi int, gj int, num byte) bool {
	if(g[gi][gj] != EMPTY_SYMBOL) {
		return false
	}

	block_i := gi / SODOKU_SIZE
	block_j := gj / SODOKU_SIZE

	// checking is valid in that block
	for i := 0; i < SODOKU_SIZE; i++ {
		for j := 0; j < SODOKU_SIZE; j++ {
			if block_i*SODOKU_SIZE + i == gi && block_j*SODOKU_SIZE + j == gj {
				continue
			}
			if g[block_i*SODOKU_SIZE + i][block_j*SODOKU_SIZE + j] == num {
				return false
			}
		}
	}

	// checking is valid in that row
	for j := 0; j < SODOKU_SIZE * SODOKU_SIZE; j++ {
		if j == gj {
			continue
		}
		if g[gi][j] == num {
			return false
		}
	}

	// checking is valid in that col
	for i := 0; i < SODOKU_SIZE * SODOKU_SIZE; i++ {
		if i == gi {
			continue
		}
		if g[i][gj] == num {
			return false
		}
	}

	return true
}

func Print(g Graph) {

	printLine := func() {
		fmt.Print("+")
		for i := 0; i < SODOKU_SIZE; i++ {
			for j := 0; j < SODOKU_SIZE * 2 + 1; j++ {
				fmt.Print("-")
			}
			fmt.Print("+")
		}
		fmt.Println()
	}

	printLine()
	for i := 0; i < SODOKU_SIZE * SODOKU_SIZE; i++ {
		for j := 0; j < SODOKU_SIZE * SODOKU_SIZE; j++ {
			if j % SODOKU_SIZE == 0 {
				fmt.Print("| ")
			}
			fmt.Print(g[i][j], " ")
		}
		fmt.Println("|")
		if (i + 1) % SODOKU_SIZE == 0 {
			printLine()
		}
	}
}

func GetSolution(g Graph) (Graph, bool) {
	var (
		try_i int = SODOKU_SIZE * SODOKU_SIZE
		try_j int = SODOKU_SIZE * SODOKU_SIZE
	)
	for i := 0; i < SODOKU_SIZE * SODOKU_SIZE; i++ {
		flag := false
		for j := 0; j < SODOKU_SIZE * SODOKU_SIZE; j++ {
			if g[i][j] == EMPTY_SYMBOL {
				flag = true
				try_i = i
				try_j = j
				break
			}
		}
		if flag {
			break
		}
	}

	if try_i == SODOKU_SIZE * SODOKU_SIZE && try_j == SODOKU_SIZE * SODOKU_SIZE {
		return g, true
	}

	for _, symbol := range symbols {
		if IsDoable(g, try_i, try_j, symbol) {
			g[try_i][try_j] = symbol
			sol, isOk := GetSolution(g)
			if isOk {
				return sol, true
			}
			g[try_i][try_j] = EMPTY_SYMBOL
		}
	}

	return g, false
}

func InputGraph(g *Graph) {
	for i := 0; i < SODOKU_SIZE * SODOKU_SIZE; i++ {
		for j := 0; j < SODOKU_SIZE * SODOKU_SIZE; j++ {
			fmt.Scan(&(*g)[i][j])
			// checking symbol is in the symbol list
			if (*g)[i][j] == EMPTY_SYMBOL {
				continue
			}
			flag := false
			for _, k := range symbols {
				if k == (*g)[i][j] {
					flag = true
					break
				}
			}
			
			if !flag {
				panic("The graph contains invalid symbol")
			}
		}
	}
}

func InitSymbols() {
	symbols = make([]byte, SODOKU_SIZE * SODOKU_SIZE)
	
	for i := 0; i < SODOKU_SIZE * SODOKU_SIZE; i++ {
		fmt.Scan(&symbols[i])
	}
}

func main() {
	graph := MakeGraph()

	fmt.Println("Please enter the symbols represented", 1, "to", SODOKU_SIZE * SODOKU_SIZE)
	InitSymbols()

	fmt.Println("Now enter the sudoku graph:")
	fmt.Println("(The empty symbol is", EMPTY_SYMBOL, ")")
	InputGraph(&graph)
	sol, _ := GetSolution(graph)


	fmt.Println("The answer:")
	Print(sol)

	var end string
	fmt.Scanln(&end)
	fmt.Scanln(&end)
}