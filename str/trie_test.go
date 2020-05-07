package str

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("1")
	t.Log(trie.IsExit("1"))
	t.Log(trie.IsExit("2"))
}
