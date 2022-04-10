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
	"bytes"
	"dsa/pkg/linkedlist"
	"encoding/binary"
	"fmt"
)

const (
	growLoadFactor   float32 = 0.75
	shrinkLoadFactor float32 = 0.25
	defaultCapacity  int     = 16
)

type Entry[K comparable, V comparable] struct {
	HashKey uint
	Key     K
	Value   V
}

type Bucket[K comparable, V comparable] struct {
	Entries linkedlist.LinkedList[Entry[K, V]]
}

type HashMap[K comparable, V comparable] struct {
	Capacity int
	Size     int
	Buckets  []Bucket[K, V]
}

func NewHashMap[K comparable, V comparable]() *HashMap[K, V] {
	h := &HashMap[K, V]{
		Capacity: defaultCapacity,
		Size:     0,
		Buckets:  make([]Bucket[K, V], defaultCapacity),
	}
	return h
}

func (h *HashMap[K, V]) Get(key K) (V, error) {
	hashKey, bucketIdx, err := hashFunc(len(h.Buckets), key)
	if err != nil {
		return *new(V), err
	}
	bucket := h.Buckets[bucketIdx]
	cur := bucket.Entries.Head
	for cur != nil {
		if cur.Value.HashKey == hashKey && cur.Value.Key == key {
			return cur.Value.Value, nil
		}
		cur = cur.Next
	}
	return *new(V), fmt.Errorf("key not found")
}

func (h *HashMap[K, V]) Remove(key K) (V, error) {
	result := *new(V)
	hashKey, bucketIdx, err := hashFunc(len(h.Buckets), key)
	if err != nil {
		return result, err
	}
	bucket := h.Buckets[bucketIdx]
	cur := bucket.Entries.Head
	idx := 0
	found := false

	for cur != nil {
		if cur.Value.HashKey == hashKey && cur.Value.Key == key {
			bucket.Entries.RemoveAt(idx)
			h.Size--
			found = true
			result = cur.Value.Value
			break
		}
		cur = cur.Next
		idx++
	}

	if h.loadFactor() <= shrinkLoadFactor {
		h.shrink()
	}
	if !found {
		return result, fmt.Errorf("key not found")
	}
	return result, nil

}

func (h *HashMap[K, V]) Put(key K, value V) error {
	if h.loadFactor() >= growLoadFactor {
		h.grow()
	}

	hashKey, bucketIdx, err := hashFunc(len(h.Buckets), key)
	if err != nil {
		return err
	}
	bucket := h.Buckets[bucketIdx]
	bucket.Entries.Insert(Entry[K, V]{
		Key:     key,
		Value:   value,
		HashKey: hashKey,
	}, 0)
	h.Buckets[bucketIdx] = bucket
	h.Size++
	return nil
}

func (h *HashMap[K, V]) ContainsKey(key K) (bool, error) {
	hashKey, bucketIdx, err := hashFunc(len(h.Buckets), key)
	if err != nil {
		return false, err
	}
	bucket := h.Buckets[bucketIdx]
	cur := bucket.Entries.Head
	for cur != nil {
		if cur.Value.HashKey == hashKey && cur.Value.Key == key {
			return true, nil
		}
		cur = cur.Next
	}
	return false, nil
}

func (h *HashMap[K, V]) loadFactor() float32 {
	return float32(h.Size) / float32(len(h.Buckets))
}

func (h *HashMap[K, V]) rehash(blockSize int) {
	buckets := make([]Bucket[K, V], blockSize)
	for _, bucket := range h.Buckets {
		cur := bucket.Entries.Head
		for cur != nil {
			bucketIdx := cur.Value.HashKey % uint(blockSize)
			buckets[bucketIdx].Entries.Insert(cur.Value, 0)
			cur = cur.Next
		}
	}
	h.Buckets = buckets
}

func (h *HashMap[K, V]) grow() {
	blockSize := len(h.Buckets) * 2
	if blockSize <= defaultCapacity {
		blockSize = defaultCapacity
	}
	h.rehash(blockSize)
}

func (h *HashMap[K, V]) shrink() {
	blockSize := len(h.Buckets) / 2
	if blockSize <= defaultCapacity {
		blockSize = defaultCapacity
	}
	h.rehash(blockSize)
}

func write(buf *bytes.Buffer, v interface{}) error {
	switch val := v.(type) {
	case string:
		buf.WriteString(val)
		return nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return binary.Write(buf, binary.LittleEndian, v)
	default:
		return fmt.Errorf("unsupported type: %T", v)
	}
}

func hashFunc(blockSize int, key interface{}) (hashKey uint, bucketIdx uint, err error) {
	buf := bytes.Buffer{}
	err = nil
	if err := write(&buf, key); err != nil {
		return 0, 0, err
	}
	hashKey = djb2Hash(&buf)
	bucketIdx = hashKey % uint(blockSize)
	return
}

func djb2Hash(buf *bytes.Buffer) uint {
	var h uint = 5381
	for _, r := range buf.Bytes() {
		h = (h << 5) + h + uint(r)
	}
	return h
}
