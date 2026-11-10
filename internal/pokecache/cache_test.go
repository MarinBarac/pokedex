package pokecache

import (
	"testing"
	"time"
)

func TestAddingValuesInCache(t *testing.T) {
	type inputType struct {
		key string
		val []byte
	}

	existingInput := inputType{
		key: "key1",
		val: make([]byte, 1),
	}

	cases := map[string]struct {
		startingCache *Cache
		input         inputType
	}{
		"adding new value": {
			startingCache: NewCache(5*time.Second, nil),
			input:         existingInput,
		},
		"adding existing value": {
			startingCache: NewCache(5*time.Second, &map[string]cacheEntry{
				"key1": cacheEntry{
					val:       make([]byte, 1),
					createdAt: time.Now(),
				},
			}),
			input: inputType{
				key: "key1",
				val: make([]byte, 1),
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			c.startingCache.Add(c.input.key, c.input.val)

			newVal := c.startingCache.entries[c.input.key].val

			if len(newVal) != len(c.input.val) {
				t.Errorf("%s: Wrong length!\nexpected: %#v, got: %#v", name, newVal, c.input.val)
			}

			for i := 0; i < len(c.input.val); i++ {
				if newVal[i] != c.input.val[i] {
					t.Errorf("%s: Wrong element at index %v\nexpected: %#v, got: %#v", name, i, newVal, c.input.val)
				}
			}

		})
	}

}

func TestGettingCacheValue(t *testing.T) {
	type expectedType struct {
		val []byte
		ok  bool
	}

	entryKey := "key"
	entries := map[string]cacheEntry{
		entryKey: cacheEntry{
			val:       make([]byte, 0),
			createdAt: time.Now(),
		},
	}

	cases := map[string]struct {
		startingCache *Cache
		input         string
		expected      expectedType
	}{
		"missing key": {
			startingCache: NewCache(5*time.Second, nil),
			input:         "key",
			expected: expectedType{
				val: nil,
				ok:  false,
			},
		},
		"existing key": {
			startingCache: NewCache(5*time.Second, &entries),
			input:         entryKey,
			expected: expectedType{
				val: entries[entryKey].val,
				ok:  true,
			},
		},
	}

	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			entry, ok := c.startingCache.Get(c.input)

			if ok != c.expected.ok {
				t.Errorf("%s:\nExpected: %#v, %#v\nGot: %#v, %#v", name, c.expected.val, c.expected.ok, entry, ok)
			}

			if len(entry) != len(c.expected.val) {
				t.Errorf("%s: Wrong length!\nExpected: %#v, %#v\nGot: %#v, %#v", name, c.expected.val, c.expected.ok, entry, ok)
			}

			for i := 0; i < len(entry); i++ {
				if entry[i] != c.expected.val[i] {
					t.Errorf("%s: Wrong element at index %v\nExpected: %#v, %#v\nGot: %#v, %#v", name, i, c.expected.val, c.expected.ok, entry, ok)
				}
			}
		})
	}
}

func TestCacheExpiration(t *testing.T) {
	cache := NewCache(100*time.Millisecond, nil)

	key := "key"
	val := "value"

	cache.Add(key, []byte(val))
	entry, ok := cache.Get(key)

	if ok == false {
		t.Fatalf("Fresh entry was immediately removed: %#v", cache)
	}

	if string(entry) != val {
		t.Fatalf("Get returned %q; expected %q", string(entry), val)
	}

	deadline := time.Now().Add(time.Second)
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for now := range ticker.C {
		if now.After(deadline) {
			t.Fatalf("Cache didn't clear stale value: %#v", cache)
		}

		_, ok := cache.Get(key)

		if ok == false {
			return
		}
	}
}
