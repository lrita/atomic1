package atomic1

import (
	"sync/atomic"
)

// AtomicBool a bool variable which can be changed atomic.
type AtomicBool struct {
	// noCopy must place ahead "v", due to https://github.com/golang/go/issues/9401
	noCopy noCopy
	v      uint32
}

// Set sets the AtomicBool to v
func (b *AtomicBool) Set(v bool) {
	var i uint32
	if v {
		i = 1
	}
	atomic.StoreUint32(&b.v, i)
}

// CAS compare and swap it from !v to v
func (b *AtomicBool) CAS(v bool) bool {
	var i, o uint32
	if v {
		i = 1
	} else {
		o = 1
	}
	return atomic.CompareAndSwapUint32(&b.v, o, i)
}

// Get return the value of this AtomicBool
func (b *AtomicBool) Get() bool {
	return atomic.LoadUint32(&b.v) == 1
}
