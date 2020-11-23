package atomcheck

import "sync"

//Designed to verify that something is running.
type Check struct {
	val bool
	mu sync.RWMutex
}

//Checks its value, if true - returns false, if false - sets its value to true and returns true
func (c *Check) CheckAndStart() bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if c.val {
		return false
	}
	c.val = true
	return true
}

//Sets its value to false
func (c *Check) End() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.val = false
}

//Return current value
func (c *Check) Current() bool {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.val
}
