package atomic1

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestAtomicUint64(t *testing.T) {
	var (
		assert = assert.New(t)
		u      AtomicUint64
	)
	assert.True(u.Get() == 0)
	u.Set(100)
	assert.True(u.Get() == 100)
	assert.True(u.CAS(100, 200))
	assert.False(u.CAS(100, 200))
	u.SIL(100)
	assert.True(u.Get() == 200)
	u.SIL(300)
	assert.True(u.Get() == 300)
	u.Inc()
	assert.True(u.Get() == 301)
}

func TestAtomicUint64Aligned(t *testing.T) {
	var (
		assert = assert.New(t)
		o      int64
		d      = struct {
			x uint8
			c [1]byte
			u AtomicUint64
		}{}
		u AtomicUint64
	)
	assert.Equal(unsafe.Sizeof(o), unsafe.Sizeof(u))
	assert.Equal(unsafe.Sizeof(o), unsafe.Alignof(d.u))
}
