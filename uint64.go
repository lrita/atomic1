package atomic1

import (
	"sync/atomic"
)

// AtomicUint64 a uint64 variable which can be changed atomic.
type AtomicUint64 struct {
	// noCopy must place ahead "v", due to https://github.com/golang/go/issues/9401
	noCopy noCopy
	v      uint64
}

// Set sets the AtomicUint64 to v
func (u *AtomicUint64) Set(v uint64) {
	atomic.StoreUint64(&u.v, v)
}

// Get return the value of this AtomicUint64
func (u *AtomicUint64) Get() uint64 {
	return atomic.LoadUint64(&u.v)
}

// CAS compare and swap it from old to n
func (u *AtomicUint64) CAS(old, n uint64) bool {
	return atomic.CompareAndSwapUint64(&u.v, old, n)
}

// SIL set the value to n if n is larger than origin.
func (u *AtomicUint64) SIL(n uint64) {
	for {
		o := u.Get()
		if o >= n {
			return
		}
		if u.CAS(o, n) {
			return
		}
	}
}

// Inc increase this AtomicUint64 value+1, and returns the new value.
func (u *AtomicUint64) Inc() uint64 {
	return atomic.AddUint64(&u.v, 1)
}
