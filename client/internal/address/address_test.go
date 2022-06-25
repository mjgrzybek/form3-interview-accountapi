package address

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOf(t *testing.T) {
	t.Run("struct{}", func(t *testing.T) {
		assert.Equal(t, struct{}{}, *Of(struct{}{}))
	})

	t.Run("int64", func(t *testing.T) {
		assert.Equal(t, 23, *Of(23))

		const constantInt = 27
		assert.Equal(t, 27, *Of(constantInt))

		mutableInt := int64(4)
		assert.Equal(t, int64(4), *Of(mutableInt))
	})
}
