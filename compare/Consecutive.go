package consecutive

import "fmt"
import . "Dijkstra/NetGen"

/*
Copyright Sérgio Freitas da Silva Júnior (C) - 2019
Universidade de Brasília - Matrícula 190037971

Esse código foi escrito a partir da implementação em paralelo para compará-la uma não-paralela.
*/

func calculateNeighbours(cur *Node) []*Node{
	rett := make([]*Node, 0)
	for _, c := range cur.Conn {
		if !c.To.Mark && (c.To.Dist > (cur.Dist + 1)) {
			c.To.Dist = cur.Dist + 1
			c.To.Mark = true
			rett = append(rett, c.To)
		}
	}
	return rett
};

/*Consecutive ...
Executa a primeira versão do programa (corrigida para evitar um teste desnecessário)
*/
func Consecutive(p bool, nodes []*Node, start int,target *Node) {

	//Bordas da busca atual (número de threads que se iniciarão)
	curNodes := make([]*Node, 1)

	//Definimos o início como distância 0
	nodes[start].Dist = 0
	curNodes[0] = nodes[start] //Definimos o nó incial

	//Nova borda de busca (atualizada a cada thread completa)
	var nNodes []*Node

	//Destino
	var outNode = target

	//Para sempre
	for {
		//Se o fim está na linha de busca, termine
		if NodeInS(outNode, curNodes) {
			break
		}
		nNodes = make([]*Node, 0)

		if len(curNodes)==0{
			goto end
		}

		if p{
			fmt.Print("\n---\n")
			Pnodes(curNodes, outNode)
		}

		for _, edge := range curNodes {
			for _, nedge := range calculateNeighbours(edge) { //Calculamos as distâncias dos vizinhos
				for _, cn := range nNodes {
					if cn.Id == nedge.Id { //Evitamos valores duplicados
						goto nextNode
					}
				}
				nNodes = append(nNodes, nedge)
				nextNode:
			}
		}
		//A nova borda agora é a borda atual
		curNodes = nNodes
	}
	if p{
		fmt.Print("Djikstra Finished.\nCurrent nodes:\n")
		Pnodes(curNodes, outNode)
		BacktrackLog(outNode)
	}
	end:
}