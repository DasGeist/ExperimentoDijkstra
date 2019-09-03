package netgen

import "time"
import "fmt"
import "math/rand"

//Nó. Tem um identificador para resolução de conflitos entre threads.
type Node struct {
	Id   int
	Dist int
	Mark bool
	Conn []Con
}

//Uma [con]exão com [dist]ância e destinatário [to]
type Con struct {
	Dist int
	To   *Node
}

//Construtor de conexão
func ConC(dist int, to *Node) *Con {
	var tcon = new(Con)
	tcon.Dist = dist
	tcon.To = to
	return tcon
}

//Cria duas conexões (ida e volta) entre nós
func Connect(a *Node, b *Node) {
	a.Conn = append(a.Conn, *ConC(1, b))
	b.Conn = append(b.Conn, *ConC(1, a))
}

//Contador dos identificadores globais
var idc int

//Construtor de Nós
func NNode() *Node {
	var tempNode = new(Node)
	tempNode.Id = idc
	tempNode.Dist = int((^uint(0)) >> 1)
	tempNode.Mark = false
	tempNode.Conn = make([]Con, 0)
	idc++
	return tempNode
}

//Checa se um nó está presente numa slice
func NodeInS(a *Node, many []*Node) bool {
	for _, b := range many {
		if b.Id == a.Id {
			return true
		}
	}
	return false
}

//Imprime os vizinhos do nó atual
func Pnodes(curNodes []*Node, outNode *Node) {
	for _, node := range curNodes {
		fmt.Printf("Node %d, distance %d", node.Id, node.Dist)
		if node.Id == outNode.Id {
			fmt.Print(" [T]")
		}
		fmt.Print("\n")
	}
}

func ResetCounter(){
	idc=0
}

func ResetGraph(graph []*Node){
	for _,node:=range graph{
		node.Dist = int((^uint(0)) >> 1)
		node.Mark = false
	}
}

//Mostra de trás para frente o caminho com a menor distância até o 0
//(assume que existe um caminho. Se não houver, comportamento indefinido (ponteiro vazio))
func BacktrackLog(a *Node) {
	fmt.Print("\n/-\\-/-\\-/-\\-/-\\\nShortest Path:\n")
	ccon := new(Node)
	for a.Dist != 0 {
		ccon.Dist = int(^uint(0) >> 1)
		fmt.Printf("Node %d<-", a.Id)
		for _, nei := range a.Conn {
			if nei.To.Dist < ccon.Dist {
				ccon = nei.To
			}
		}
		a = ccon
	}
	fmt.Printf("Node %d", a.Id)
	fmt.Print("\n\\-/-\\-/-\\-/-\\-/\n")
}

func Connected(who *Node, toWhom *Node) bool{
	for _,con:=range who.Conn{
		if(con.To.Id==toWhom.Id){
			return true
		}
	}
	return false
}

func GenerateGraph(count int, baseConnections int,extraConnections int) []*Node{
	graph:=make([]*Node,count)
	for i := 0; i < count; i++ {
		graph[i] = NNode()
	}
	ccount:=0
	curnodes:=make([]*Node,1)
	var nnodes []*Node
	curnodes[0]=graph[0]
	//Criamos um grafo "hierárquico" (cada camada só se conecta com uma outra camada)
	//Assim garantimos que sempre haverá um caminho entre quaisquer dois nós
	for{
		nnodes=make([]*Node,0)
		if(ccount==count-1){
			break
		}
		for _,node := range curnodes{
			for i:=0;i<baseConnections;i++{
				if(ccount==count-1){
					goto extra
				}
				ccount++
				Connect(node,graph[ccount])
				nnodes=append(nnodes,graph[ccount])
			}
		}
		curnodes=nnodes
	}
	//Quebramos a hierarquia fazendo [extraConnections] conexões aleatórias em cada nó
	extra:
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	var addr int
	for _,node := range graph{
		for i:=0;i<extraConnections;i++{
			for{
				addr=r1.Intn(count)
				if(graph[addr].Id!=node.Id && !Connected(node,graph[addr])){
					Connect(node,graph[addr])
					break
				}
			}
		}
	}
	return graph
}