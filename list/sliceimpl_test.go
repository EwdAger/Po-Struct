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

func TestAppend(t *testing.T) {
	l := List("", "abc", 123, "def")
	l.Append(10)

	assert.Equal(t, l.Cap(), 6)

	l2 := List("", 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	l2.Append(11)
	assert.Equal(t, l2.Cap(), 17)
}
