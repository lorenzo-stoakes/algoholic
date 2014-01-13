package algoholic

import "testing"

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
