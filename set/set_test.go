// @Title  set_test
// @Description
// @Author  EwdAger
// @Update  2021/7/17 13:42

package set

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	s := Set("a", "b", "c", 1, 2, 3)
	assert.Equal(t, Equal(s, s), true)
}

func TestSubAndSuper(t *testing.T) {
	a := Set("a", "b", "c", 1, 2, 3)
	b := Set(1, 3)

	assert.Equal(t, IsSubSet(b, a), true)
	assert.Equal(t, IsSuperSet(a, b), true)

	assert.Equal(t, b.IsSubSet(a), true)
	assert.Equal(t, a.IsSuperSet(b), true)
}

func TestUnion(t *testing.T) {
	a := Set("a", "b", "c", 1, 2, 3)
	b := Set(1, 3)

	assert.Equal(t, Equal(Union(a, b), a), true)
	assert.Equal(t, Equal(a.Union(b), a), true)

}

func TestDifference(t *testing.T) {
	a := Set("a", "b", "c", 1, 2, 3)
	b := Set(1, 2, 3)
	c := Set("a", "b", "c")

	assert.Equal(t, Equal(Difference(a, b), c), true)
	assert.Equal(t, Equal(a.Difference(b), c), true)

}

func TestIntersection(t *testing.T) {
	a := Set("a", "b", "c", 1, 2, 3)
	b := Set(1, 2, 3)

	assert.Equal(t, Equal(Intersection(a, b), b), true)
	assert.Equal(t, b.Equal(a.Intersection(b)), true)

}

func TestRange(t *testing.T) {
	a := Set("a", "b", "c", 1, 2, 3)

	a.Range(func(key interface{}) {
		fmt.Println(key)
	})

	a.Range(func(key interface{}) {
		if _, ok := key.(int); ok && key == 1 {
			fmt.Println(key)
		}
	})

}
