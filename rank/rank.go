package rank

import ("gfzy/score"
		"sort")

type node struct {
	File string
	Score int
}

func Rank(raw []string, input string) []string{
	nodes := initNodes(raw,input)
	temp := make([]string,0)
	//sort nodes
//	nodes = sortNodes(nodes)

	 sort.SliceStable(nodes,func(i,j int) bool { return nodes[i].Score > nodes[j].Score})
	//after sort
	for _, v:= range nodes{
		temp = append(temp,v.File)
	}
	//return nodes
	return temp

}

//sort
//rebalance

/*func sortNodes(nodes []node) []node {
//	return sort.SliceStable(nodes,func(i,j int) bool { return nodes[i].Score < nodes[j].Score})
	return nodes
}*/

func initNodes(raw []string, input string) []node {
	nodes := make([]node,0)
	var temp node
	for _, v:= range raw {
		temp.File = v
		temp.Score = score.Accumulator(v, input)
		nodes = append(nodes,temp)
	}
	return nodes
}

/*func main(){
	test := make([]string,5)
	test[0] = "testing"
	test[1] = "tested"
	test[2] = "rested"
	test[3] = "resting"
	test[4] = "bested"
	input := "rsg"
	fmt.Println(test)
	fmt.Println("after sorting based on score")
	fmt.Println(Rank(test,input))

}
*/
