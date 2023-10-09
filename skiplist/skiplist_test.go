package skiplist

import (
	"testing"
)

func TestHelloName(t *testing.T) {
	block5 := skipListNode[int]{
		value: 5,
		next: []*skipListNode[int]{
			nil,
			nil,
		},
	}

	block3 := skipListNode[int]{
		value: 3,
		next: []*skipListNode[int]{
			&block5,
		},
	}

	block1 := skipListNode[int]{
		value: 1,
		next: []*skipListNode[int]{
			&block3,
			&block5,
		},
	}

	/*
		|*| -> |1| --------> |5| -> |nil|
		|*| -> |1| -> |3| -> |5| -> |nil|
	*/

	s := SkipList[int]{
		head: []*skipListNode[int]{
			&block1,
			&block1,
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
