package main

import (
	"fmt"
	"time"
)

const HolyBible = `In the beginning God created the heaven and the earth.`

var randomFunction = RandomVFastrand

func main() {
	index := 0
	longuer := 0
	var iteration int64 = 0
	startedAt := time.Now()

	go func() {
		for range time.NewTicker(time.Second).C {
			fmt.Printf("Tick: %d randomized character in %s\n", iteration, time.Since(startedAt))
		}
	}()

	for {
		if HolyBible[index] == RandomizedCharacter(randomFunction) {
			index++
		} else if index != 0 {
			if index > longuer {
				fmt.Printf("Randomized Bible: %s\n", HolyBible[:index])
				longuer = index
			}
			index = 0
		}

		iteration++
	}
}
