// @Title  slice implement
// @Description  切片实现
// @Author  EwdAger
// @Update  2021/7/6 13:14

package list

type Lists interface {
	Init(values ...interface{})
	String() string
	Size() int
	Cap() int
	GetItem(index int) (interface{}, error)
	SetItem(index int, value interface{}) error
	Append(v interface{})
	Insert() ([]interface{}, error)
	GetSlice() ([]interface{}, error)
	Sort() ([]interface{}, error)
	Reverse() ([]interface{}, error)
	AsArray() ([]interface{}, error)
	Clear() error
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
