package atomic1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type e1 struct{}

func (e *e1) Error() string {
	return "e1"
}

type e2 struct{}

func (e *e2) Error() string {
	return "e2"
}

func TestAtomicError(t *testing.T) {
	var (
		err    error
		assert = assert.New(t)
		b      AtomicError
	)
	assert.NoError(b.Get())

	b.Set(&e1{})
	assert.EqualError(b.Get(), (&e1{}).Error())

	b.Set(nil)
	assert.NoError(b.Get())

	b.Set(&e2{})
	assert.EqualError(b.Get(), (&e2{}).Error())

	err = &b

	assert.EqualError(err, (&e2{}).Error())
}
