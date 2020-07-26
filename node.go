package skiplist

import (
	"sync/atomic"
	"unsafe"
)

type Node struct {
	keyHash  uintptr
	previous unsafe.Pointer
	next     unsafe.Pointer
	key      interface{}
	value    unsafe.Pointer
	deleted  uintptr
}

func (node *Node) Value() (value interface{}) {
	return *(*interface{})(atomic.LoadPointer(&node.value))
}

func (e *Node) Next() *Node {
	return (*Node)(atomic.LoadPointer(&e.next))
}

func (e *Node) Previous() *Node {
	return (*Node)(atomic.LoadPointer(&e.previous))
}

func (e *Node) setValue(value unsafe.Pointer) {
	atomic.StorePointer(&e.value, value)
}

func (e *Node) casValue(from interface{}, to unsafe.Pointer) bool {
	old := atomic.LoadPointer(&e.value)
	if *(*interface{})(old) != from {
		return false
	}
	return atomic.CompareAndSwapPointer(&e.value, old, to)
}

type Index struct {
	node  unsafe.Pointer
	down  unsafe.Pointer
	right unsafe.Pointer
}
