package main

import "fmt"

func main() {
	people := map[string][]string{
		`bond_james`:      []string{`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: []string{`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           []string{`Being evil`, `Ice cream`, `Sunsets`},
	}

	for k, v := range people {
		fmt.Println(`This is the record for`, k)

		for _, v2 := range v {
			fmt.Println(v2)
		}
	}
}
