package algoholic

import (
	"sort"
	"testing"
)

func checkTrieHasChildren(t *testing.T, root *Trie, search, children string) {
	trie := root.FindTrie(search)

	if trie == nil {
		t.Fatalf("Could not find '%s' in trie.", search)
	}

	expectedLen := len(children)
	if actualLen := len(trie.Children); actualLen != expectedLen {
		t.Fatalf("Expected trie %s to have %d children, got %d.", search,
			expectedLen, actualLen)
	}

	var missingRunes []rune

	for _, chr := range children {
		if _, has := trie.Children[chr]; !has {
			missingRunes = append(missingRunes, chr)
		}
	}

	if len(missingRunes) > 0 {
		t.Fatalf("Trie missing expected children %s.", string(missingRunes))
	}
}

func checkStringSlicesEqual(t *testing.T, slice1, slice2 []string) {
	if len(slice1) != len(slice2) {
		t.Fatalf("Slice lengths %d and %d differ.", len(slice1), len(slice2))
	}

	sort.StringSlice(slice1).Sort()
	sort.StringSlice(slice2).Sort()

	for i, str1 := range slice1 {
		str2 := slice2[i]
		if str1 != str2 {
			t.Fatalf("Slice mismatch at index %d, %s != %s.", i, str1, str2)
		}
	}
}
