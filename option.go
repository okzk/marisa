package marisa

/*
#include "marisa_glue.h"
*/
import "C"
import "errors"

const (
	MARISA_MIN_NUM_TRIES     = C.MARISA_MIN_NUM_TRIES
	MARISA_MAX_NUM_TRIES     = C.MARISA_MAX_NUM_TRIES
	MARISA_DEFAULT_NUM_TRIES = C.MARISA_DEFAULT_NUM_TRIES
)

type CacheLevel int

const (
	MARISA_HUGE_CACHE    CacheLevel = C.MARISA_HUGE_CACHE
	MARISA_LARGE_CACHE              = C.MARISA_LARGE_CACHE
	MARISA_NORMAL_CACHE             = C.MARISA_NORMAL_CACHE
	MARISA_SMALL_CACHE              = C.MARISA_SMALL_CACHE
	MARISA_TINY_CACHE               = C.MARISA_TINY_CACHE
	MARISA_DEFAULT_CACHE            = C.MARISA_NORMAL_CACHE
)

type TailMode int

const (
	MARISA_TEXT_TAIL    TailMode = C.MARISA_TEXT_TAIL
	MARISA_BINARY_TAIL           = C.MARISA_BINARY_TAIL
	MARISA_DEFAULT_TAIL          = C.MARISA_TEXT_TAIL
)

type NodeOrder int

const (
	MARISA_LABEL_ORDER   NodeOrder = C.MARISA_LABEL_ORDER
	MARISA_WEIGHT_ORDER            = C.MARISA_WEIGHT_ORDER
	MARISA_DEFAULT_ORDER           = C.MARISA_WEIGHT_ORDER
)

type BuildOption struct {
	NumTries   int
	CacheLevel CacheLevel
	TailMode   TailMode
	NodeOrder  NodeOrder
}

func NewBuildOption() *BuildOption {
	return &BuildOption{
		NumTries:   MARISA_DEFAULT_NUM_TRIES,
		CacheLevel: MARISA_DEFAULT_CACHE,
		TailMode:   MARISA_DEFAULT_TAIL,
		NodeOrder:  MARISA_DEFAULT_ORDER,
	}
}

func (b *BuildOption) validate() error {
	if b.NumTries^C.MARISA_NUM_TRIES_MASK != 0 {
		return errors.New("Invalid NumTries")
	}
	if b.CacheLevel^C.MARISA_CACHE_LEVEL_MASK != 0 {
		return errors.New("Invalid CacheLevel")
	}
	if b.TailMode^C.MARISA_TAIL_MODE_MASK != 0 {
		return errors.New("Invalid TailMode")
	}
	if b.NodeOrder^C.MARISA_NODE_ORDER_MASK != 0 {
		return errors.New("Invalid NodeOrder")
	}

	return nil
}

func (b *BuildOption) value() int {
	return b.NumTries | int(b.CacheLevel) | int(b.TailMode) | int(b.NodeOrder)
}
