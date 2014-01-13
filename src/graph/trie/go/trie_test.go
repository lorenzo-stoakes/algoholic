package algoholic

import (
	"testing"
)

var testSet = map[string]int{
	"aha":    15,
	"ahaha":  20,
	"ahmah":  30,
	"zxy":    12,
	"foo":    50,
	"foobar": 100,
}

/*
Results in:-

root -> a -> h -> a[15] -> h      -> a[20]
               -> m     -> a[30]
     -> f -> o -> o[50]
                        -> b -> a -> r[100]
     -> z -> x -> y[12]
*/

func TestNewTrieFromMap(t *testing.T) {
	root := NewTrieFromMap(testSet)

	checkTrieHasChildren(t, root, "", "afz")
	checkTrieHasChildren(t, root, "ah", "am")
	checkTrieHasChildren(t, root, "foo", "b")

	for word, expected := range testSet {
		if actual, has := root.Find(word); has {
			if actual != expected {
				t.Fatalf("Expected trie entry '%s' to have value %d, actually %d.",
					word, expected, actual)
			}
		} else {
			t.Fatalf("Couldn't find inserted word '%s' in trie.", word)
		}
	}
}

func TestTrieWalk(t *testing.T) {
	root := NewTrieFromMap(testSet)

	actualSet := root.Walk()

	if len(actualSet) != len(testSet) {
		t.Fatal("Mismatched set lengths.")
	}

	for word, expected := range testSet {
		if actual, has := actualSet[word]; has {
			if actual != expected {
				t.Fatalf("Word '%s' was expected to have value %d, but has %d.",
					word, expected, actual)
			}
		} else {
			t.Fatalf("Word '%s' from input set not found in Walk() output set.", word)
		}
	}
}
