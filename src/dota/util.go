package dota

import (
	"math/rand"
	"time"
)

func RandInt(min, max int) int {
	if max <= min {
		return 0
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(max-min+1) + min
}

func RandInt64(min, max int64) int64 {
	if max <= min {
		return 0
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(max-min+1) + min
}

func RandIntn(n int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(n)
}

func RandInt64n(n int64) int64 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Int63n(n)
}

func RandPerm(n int) []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Perm(n)
}

func RandSlice(s []int, length int) []int {
	r := make([]int, length)
	perm := RandPerm(length)
	for i, v := range perm {
		r[v] = s[i]
	}
	return r
}
