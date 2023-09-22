study-join-go
=============

The best algorithm to join strings with a separator

## main.go

```main.go
package studyjoin

import (
    "strings"
)

func join1(n int) string {
    var buffer strings.Builder
    buffer.Grow(n * 4)
    sep := ""

    for i := 0; i < n; i++ {
        buffer.WriteString(sep)
        buffer.WriteString("foo")
        sep = ","
    }
    return buffer.String()
}

func join2(n int) string {
    var buffer strings.Builder
    buffer.Grow(n * 4)

    for i := 0; i < n; i++ {
        if i > 0 {
            buffer.WriteString(",")
        }
        buffer.WriteString("foo")
    }
    return buffer.String()
}

func join3(n int) string {
    var buffer strings.Builder
    buffer.Grow(n * 4)

    buffer.WriteString("foo")
    for i := 1; i < n; i++ {
        buffer.WriteString(",")
        buffer.WriteString("foo")
    }
    return buffer.String()
}

func join4(n int) string {
    var buffer strings.Builder
    buffer.Grow(n * 4)

    for i := 0; ; {
        buffer.WriteString("foo")
        i++
        if i >= n {
            break
        }
        buffer.WriteString(",")
    }
    return buffer.String()
}
```

## main\_test.go

```main_test.go
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
```

## Benchmark

`go test -bench . -benchmem`

```go test -bench . -benchmem |
goos: windows
goarch: amd64
pkg: github.com/hymkor/study-join
cpu: Intel(R) Core(TM) i5-6500T CPU @ 2.50GHz
Benchmark1-4   	171278196	         6.783 ns/op	       4 B/op	       0 allocs/op
Benchmark2-4   	286658004	         5.637 ns/op	       4 B/op	       0 allocs/op
Benchmark3-4   	289203525	         4.104 ns/op	       4 B/op	       0 allocs/op
Benchmark4-4   	323644233	         3.937 ns/op	       4 B/op	       0 allocs/op
PASS
ok  	github.com/hymkor/study-join	8.575s
```
