package studyjoin

import (
	"strings"
)

func join1(n int) string {
	var buffer strings.Builder
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

	buffer.WriteString("foo")
	for i := 1; i < n; i++ {
		buffer.WriteString(",")
		buffer.WriteString("foo")
	}
	return buffer.String()
}
