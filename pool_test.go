// Copyright (c) Jeevanandam M. (https://github.com/jeevatkm)
// go-aah/pool source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package pool

import (
	"bytes"
	"testing"
	"time"

	"aahframework.org/test.v0/assert"
)

func TestBufferPool(t *testing.T) {
	bufPool := NewPool(10, func() interface{} {
		return &bytes.Buffer{}
	})

	// get from pool and put it back to pool
	for idx := 0; idx < 5; idx++ {
		buf := bufPool.Get().(*bytes.Buffer)
		assert.NotNil(t, buf)

		time.Sleep(5 * time.Millisecond)
		bufPool.Put(buf)
	}

	// count
	assert.Equal(t, 1, bufPool.Count())

	// clean up
	bufPool.Drain()
}
