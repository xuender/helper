package cont_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xuender/helper/cont"
)

func TestSet_AddAll(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	nums := cont.NewSet(1, 2, 3)

	nums.AddSet(cont.NewSet(3, 4, 5))
	ass.Len(nums, 5)
}

func TestSet_Values(t *testing.T) {
	t.Parallel()

	ass := assert.New(t)
	nums := cont.NewSet(1)

	for num := range nums.Values() {
		ass.Equal(1, num)
	}
}
