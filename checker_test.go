package atomcheck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheck(t *testing.T) {
	var c Check
	assert.Equal(t, true, c.CheckAndStart())
	assert.Equal(t, false, c.CheckAndStart())
	c.End()
	assert.Equal(t, true, c.CheckAndStart())
}

func TestSwitch_Flip(t *testing.T) {
	var s Switch
	assert.Equal(t, false, s.Flip())
	assert.Equal(t, true, s.Flip())
	assert.Equal(t, false, s.Flip())
	assert.Equal(t, true, s.Flip())
}

func TestSwitch_Set(t *testing.T) {
	var s Switch
	assert.Equal(t, false, s.Set(true))
	assert.Equal(t, true, s.Set(true))
	assert.Equal(t, true, s.Set(false))
	assert.Equal(t, false, s.Set(false))
	assert.Equal(t, false, s.Set(false))
}

func BenchmarkCheck(b *testing.B) {
	var s Check
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if s.CheckAndStart() {
				s.End()
			}
		}
	})
}
