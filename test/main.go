package main

import (
	"fmt"
	"strings"

	littlealchemy "github.com/alancnet/go-little-alchemy"
)

// Discovers all elements obtainable in Little Alchemy
func main() {
	elements := littlealchemy.GetBaseElements()
	attempts := make(map[[2]string]bool)
	discoveries := make(map[string]bool)

	fmt.Println("Base elements are " + strings.Join(elements, ", ") + ".")

	var count = len(elements)
	for count > 0 {
		count = 0
		for i, element1 := range elements {
			for _, element2 := range elements[i:] {
				recipe := [2]string{element1, element2}
				if !attempts[recipe] {
					attempts[recipe] = true
					results := littlealchemy.Combine(element1, element2)
					for _, discovery := range results {
						fmt.Println("Combining " + element1 + " and " + element2 + " yields " + discovery)
						if !discoveries[discovery] {
							discoveries[discovery] = true
							elements = append(elements, discovery)
							count++
						}
					}
				}
			}
		}
	}

}
