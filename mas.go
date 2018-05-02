package main

import (
	"fmt"
)

// import "fmt"

func main () {
	// var scores [10] int
	// scores[0] = 339
	// fmt.Printf("%v",scores)
	// scores := [4]int{9001, 9333, 212, 33}
	// for index, value := range scores {
	// 	fmt.Println(index, value)
	// }

	// scores := make([]int, 10)
	scores := make([]int, 0, 10)
	scores = scores[0: 8]
	scores[7] = 9033
	// scores = append(scores, 5)
	fmt.Println(scores)

}