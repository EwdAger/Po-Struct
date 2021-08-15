// @Title  set
// @Description
// @Author  EwdAger
// @Update  2021/7/17 13:29

package hashset

import (
	"Po-Struct/set"
	"fmt"
	"strings"
)

// 检查是否 implemented
func assertListImplementation() {
	var _ set.Set = &HashSet{}
}

// New 初始化一个新的 Set
func New(val ...interface{}) *HashSet {
	h := &HashSet{}
	h.items = make(map[interface{}]struct{}, len(val))

	for _, v := range val {
		h.items[v] = struct{}{}
	}

	return h
}

type HashSet struct {
	items map[interface{}]struct{}
}

// Contains 当前 Set 是否包含 value
func (h *HashSet) Contains(val ...interface{}) bool {
	for _, k := range val {
		if _, ok := h.items[k]; !ok {
			return false
		}
	}
	return true
}

// Values 将当前 Set 中的所有元素以切片的形式返回
func (h *HashSet) Values() (res []interface{}) {
	for k := range h.items {
		res = append(res, k)
	}

	return res
}

// Size 返回大小
func (h HashSet) Size() int {
	return len(h.items)
}

// String 序列化
func (h HashSet) String() string {
	strSli := make([]string, h.Size())
	i := 0

	for k := range h.items {
		strSli[i] = fmt.Sprintf("%v", k)
		i++
	}

	res := strings.Join(strSli, ", ")
	return fmt.Sprintf("{%s}", res)
}

// Copy 返回一个副本 Set
func (h *HashSet) Copy() *HashSet {
	newMap := make(map[interface{}]struct{}, h.Size())

	for k := range h.items {
		newMap[k] = struct{}{}
	}

	res := &HashSet{}
	res.items = newMap

	return res
}

// IsSubSet a 是否为 b 子集，即 a 中元素是否都在 b 中， b >= a
func (h *HashSet) IsSubSet(b *HashSet) bool {
	return IsSubSet(h, b)
}

// IsSuperSet a 是否为 b 超集，即 b 中元素是否都在 a 中， b <= a
func (h *HashSet) IsSuperSet(b *HashSet) bool {
	return IsSuperSet(h, b)
}

// Equal a b 长度和元素全部相等
func (h *HashSet) Equal(b *HashSet) bool {
	return Equal(h, b)
}

// Union 并集，返回 a 与 b 并集
func (h *HashSet) Union(b *HashSet) *HashSet {
	return Union(h, b)
}

// Difference 差集，返回 a 与 b 的差集
func (h *HashSet) Difference(b *HashSet) *HashSet {
	return Difference(h, b)
}

// Intersection 交集，返回 a 与 b 的交集
func (h *HashSet) Intersection(b *HashSet) *HashSet {
	return Intersection(h, b)
}

// Range 使用传入的 func 遍历当前 Set
func (h *HashSet) Range(fun func(key interface{})) {
	for key := range h.items {
		fun(key)
	}
}

// Equal a b 长度和元素全部相等
func Equal(a *HashSet, b *HashSet) bool {
	if a.Size() != b.Size() || !IsSubSet(a, b) {
		return false
	}

	return true
}

// IsSubSet a 是否为 b 子集，即 a 中元素是否都在 b 中， b >= a
func IsSubSet(a *HashSet, b *HashSet) bool {
	for k := range a.items {
		if _, ok := b.items[k]; !ok {
			return false
		}
	}
	return true
}

// IsSuperSet a 是否为 b 超集，即 b 中元素是否都在 a 中， b <= a
func IsSuperSet(a *HashSet, b *HashSet) bool {
	return IsSubSet(b, a)
}

// Union 并集，返回 a 与 b 并集
func Union(a *HashSet, b *HashSet) *HashSet {
	newSet := a.Copy()

	for k := range b.items {
		newSet.items[k] = struct{}{}
	}
	return newSet
}

// Difference 差集，返回 a 与 b 的差集
func Difference(a *HashSet, b *HashSet) *HashSet {
	newSet := a.Copy()

	for k := range b.items {
		if _, ok := newSet.items[k]; ok {
			delete(newSet.items, k)
		}
	}
	return newSet

}

// Intersection 交集，返回 a 与 b 的交集
func Intersection(a *HashSet, b *HashSet) *HashSet {
	newSet := a.Copy()

	for k := range b.items {
		if _, ok := newSet.items[k]; !ok {
			delete(newSet.items, k)
		}
	}

	for k := range newSet.items {
		if _, ok := b.items[k]; !ok {
			delete(newSet.items, k)
		}
	}

	return newSet

}
