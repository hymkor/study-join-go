package studyjoin

import (
	"testing"
)

func Test1(t *testing.T) {
	if join1(3) != "foo,foo,foo" {
		t.Fatal()
	}
}

func Benchmark1(B *testing.B) {
	join1(B.N)
}

func Test2(t *testing.T) {
	if join2(3) != "foo,foo,foo" {
		t.Fatal()
	}
}

func Benchmark2(B *testing.B) {
	join2(B.N)
}

func Test3(t *testing.T) {
	if join3(3) != "foo,foo,foo" {
		t.Fatal()
	}
}

func Benchmark3(B *testing.B) {
	join3(B.N)
}

func Test4(t *testing.T) {
	if join4(3) != "foo,foo,foo" {
		t.Fatal()
	}
}

func Benchmark4(B *testing.B) {
	join4(B.N)
}
