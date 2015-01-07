package unik

import (
	"testing"
)

func TestGenerateID(t *testing.T) {
	bson := NewBSON()
	bson.Generate()
	sf := NewSnowflake(0)
	sf.Generate()
	u1 := NewUUIDV1()
	u1.Generate()
	u4 := NewUUIDV4()
	u4.Generate()
}

func BenchmarkBSON(b *testing.B) {
	id := NewBSON()
	for i := 0; i < b.N; i++ {
		id.Generate()
	}
}

func BenchmarkSnowflake(b *testing.B) {
	id := NewSnowflake(0)
	for i := 0; i < b.N; i++ {
		id.Generate()
	}
}

func BenchmarkUUIDV1(b *testing.B) {
	id := NewUUIDV1()
	for i := 0; i < b.N; i++ {
		id.Generate()
	}
}

func BenchmarkUUIDV4(b *testing.B) {
	id := NewUUIDV4()
	for i := 0; i < b.N; i++ {
		id.Generate()
	}
}
