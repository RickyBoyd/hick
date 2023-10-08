package skiplist

import (
	"testing"
)

func TestHelloName(t *testing.T) {
	bottom5 := skipListNode[int]{
		value: 5,
		next:  nil,
		down:  nil,
	}

	s := SkipList[int]{
		head: &skipListNode[int]{
			value: 1,
			next: &skipListNode[int]{
				value: 5,
				down:  &bottom5,
				next:  nil,
			},
			down: &skipListNode[int]{
				value: 1,
				next: &skipListNode[int]{
					value: 3,
					next:  &bottom5,
					down:  nil,
				},
				down: nil,
			},
		},
	}
	if !s.Find(5) {
		t.Fatal("could not find 5")
	}

	if s.Find(1000) {
		t.Fatal("found something that doesn't exist")
	}

	if !s.Find(1) {
		t.Fatal("could not find 1")
	}
}
