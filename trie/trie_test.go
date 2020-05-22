package trie

import "testing"

func TestTrie_Find(t *testing.T) {
	tests := []struct {
		str    string
		find   string
		result bool
	}{
		{
			"abcdefg",
			"abcdefg",
			true,
		},
		{
			"zxcvbn",
			"zxcvbn",
			true,
		},
		{
			"zxcvbn",
			"zxcvb",
			false,
		},
		{
			"zxcvbn",
			"zxcvbnb",
			false,
		},
		{
			"",
			"zxcvbnb",
			false,
		},
		{
			"z",
			"zxcvbnb",
			false,
		},
		{
			"a",
			"zxcvbnb",
			false,
		},
		{
			"陈太汉好帅",
			"陈太汉好帅",
			true,
		},
		{
			"陈太汉好帅",
			"陈太汉好帅a",
			false,
		},
	}

	for index, test := range tests {
		trie := NewTrie()
		trie.Add(test.str)
		result := trie.Find(test.find)
		if result != test.result {
			t.Fatal("with error ", index)
		}
	}
}

func TestTrie_StartWith(t *testing.T) {
	tests := []struct {
		str    string
		with   string
		result bool
	}{
		{
			"abcdefg",
			"abcdefg",
			true,
		},
		{
			"zxcvbn",
			"zxcvbn",
			true,
		},
		{
			"zxcvbn",
			"zxcvb",
			true,
		},
		{
			"zxcvbn",
			"zx",
			true,
		},
		{
			"zxcvbn",
			"z",
			true,
		},
		{
			"zxcvbn",
			"zxcvbnb",
			false,
		},
		{
			"",
			"zxcvbnb",
			false,
		},
		{
			"z",
			"zxcvbnb",
			false,
		},
		{
			"a",
			"zxcvbnb",
			false,
		},
		{
			"陈太汉好帅",
			"陈太汉好帅",
			true,
		},
		{
			"陈太汉好帅",
			"陈太汉好",
			true,
		},
	}

	for index, test := range tests {
		trie := NewTrie()
		trie.Add(test.str)
		result := trie.StartWith(test.with)
		if result != test.result {
			t.Fatal("with error ", index)
		}
	}
}

func TestTrie_Remove(t *testing.T) {
	tests := []struct {
		str    string
		find   string
		result bool
	}{
		{
			"abcdefg",
			"abcdefg",
			true,
		},
		{
			"zxcvbn",
			"zxcvbn",
			true,
		},
		{
			"zxcvbn",
			"zxcvb",
			false,
		},
		{
			"azxcvbn",
			"azxcvbnb",
			false,
		},
		{
			"",
			"azxcvbnb",
			false,
		},
		{
			"z",
			"azxcvbnb",
			false,
		},
		{
			"a",
			"azxcvbnb",
			false,
		},
		{
			"陈太汉好帅",
			"陈太汉好帅",
			true,
		},
		{
			"陈太汉好帅",
			"陈太汉好帅a",
			false,
		},
	}
	trie := NewTrie()
	for index, test := range tests {
		trie.Add(test.str)
		result := trie.Find(test.find)
		if result != test.result {
			t.Fatal("Get error ", index)
		}
		trie.Remove(test.find)
		result = trie.Find(test.find)
		if result {
			t.Fatal("remove error ", index)
		}
	}
}

func TestTrie_GetTopN(t *testing.T) {
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
	trie := NewTrie()
	for _, test := range tests {
		for i := 0; i < test.count; i++ {
			trie.Add(test.str)
		}
	}
	for n := 0; n < 10; n++ {
		t.Log(trie.GetTopN(n))
	}

}
