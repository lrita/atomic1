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
