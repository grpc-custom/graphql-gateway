package cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCache(t *testing.T) {
	cache := New(100)
	cache.Set("aaa", "12345", time.Second)
	data, ok := cache.Get("aaa")
	assert.Equal(t, "12345", data)
	assert.True(t, ok)
	time.Sleep(3 * time.Second)
	data, ok = cache.Get("aaa")
	assert.Nil(t, data)
	assert.False(t, ok)
}
