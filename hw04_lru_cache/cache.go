package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

func (c *lruCache) removeItemByValue(val interface{}) {
	filtered := make(map[Key]*ListItem, c.capacity)

	for k, i := range c.items {
		if i.Value == val {
			continue
		}

		filtered[k] = i
	}

	c.items = filtered
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	ci := cacheItem{
		value: value,
		key:   key,
	}
	if i, ok := c.items[key]; ok {
		i.Value = ci
		c.queue.MoveToFront(i)
		return true
	}

	if c.queue.Len() < c.capacity {
		i := c.queue.PushFront(ci)
		c.items[key] = i
		return false
	}

	last := c.queue.Back()
	c.queue.Remove(last)
	delete(c.items, last.Value.(cacheItem).key)

	i := c.queue.PushFront(ci)
	c.items[key] = i

	return false
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	if i, ok := c.items[key]; ok {
		c.queue.MoveToFront(i)
		return i.Value.(cacheItem).value, true
	}

	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}
