package internal

import "math/rand/v2"

const min = 0
const max = 6

func GetRND() int {
	return min + rand.IntN(max-min)
}
