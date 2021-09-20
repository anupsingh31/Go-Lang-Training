package main

import "fmt"

func main() {
	data := []string{"a", "c", "e", "a", "e"}
	counts := make(map[string]int)
	for _, item := range data {
		counts[item]++
	}

	letters := []string{"a", "b", "c", "d", "e"}
	for _, letter := range letters {
		count, ok := counts[letter]
		if !ok {
			fmt.Println(letter, "no founds")
		} else {
			fmt.Printf("%s: %d\n", letter, count)
		}
	}
}
