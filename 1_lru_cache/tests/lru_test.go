package internal

import (
	"testing"

	"devgym-exc.com/lru/internal"
	"github.com/stretchr/testify/assert"
)

// TestLruGetInvalidKey check for invalid keys
func TestLruGetInvalidKeyShouldReturnFalse(t *testing.T) {
	lru := internal.NewLRU(3)
	assert.Equal(t, -1, lru.Get("Mark"))
}

// TestLruInsertValidKeys saves keys
func TestLruInsertValidKeys(t *testing.T) {
	lru := internal.NewLRU(3)
	lru.Set("Mark", 123)
	lru.Set("Mark2", 1244)
	lru.Set("Mark3", 125)

	assert.Equal(t, 123, lru.Get("Mark"))
	assert.Equal(t, 1244, lru.Get("Mark2"))
	assert.Equal(t, 125, lru.Get("Mark3"))
}

// TestLruInsertMoreKeysThanSize saves more keys than its size
func TestLruInsertMoreKeysThanSize(t *testing.T) {
	lru := internal.NewLRU(3)
	lru.Set("Mark", 123)
	lru.Set("Mark2", 1244)
	lru.Set("Mark3", 125)
	lru.Set("Mark4", 111)

	assert.Equal(t, -1, lru.Get("Mark"))
	assert.Equal(t, 1244, lru.Get("Mark2"))
	assert.Equal(t, 125, lru.Get("Mark3"))
	assert.Equal(t, 111, lru.Get("Mark4"))
}
