// @Title  doublelink_test
// @Description
// @Author  EwdAger
// @Update  2021/8/3 14:26

package doublelink

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewAndString(t *testing.T) {
	l := New("abc", 123, "def")
	assert.Equal(t, fmt.Sprint(l), "[abc, 123, def]")

	l2 := New()
	assert.Equal(t, fmt.Sprint(l2), "[]")

	l3 := New(1)
	assert.Equal(t, fmt.Sprint(l3), "[1]")

}

func TestSizes(t *testing.T) {
	l := New("abc", 123, "def")
	assert.Equal(t, l.Size(), 3)

	l2 := New()
	assert.Equal(t, l2.Size(), 0)
}

func TestGetItem(t *testing.T) {
	l := New("abc", 123, "def")

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

func TestSetItem(t *testing.T) {
	l := New("abc", 123, "def")

	if assert.Nil(t, l.SetItem(1, 456)) {
		v, _ := l.GetItem(1)
		assert.Equal(t, v, 456)
	}
	assert.NotNil(t, l.SetItem(4, 456))
	assert.NotNil(t, l.SetItem(10, 456))
}

func TestInsert(t *testing.T) {
	l := New("abc", 123, "def")

	l.Insert(0, "a")
	assert.Equal(t, fmt.Sprint(l), "[a, abc, 123, def]")

	l.Insert(1, "b")
	assert.Equal(t, fmt.Sprint(l), "[a, b, abc, 123, def]")

	l.Insert(-1, "g")
	assert.Equal(t, fmt.Sprint(l), "[a, b, abc, 123, def, g]")

	l.Insert(10, "h")
	assert.Equal(t, fmt.Sprint(l), "[a, b, abc, 123, def, g, h]")

}

func TestExtend(t *testing.T) {
	l1 := New(1, 2, 3)
	l2 := New(4, 5, 6)
	l3 := []interface{}{7, 8, 9}

	if err := l1.Extend(l2); assert.Nil(t, err) {
		assert.Equal(t, fmt.Sprint(l1), "[1, 2, 3, 4, 5, 6]")
	}

	if err := l1.Extend(l3); assert.NotNil(t, err) {
	}

}

func TestReverse(t *testing.T) {
	l1 := New(1, 2, 3)
	l1.Reverse()
	assert.Equal(t, fmt.Sprint(l1), "[3, 2, 1]")

}

func TestPop(t *testing.T) {
	l1 := New(1, 2, 3)

	if item, err := l1.Pop(0); assert.Nil(t, err) {
		assert.Equal(t, item, 1)
	}

	if item, err := l1.Pop(-1); assert.Nil(t, err) {
		assert.Equal(t, item, 3)
	}

	if _, err := l1.Pop(10); assert.NotNil(t, err) {
	}

}

func TestRange(t *testing.T) {
	l := New(1, 2, 3)

	l.Range(func(idx int, val interface{}) {
		fmt.Println(val)
	})

	l.Range(func(idx int, val interface{}) {
		if _, ok := val.(int); ok && val == 1 {
			fmt.Println(val)
		}
	})
}
