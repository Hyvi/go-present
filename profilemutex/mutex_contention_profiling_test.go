//  https://making.pusher.com/go-tool-trace/
// rykyll 关于pprof的文档
package gomutex

import (
	"fmt"
	"sync"
	"testing"
)

func factors(nr int64) []int64 {
	fs := make([]int64, 1)
	if nr < 1 {
		fmt.Println("\nFactors of", nr, "not computed")
		// err := errors.New("\nFactors of" + string(nr) + "not computed")
		return fs
	}
	// fmt.Printf("\nFactors of %d: ", nr)
	fs[0] = 1
	apf := func(p int64, e int) {
		n := len(fs)
		for i, pp := 0, p; i < e; i, pp = i+1, pp*p {
			for j := 0; j < n; j++ {
				fs = append(fs, fs[j]*pp)
			}
		}
	}
	e := 0
	for ; nr&1 == 0; e++ {
		nr >>= 1
	}
	apf(2, e)
	for d := int64(3); nr > 1; d += 2 {
		if d*d > nr {
			d = nr
		}
		for e = 0; nr%d == 0; e++ {
			nr /= d
		}
		if e > 0 {
			apf(d, e)
		}
	}
	// fmt.Println(fs)
	// fmt.Println("Number of factors =", len(fs))
	return fs
}

var mu sync.Mutex

func FactorsCount(n int64) {
	m := make([]int64, n+1)
	mu.Lock()
	for _, f := range factors(n) {
		m[f]++
	}
	mu.Unlock()
}

func TestFactorsCount(t *testing.T) {

	for i := 1; i < 10000; i++ {
		go FactorsCount(int64(i))
	}
}

func BenchmarkFactorsCount(b *testing.B) {
	for i := 1; i < b.N; i++ {
		FactorsCount(int64(i))
	}
}

func TestLock(t *testing.T) {
	var mu sync.Mutex
	var items = make(map[int]struct{})

	for i := 0; i < 1000*1000; i++ {
		go func(i int) {
			mu.Lock()
			defer mu.Unlock()

			items[i] = struct{}{}
		}(i)
	}

}
