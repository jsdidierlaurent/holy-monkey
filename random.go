package main

import (
	"math/rand"
	"time"

	vFastrand "github.com/valyala/fastrand"
	nFastrand "gitlab.com/NebulousLabs/fastrand"
	"lukechampine.com/frand"
)

// Alphabet is used as base for randomize character
const Alphabet = `ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz.,;:!?'" `

// Computed
var alphabetLength uint32 = 0
var vFastrandRNG *vFastrand.RNG

func init() {
	// Init computed value
	alphabetLength = uint32(len(Alphabet))

	// Init random Seed
	rand.Seed(time.Now().UnixNano())
	vFastrandRNG = &vFastrand.RNG{}
}

func RandomizedCharacter(randomFunction func() uint32) byte {
	return Alphabet[randomFunction()]
}

func RandomGoSDK() uint32 {
	return uint32(rand.Int31n(int32(alphabetLength)))
}

func RandomFrand() uint32 {
	return uint32(frand.Intn(int(alphabetLength)))
}

func RandomVFastrand() uint32 {
	return vFastrandRNG.Uint32n(alphabetLength)
}

func RandomUFastrand() uint32 {
	return uint32(nFastrand.Intn(int(alphabetLength)))
}
