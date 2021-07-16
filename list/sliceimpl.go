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

func (l *sliceImpl) SetItem(index int, val interface{}) error {
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
		l.items[l.Size()-index] = val
	} else {
		l.items[index] = val
	}
	return nil

}

func (l *sliceImpl) Append(val ...interface{}) {
	if need := needGrow(l, len(val)); need == true {
		err := grow(l, len(val))
		if err != nil {
			return
		}
	}

	l.items = append(l.items, val...)

}

// Insert 支持倒叙索引插入；支持越界索引插入
func (l *sliceImpl) Insert(index int, val interface{}) {
	if need := needGrow(l, 1); need == true {
		err := grow(l, 1)
		if err != nil {
			return
		}
	}

	if index >= l.Size() {
		l.Append(val)
	} else if index < 0 && -index >= l.Size() {
		insertLocation(l, 0, val)
	} else if index < 0 {
		insertLocation(l, l.Size()+index+1, val)
	} else {
		insertLocation(l, index, val)
	}

}

// GetSlice 遵循 Slice 子切片规则，左闭右开；支持倒叙索引；支持越界索引
func (l sliceImpl) GetSlice(left int, right int) interface{} {

	if left > right && right >= 0 {
		return &sliceImpl{}
	}

	// 处理 left 和 right 越界的情况
	if left >= l.Size() {
		return &sliceImpl{}
	} else if left < 0 && -left >= l.Size() {
		left = 0
	} else if left < 0 {
		left = l.Size() + left + 1
	}

	if right > l.Size() || right == 0 {
		right = l.Size()
	} else if right < 0 && -right >= l.Size() {
		return &sliceImpl{}
	} else if right < 0 {
		right = l.Size() + right
	}
	res := &sliceImpl{l.items[left:right]}

	return res

}

func (l *sliceImpl) Extend(b interface{}) error {

	sli2, ok := b.(*sliceImpl)
	if ok != true {
		return errors.New("List can't concat non-list type objects\n")
	}

	items := sli2.items
	l.Append(items...)

	return nil
}

func (l *sliceImpl) Reverse() {
	left, right := 0, l.Size()-1

	for left < right {
		leftItem, _ := l.GetItem(left)
		rightItem, _ := l.GetItem(right)
		l.SetItem(left, rightItem)
		l.SetItem(right, leftItem)

		left += 1
		right -= 1
	}

}

// AsArray 没办法直接转换数组，目前只能返回 slice
func (l sliceImpl) AsArray() []interface{} {
	return l.items

}

func (l *sliceImpl) Pop(index int) (interface{}, error) {

	popItem, err := l.GetItem(index)
	if err != nil {
		return nil, errors.New("pop index out of range")
	}

	reverseIndex := false

	if index < 0 {
		reverseIndex = true
		index = -index
	}

	if reverseIndex == false {
		l.items = append(l.items[:index], l.items[index:]...)
	} else {
		l.items = append(l.items[:l.Size()-index-1], l.items[l.Size()-index-1:]...)
	}

	return popItem, nil

}

func (l *sliceImpl) Clear() {
	var newSli []interface{}
	l.items = newSli

}

// Sort 没有泛型排序是真的不好写
func (l sliceImpl) Sort() ([]interface{}, error) {
	panic("implement me")
}

// 判断是否需要扩容
// nums 为本次扩容需要添加的元素个数
func needGrow(l *sliceImpl, nums int) bool {
	return len(l.items)+nums > cap(l.items)
}

// 每次扩容扩大 1/8 再加上 3 或 6 的余量
// nums 为本次扩容需要添加的元素个数
func grow(l *sliceImpl, nums int) error {
	nowCap := cap(l.items)

	// 扩容的余量，防止容量很小时频繁扩容, 至于为啥是 3 和 6，这里直接抄了 Python 的源码
	// new_allocated = (size_t)newsize + (newsize >> 3) + (newsize < 9 ? 3 : 6);
	var over int
	if nowCap < 9 {
		over = 3
	} else {
		over = 6
	}

	newCap := nowCap + (nowCap >> 3) + over + nums
	newSli := make([]interface{}, l.Size(), newCap)

	copy(newSli, l.items)
	l.items = newSli

	// 感觉只有 oom 的时候会出 error，但是 oom 的时候就应该 panic 了，所以暂时不做错误处理吧
	return nil
}

// 具体插入到某一位置的逻辑，index 为正数且小于 slice 的长度
func insertLocation(l *sliceImpl, index int, val interface{}) {
	l.Append(struct{}{})

	var tmp interface{}
	nowItem, _ := l.GetItem(index)
	for i := index + 1; i < l.Size(); i++ {
		tmp, _ = l.GetItem(i)
		l.SetItem(i, nowItem)

		nowItem = tmp
	}
	l.SetItem(index, val)

}
