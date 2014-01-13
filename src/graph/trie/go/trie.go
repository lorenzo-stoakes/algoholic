package algoholic

type Trie struct {
	Char     rune
	Children map[rune]*Trie
	Terminal bool
	Value    interface{}
}

const RootTrie = rune(0)

func NewTrie(chr rune) *Trie {
	return &Trie{chr, make(map[rune]*Trie), false, nil}
}

func NewRootTrie() *Trie {
	return NewTrie(RootTrie)
}

func NewTrieFromMap(strMap map[string]interface{}) *Trie {
	ret := NewRootTrie()

	for str, val := range strMap {
		ret.Insert(str, val)
	}

	return ret
}

func NewTrieFromStrings(strs []string) *Trie {
	ret := NewRootTrie()

	for _, str := range strs {
		ret.Insert(str, nil)
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

func (trie *Trie) Find(str string) (val interface{}, has bool) {
	trie = trie.FindTrie(str)

	if trie == nil || !trie.Terminal {
		// Not found.
		return
	}

	has = true
	val = trie.Value

	return
}

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
		next := NewTrie(chr)
		trie.Children[chr] = next
		trie = next
	}
	trie.Terminal = true
	trie.Value = val
}

func (trie *Trie) doWalk(trieMap map[string]interface{}, prev []rune) {
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

func (trie *Trie) Walk() map[string]interface{} {
	ret := make(map[string]interface{})
	trie.doWalk(ret, nil)
	return ret
}
