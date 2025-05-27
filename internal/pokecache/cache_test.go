package pokecache

import (
	"testing"
	"time"
)

func TestCacheAddAndGet(t *testing.T) {
	c := NewCache(2 * time.Second)
	key := "https://pokeapi.co/api/v2/location-area?offset=0&limit=20"
	value := []byte("test-data")

	c.Add(key, value)
	if got, ok := c.Get(key); !ok {
		t.Fatalf("key %s not found in cache", key)
	} else if string(got) != string(value) {
		t.Fatalf("got %s, want %s", string(got), string(value))
	}
}
