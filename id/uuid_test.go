package id

import (
	"fmt"
	"testing"
)

func TestNewUUID(t *testing.T) {
	t.Log(NewUUID(1, nil))
	t.Log(NewUUID(2, nil))
	t.Log(NewUUID(3, []byte("nil")))
	t.Log(NewUUID(4, nil))
	t.Log(NewUUID(5, []byte("nil")))
}

func BenchmarkNewUUID_Version1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewUUID(1, nil)
	}
}
func BenchmarkNewUUID_Version2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewUUID(2, nil)
	}
}
func BenchmarkNewUUID_Version3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewUUID(3, []byte("googege"))
	}
}
func BenchmarkNewUUID_Version4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewUUID(4, nil)
	}
}
func BenchmarkNewUUID_Version5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewUUID(5, []byte("googege"))
	}
}
