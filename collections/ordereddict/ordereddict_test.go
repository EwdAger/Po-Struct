// @Title  ordereddict_test
// @Description
// @Author  EwdAger
// @Update  2021/8/3 13:11

package ordereddict

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetAndGet(t *testing.T) {
	ordDict := New()
	ordDict.Set("a", 100)
	ordDict.Set("b", 200)
	ordDict.Set("c", 300)

	val, _ := ordDict.Get("a")
	assert.Equal(t, val, 100)

	val, ok := ordDict.Get("d")
	assert.Nil(t, val)
	assert.False(t, ok, true)

}

func TestString(t *testing.T) {
	ordDict := New()

	ordDict.Set("b", 200)
	ordDict.Set("a", 100)
	ordDict.Set("c", 300)

	assert.Equal(t, ordDict.String(), "{b: 200, a: 100, c: 300}")
}
