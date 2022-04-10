/*
 * MIT License
 *
 * Copyright (c) 2022 Ludovic Cleroux
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package hashmap

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func TestHashMap_Get(t *testing.T) {
	h := NewHashMap[string, int]()
	if err := h.Put("bla", 1); !assert.NoError(t, err) {
		return
	}
	v, err := h.Get("bla")
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, 1, v)
}

func TestHashMap_GetNotExisting(t *testing.T) {
	h := NewHashMap[string, int]()
	_, err := h.Get("bla")
	if !assert.Error(t, err) {
		return
	}
}

func TestHashMap_ContainsKey(t *testing.T) {
	h := NewHashMap[string, int]()
	if err := h.Put("bla", 1); !assert.NoError(t, err) {
		return
	}
	containsKey, err := h.ContainsKey("bla")
	if !assert.NoError(t, err) {
		return
	}
	assert.True(t, containsKey)
}

func TestHashMap_ContainsKey_NotExisting(t *testing.T) {
	h := NewHashMap[string, int]()
	containsKey, err := h.ContainsKey("bla")
	if !assert.NoError(t, err) {
		return
	}
	assert.False(t, containsKey)
}

func TestHashMap_Remove(t *testing.T) {
	h := NewHashMap[string, int]()
	if err := h.Put("bla", 1); !assert.NoError(t, err) {
		return
	}
	v, err := h.Remove("bla")
	if !assert.NoError(t, err) {
		return
	}
	assert.Equal(t, 1, v)
}

func TestHashMap_Remove_NotExisting(t *testing.T) {
	h := NewHashMap[string, int]()
	_, err := h.Remove("bla")
	if !assert.Error(t, err) {
		return
	}
}

func TestHashMap_Grow_Shrink(t *testing.T) {
	h := NewHashMap[string, int]()
	iterations := 100000
	for i := 0; i < iterations; i++ {
		if err := h.Put(strconv.Itoa(i), i); !assert.NoError(t, err) {
			return
		}
	}
	fmt.Printf("HashMap Size: %v\n", h.Size)
	fmt.Printf("HashMap Bucket Count: %v\n", len(h.Buckets))
	for i := 0; i < iterations; i++ {
		val, err := h.Get(strconv.Itoa(i))
		if !assert.NoError(t, err) {
			return
		}
		assert.Equal(t, i, val)
	}
	for i := 0; i < iterations; i++ {
		if _, err := h.Remove(strconv.Itoa(i)); !assert.NoError(t, err) {
			return
		}
	}
	assert.Equal(t, 0, h.Size)
	fmt.Printf("HashMap Size: %v\n", h.Size)
	fmt.Printf("HashMap Bucket Count: %v\n", len(h.Buckets))
}

func BenchmarkHashMap_Put(b *testing.B) {
	h := NewHashMap[string, int]()
	for i := 0; i < b.N; i++ {
		if err := h.Put(strconv.Itoa(i), i); !assert.NoError(b, err) {
			return
		}
	}
}

func BenchmarkHashMap_Get(b *testing.B) {
	h := NewHashMap[string, int]()
	for i := 0; i < b.N; i++ {
		if err := h.Put(strconv.Itoa(i), i); !assert.NoError(b, err) {
			return
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := h.Get(strconv.Itoa(i)); !assert.NoError(b, err) {
			return
		}
	}
}

func BenchmarkMap_Put(b *testing.B) {
	h := map[string]int{}
	for i := 0; i < b.N; i++ {
		h[strconv.Itoa(i)] = i
	}
}

func BenchmarkMap_Get(b *testing.B) {
	h := map[string]int{}
	for i := 0; i < b.N; i++ {
		h[strconv.Itoa(i)] = i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = h[strconv.Itoa(i)]
	}
}
