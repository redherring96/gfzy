package main

import ("strings"
		"fmt")

//package to calculate scores and rank them
//What metrics to use to calcualte scores?
//1.ratio to matched characters to unmatched characters
//2.contiguous matches (ie 'ae' match in string "docs/sae"
//3.match after a punctuation mark (ie 's' after "docs/sae"
//4.Camel Case (ie 'sa' in "docs/Sae"
//5.matches in close proximity to the start


//findContig() accepts two strings and finds contiguous matches from the
//input string in the focus string and will return a score used in the
//ranking calculations
//Accumulator() is the only outward facing function of the package. Calling Accumulator() will accumulate the score of the focus argument for the given input argument and will return the score
func Accumulator(focus,input string) int {
	score := findMatch(focus,input) + findContig(focus,input) + findSepar(focus,input)
	return score
}
//findMatch() will find all matches of the given input argument to the
//the given focus argument and will return a score consisting of the sum of unmatched characters (-1) to matched characters (+1)
func findMatch(focus, input string) int {
	score := 0
	score -= len(focus)//decrement unmatched characters
	for i:=0; i < len(input); i++ {//iterate through input
		if strings.Contains(strings.ToLower(focus),strings.ToLower(string(input[i]))) {//test each element of input
			score +=1 
		}
	}
	return score
}

func findContig(focus string, input string) int {
	score := 0
	for i := 0; i < (len(input)-1); i++ {
		j := i + 2
		if strings.Contains(strings.ToLower(focus),strings.ToLower(input[i:j])) {
			//TODO: add camelCase,
			score += findCamel(focus,input[i:j])
			score += 5 //TODO: add global contigBonus const

		}
	}
	return score
}

func findCamel(focus, sub string) int {
	score := 0
	camel := sub[0:1] 
	cc := strings.Replace(sub,camel,strings.ToUpper(camel),1)
	if strings.Contains(focus,cc){
		score += 10
	}
	return score
}

//findSepar() accepts a focus and an input string and will return a score
//based on any input characters after a separator character('+',' ')

func findSepar(focus string, input string) int {

	space := make([]string,0)
	space = append(space," ")
	slash := make([]string,0)
	slash = append(slash,"/")
	score := 0


	if strings.ContainsAny(focus," ") {

		temp := make([]string,0)

		for i:=0; i < len(input); i++ {
			temp = append(temp,(" " + string(input[i])))
		}
		for _, inpt := range temp {
			score = score + findContig(focus, inpt)
		}
	}

	if strings.ContainsAny(focus,"/") {

		temp := make([]string,0)

		for i:=0; i < len(input); i++ {
			temp = append(temp,("/" + string(input[i])))
		}

		for _, inpt := range temp {
			fmt.Println(inpt)
			score = score + findContig(focus, inpt)
		}
	}

	return score
}


func main() {
	fmt.Println(findSepar("a Bc", "abc"))
	fmt.Println(findContig("a Bc","abc"))
	fmt.Println(findMatch("a Bc","abc"))
	fmt.Println(findCamel("a Bc","abc"))
	fmt.Println(Accumulator("a Bc", "abc"))
}
