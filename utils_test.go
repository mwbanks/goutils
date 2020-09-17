package goutils_test

import (
	"testing"

	"github.com/mwbanks/goutils"
)

const n = 16

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		goutils.RandString(n)

	}
}
