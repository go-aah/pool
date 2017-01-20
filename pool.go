// Copyright (c) Jeevanandam M (https://github.com/jeevatkm)
// go-aah/pool source code and usage is governed by a MIT style
// license that can be found in the LICENSE file.

package pool

// Pool holds the bounded channel for interface{}.
type Pool struct {
	// c bounded channel
	c chan interface{}

	// New optionally specifies a function to generate
	// a value when Get would otherwise return nil.
	// It may not be changed concurrently with calls to Get.
	New func() interface{}
}

// NewPool method creates a new Pool bounded to the given size.
func NewPool(size int, fn func() interface{}) (p *Pool) {
	return &Pool{
		c:   make(chan interface{}, size),
		New: fn,
	}
}

// Get method gets a value from the Pool, or creates a new one if none are
// available in the pool.
func (p *Pool) Get() (v interface{}) {
	select {
	case v = <-p.c:
	// reuse from pool
	default:
		// create new one
		if p.New != nil {
			v = p.New()
		}
	}
	return
}

// Put method returns given value into Pool.
func (p *Pool) Put(v interface{}) {
	select {
	case p.c <- v:
	default: // Discard the value if the pool is full.
	}
}

// Count method returns current count of pool.
func (p *Pool) Count() int {
	return len(p.c)
}

// Drain method drains all the values for the channel
func (p *Pool) Drain() {
	for {
		select {
		case <-p.c:
			// draining it
		default:
			return
		}
	}
}
