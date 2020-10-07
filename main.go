package main

import (
	"fmt"
	"time"
)

const HolyBible = `In the beginning God created the heaven and the earth.`

// fastest random function
var randomFunction = RandomVFastrand

func main() {
	// Init counters
	bibleIndex := 0
	longerIndex := 0
	var iteration int64 = 0

	// Init logger
	startedAt := time.Now()
	go func() {
		for range time.NewTicker(time.Second).C {
			fmt.Printf("Tick: %d randomized character in %s\n", iteration, time.Since(startedAt))
		}
	}()

	// Write the bible with super fast monkey !!
	for {
		if HolyBible[bibleIndex] == RandomizedCharacter(randomFunction) {
			bibleIndex++
		} else if bibleIndex != 0 {
			if bibleIndex > longerIndex {
				fmt.Printf("Randomized Bible: %s\n", HolyBible[:longerIndex])
				longerIndex = bibleIndex
			}
			bibleIndex = 0
		}

		iteration++
	}
}
