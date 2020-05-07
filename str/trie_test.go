package str

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := NewTrie()
	trie.Insert("1")
	fmt.Println(trie.IsExit("1"))
	fmt.Println(trie.IsExit("2"))
}
