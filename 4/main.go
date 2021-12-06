package main

import "fmt"

type result struct {
	List []string
}

var results []result

func main() {
	strings := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	bools := []bool{true, true, true, true, true, true, true}
	resultCounter := 0
	for x := range strings {
		if bools[x] {
			results = append(results, result{
				List: []string{
					strings[x],
				},
			})
			for y := x + 1; y < len(strings); y++ {
				if compareWords(strings[x], strings[y]) {
					results[resultCounter].List = append(results[resultCounter].List, strings[y])
					bools[y] = false
				}
			}
			resultCounter++
		}

	}
	for _, x := range results {
		fmt.Println(x)
	}

}

func compareWords(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}
	counter := 0
	for _, x := range b {
		for _, y := range a {
			if x == y {
				counter++
			}
		}
	}
	return counter == len(a)
}
