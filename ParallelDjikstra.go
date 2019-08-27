package main

import "fmt"

/*
Código por Sérgio Freitas da Silva Júnior
Universidade de Brasília - Matrícula 190037971

Escrito no período 26/08/2019 ~ 27/08/2019
Para eu mesmo no futuro:
Isso foi antes de começar o conteúdo efetivo de ED
Essa é a versão antiga. Você deve ter feito a nova no final do semestre.
Espero que tenha encontrado erros ridículos e feito algo muito melhor!

^_^

*/

//Uma [con]exão com [dist]ância e destinatário [to]
type con struct {
	dist int
	to   *node
}

//Construtor de conexão
func conC(dist int, to *node) *con {
	var tcon = new(con)
	tcon.dist = dist
	tcon.to = to
	return tcon
}

//Cria duas conexões (ida e volta) entre nós
func connect(a *node, b *node) {
	a.conn = append(a.conn, *conC(1, b))
	b.conn = append(b.conn, *conC(1, a))
}

//Nó. Tem um identificador para resolução de conflitos entre threads.
type node struct {
	id   int
	dist int
	mark bool
	conn []con
}

//Contador dos identificadores globais
var idc int

//Construtor de Nós
func nNode() *node {
	var tempNode = new(node)
	tempNode.id = idc
	tempNode.dist = int((^uint(0)) >> 1)
	tempNode.mark = false
	tempNode.conn = make([]con, 0)
	idc++
	return tempNode
}

//Aqui acontece a magia.
//Define as distâncias dos vizinhos, se marca e retorna a lista de alterações para a thread principal
func calculateNeighbours(cur *node, ret chan []*node) {
	rett := make([]*node, 0)
	for _, c := range cur.conn {
		if !c.to.mark && (c.to.dist > (cur.dist + 1)) {
			c.to.dist = cur.dist + 1
			rett = append(rett, c.to)
		}
	}
	cur.mark = true
	ret <- rett
}

//Checa se um nó está presente numa slice
func nodeInS(a *node, many []*node) bool {
	for _, b := range many {
		if b.id == a.id {
			return true
		}
	}
	return false
}

//Imprime os vizinhos do nó atual
func pnodes(curNodes []*node, outNode *node) {
	for _, node := range curNodes {
		fmt.Printf("Node %d, distance %d", node.id, node.dist)
		if node.id == outNode.id {
			fmt.Print(" [T]")
		}
		fmt.Print("\n")
	}
}

//Mostra de trás para frente o caminho com a menor distância até o 0
//(assume que existe um caminho. Se não houver, comportamento indefinido (ponteiro vazio))
func backtrackLog(a *node) {
	fmt.Print("\n/-\\-/-\\-/-\\-/-\\\nShortest Path:\n")
	ccon := new(node)
	for a.dist != 0 {
		ccon.dist = int(^uint(0) >> 1)
		fmt.Printf("Node %d<-", a.id)
		for _, nei := range a.conn {
			if nei.to.dist < ccon.dist {
				ccon = nei.to
			}
		}
		a = ccon
	}
	fmt.Printf("Node %d", a.id)
	fmt.Print("\n\\-/-\\-/-\\-/-\\-/\n")
}

func main() {
	//Graph A (exemplo)
	/*
			{Origem} - [destino]
			{0}      8 -[5]
			 | \   / |   |
		     |   6 - 7 - 4
			 1 - 2 - 3 -/
	*/
	nodes := make([]*node, 9)
	for i := 0; i < 9; i++ {
		nodes[i] = nNode()
	}

	connect(nodes[0], nodes[1])
	connect(nodes[0], nodes[6])

	connect(nodes[1], nodes[2])

	connect(nodes[2], nodes[3])

	connect(nodes[3], nodes[4])

	connect(nodes[4], nodes[5])
	connect(nodes[4], nodes[7])

	connect(nodes[5], nodes[8])

	connect(nodes[6], nodes[7])
	connect(nodes[6], nodes[8])

	connect(nodes[7], nodes[8])
	//End of Graph A

	//Bordas da busca atual (número de threads que se iniciarão)
	curNodes := make([]*node, 1)

	//Definimos o início como distância 0
	nodes[0].dist = 0
	curNodes[0] = nodes[0] //Definimos o nó incial

	//Nova borda de busca (atualizada a cada thread completa)
	var nNodes []*node

	//Destino
	var outNode = nodes[5]
	//Número de threads completas (devemos esperar todas para o próximo ciclo)
	var ans int

	//Canal para comunicação entre a principal e as outras threads
	comm := make(chan []*node)

	//Para sempre
	for {
		//Se o fim está na linha de busca, termine
		if nodeInS(outNode, curNodes) {
			break
		}
		//Ninguém respondeu ainda (eu nem perguntei, ora)
		ans = 0
		nNodes = make([]*node, 0)

		//Para cada nó da borda, iniciamos uma thread
		for _, edge := range curNodes {
			if !edge.mark {
				go calculateNeighbours(edge, comm)
			}
		}
		//Imprimimos os nós que "acordamos"
		fmt.Print("\n---\n")
		pnodes(curNodes, outNode)
		//Para cada resposta
		for ans < len(curNodes) {
			ans++
			for _, nedge := range <-comm { //Esperamos a lista de atualizações
				for _, cn := range nNodes {
					if cn.id == nedge.id { //Se houver conflito
						if nedge.dist < cn.dist {
							cn.dist = nedge.dist //Escolhemos a menor distância
							fmt.Print("Conflict solved.")
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
	fmt.Print("Djikstra Finished.\nCurrent nodes:\n")
	pnodes(curNodes, outNode)
	backtrackLog(outNode)
}
