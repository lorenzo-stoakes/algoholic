package algoholic

// Trie of characters mapped to int.
type Trie struct {
	Char     rune
	Children map[rune]*Trie
	Terminal bool
	Value    int
}

const RootTrie = rune(0)

func NewTrie(chr rune) *Trie {
	return &Trie{chr, make(map[rune]*Trie), false, 0}
}

func NewTrieFromMap(strMap map[string]int) *Trie {
	ret := NewTrie(RootTrie)

	for str, val := range strMap {
		ret.Insert(str, val)
	}

	return ret
}

func (trie *Trie) FindTrie(str string) *Trie {
	if len(str) == 0 {
		return trie
	}

	if next := trie.Children[rune(str[0])]; next != nil {
		return next.FindTrie(str[1:])
	}

	return nil
}

func (trie *Trie) Find(str string) (val int, has bool) {
	trie = trie.FindTrie(str)

	if trie == nil || !trie.Terminal {
		// Not found.
		return
	}

	has = true
	val = trie.Value

	return
}

func (trie *Trie) Insert(str string, val int) {
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
		next := NewTrie(chr)
		trie.Children[chr] = next
		trie = next
	}
	trie.Terminal = true
	trie.Value = val
}
