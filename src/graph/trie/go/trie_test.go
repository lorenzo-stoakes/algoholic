package algoholic

import "testing"

var testSet = map[string]interface{}{
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
			t.Fatalf("Word '%s' from input set not found in ToMap() output set %v.", word,
				actualSet)
		}
	}
}

func TestFindTrieSuffixes(t *testing.T) {
	root := NewTrieFromMap(testSet)

	ahs := root.FindSuffixes("ah")
	checkStringSlicesEqual(t, ahs, []string{"aha", "ahaha", "ahmah"})

	ahas := root.FindSuffixes("aha")
	checkStringSlicesEqual(t, ahas, []string{"aha", "ahaha"})

	foos := root.FindSuffixes("foo")
	checkStringSlicesEqual(t, foos, []string{"foo", "foobar"})

	zxys := root.FindSuffixes("zxy")
	checkStringSlicesEqual(t, zxys, []string{"zxy"})

	meshuggahAreVeryHeavys := root.FindSuffixes("Meshuggah really are VERY heavy.")
	if len(meshuggahAreVeryHeavys) > 0 {
		t.Fatalf("Invalid suffix search succeeded.")
	}
}

func BenchmarkTrieFind(b *testing.B) {
	root := NewTrieFromMap(testSet)

	var strs []string
	for str := range testSet {
		strs = append(strs, str)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		root.Find(strs[i%len(strs)])
	}
}
