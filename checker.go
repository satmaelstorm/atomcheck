package atomcheck

import "sync"

type Check struct {
	val bool
	mu sync.Mutex
}

func (c *Check) CheckAndStart() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.val {
		return false
	}
	c.val = true
	return true
}

func (c *Check) End() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val = false
}
