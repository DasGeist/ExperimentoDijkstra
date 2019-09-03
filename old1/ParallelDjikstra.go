package old1

import "fmt"
import . "Dijkstra/NetGen"

/*
Copyright Sérgio Freitas da Silva Júnior (C) - 2019
Universidade de Brasília - Matrícula 190037971

Escrito no período 26/08/2019 ~ 27/08/2019
(Pequena correção adicional em 3/09/2019)
Para eu mesmo no futuro:
Isso foi antes de começar o conteúdo efetivo de ED
Essa é a versão antiga. Você deve ter feito a nova no final do semestre.
Espero que tenha encontrado erros ridículos e feito algo muito melhor!

^_^

*/

//Aqui acontece a magia.
//Define as distâncias dos vizinhos, se marca e retorna a lista de alterações para a thread principal
func calculateNeighbours(cur *Node, ret chan []*Node) {
	rett := make([]*Node, 0)
	for _, c := range cur.Conn {
		if !c.To.Mark && (c.To.Dist > (cur.Dist + 1)) {
			c.To.Dist = cur.Dist + 1
			rett = append(rett, c.To)
		}
	}
	cur.Mark = true
	ret <- rett
}

//Versão corrigida
func calculateNeighboursC(cur *Node, ret chan []*Node) {
	rett := make([]*Node, 0)
	for _, c := range cur.Conn {
		if !c.To.Mark && (c.To.Dist > (cur.Dist + 1)) {
			c.To.Dist = cur.Dist + 1
			c.To.Mark = true
			rett = append(rett, c.To)
		}
	}
	ret <- rett
}

/*FirstParallel ...
Executa a primeira versão do programa
*/
func FirstParallel(p bool,nodes []*Node,start int,target *Node) {

	//Bordas da busca atual (número de threads que se iniciarão)
	curNodes := make([]*Node, 1)

	//Definimos o início como distância 0
	nodes[start].Dist = 0
	curNodes[0] = nodes[start] //Definimos o nó incial

	//Nova borda de busca (atualizada a cada thread completa)
	var nNodes []*Node

	//Destino
	var outNode = target
	//Número de threads completas (devemos esperar todas para o próximo ciclo)
	var ans int

	//Canal para comunicação entre a principal e as outras threads
	comm := make(chan []*Node)

	//Para sempre
	for {
		//Se o fim está na linha de busca, termine
		if NodeInS(outNode, curNodes) {
			break
		}
		//Ninguém respondeu ainda (eu nem perguntei, ora)
		ans = 0
		nNodes = make([]*Node, 0)

		//Para cada nó da borda, iniciamos uma thread
		for _, edge := range curNodes {
			if !edge.Mark {
				go calculateNeighbours(edge, comm)
			}
		}
		//Imprimimos os nós que "acordamos"
		if p{
			fmt.Print("\n---\n")
			Pnodes(curNodes, outNode)
		}
		//Para cada resposta
		for ans < len(curNodes) {
			ans++
			for _, nedge := range <-comm { //Esperamos a lista de atualizações
				for _, cn := range nNodes {
					if cn.Id == nedge.Id { //Se houver conflito
						if nedge.Dist < cn.Dist {
							cn.Dist = nedge.Dist //Escolhemos a menor distância
							if(p){
								fmt.Print("Conflict solved.")
							}
						}
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
	if(p){
		fmt.Print("Dijkstra Finished.\nCurrent nodes:\n")
		Pnodes(curNodes, outNode)
		BacktrackLog(outNode)
	}
}

/*FirstParallelC ...
Executa a primeira versão do programa (corrigida para evitar um teste desnecessário (assume que as distâncias serão todas unitárias))
*/
func FirstParallelC(p bool, nodes []*Node, start int,target *Node) {
	//Bordas da busca atual (número de threads que se iniciarão)
	curNodes := make([]*Node, 1)

	//Definimos o início como distância 0
	nodes[start].Dist = 0
	curNodes[0] = nodes[start] //Definimos o nó incial

	//Nova borda de busca (atualizada a cada thread completa)
	var nNodes []*Node

	//Destino
	var outNode = target
	//Número de threads completas (devemos esperar todas para o próximo ciclo)
	var ans int

	//Canal para comunicação entre a principal e as outras threads
	comm := make(chan []*Node)

	//Para sempre
	for {
		//Se o fim está na linha de busca, termine
		if NodeInS(outNode, curNodes) {
			break
		}
		//Ninguém respondeu ainda (eu nem perguntei, ora)
		ans = 0
		nNodes = make([]*Node, 0)

		//Para cada nó da borda, iniciamos uma thread
		if len(curNodes)==0{
			goto end
		}
		for _, edge := range curNodes {
			go calculateNeighboursC(edge, comm)
		}
		//Imprimimos os nós que "acordamos"
		if p{
			fmt.Print("\n---\n")
			Pnodes(curNodes, outNode)
		}
		//Para cada resposta
		for ans < len(curNodes) {
			ans++
			for _, nedge := range <-comm { //Esperamos a lista de atualizações
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