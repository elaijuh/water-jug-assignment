package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {

	t.Run("test gcd func", func(t *testing.T) {
		s := NewProblemResolver(2, 10, 4)
		assert.Equal(t, 1, s.gcd(5, 3))
		assert.Equal(t, 2, s.gcd(2, 10))
	})

	t.Run("test Resolvable func", func(t *testing.T) {
		s := NewProblemResolver(2, 10, 4)
		assert.True(t, s.Resolvable())
		s = NewProblemResolver(17, 23, 11)
		assert.True(t, s.Resolvable())
		s = NewProblemResolver(2, 10, 3)
		assert.False(t, s.Resolvable())
		s = NewProblemResolver(3, 5, 10)
		assert.False(t, s.Resolvable())
	})

	t.Run("test act func", func(t *testing.T) {
		// x=2 y=10 z=4
		s := NewProblemResolver(2, 10, 4)
		assert.Equal(t, 4, s.act('x').steps)
		assert.Equal(t, 6, s.act('y').steps)
	})
}
