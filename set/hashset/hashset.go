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
	var _ set.Set = &hashSet{}
}

func New(val ...interface{}) *hashSet {
	h := &hashSet{}
	h.items = make(map[interface{}]struct{}, len(val))

	for _, v := range val {
		h.items[v] = struct{}{}
	}

	return h
}

type hashSet struct {
	items map[interface{}]struct{}
}

func (h *hashSet) Contains(val ...interface{}) bool {
	for _, k := range val {
		if _, ok := h.items[k]; !ok {
			return false
		}
	}
	return true
}

func (h *hashSet) Values() (res []interface{}) {
	for k, _ := range h.items {
		res = append(res, k)
	}

	return res
}

func (h hashSet) Len() int {
	return len(h.items)
}

func (h hashSet) String() string {
	strSli := make([]string, h.Len())
	i := 0

	for k, _ := range h.items {
		strSli[i] = fmt.Sprintf("%v", k)
		i += 1
	}

	res := strings.Join(strSli, ", ")
	return fmt.Sprintf("{%s}", res)
}

func (h *hashSet) Copy() *hashSet {
	newMap := make(map[interface{}]struct{}, h.Len())

	for k, _ := range h.items {
		newMap[k] = struct{}{}
	}

	res := &hashSet{}
	res.items = newMap

	return res
}

func (h *hashSet) IsSubSet(b *hashSet) bool {
	return IsSubSet(h, b)
}

func (h *hashSet) IsSuperSet(b *hashSet) bool {
	return IsSuperSet(h, b)
}

func (h *hashSet) Equal(b *hashSet) bool {
	return Equal(h, b)
}

func (h *hashSet) Union(b *hashSet) *hashSet {
	return Union(h, b)
}

func (h *hashSet) Difference(b *hashSet) *hashSet {
	return Difference(h, b)
}

func (h *hashSet) Intersection(b *hashSet) *hashSet {
	return Intersection(h, b)
}

func (h *hashSet) Range(fun func(key interface{})) {
	for key, _ := range h.items {
		fun(key)
	}
}

// Equal a b 长度和元素全部相等
func Equal(a *hashSet, b *hashSet) bool {
	if a.Len() != b.Len() || !IsSubSet(a, b) {
		return false
	}

	return true
}

// IsSubSet a 是否为 b 子集，即 a 中元素是否都在 b 中， b >= a
func IsSubSet(a *hashSet, b *hashSet) bool {
	for k, _ := range a.items {
		if _, ok := b.items[k]; !ok {
			return false
		}
	}
	return true
}

// IsSuperSet a 是否为 b 超集，即 b 中元素是否都在 a 中， b <= a
func IsSuperSet(a *hashSet, b *hashSet) bool {
	return IsSubSet(b, a)
}

// Union 并集，返回 a 与 b 并集
func Union(a *hashSet, b *hashSet) *hashSet {
	newSet := a.Copy()

	for k, _ := range b.items {
		newSet.items[k] = struct{}{}
	}
	return newSet
}

// Difference 差集，返回 a 与 b 的差集
func Difference(a *hashSet, b *hashSet) *hashSet {
	newSet := a.Copy()

	for k, _ := range b.items {
		if _, ok := newSet.items[k]; ok {
			delete(newSet.items, k)
		}
	}
	return newSet

}

// Intersection 交集，返回 a 与 b 的交集
func Intersection(a *hashSet, b *hashSet) *hashSet {
	newSet := a.Copy()

	for k, _ := range b.items {
		if _, ok := newSet.items[k]; !ok {
			delete(newSet.items, k)
		}
	}

	for k, _ := range newSet.items {
		if _, ok := b.items[k]; !ok {
			delete(newSet.items, k)
		}
	}

	return newSet

}
