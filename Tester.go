package main

import old1 "Dijkstra/old1"
import consecutive "Dijkstra/compare"
import . "Dijkstra/NetGen"
import "fmt"

func main() {
	//Graph A (exemplo)
	/*
			{Origem} - [destino]
			{0}      8 -[5]
			 | \   / |   |
		     |   6 - 7 - 4
			 1 - 2 - 3 -/
	*/
	ResetCounter()
	graphA := make([]*Node, 9)
	for i := 0; i < 9; i++ {
		graphA[i] = NNode()
	}

	Connect(graphA[0], graphA[1])
	Connect(graphA[0], graphA[6])

	Connect(graphA[1], graphA[2])

	Connect(graphA[2], graphA[3])

	Connect(graphA[3], graphA[4])

	Connect(graphA[4], graphA[5])
	Connect(graphA[4], graphA[7])

	Connect(graphA[5], graphA[8])

	Connect(graphA[6], graphA[7])
	Connect(graphA[6], graphA[8])

	Connect(graphA[7], graphA[8])
	//End of Graph A
	fmt.Print("Generating Graph B\n")
	ResetCounter()
	graphB:=GenerateGraph(2000,1,1)

	fmt.Print("Testing Graph A\n")
	old1.FirstParallel(true,graphA,0,graphA[5])
	ResetGraph(graphA)
	old1.FirstParallelC(true,graphA,0,graphA[5])
	ResetGraph(graphA)
	consecutive.Consecutive(true,graphA,0,graphA[5])

	fmt.Print("Testing Graph B\n")
	old1.FirstParallel(true,graphB,0,graphB[70])
	ResetGraph(graphB)
	old1.FirstParallelC(true,graphB,0,graphB[70])
	ResetGraph(graphB)
	consecutive.Consecutive(true,graphB,0,graphB[70])
}
