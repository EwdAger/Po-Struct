// @Title  collections
// @Description
// @Author  EwdAger
// @Update  2021/8/3 12:47

package collections

type OrderedDict interface {
	PopItem(last bool) interface{}
	MoveToFirst(key interface{}, last bool)
	Set(key interface{}, value interface{})
	Get(key interface{}) (interface{}, bool)
	String() string
	Contains(elements ...interface{}) bool
	Range(func(key interface{}, value interface{}))
}
