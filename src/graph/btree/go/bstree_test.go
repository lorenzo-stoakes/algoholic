package algoholic

import (
	"math/rand"
	"sort"
	"testing"
)

const MAX = 1e3

func TestBinarySearchTreeFulfillsSearchProperty(t *testing.T) {
	root := generateRandomTree(MAX)

	assertSearchProperty(t, root)
}

func TestBinarySearchTreeContainsNoDuplicates(t *testing.T) {
	ns := rand.Perm(MAX)
	ns = append(ns, ns...)
	// Twice, for good measure ;-)
	ns = append(ns, ns...)

	root := NewBinarySearchTree(ns)

	if length := root.Count(); length != MAX {
		t.Fatalf("Expected slice of length %d, got %d.", MAX, length)
	}

	seen := make(map[int]bool)

	root.Walk(func(n int) {
		if seen[n] {
			t.Fatalf("Duplicate entry for %d.", n)
		}
		seen[n] = true
	})
}

func TestWalkingBinarySearchTreeSortsInput(t *testing.T) {
	ns := rand.Perm(MAX)
	root := NewBinarySearchTree(ns)

	var sorted []int

	root.Walk(func(n int) {
		sorted = append(sorted, n)
	})

	if len(sorted) != len(ns) {
		t.Fatalf("Walking a tree generated from %d numbers results in %d 'steps'.",
			len(ns), len(sorted))
	}

	if !sort.IntsAreSorted(sorted) {
		t.Fatalf("Traversing the tree did not result in a sorted output.")
	}
}

func TestCanFindBinarySearchTreeValues(t *testing.T) {
	ns := rand.Perm(MAX)
	root := NewBinarySearchTree(ns)

	for _, n := range ns {
		if node := root.Find(n); node == nil {
			t.Fatalf("Couldn't find inserted value %d.", n)
		} else if node.Val != n {
			t.Fatalf("Found node with value %d when searching for %d.", node.Val, n)
		}
	}

	if node := root.Find(-1); node != nil {
		t.Fatalf("Got value %d when searching for non-existent -1 node.", node.Val)
	}
}

func TestCanInsertCorrectlyIntoBinarySearchTree(t *testing.T) {
	root := generateRandomTree(MAX)

	lessThan := rand.Perm(MAX)
	for i, n := range lessThan {
		n -= MAX
		lessThan[i] = n

		root.Insert(n)
	}

	moreThan := rand.Perm(MAX)
	for i, n := range moreThan {
		n += MAX
		moreThan[i] = n

		root.Insert(n)
	}

	if length := root.Count(); length != 3*MAX {
		t.Fatalf("Inserted %d elements to tree of %d nodes, tree now contains %d nodes.",
			2*MAX, MAX, length)
	}
}
