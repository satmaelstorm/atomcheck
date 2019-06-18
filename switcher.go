package atomcheck

import "sync"

//Switch changes its value and returns the previous value.
type Switch struct {
	val bool
	mu sync.Mutex
}

//Flip value from true to false or from false to true and returns the previous value.
func (s *Switch) Flip() bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	v := s.val
	s.val = !s.val
	return v
}

//Set value and returns the previous value.
func (s *Switch) Set(v bool) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	r := s.val
	s.val = v
	return r
}
