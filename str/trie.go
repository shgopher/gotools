package str

import "github.com/googege/godata"

type Trie struct {
	Value *godata.Trie
}

func (t *Trie) IsExit(v string) bool {
	b, _ := t.Value.StartsWith(v)
	return b
}
func (t *Trie) Insert(v string) {
	t.Value.Insert(v)
}
// return a new Trie.
func NewTrie() *Trie {
	return &Trie{Value: godata.NewTrie()}
}
