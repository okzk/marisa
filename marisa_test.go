package marisa

import (
	"fmt"
)

func printKey(id uint64, key string) error {
	fmt.Println(id, key)
	return nil
}

func panicIfErrorOccurred(err error) {
	if err != nil {
		panic(err)
	}
}

func ExampleSaveAndLoad() {
	keyset := NewKeyset()
	defer keyset.Dispose()

	keyset.PushBack("ho")
	keyset.PushBack("hoge")
	keyset.PushBack("hogehoge")
	keyset.PushBack("mogemoge")

	trie, err := keyset.Build()
	panicIfErrorOccurred(err)
	defer trie.Dispose()

	err = trie.Save("./tmp.dic")
	panicIfErrorOccurred(err)

	trieLoaded, err := LoadTrieFromFile("./tmp.dic")
	panicIfErrorOccurred(err)
	defer trieLoaded.Dispose()
	fmt.Println(trieLoaded.ReverseLookup(0))

	trieMmap, err := MmapTrie("./tmp.dic")
	panicIfErrorOccurred(err)
	defer trieMmap.Dispose()
	fmt.Println(trieMmap.ReverseLookup(0))

	// Output:
	// ho <nil>
	// ho <nil>
}

func ExampleBuildWithOption() {
	keyset := NewKeyset()
	defer keyset.Dispose()

	keyset.PushBack("ho")
	keyset.PushBack("hoge")
	keyset.PushBackWithWeight("hogehoge", 1.0)
	keyset.PushBackWithWeight("mogemoge", 1.0)

	option := NewBuildOption()
	option.NumTries = 1
	option.NodeOrder = MARISA_LABEL_ORDER
	option.CacheLevel = MARISA_SMALL_CACHE
	option.TailMode = MARISA_BINARY_TAIL

	trie, err := keyset.Build()
	panicIfErrorOccurred(err)
	defer trie.Dispose()

	fmt.Println(trie.Size())

	// Output:
	// 4 <nil>
}

func ExampleLookup() {
	keyset := NewKeyset()
	defer keyset.Dispose()

	keyset.PushBack("ho")
	keyset.PushBack("hoge")
	keyset.PushBack("hogehoge")
	keyset.PushBack("mogemoge")

	trie, err := keyset.Build()
	panicIfErrorOccurred(err)
	defer trie.Dispose()

	fmt.Println(trie.Lookup("hoge"))
	fmt.Println(trie.Lookup("foo"))

	// Output:
	// 2 <nil>
	// 0 not found
}

func ExampleReverseLookup() {
	keyset := NewKeyset()
	defer keyset.Dispose()

	keyset.PushBack("ho")
	keyset.PushBack("hoge")
	keyset.PushBack("hogehoge")
	keyset.PushBack("mogemoge")

	trie, err := keyset.Build()
	panicIfErrorOccurred(err)
	defer trie.Dispose()

	fmt.Println(trie.ReverseLookup(0))

	// Output:
	// ho <nil>
}

func ExamplePredictiveSearch() {
	keyset := NewKeyset()
	defer keyset.Dispose()

	keyset.PushBack("ho")
	keyset.PushBack("hoge")
	keyset.PushBack("hogehoge")
	keyset.PushBack("mogemoge")

	trie, err := keyset.Build()
	panicIfErrorOccurred(err)
	defer trie.Dispose()

	err = trie.PredictiveSearch("h", printKey)
	panicIfErrorOccurred(err)

	// Output:
	// 0 ho
	// 2 hoge
	// 3 hogehoge
}

func ExampleCommonPrefixSearch() {
	keyset := NewKeyset()
	defer keyset.Dispose()

	keyset.PushBack("ho")
	keyset.PushBack("hoge")
	keyset.PushBack("hogehoge")
	keyset.PushBack("mogemoge")

	trie, err := keyset.Build()
	panicIfErrorOccurred(err)
	defer trie.Dispose()

	err = trie.CommonPrefixSearch("hogehogehogehogehoge", printKey)
	panicIfErrorOccurred(err)

	// Output:
	// 0 ho
	// 2 hoge
	// 3 hogehoge
}
