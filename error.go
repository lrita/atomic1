package atomic1

import (
	"sync/atomic"
)

// AtomicError a error variable which can be changed atomic.
type AtomicError struct {
	// noCopy must place ahead "v", due to https://github.com/golang/go/issues/9401
	noCopy noCopy
	e      atomic.Value
}

// Set sets the AtomicError to v
func (e *AtomicError) Set(v error) {
	e.e.Store(struct{ error }{v})
}

// Get return the value of this AtomicError
func (e *AtomicError) Get() error {
	v, _ := e.e.Load().(struct{ error })
	return v.error
}

// Error implements the error interface.
func (e *AtomicError) Error() string {
	v := e.Get()
	if v == nil {
		return ""
	}
	return v.Error()
}
