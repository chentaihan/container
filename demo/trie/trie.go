package main

import (
	"fmt"
	"github.com/chentaihan/container/trie"
)

func main() {
	tests := []struct {
		str   string
		count int
	}{
		{
			"1-陈太汉好帅aaaa",
			50,
		},
		{
			"2-123456789",
			40,
		},
		{
			"3-zxcvbnbv",
			30,
		},
		{
			"4-zx12112zzx",
			20,
		},
		{
			"5-2附加费1zzx",
			10,
		},
		{
			"6-dsd?????",
			6,
		},
		{
			"7-dsd的a啊？",
			1,
		},
	}
	trie := trie.NewTrie()
	for _, test := range tests {
		for i := 0; i < test.count; i++ {
			trie.Add(test.str)
		}
	}
	for n := 0; n < 10; n++ {
		fmt.Println(trie.GetTopN(n))
	}
	for _, test := range tests {
		if !trie.Find(test.str) {
			fmt.Println("find error")
		}
		if !trie.StartWith(test.str) {
			fmt.Println("StartWith error")
		}
		trie.Remove(test.str)
	}
}
