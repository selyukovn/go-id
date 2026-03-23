package like_uuid

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GenerateUniqueRandom(t *testing.T) {
	t.Run("success unique", func(t *testing.T) {
		n := 1_000_000
		prevIds := make(map[Id]struct{}, n)

		for i := 0; i < n; i++ {
			id, err := GenerateUniqueRandom()

			assert.NotEqual(t, IdNil, id)
			assert.NoError(t, err)
			assert.False(t, func() bool { _, ok := prevIds[id]; return ok }() /* assert.NotContains -- O(n) */)

			prevIds[id] = struct{}{}
		}
	})
}
