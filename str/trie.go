package str

import "github.com/googege/godata"

type Trie struct {
	Value *godata.Trie
}

func (t *Trie) IsExit(v interface{}) bool {
	vt := v.(string)
	b, _ := t.Value.StartsWith(vt)
	return b
}
func (t *Trie) Insert(v interface{}) {
	t.Value.Insert(v.(string))
}
// return a new Trie.
func NewTrie() Str {
	return &Trie{Value: godata.NewTrie()}
}
