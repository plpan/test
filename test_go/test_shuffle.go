package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func shuffle(slice []int) {
	for i := 0; i < len(slice); i++ {
		a := rand.Intn(len(slice))
		b := rand.Intn(len(slice))
		slice[a], slice[b] = slice[b], slice[a]
	}
}

func lean() {
	fmt.Println("--not balance--")
	var cntMap = map[string]int{}
	for i := 0; i < 1000000; i++ {
		var a = []int{0, 1, 2, 3}
		shuffle(a)
		for j := 0; j < len(a); j++ {
			key := fmt.Sprintf("%d_%d", a[j], j)
			cntMap[key]++
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			key := fmt.Sprintf("%d_%d", i, j)
			fmt.Printf("%d ", cntMap[key])
		}
		fmt.Println()
	}
}

func normal() {
	fmt.Println("--balance--")
	var cntMap = map[string]int{}
	var a = []int{0, 1, 2, 3}
	for i := 0; i < 1000000; i++ {
		shuffle(a)
		for j := 0; j < len(a); j++ {
			key := fmt.Sprintf("%d_%d", a[j], j)
			cntMap[key]++
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			key := fmt.Sprintf("%d_%d", i, j)
			fmt.Printf("%d ", cntMap[key])
		}
		fmt.Println()
	}
}

func main() {
	// shuffle有倾斜，因为数据都倾向于留在原来的位置上
	lean()

	// shuffle无倾斜，因为多次操作真正随机了
	normal()
}
