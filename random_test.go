package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkRandomGoSDK(b *testing.B)     { benchmarkRandom(b, RandomGoSDK) }
func BenchmarkRandomFrand(b *testing.B)     { benchmarkRandom(b, RandomFrand) }
func BenchmarkRandomVFastrand(b *testing.B) { benchmarkRandom(b, RandomVFastrand) }
func BenchmarkRandomUFastrand(b *testing.B) { benchmarkRandom(b, RandomUFastrand) }
func benchmarkRandom(b *testing.B, randomFunction func() uint32) {
	for n := 0; n < b.N; n++ {
		randomFunction()
	}
}

func BenchmarkRandomizedCharacterGoSDK(b *testing.B)     { benchmarkRandomizedChar(b, RandomGoSDK) }
func BenchmarkRandomizedCharacterFrand(b *testing.B)     { benchmarkRandomizedChar(b, RandomFrand) }
func BenchmarkRandomizedCharacterVFastrand(b *testing.B) { benchmarkRandomizedChar(b, RandomVFastrand) }
func BenchmarkRandomizedCharacterUFastrand(b *testing.B) { benchmarkRandomizedChar(b, RandomUFastrand) }
func benchmarkRandomizedChar(b *testing.B, randomFunction func() uint32) {
	for n := 0; n < b.N; n++ {
		RandomizedCharacter(randomFunction)
	}
}

const iteration = 20

func TestRandomGoSDK(t *testing.T)     { testRandom(t, RandomGoSDK) }
func TestRandomFrand(t *testing.T)     { testRandom(t, RandomFrand) }
func TestRandomVFastrand(t *testing.T) { testRandom(t, RandomVFastrand) }
func TestRandomUFastrand(t *testing.T) { testRandom(t, RandomUFastrand) }
func testRandom(t *testing.T, randomFunction func() uint32) {
	var previous = randomFunction()
	for i := 0; i < iteration; i++ {
		if previous != randomFunction() {
			return
		}
	}

	assert.False(t, false, fmt.Sprintf("random is the same after %d iteration", iteration))
}
