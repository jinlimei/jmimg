package main

import (
	"math/rand/v2"
	"path/filepath"
	"time"
)

var space = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func fileGen(input string) string {
	ext := filepath.Ext(input)

	var (
		sLen = len(input)
		out  = make([]rune, 8)
		rng  = rand.New(rand.NewPCG(8_675_309, uint64(time.Now().Unix())))
	)

	out[0] = 'j'

	for i := 1; i < 8; i++ {
		out[i] = space[rng.IntN(sLen)]
	}

	return string(out) + ext
}
