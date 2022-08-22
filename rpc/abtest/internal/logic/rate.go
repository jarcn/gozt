package logic

import (
	"math/rand"
	"sync/atomic"
	"time"
)

var base int = 10
var opts int64
var source = make([]int, 10)

func rate() int64 {
	return int64(source[atomic.AddInt64(&opts, 1)%int64(base)])
}

func init() {
	atomic.AddInt64(&opts, 0)
	for i := 0; i < 10; i++ {
		source = append(source, i)
	}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(source), func(i, j int) { source[i], source[j] = source[j], source[i] })
}
