package main_test

import "testing"
import ."Dijkstra/NetGen"
import old1 "Dijkstra/old1"
import consecutive "Dijkstra/compare"

var graphB []*Node

func init(){
	graphB = GenerateGraph(2000,1,1)
}

func BenchmarkFirstParallel(b *testing.B){
	for i:=0;i<b.N;i++{
		ResetGraph(graphB)
		old1.FirstParallel(false,graphB,0,graphB[70])
	}
}

func BenchmarkFirstParallelC(b *testing.B){
	for i:=0;i<b.N;i++{
		ResetGraph(graphB)
		old1.FirstParallelC(false,graphB,0,graphB[70])
	}
}

func BenchmarkConsecutive(b *testing.B){
	for i:=0;i<b.N;i++{
		ResetGraph(graphB)
		consecutive.Consecutive(false,graphB,0,graphB[70])
	}
}