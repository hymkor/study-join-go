package studyjoin

import (
	"testing"
)

func Benchmark1(B *testing.B) {
	join1(B.N)
}

func Benchmark2(B *testing.B) {
	join2(B.N)
}

func Benchmark3(B *testing.B) {
	join3(B.N)
}
