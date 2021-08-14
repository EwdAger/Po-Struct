// @Title  ordereddict
// @Description
// @Author  EwdAger
// @Update  2021/8/3 13:11

package ordereddict

import (
	"Po-Struct/collections"
	"Po-Struct/list/doublelink"
	"fmt"
	"strings"
)

func assertOrderedDictImplementation() {
	var _ collections.OrderedDict = &OrderedDict{}
}

func New() *OrderedDict {
	return &OrderedDict{
		items: map[interface{}]interface{}{},
		order: doublelink.New(),
	}
}

type OrderedDict struct {
	items map[interface{}]interface{}
	order *doublelink.DLink
}

func (o *OrderedDict) PopItem(last bool) interface{} {
	var key interface{}

	if last == true {
		key, _ = o.order.Pop(-1)
	} else {
		key, _ = o.order.Pop(0)
	}

	val := o.items[key]
	delete(o.items, key)

	return val
}

func (o *OrderedDict) MoveToFirst(key interface{}, last bool) {
	if last == true {
		o.order.Append(key)
	} else {
		o.order.Insert(0, key)
	}
}

func (o *OrderedDict) Set(key interface{}, value interface{}) {
	if o.items == nil {
		o.items = make(map[interface{}]interface{})
	}

	if ok := o.Contains(key); !ok {
		o.items[key] = value
	}

	o.MoveToFirst(key, true)
}

func (o OrderedDict) Get(key interface{}) (interface{}, bool) {
	if v, ok := o.items[key]; !ok {
		return nil, false
	} else {
		return v, true
	}
}

func (o OrderedDict) String() string {
	orders := o.order.AsArray()
	strDict := make([]string, o.order.Size())

	for i, key := range orders {
		strDict[i] = fmt.Sprintf("%v: %v", key, o.items[key])
	}

	res := strings.Join(strDict, ", ")

	return fmt.Sprintf("{%s}", res)
}

func (o OrderedDict) Contains(elements ...interface{}) bool {
	res := true

	for _, elem := range elements {
		if _, ok := o.items[elem]; !ok {
			res = false
		}
	}

	return res

}

func (o OrderedDict) Range(f func(key interface{}, value interface{})) {
	panic("implement me")
}
