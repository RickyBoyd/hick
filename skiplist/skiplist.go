package skiplist

import (
	"cmp"
	"errors"
)

type List[T cmp.Ordered] interface {
	Find(t T) bool
	Insert(t T) error
	Delete(t T) error
}

type skipListNode[T cmp.Ordered] struct {
	value T
	next  *skipListNode[T]
	down  *skipListNode[T]
}

type SkipList[T cmp.Ordered] struct {
	head *skipListNode[T]
}

func (self *SkipList[T]) Insert(t T) error {
	return errors.New("boo")
}

func (self *SkipList[T]) Find(t T) bool {
	node := self.findNode(t)
	return node != nil
}

func (self *SkipList[T]) findNode(t T) *skipListNode[T] {
	currentNode := self.head
	for currentNode != nil {
		if currentNode.value == t {
			return currentNode
		}

		if currentNode.next == nil {
			// then we're at the end, try to go down
			currentNode = currentNode.down
			continue
		}
		if !cmp.Less(currentNode.next.value, t) {
			// then go there
			currentNode = currentNode.next
			continue
		}

		currentNode = currentNode.down
	}

	return currentNode
}

func (self *SkipList[T]) Delete(t T) error {
	return errors.New("boo")
}
