package atomic1

import (
	"sync/atomic"
)

// AtomicInt64 a int64 variable which can be changed atomic.
type AtomicInt64 struct {
	// noCopy must place ahead "v", due to https://github.com/golang/go/issues/9401
	noCopy noCopy
	v      int64
}

// Set sets the AtomicInt64 to v
func (u *AtomicInt64) Set(v int64) {
	atomic.StoreInt64(&u.v, v)
}

// Get return the value of this AtomicInt64
func (u *AtomicInt64) Get() int64 {
	return atomic.LoadInt64(&u.v)
}

// CAS compare and swap it from old to n
func (u *AtomicInt64) CAS(old, n int64) bool {
	return atomic.CompareAndSwapInt64(&u.v, old, n)
}

// ANL add the value with i and keep it is not larger than max.
func (u *AtomicInt64) ANL(i, max int64) {
	for {
		o := u.Get()
		if o >= max {
			return
		}
		n := o + i
		if n > max {
			n = max
		}
		if u.CAS(o, n) {
			return
		}
	}
}

// SNL subtraction the value with i and keep it is not little than min.
func (u *AtomicInt64) SNL(i, min int64) {
	for {
		o := u.Get()
		if o <= min {
			return
		}
		n := o - i
		if n < min {
			n = min
		}
		if u.CAS(o, n) {
			return
		}
	}
}

// SIL set the value to n if n is larger than origin.
func (u *AtomicInt64) SIL(n int64) bool {
	for {
		o := u.Get()
		if o >= n {
			return false
		}
		if u.CAS(o, n) {
			return true
		}
	}
}

// Add atomically adds the d to this value, and returns the new value.
func (u *AtomicInt64) Add(d int64) int64 {
	return atomic.AddInt64(&u.v, d)
}
