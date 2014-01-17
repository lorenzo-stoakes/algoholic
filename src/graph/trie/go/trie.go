package algoholic

// A trie is a data structure which stores words in a tree by letter, when the letters form a
// valid string, the node is marked terminal and a value can be stored.

type Trie struct {
	Char     rune
	Parent   *Trie
	Children map[rune]*Trie
	Terminal bool
	Value    interface{}
}

// A value indicating that the node is the root node, thus it doesn't denote a character
// itself.
const RootTrie = rune(0)

// Create a new trie with the specified character and parent.
func NewTrie(parent *Trie, chr rune) *Trie {
	return &Trie{chr, parent, make(map[rune]*Trie), false, nil}
}

func NewRootTrie() *Trie {
	return NewTrie(nil, RootTrieChar)
}

// Create a new trie with strings mapped to specified values.
func NewTrieFromMap(strMap map[string]interface{}) *Trie {
	ret := NewRootTrie()

	for str, val := range strMap {
		ret.Insert(str, val)
	}

	return ret
}

// Create a new trie with strings whose values we don't care about.
func NewTrieFromStrings(strs []string) *Trie {
	ret := NewRootTrie()

	for _, str := range strs {
		ret.Insert(str, nil)
	}

	return ret
}

// Find the specified string and return its trie node.
// O(m) worst-case where m is the length of the string searched for.
// Note this returns non-terminal nodes.
func (trie *Trie) FindTrie(str string) *Trie {
	if len(str) == 0 {
		return trie
	}

	if next := trie.Children[rune(str[0])]; next != nil {
		return next.FindTrie(str[1:])
	}

	return nil
}

// Find the specified string and return its value.
// O(m) worst-case where m is the length of the string searched for.
func (trie *Trie) Find(str string) (val interface{}, has bool) {
	ret := trie.FindTrie(str)

	if ret == nil || !ret.Terminal {
		// Not found.
		return
	}

	has = true
	val = ret.Value

	return
}

// Find all valid strings that consist of suffixes of the input prefix.
// O(m) worst-case where m is the length of the longest returned string.
func (trie *Trie) FindSuffixes(prefix string) []string {
	trie = trie.FindTrie(prefix)

	if trie == nil {
		return nil
	}

	var ret []string

	for suffix := range trie.Walk() {
		ret = append(ret, prefix+suffix[1:])
	}

	return ret
}

// Insert string, value pair into the specified trie.
// O(m) worst-case where m is the length of the inserted string.
func (trie *Trie) Insert(str string, val interface{}) {
	var (
		i   int
		chr rune
	)

	// Search through existing nodes.
	for i, chr = range str {
		if next, has := trie.Children[chr]; has {
			trie = next
		} else {
			break
		}
	}

	// Insert nodes as necessary.
	for _, chr = range str[i:] {
		next := NewTrie(trie, chr)
		trie.Children[chr] = next
		trie = next
	}
	trie.Terminal = true
	trie.Value = val
}

// Recursively walk through all children of the input trie, adding string, value pairs to
// trieMap as the walk is performed. Trie is traversed in pre-order.
// O(n) where n is the number of nodes in the input trie.
func (trie *Trie) doWalk(trieMap map[string]interface{}, prev []rune) {
	// TODO: Use something other than a hash for map to allow alphabetical output ordering.

	if trie.Char != RootTrie {
		prev = append(prev, trie.Char)
	}

	if trie.Terminal {
		str := string(prev)
		trieMap[str] = trie.Value
	}

	for _, child := range trie.Children {
		child.doWalk(trieMap, prev)
	}
}

// Recursively walk through all children of the input trie, returning a map of string, value
// pairs.
// O(n) where n is the number of nodes in the input trie.
func (trie *Trie) Walk() map[string]interface{} {
	ret := make(map[string]interface{})
	trie.doWalk(ret, nil)
	return ret
}
