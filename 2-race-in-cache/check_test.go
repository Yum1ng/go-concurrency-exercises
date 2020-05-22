//////////////////////////////////////////////////////////////////////
//
// DO NOT EDIT THIS PART
// Your task is to edit `main.go`
//

package main

import "testing"

func TestMain(t *testing.T) {
	cache := run()

	cacheLen := len(cache.cache)
	pagesLen := cache.pages.Len()
	if cacheLen != CacheSize {
		t.Errorf("Incorrect cache size cacheLen %v, pagesLen %v", cacheLen, pagesLen)
	}
	if pagesLen != CacheSize {
		t.Errorf("Incorrect pages size pagesLen %v, cacheLen %v", pagesLen, cacheLen)
	}
}
