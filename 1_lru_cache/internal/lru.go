package internal

import (
	"slices"
)

// Lru structure for caching results
type Lru struct {
	Size  int
	Cache map[string]any
	Keys  []string
}

// Get return index from a key
func (l *Lru) Get(k string) any {
	if slices.Contains(l.Keys, k) {
		return l.Cache[k]
	}

	return -1
}

// Set adds key:value if Keys is under capacity
func (l *Lru) Set(k string, v any) {
	if len(l.Keys) >= l.Size {
		lp := len(l.Keys) - 1
		lk := l.Keys[lp]
		l.Keys = slices.Delete(l.Keys, lp, lp+1)
		delete(l.Cache, lk)
	}

	l.Cache[k] = v

	nk := []string{}
	nk = append(nk, k)

	for _, s := range l.Keys {
		if s != k {
			nk = append(nk, s)
		}
	}
	l.Keys = nk

	return
}

// NewLRU initializer
func NewLRU(size int) *Lru {
	return &Lru{
		Size:  size,
		Cache: make(map[string]any),
		Keys:  []string{},
	}
}
