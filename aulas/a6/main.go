package main

import (
	"fmt"
)

func main() {
	sl0 := []int{1, 2, 3, 4, 5, 6, 7}
	sl1 := make([]int, 2)
	sl2 := sl0[:]
	sl3 := sl0[2:]
	sl4 := sl0[2:5]
	sl5 := sl0[:5]
	sl6 := sl0[:7]
	sl7 := sl0[0:]
	sl8 := sl1[:]

	fmt.Printf("SL0: %v\nSL1: %v\nSL2: %v\nSL3: %v\nSL4: %v\nSL5: %v\nSL6: %v\nSL7: %v\n", sl0, sl1, sl2, sl3, sl4, sl5, sl6, sl7)

	fmt.Println("\nDepois da alteração... [SL7]\n")

	sl7[3] = 2

	fmt.Printf("SL0: %v\nSL1: %v\nSL2: %v\nSL3: %v\nSL4: %v\nSL5: %v\nSL6: %v\nSL7: %v\n", sl0, sl1, sl2, sl3, sl4, sl5, sl6, sl7)

	fmt.Println("\nDepois da alteração... [SL5]\n")

	sl5[3] = 99

	fmt.Printf("SL0: %v\nSL1: %v\nSL2: %v\nSL3: %v\nSL4: %v\nSL5: %v\nSL6: %v\nSL7: %v\n", sl0, sl1, sl2, sl3, sl4, sl5, sl6, sl7)
}
