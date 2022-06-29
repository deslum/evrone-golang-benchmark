package benchstat_example

import (
	"strings"
	"testing"
)

// BenchmarkParallelWSDefault benchmark WriteBuffer default.
func BenchmarkParallelWSDefault(b *testing.B) {
	res := Fill()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			WSDefault(res)
		}
	})
}

// BenchmarkParallelWSCustom benchmark WriteBuffer with grow.
func BenchmarkParallelWSCustom(b *testing.B) {
	res := Fill()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			WSCustom(res)
		}
	})
}

// BenchmarkParallelJoin benchmark with classic Join func.
func BenchmarkParallelJoin(b *testing.B) {
	res := Fill()

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			strings.Join(res, "")
		}
	})
}
