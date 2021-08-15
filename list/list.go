// @Title  slice implement
// @Description  切片实现
// API 及实现参考 https://github.com/python/cpython/blob/main/Objects/listobject.c
// @Author  EwdAger
// @Update  2021/7/6 13:14

package list

type Lists interface {
	String() string
	Size() int
	GetItem(index int) (interface{}, error)
	SetItem(index int, val interface{}) error
	Append(val ...interface{})
	Insert(index int, val interface{})
	Extend(b interface{}) error
	Sort() ([]interface{}, error)
	Reverse()
	AsArray() []interface{}
	Pop(index int) (interface{}, error)
	Clear()
	Range(func(idx int, val interface{}))
}

// Equal 判断两个 Lists 是否相等
func Equal(a interface{}, b interface{}) bool {

	if a.(Lists).Size() != b.(Lists).Size() {
		return false
	}

	for i := 0; i < a.(Lists).Size(); i++ {
		left, _ := a.(Lists).GetItem(i)
		right, _ := b.(Lists).GetItem(i)
		if left != right {
			return false
		}
	}

	return true
}

// Is 判断两个 Lists 是否为同一个
func Is(a interface{}, b interface{}) bool {
	return a == b
}
