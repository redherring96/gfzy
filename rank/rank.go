package rank

import ("fmt")

type node struct {
	File string
	Score int
}

func Rank(raw []string, input string) []string{
	nodes =: initNodes(raw,input) 
	//sort nodes
	//return nodes

}

//sort
//rebalance

func initNodes(raw []string, input string) []nodes {
	nodes := make([]node,0)
	var temp node
	for _, v:= range raw {
		temp.File = v
		temp.Score = score.Acculmulator(v, input)
		nodes = append(nodes,temp)
	}
	return nodes
}
