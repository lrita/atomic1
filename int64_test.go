package atomic1

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestAtomicInt64(t *testing.T) {
	var (
		assert = assert.New(t)
		u      AtomicInt64
	)
	assert.True(u.Get() == 0)
	u.Set(100)
	assert.True(u.Get() == 100)
	assert.True(u.CAS(100, 200))
	assert.False(u.CAS(100, 200))
	assert.False(u.SIL(100))
	assert.True(u.Get() == 200)
	assert.True(u.SIL(300))
	assert.True(u.Get() == 300)
	u.Add(1)
	assert.True(u.Get() == 301)
	u.Add(-302)
	assert.True(u.Get() == -1)
	u.ANL(100, 2)
	u.ANL(100, 2)
	assert.True(u.Get() == 2)
	u.SNL(100, 0)
	u.SNL(100, 0)
	assert.True(u.Get() == 0)
}

func TestAtomicInt64Aligned(t *testing.T) {
	var (
		assert = assert.New(t)
		o      int64
		d      = struct {
			x uint8
			c [1]byte
			u AtomicInt64
		}{}
		u AtomicInt64
	)
	assert.Equal(unsafe.Sizeof(o), unsafe.Sizeof(u))
	assert.Equal(unsafe.Sizeof(o), unsafe.Alignof(d.u))

	assert.True(uintptr(unsafe.Pointer(&d.u))%unsafe.Sizeof(o) == 0)
}

func BenchmarkInt64Add(b *testing.B) {
	var u AtomicInt64
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			u.Add(1)
		}
	})
	assert.EqualValues(b, b.N, u.Get())
}

func BenchmarkInt64ANL(b *testing.B) {
	var u AtomicInt64
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			u.ANL(1, 100)
			u.Add(-1)
		}
	})
}

func BenchmarkInt64SNL(b *testing.B) {
	var u AtomicInt64
	u.Set(1000)
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			u.SNL(1, 0)
			u.Add(1)
		}
	})
}

func BenchmarkInt64CAS(b *testing.B) {
	var u AtomicInt64
	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			u.CAS(1, 2)
		}
	})
}
