package main

import (
	"testing"
)

func BenchmarkIndex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Index("enron_mail_20110402.tgz", 1_000)
	}
}
