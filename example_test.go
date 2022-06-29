package benchstat_example

import (
	"strconv"
	"strings"
	"testing"
)

const epochCount = 10000

// Fill generate slice string data.
func Fill() []string {
	var result = make([]string, epochCount)

	for i := 0; i < epochCount; i++ {
		result[i] = strconv.Itoa(i)
	}

	return result
}

// BenchmarkWSDefault benchmark WriteBuffer default.
func BenchmarkWSDefault(b *testing.B) {
	res := Fill()

	b.Run("BenchmarkWSDefault", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			WSDefault(res)
		}
	})
}

// BenchmarkWSDefault benchmark WriteBuffer custom with grow.
func BenchmarkWSCustom(b *testing.B) {
	res := Fill()

	b.Run("BenchmarkWSCustom", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			WSCustom(res)
		}
	})
}

// BenchmarkWBDefault benchmark classic join command.
func BenchmarkJoin(b *testing.B) {
	res := Fill()

	b.Run("BenchmarkJoin", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			strings.Join(res, "")
		}
	})
}
