package atomic1

import (
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestAtomicBool(t *testing.T) {
	var (
		assert = assert.New(t)
		b      AtomicBool
	)
	assert.False(b.Get())
	b.Set(false)
	assert.False(b.Get())
	b.Set(true)
	assert.True(b.Get())
	assert.False(b.CAS(true))
	assert.True(b.CAS(false))
	assert.False(b.Get())
}

func TestAtomicBoolAligned(t *testing.T) {
	var (
		assert = assert.New(t)
		o      uint32
		d      = struct {
			x uint8
			u AtomicBool
		}{}
		u AtomicBool
	)
	assert.Equal(unsafe.Sizeof(o), unsafe.Sizeof(u))
	assert.Equal(unsafe.Sizeof(o), unsafe.Alignof(d.u))
}
