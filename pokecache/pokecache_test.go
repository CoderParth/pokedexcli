package pokecache

import (
	"fmt"
	"testing"
	"time"
)

// TestAddGetCache creates a new cache, adds a new cache, and checks if a valid cache exists.
func TestAddGetCache(t *testing.T) {
	const interval = 5 * time.Second
	testCases := []struct {
		key string
		val []byte
	}{
		{
			key: "https://example.com",
			val: []byte("testdata"),
		},
		{
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for i, c := range testCases {
		t.Run(fmt.Sprintf("Test Case %v", i), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(c.key, c.val)
			val, ok := cache.Get(c.key)
			if !ok {
				t.Errorf("Expected to find key: %v", c.key)
				t.Errorf("Key Not Found")
				return
			}
			if string(val) != string(c.val) {
				t.Errorf("Value does ont match")
				t.Errorf("Expected: %v", string(c.val))
				t.Errorf("Received: %v", string(val))
				return
			}
		})
	}
}

// TestReapLoop creates a new cache with an interval, adds a new cache, and checks if the cache is deleted after that interval
func TestReapLoop(t *testing.T) {
	const interval = 2 * time.Second
	const waitTime = interval + 2*time.Second
	cache := NewCache(interval)
	k, v := "example.com", []byte("testdata")
	cache.Add(k, v)

	_, ok := cache.Get(k)
	if !ok {
		t.Errorf("Expected to find key:  %v", k)
		return
	}

	time.Sleep(waitTime)

	_, ok = cache.Get(k)

	if ok {
		t.Errorf("Expected to not find key")
		return
	}
}
