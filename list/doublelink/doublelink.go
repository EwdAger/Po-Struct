// @Title  doubleline
// @Description
// @Author  EwdAger
// @Update  2021/8/3 13:14

package doublelink

import (
	"Po-Struct/list"
	"errors"
	"fmt"
	"strings"
)

func assertListImplementation() {
	var _ list.Lists = &DLink{}
}

func New(val ...interface{}) *DLink {
	l := &DLink{}
	if len(val) > 0 {
		l.Append(val...)
	}
	return l
}

type node struct {
	val  interface{}
	prev *node
	next *node
}

type DLink struct {
	head *node
	tail *node
	size int
}

func (d DLink) String() string {
	if d.head == nil {
		return "[]"
	}

	strSli := make([]string, d.Size())

	for h, i := d.head, 0; i < d.Size(); i, h = i+1, h.next {
		strSli[i] = fmt.Sprintf("%v", h.val)
	}

	res := strings.Join(strSli, ", ")
	return fmt.Sprintf("[%s]", res)
}

func (d DLink) Size() int {
	return d.size
}

func (d DLink) GetItem(index int) (res interface{}, error error) {

	head, err := d.getNodeByIndex(index)
	if err != nil {
		return nil, err
	}

	return head.val, nil

}

func (d DLink) SetItem(index int, val interface{}) error {

	head, err := d.getNodeByIndex(index)
	if err != nil {
		return err
	}

	head.val = val
	return nil

}

func (d *DLink) Append(val ...interface{}) {
	for _, v := range val {
		elem := &node{val: v, prev: d.tail}
		if d.size == 0 {
			d.head = elem
			d.tail = elem
		} else {
			d.tail.next = elem
			d.tail = elem
		}
		d.size++
	}
	d.head.prev = d.tail
	d.tail.next = d.head
}

func (d *DLink) Insert(index int, val interface{}) {
	if index >= d.Size() || -index > d.Size() {
		index = -1
	}

	head, _ := d.getNodeByIndex(index)

	var elem *node

	if index >= 0 {
		elem = &node{val: val, prev: head.prev, next: head}
		head.prev = elem
		elem.prev.next = elem
	} else {
		elem = &node{val: val, prev: head, next: head.next}
		head.next = elem
		elem.next.prev = elem
	}

	d.size++

	if index == 0 {
		d.head = elem
	}

}

func (d *DLink) Extend(b interface{}) error {
	d2, ok := b.(*DLink)
	if ok != true {
		return errors.New("DLink can't concat non-DLink type objects\n")
	}

	d.tail.next = d2.head
	d2.head.prev = d.tail
	d.head.prev = d2.head
	d2.tail.next = d.head

	d.size += d2.size

	return nil

}

func (d DLink) Sort() ([]interface{}, error) {
	panic("implement me")
}

func (d *DLink) Reverse() {

	head := d.head
	for i := 0; i < d.Size(); i++ {
		temp := head.next
		head.next, head.prev = head.prev, head.next
		head = temp
	}
	d.head, d.tail = d.tail, d.head

}

func (d DLink) AsArray() []interface{} {

	strSli := make([]interface{}, d.Size())

	for h, i := d.head, 0; i < d.Size(); i, h = i+1, h.next {
		strSli[i] = h.val
	}

	return strSli
}

func (d *DLink) Pop(index int) (interface{}, error) {
	head, err := d.getNodeByIndex(index)
	if err != nil {
		return nil, errors.New("pop index out of range")
	}

	head.prev.next = head.next
	head.next.prev = head.prev

	return head.val, err

}

func (d *DLink) Clear() {
	d.head = nil
	d.tail = nil
	d.size = 0
}

func (d *DLink) Range(f func(idx int, val interface{})) {

	head := d.head
	for i := 0; i < d.Size(); i++ {
		f(i, head.val)
		head = head.next
	}
}

func (d *DLink) getNodeByIndex(index int) (*node, error) {
	if index >= d.Size() || -index > d.Size() {
		return nil, errors.New("list index out of range")
	}

	head := d.head
	for index != 0 {
		if index < 0 {
			head = head.prev
			index++
		} else {
			head = head.next
			index--
		}
	}

	return head, nil
}
