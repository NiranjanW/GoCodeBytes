package Concurrency

import "fmt"


//result := []char, int
//result = map[string]int
var s = "aaabbc"

func main() {
	for i , rune :=range (s){
		fmt.Print("%d  : %c \n" , i ,rune)
	}
}

func noTImes( inp []int) int {

	var result int = 0
	for i:=0 ; i <= len(inp) ; i++ {
		for j:=1 ; j<= len(inp) -i ; j++ {
			if inp[i] == inp[j] {
				result ++
			}
		}
	}
	return result
}
