// @Title  set_test
// @Description
// @Author  EwdAger
// @Update  2021/7/17 13:42

package hashset

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContains(t *testing.T) {
	a := New("a", "b", "c", 1, 2, 3)
	assert.Equal(t, a.Contains(1, 2, 3), true)
	assert.Equal(t, a.Contains(1, 2, 3, 4), false)

}

func TestValues(t *testing.T) {
	a := New("a", "b", "c", 1, 2, 3)
	fmt.Println(a.Values())
	assert.Equal(t, a.Contains(a.Values()...), true)
}

func TestInit(t *testing.T) {
	s := New("a", "b", "c", 1, 2, 3)
	assert.Equal(t, Equal(s, s), true)
}

func TestSubAndSuper(t *testing.T) {
	a := New("a", "b", "c", 1, 2, 3)
	b := New(1, 3)

	assert.Equal(t, IsSubSet(b, a), true)
	assert.Equal(t, IsSuperSet(a, b), true)

	assert.Equal(t, b.IsSubSet(a), true)
	assert.Equal(t, a.IsSuperSet(b), true)
}

func TestUnion(t *testing.T) {
	a := New("a", "b", "c", 1, 2, 3)
	b := New(1, 3)

	assert.Equal(t, Equal(Union(a, b), a), true)
	assert.Equal(t, Equal(a.Union(b), a), true)

}

func TestDifference(t *testing.T) {
	a := New("a", "b", "c", 1, 2, 3)
	b := New(1, 2, 3)
	c := New("a", "b", "c")

	assert.Equal(t, Equal(Difference(a, b), c), true)
	assert.Equal(t, Equal(a.Difference(b), c), true)

}

func TestIntersection(t *testing.T) {
	a := New("a", "b", "c", 1, 2, 3)
	b := New(1, 2, 3)

	assert.Equal(t, Equal(Intersection(a, b), b), true)
	assert.Equal(t, b.Equal(a.Intersection(b)), true)

}

func TestRange(t *testing.T) {
	a := New("a", "b", "c", 1, 2, 3)

	a.Range(func(key interface{}) {
		fmt.Println(key)
	})

	a.Range(func(key interface{}) {
		if _, ok := key.(int); ok && key == 1 {
			fmt.Println(key)
		}
	})

}
