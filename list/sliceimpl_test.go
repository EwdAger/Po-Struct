// @Title  sliceimpl_test
// @Description
// @Author  EwdAger
// @Update  2021/7/15 20:48

package list

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAndString(t *testing.T) {
	l := List("", "abc", 123, "def")
	fmt.Println(l)

	l2 := List("")
	fmt.Println(l2)
}

func TestSizes(t *testing.T) {
	l := List("", "abc", 123, "def")
	assert.Equal(t, l.Size(), 3)
}

func TestGetItem(t *testing.T) {
	l := List("", "abc", 123, "def")

	if res, err := l.GetItem(0); assert.Nil(t, err) {
		assert.Equal(t, res, "abc")
	}

	if res, err := l.GetItem(1); assert.Nil(t, err) {
		assert.Equal(t, res, 123)
	}

	if res, err := l.GetItem(-1); assert.Nil(t, err) {
		assert.Equal(t, res, "def")
	}

	if res, err := l.GetItem(-3); assert.Nil(t, err) {
		assert.Equal(t, res, "abc")
	}

	if _, err := l.GetItem(-4); assert.NotNil(t, err) {
	}

	if _, err := l.GetItem(3); assert.NotNil(t, err) {
	}
}

func TestEqual(t *testing.T) {
	l1 := List("", "abc", 123, "def")
	l2 := List("", "abc", 123, "def")

	assert.Equal(t, Equal(l1, l2), true)

}

func TestIs(t *testing.T) {
	l1 := List("", "abc", 123, "def")
	l2 := List("", "abc", 123, "def")

	assert.Equal(t, Is(&l1, &l2), false)
	assert.Equal(t, Is(&l1, &l1), true)
}

func TestAppend(t *testing.T) {
	l := List("", "abc", 123, "def")
	l.Append(10)

	assert.Equal(t, l.Cap(), 7)

	l2 := List("", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	l2.Append(11)
	assert.Equal(t, l2.Cap(), 18)

	l3 := List("", 1)
	l3.Append(2, 3, 4, 5, 6)
	assert.Equal(t, l3.Cap(), 9)
}

func TestInsert(t *testing.T) {
	l := List("", 1, 2, 4)

	l.Insert(2, 3)
	assert.Equal(t, fmt.Sprint(l), "[1, 2, 3, 4]")

	l.Insert(0, 0)
	assert.Equal(t, fmt.Sprint(l), "[0, 1, 2, 3, 4]")

	l.Insert(-1, 5)
	assert.Equal(t, fmt.Sprint(l), "[0, 1, 2, 3, 4, 5]")

	l.Insert(-6, -1)
	assert.Equal(t, fmt.Sprint(l), "[-1, 0, 1, 2, 3, 4, 5]")

	l.Insert(100, 6)
	assert.Equal(t, fmt.Sprint(l), "[-1, 0, 1, 2, 3, 4, 5, 6]")

	l.Insert(-100, -2)
	assert.Equal(t, fmt.Sprint(l), "[-2, -1, 0, 1, 2, 3, 4, 5, 6]")
}

func TestGetSlice(t *testing.T) {
	l := List("", 1, 2, 3)

	l2 := l.GetSlice(0, l.Size())
	assert.Equal(t, fmt.Sprint(l2), "[1, 2, 3]")

	l3 := l.GetSlice(1, l.Size())
	assert.Equal(t, fmt.Sprint(l3), "[2, 3]")

	l4 := l.GetSlice(1, 2)
	assert.Equal(t, fmt.Sprint(l4), "[2]")

	l5 := l.GetSlice(0, -1)
	assert.Equal(t, fmt.Sprint(l5), "[1, 2]")

	l6 := l.GetSlice(-3, -2)
	assert.Equal(t, fmt.Sprint(l6), "[1]")

	l7 := l.GetSlice(4, 6)
	assert.Equal(t, fmt.Sprint(l7), "[]")

	l8 := l.GetSlice(4, 0)
	assert.Equal(t, fmt.Sprint(l8), "[]")
}

func TestExtend(t *testing.T) {
	l1 := List("", 1, 2, 3)
	l2 := List("", 4, 5, 6)
	l3 := []interface{}{7, 8, 9}

	if err := l1.Extend(l2); assert.Nil(t, err) {
		assert.Equal(t, fmt.Sprint(l1), "[1, 2, 3, 4, 5, 6]")
		assert.Equal(t, l1.Cap(), 9)
	}

	if err := l1.Extend(l3); assert.NotNil(t, err) {
	}

}

func TestReverse(t *testing.T) {
	l1 := List("", 1, 2, 3)

	l1.Reverse()
	assert.Equal(t, fmt.Sprint(l1), "[3, 2, 1]")
}

func TestPop(t *testing.T) {
	l1 := List("", 1, 2, 3)

	if item, err := l1.Pop(0); assert.Nil(t, err) {
		assert.Equal(t, item, 1)
	}

	if item, err := l1.Pop(-1); assert.Nil(t, err) {
		assert.Equal(t, item, 3)
	}

	if _, err := l1.Pop(10); assert.NotNil(t, err) {
	}

}
