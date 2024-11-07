package lru

import "container/list"

// 实现key value 快速的寻找（map），增删改查
// lru, 它是有序的，可以用list 实现


type Cache struct {
	ll *list.List
	cache map[string]*list.Element
	maxBytes int64
	nbytes int64
	OnEvicted func(key string, value Value)
}

type entry struct {
	key string 
	value Value
}

type Value interface {
	Len() int
}

func New(maxBytes int64, onEvicted func(key string, value Value)) *Cache {
	return &Cache{
		ll: list.New(),
		cache: make(map[string]*list.Element),
		maxBytes: maxBytes,
		OnEvicted: onEvicted,
	}
}

func (c *Cache) Get(key string) (value Value, ok bool) {
	if ele, ok := c.cache[key]; ok {
		c.ll.MoveToFront(ele)
		kv := ele.Value.(*entry)
		value = kv.value
		return value, true
	}
	return nil, false
}

func (c *Cache) RemoveOldest() {
	for c.maxBytes != 0 && c.nbytes >= c.maxBytes {
		ele := c.ll.Back()
		kv := ele.Value.(*entry)
		key := kv.key
		delete(c.cache, key)
		c.ll.Remove(ele)
		c.OnEvicted(key, kv.value)
		c.nbytes -= int64(len(kv.key) + kv.value.Len())
	}
}

func (c *Cache) Add(key string, value Value) {
	if ele, ok := c.cache[key]; ok {
		kv := ele.Value.(*entry)
		c.nbytes += (int64(value.Len()) - int64(kv.value.Len()))
		kv.value = value
		c.ll.MoveToFront(ele)
		if c.nbytes >= c.maxBytes {
			c.RemoveOldest()
		}
		return
	}
	entry := &entry{
		key: key,
		value: value,
	}
	ele := c.ll.PushFront(entry)
	c.cache[key] = ele
	c.nbytes += (int64(len(key)) + int64(value.Len()))
	if c.nbytes >= c.maxBytes {
		c.RemoveOldest()
	}
}

func (c *Cache) Len() int {
	return c.ll.Len()
}

