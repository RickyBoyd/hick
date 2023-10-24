package skiplist

import (
	"cmp"
	"errors"
	"math/rand"
)

type List[T cmp.Ordered] interface {
	Find(t T) bool
	Insert(t T) error
	Delete(t T) error
}

type skipListNode[T cmp.Ordered] struct {
	value T
	next  []*skipListNode[T]
}

func (self *skipListNode[T]) Height() int {
	return len(self.next)
}

func setNext[T cmp.Ordered](setOn *skipListNode[T], setTo *skipListNode[T], height int) {
	tmp := setOn.next[height]
	setOn.next[height] = setTo
	setTo.next[height] = tmp
}

type SkipList[T cmp.Ordered] struct {
	head       []*skipListNode[T]
	headHeight int
}

func (self *SkipList[T]) Insert(t T) error {
	newHeight := generateHeight()
	lastSeenByHeight := make([]*skipListNode[T], newHeight)
	node, height := self.searchHead(t)
	lastSeenByHeight[height] = node

	return errors.New("boo")
}

func (self *SkipList[T]) Find(t T) bool {
	node := self.findNode(t)
	return node != nil
}

func (self *SkipList[T]) searchHead(t T) (*skipListNode[T], int) {
	for currentHeight := len(self.head) - 1; currentHeight >= 0; currentHeight-- {
		headNode := self.head[currentHeight]
		if headNode.value <= t {
			return headNode, currentHeight
		}
	}
	return nil, -1
}

func (self *SkipList[T]) findNode(t T) *skipListNode[T] {
	currentNode, currentHeight := self.searchHead(t)
	if currentNode == nil {
		return nil
	}
	if currentNode.value == t {
		return currentNode
	}

	for currentNode != nil {
		for ; currentHeight >= 0; currentHeight-- {
			nextNode := currentNode.next[currentHeight]
			if nextNode == nil {
				continue
			}
			if nextNode.value == t {
				return nextNode
			}
			if t < nextNode.value {
				currentNode = nextNode
				break
			}
		}
		if currentHeight < 0 {
			return nil
		}
	}

	return currentNode
}

func (self *SkipList[T]) Delete(t T) error {
	return errors.New("boo")
}

func generateHeight() int {
	height := 1
	for rand.Float32() < 0.5 {
		height += 1
	}
	return height
}
