package main

import (
	"fmt"
	"math/rand/v2"
	"time"
)

var space = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func fileGen(input string, mimeType string) string {
	if len(input) == 0 {
		panic("fileGen: input is empty")
	}

	fmt.Println("FileGen, input", input, mimeType)

	var (
		sLen = len(input)
		out  = make([]rune, 8)
		rng  = rand.New(rand.NewPCG(8_675_309, uint64(time.Now().Unix())))
	)

	out[0] = 'j'

	for pos := 1; pos < 8; pos++ {
		out[pos] = space[rng.IntN(sLen)]
	}

	return string(out) + mimeToExt(mimeType)
}

func mimeToExt(mimeType string) string {
	var ext string

	switch mimeType {
	case "image/jpeg":
		ext = ".jpg"
	case "image/png":
		ext = ".png"
	case "image/gif":
		ext = ".gif"
	case "image/bmp":
		ext = ".bmp"
	case "image/tiff":
		ext = ".tiff"
	case "image/webp":
		ext = ".webp"
	}

	return ext
}
