package main

import "sync"

type mapWithMyutex[K comparable, V any] struct {
	mu          sync.Mutex
	internalMap map[K]V
}

func (c *mapWithMyutex[K, V]) read(variable K) V {

	c.mu.Lock()
	defer c.mu.Unlock()
	return c.internalMap[variable]
}

func (c *mapWithMyutex[K, V]) write(variable K, value V) {

	c.mu.Lock()
	defer c.mu.Unlock()
	c.internalMap[variable] = value
}

func (c *mapWithMyutex[K, V]) copyMap() map[K]V {
	c.mu.Lock()
	defer c.mu.Unlock()
	newMap := make(map[K]V)
	for key, value := range c.internalMap {
		newMap[key] = value
	}
	return newMap
}

func (c *mapWithMyutex[K, V]) hasKey(variable K) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	_, exists := c.internalMap[variable]
	return exists
}
