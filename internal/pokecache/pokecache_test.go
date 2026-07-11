package pokecache

import (
	"testing"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := map[string]struct {
		cacheKey string
		cacheVal []byte
	}{
		"test01": {
			cacheKey: "https://example.com",
			cacheVal: []byte("testdata"),
		},
		"test02": {
			cacheKey: "https://example.com/path",
			cacheVal: []byte("moretestdata"),
		},
	}

	for name, testCase := range cases {
		t.Run(name, func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(testCase.cacheKey, testCase.cacheVal)
			val, ok := cache.Get(testCase.cacheKey)
			if !ok {
				t.Fatalf("expected to find key")
			}
			if string(val) != string(testCase.cacheVal) {
				t.Fatalf("expected to find value")
			}
		})
	}
}

func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5*time.Millisecond
	cache := NewCache(baseTime)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Fatalf("expected to find key")
	}

	time.Sleep(waitTime)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Fatalf("expected to not find key")
	}
}
