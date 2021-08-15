// @Title  set
// @Description
// @Author  EwdAger
// @Update  2021/8/3 10:27

package set

type Set interface {
	Size() int
	String() string
	Contains(val ...interface{}) bool
	Values() []interface{}
}
