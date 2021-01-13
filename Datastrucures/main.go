package main

import (
"fmt"
"math"
)

//https://github.com/golang/go/blob/b4c8b67adcd39da54f210bef5c201b1df8124d73/src/runtime/map.go#L571
func main() {
	arrayLen := 16

	B := math.Ceil(math.Log2(float64(arrayLen)))

	hash := 4578
	fmt.Printf("Hash:  %b\n", hash)

	mask := 1 << int(B) - 1
	fmt.Printf("Mask:  %b\n", mask)

	index := hash & mask
	fmt.Printf("Index: %b\n", index)
	fmt.Printf("Index (decimal): %d", index)
}
