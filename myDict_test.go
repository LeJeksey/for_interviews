package main

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Attention: Это не классический способ написания тестов на go, но так понятней в данном кейсе

// TestMyDictSync just checking you dict
func TestMyDictSync(t *testing.T) {
	dict := NewMyDict()

	dict.Set("key1", "value1")
	dict.Set("key2", "value2")
	dict.Set("key3", "value3")
	dict.Set("key3", "value4")

	value, ok := dict.Get("key1")
	assert.True(t, ok)
	assert.Equal(t, "value1", value)

	value, ok = dict.Get("key2")
	assert.True(t, ok)
	assert.Equal(t, "value2", value)

	value, ok = dict.Get("key3")
	assert.True(t, ok)
	assert.Equal(t, "value4", value)

	value, ok = dict.Get("key4")
	assert.False(t, ok)
	assert.Nil(t, value)
}

// TestMyDictConcurrent will be panic if dict is not thread-safe
func TestMyDictConcurrent(t *testing.T) {
	dict := NewMyDict()

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			dict.Set("key", i)
		}(i)
	}
	wg.Wait()

	_, ok := dict.Get("key")
	assert.True(t, ok)
}
