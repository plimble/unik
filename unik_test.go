package unik

import (
	"fmt"
	"testing"
)

func TestGenerateID(t *testing.T) {
	bson := NewBSON()
	bson.Generate()
	sf := NewSnowflake(0)
	fmt.Println("snowflake:", sf.Generate())
	u1 := NewUUIDV1()
	fmt.Println("uuid1:", u1.Generate())
	u4 := NewUUIDV4()
	fmt.Println("uuid4:", u4.Generate())
	u1b64 := NewUUID1Base64()
	fmt.Println("uuid1 base64:", u1b64.Generate())
}

func BenchmarkBSON(b *testing.B) {
	id := NewBSON()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.Generate()
	}
}

func BenchmarkSnowflake(b *testing.B) {
	id := NewSnowflake(0)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.Generate()
	}
}

func BenchmarkUUIDV1(b *testing.B) {
	id := NewUUIDV1()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.Generate()
	}
}

func BenchmarkUUIDV4(b *testing.B) {
	id := NewUUIDV4()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.Generate()
	}
}

func BenchmarkUUID1Base64(b *testing.B) {
	id := NewUUID1Base64()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		id.Generate()
	}
}
