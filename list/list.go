// @Title  slice implement
// @Description  切片实现
// API 及实现参考 https://github.com/python/cpython/blob/main/Objects/listobject.c
// @Author  EwdAger
// @Update  2021/7/6 13:14

package list

type Lists interface {
	Init(val ...interface{})
	String() string
	Size() int
	Cap() int
	GetItem(index int) (interface{}, error)
	SetItem(index int, val interface{}) error
	Append(val ...interface{})
	Insert(index int, val interface{})
	GetSlice(left int, right int) interface{}
	Extend(b interface{}) error
	Sort() ([]interface{}, error)
	Reverse()
	AsArray() []interface{}
	Pop(index int) (interface{}, error)
	Clear()
	Range(func(idx int, val interface{}))
}

func List(types string, items ...interface{}) (res Lists) {
	if types == "" || types == "slice" {
		res = &sliceImpl{}
	} else {
		panic("// TODO")
	}
	res.Init(items...)
	return res
}

func Equal(a Lists, b Lists) bool {

	if a.Size() != b.Size() {
		return false
	}

	for i := 0; i < a.Size(); i++ {
		left, _ := a.GetItem(i)
		right, _ := b.GetItem(i)
		if left != right {
			return false
		}
	}

	return true
}

func Is(a *Lists, b *Lists) bool {
	return a == b
}
