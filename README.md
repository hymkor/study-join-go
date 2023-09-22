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
```

## main\_test.g0

```main_test.go
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
```

## Benchmark

```go test -bench . -benchmem |
goos: windows
goarch: amd64
pkg: github.com/hymkor/study-join
cpu: Intel(R) Core(TM) i5-6500T CPU @ 2.50GHz
Benchmark1-4   	150945446	         6.656 ns/op	       4 B/op	       0 allocs/op
Benchmark2-4   	319171430	         5.074 ns/op	       4 B/op	       0 allocs/op
Benchmark3-4   	308101213	         3.999 ns/op	       4 B/op	       0 allocs/op
PASS
ok  	github.com/hymkor/study-join	5.708s
```
