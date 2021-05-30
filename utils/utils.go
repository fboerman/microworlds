package utils

import (
	"math/rand"
	"time"
)

type UniqueRand struct {
	generated map[int]bool
}

func (u *UniqueRand) Intn(n int) int {
	for {
		i := rand.Intn(n)
		if !u.generated[i] {
			u.generated[i] = true
			return i
		}
	}
}

func RandomListN(N, n int) []int {
	rand.Seed(time.Now().UnixNano())
	var UR UniqueRand
	UR.generated = make(map[int]bool)
	l := make([]int, N)
	for i := 0; i < N; i++ {
		l[i] = UR.Intn(n)
	}

	return l
}

func Randomn(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}
