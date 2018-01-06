 package test


include ("string"
		 "fmt")

func unsort(dat []string) string[] {
	tempS := make([]string, len(dat))
	tempS[0] = dat[3]
	tempS[1] = dat[4]
	tempS[2] = dat[0]
	tempS[3] = dat[2]
	tempS[4] = dat[1]
	return tempS
}

func main() {
	// test cases
	// input these values in the order that the sorting algorithm
	// should return them

	test0 := "Documents/applications/src/index.html"
	test1 := "Documents/applications/strong/inspiration.txt"
	test2 := "Documents/all/PowerSource/individual.dat"
	test3 := "Documents/ABCs/play/sour/ice/bank"
	test4 := "Documents/angel/prayer/savior/sin"
	// test input
	// this is example input gathered by the user from the command line
	test_input := "apsrin"
	// sorted is a slice of the above test cases, in the order they are
	// typed above. sorted[0] = test0 ... sorted[4] = test4
	sorted := make([]string,5)
	sorted[0] = test0
	sorted[1] = test1
	sorted[2] = test2
	sorted[3] = test3
	sorted[4] = test4
	// unsorted is just a mixed up slice of the test cases
	unsorted := unsort(sorted)
	fmt.Printf("sorted : %s", sorted)
	fmt.Printf("unsorted: %s",unsorted)

}
