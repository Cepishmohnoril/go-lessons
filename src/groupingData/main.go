package main

import "fmt"

func main() {
	//Array
	var arr [5]int
	arr[3] = 42
	fmt.Println(arr)
	fmt.Println(arr[3])

	//Slice
	slice := []int{2, 4, 8, 16, 32, 64}
	fmt.Println(slice)

	//Iterate
	for i, v := range slice {
		fmt.Println(i, v)
	}

	fmt.Println(slice[2:5])
	fmt.Println(slice[2:])

	//Add to
	slice2 := []int{128, 256, 512}

	slice = append(slice, slice2...)
	slice = append(slice, 123, 456)

	fmt.Println(slice)

	//Delete
	sliceRm := []int{1, 2, 3, 456, 4, 5}

	fmt.Println(append(sliceRm[:3], sliceRm[4:]...))

	//Make

	sliceMk := make([]int, 2, 3)

	sliceMk = append(sliceMk, 3, 4)

	fmt.Println(sliceMk)
	fmt.Println(len(sliceMk))
	fmt.Println(cap(sliceMk))

	//Multy dimensional slice

	sl1 := []string{"Some", "String", "Here"}
	sl2 := []string{"Some", "Other", "String"}

	mdsl := [][]string{sl1, sl2}

	fmt.Println(mdsl)

	//Maps

	testMap := map[string]int{
		"Foo": 1,
		"Bar": 2,
	}

	fmt.Println(testMap["Foo"])

	if v, ok := testMap["Doot"]; !ok {
		fmt.Println("Zero value in map", v)
	}

	testMap["Doot"] = 3

	fmt.Println(testMap["Doot"])

	for key, value := range testMap {
		fmt.Println(key, "is key for ", value)
	}

	delete(testMap, "Doot")
	fmt.Println(testMap)
}
