// @Title  slice implement
// @Description  切片实现
// @Author  EwdAger
// @Update  2021/7/6 13:14

package list

import (
	"errors"
	"fmt"
	"strings"
)

type sliceImpl struct {
	items []interface{}
}

func (l *sliceImpl) Init(values ...interface{}) {
	l.items = values
}

func (l sliceImpl) Size() int {
	return len(l.items)

}

func (l sliceImpl) Cap() int {
	return cap(l.items)

}

func (l sliceImpl) String() string {
	strSli := make([]string, l.Size())

	for i, v := range l.items {
		strSli[i] = fmt.Sprintf("%v", v)
	}

	res := strings.Join(strSli, ", ")
	return fmt.Sprintf("[%s]", res)
}

func (l sliceImpl) GetItem(index int) (interface{}, error) {
	if index == 0 {
		return l.items[0], nil
	}

	// 倒叙索引 feature
	reverseIndex := false

	if index < 0 {
		index = -index
		reverseIndex = true
	}

	if index >= l.Size() && reverseIndex == false || index > l.Size() && reverseIndex == true {
		return nil, errors.New("list index out of range")
	}

	if reverseIndex == true {
		return l.items[l.Size()-index], nil
	} else {
		return l.items[index], nil
	}

}

func (l *sliceImpl) SetItem(index int, value interface{}) error {
	// 倒叙索引 feature
	reverseIndex := false

	if index < 0 {
		index = -index
		reverseIndex = true
	}

	if index >= l.Size() && reverseIndex == false || index > l.Size() && reverseIndex == true {
		return errors.New("list index out of range")
	}

	if reverseIndex == true {
		l.items[l.Size()-index] = value
	} else {
		l.items[index] = value
	}
	return nil

}

func (l *sliceImpl) Append(v interface{}) {
	if need := needGrow(l); need == true {
		err := grow(l)
		if err != nil {
			return
		}
	}

	l.items = append(l.items, v)

}

func (l sliceImpl) Insert() ([]interface{}, error) {
	panic("implement me")
}

func (l sliceImpl) GetSlice() ([]interface{}, error) {
	panic("implement me")
}

func (l sliceImpl) Sort() ([]interface{}, error) {
	panic("implement me")
}

func (l sliceImpl) Reverse() ([]interface{}, error) {
	panic("implement me")
}

func (l sliceImpl) AsArray() ([]interface{}, error) {
	panic("implement me")
}

func (l sliceImpl) Clear() error {
	panic("implement me")
}

// 判断是否需要扩容
func needGrow(l *sliceImpl) bool {
	return len(l.items) == cap(l.items)
}

// 每次扩容扩大 1/8 再加上 3 或 6 的余量
func grow(l *sliceImpl) error {
	nowCap := cap(l.items)

	// 扩容的余量，防止容量很小时频繁扩容, 至于为啥是 3 和 6，这里直接抄了 Python 的源码
	// new_allocated = (size_t)newsize + (newsize >> 3) + (newsize < 9 ? 3 : 6);
	var over int
	if nowCap < 9 {
		over = 3
	} else {
		over = 6
	}

	newCap := nowCap + (nowCap >> 3) + over
	newSli := make([]interface{}, l.Size(), newCap)

	copy(newSli, l.items)
	l.items = newSli

	// 感觉只有 oom 的时候会出 error，但是 oom 的时候就应该 panic 了，所以暂时不做错误处理吧
	return nil
}
