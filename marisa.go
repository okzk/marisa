package marisa

/*
#cgo CPPFLAGS: -I ./marisa-trie/include -I ./marisa-trie/lib
#cgo LDFLAGS: -lstdc++

#include <stdlib.h>
#include "marisa_glue.h"
*/
import "C"
import (
	"errors"
	"fmt"
	"unsafe"
)

var NotFound error = errors.New("not found")

type LookupCallback func(id uint64, key string) error

func toError(errCode C.marisa_error_code) error {
	if errCode != C.MARISA_OK {
		return fmt.Errorf("marisa error: %d", errCode)
	}
	return nil
}

type Keyset struct {
	keyset C.marisa_keyset
}

func NewKeyset() *Keyset {
	return &Keyset{C.marisa_keyset_new()}
}

func (k *Keyset) Dispose() {
	C.marisa_keyset_delete(k.keyset)
}

func (k *Keyset) PushBack(str string) error {
	return k.PushBackWithWeight(str, 1.0)
}

func (k *Keyset) PushBackWithWeight(str string, weight float32) error {
	// to avoid a redundant copy, use raw buffer of string.
	// buffer will be copied in marisa_keyset_push_back()
	buf := *(*unsafe.Pointer)(unsafe.Pointer(&str))

	errCode := C.marisa_keyset_push_back(k.keyset, (*C.char)(buf), C.size_t(len(str)), C.float(weight))
	return toError(errCode)
}

func (k *Keyset) build(flag int) (*Trie, error) {
	trie := C.marisa_trie_new()
	errCode := C.marisa_trie_build(trie, k.keyset, C.int(flag))
	if errCode != C.MARISA_OK {
		C.marisa_trie_delete(trie)
		return nil, toError(errCode)
	}

	return &Trie{trie}, nil
}

func (k *Keyset) Build() (*Trie, error) {
	return k.build(0)
}

func (k *Keyset) BuildWithOption(option *BuildOption) (*Trie, error) {
	if err := option.validate(); err != nil {
		return nil, err
	}
	return k.build(option.value())
}

type Trie struct {
	trie C.marisa_trie
}

func LoadTrieFromFile(file string) (*Trie, error) {
	cs := C.CString(file)
	defer C.free(unsafe.Pointer(cs))

	trie := C.marisa_trie_new()
	errCode := C.marisa_trie_load(trie, cs)
	if errCode != C.MARISA_OK {
		C.marisa_trie_delete(trie)
		return nil, toError(errCode)
	}

	return &Trie{trie}, nil
}

func MmapTrie(file string) (*Trie, error) {
	cs := C.CString(file)
	defer C.free(unsafe.Pointer(cs))

	trie := C.marisa_trie_new()
	errCode := C.marisa_trie_mmap(trie, cs)
	if errCode != C.MARISA_OK {
		C.marisa_trie_delete(trie)
		return nil, toError(errCode)
	}

	return &Trie{trie}, nil
}

func (t *Trie) Save(file string) error {
	cs := C.CString(file)
	defer C.free(unsafe.Pointer(cs))

	errCode := C.marisa_trie_save(t.trie, cs)
	return toError(errCode)
}

func (t *Trie) Dispose() {
	C.marisa_trie_delete(t.trie)
	t.trie = nil
}

func (t *Trie) Lookup(str string) (uint64, error) {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	agent := C.marisa_agent_new()
	defer C.marisa_agent_delete(agent)

	errCode := C.marisa_agent_set_query_with_str(agent, cstr, C.size_t(len(str)))
	if errCode != C.MARISA_OK {
		return 0, toError(errCode)
	}

	result := C.marisa_trie_lookup(t.trie, agent)
	if result.err != C.MARISA_OK {
		return 0, toError(errCode)
	}

	if result.found {
		return uint64(result.id), nil
	} else {
		return 0, NotFound
	}
}

func (t *Trie) ReverseLookup(id uint64) (string, error) {
	agent := C.marisa_agent_new()
	defer C.marisa_agent_delete(agent)

	errCode := C.marisa_agent_set_query_with_id(agent, C.size_t(id))
	if errCode != C.MARISA_OK {
		return "", toError(errCode)
	}

	result := C.marisa_trie_reverse_lookup(t.trie, agent)
	if result.err != C.MARISA_OK {
		return "", toError(errCode)
	}

	if result.found {
		return C.GoStringN(result.str, C.int(result.len)), nil
	} else {
		return "", NotFound
	}
}

func (t *Trie) CommonPrefixSearch(str string, callback LookupCallback) error {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	agent := C.marisa_agent_new()
	defer C.marisa_agent_delete(agent)

	errCode := C.marisa_agent_set_query_with_str(agent, cstr, C.size_t(len(str)))
	if errCode != C.MARISA_OK {
		return toError(errCode)
	}

	for {
		result := C.marisa_trie_common_prefix_search(t.trie, agent)
		if result.err != C.MARISA_OK {
			return toError(errCode)
		}
		if !result.found {
			return nil
		}

		err := callback(uint64(result.id), C.GoStringN(result.str, C.int(result.len)))
		if err != nil {
			return err
		}
	}
}

func (t *Trie) PredictiveSearch(str string, callback LookupCallback) error {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	agent := C.marisa_agent_new()
	defer C.marisa_agent_delete(agent)

	errCode := C.marisa_agent_set_query_with_str(agent, cstr, C.size_t(len(str)))
	if errCode != C.MARISA_OK {
		return toError(errCode)
	}

	for {
		result := C.marisa_trie_predictive_search(t.trie, agent)
		if result.err != C.MARISA_OK {
			return toError(errCode)
		}
		if !result.found {
			return nil
		}

		err := callback(uint64(result.id), C.GoStringN(result.str, C.int(result.len)))
		if err != nil {
			return err
		}
	}
}

func (t *Trie) NumTries() (uint64, error) {
	n := C.marisa_trie_num_tries(t.trie)
	return uint64(n.num), toError(n.err)
}

func (t *Trie) NumKeys() (uint64, error) {
	n := C.marisa_trie_num_keys(t.trie)
	return uint64(n.num), toError(n.err)
}

func (t *Trie) NumNodes() (uint64, error) {
	n := C.marisa_trie_num_nodes(t.trie)
	return uint64(n.num), toError(n.err)
}

func (t *Trie) Size() (uint64, error) {
	n := C.marisa_trie_size(t.trie)
	return uint64(n.num), toError(n.err)
}

func (t *Trie) TotalSize() (uint64, error) {
	n := C.marisa_trie_total_size(t.trie)
	return uint64(n.num), toError(n.err)
}

func (t *Trie) IoSize() (uint64, error) {
	n := C.marisa_trie_io_size(t.trie)
	return uint64(n.num), toError(n.err)
}
