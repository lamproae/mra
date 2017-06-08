package main

import (
	"fmt"
	"github.com/Workiva/go-datastructures/trie/ctrie"
)

func main() {
	t := ctrie.New(nil)
	t.Insert([]byte("foo"), "bar")
	t.Insert([]byte("1234"), "4321")
	t.Insert([]byte("liwei"), "weiliw00e")
	t.Insert([]byte("liwei1"), "weiliwe")
	t.Insert([]byte("liwei2"), "weiliwadfae")
	t.Insert([]byte("liwei3"), "weildasfkiwe")
	t.Insert([]byte("liwei4kl"), "wediliwe")

	val, ok := t.Lookup([]byte("foo"))
	fmt.Println(val, ok)

	for entry := range t.Iterator(nil) {
		fmt.Println(entry.Value, string(entry.Key))
	}
}
