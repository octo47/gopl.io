package main
import (
	"testing"
	"strings"
)

var result string

func BenchmarkJoin(b *testing.B) {
	arr := make([]string, b.N, b.N)
	for i := 0; i < len(arr); i++ {
		arr[i] = "Hello"
	}
	result = strings.Join(arr, " ")
}

func BenchmarkAppend(b *testing.B) {
	result := ""
	for i := 0; i < b.N; i++ {
		result += "Hello"
		result += " "
	}
}
