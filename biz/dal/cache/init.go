package cache

import (
	"sync"
	"time"
)

var (
	DB *Cache
)

// Cache 定义缓存结构
type Cache struct {
	sync.RWMutex
	items map[string]*Item
}

// Item 定义缓存项结构
type Item struct {
	Value      interface{}
	Expiration int64
}

func Init() {
	DB = New()
}

// New 创建一个新的缓存实例
func New() *Cache {
	cache := &Cache{
		items: make(map[string]*Item),
	}
	go cache.cleanExpired() // 启动清理过期项的goroutine
	return cache
}

// Set 设置缓存，可以指定过期时间
func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
	c.Lock()
	defer c.Unlock()

	expiration := time.Now().Add(duration).UnixNano()
	c.items[key] = &Item{
		Value:      value,
		Expiration: expiration,
	}
}

// Get 获取缓存项
func (c *Cache) Get(key string) (interface{}, bool) {
	c.RLock()
	defer c.RUnlock()

	item, exists := c.items[key]
	if !exists {
		return nil, false
	}

	if item.Expiration > 0 && time.Now().UnixNano() > item.Expiration {
		return nil, false
	}

	return item.Value, true
}

// Delete 删除缓存项
func (c *Cache) Delete(key string) {
	c.Lock()
	defer c.Unlock()
	delete(c.items, key)
}

// cleanExpired 定期清理过期的缓存项
func (c *Cache) cleanExpired() {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		c.Lock()
		for key, item := range c.items {
			if item.Expiration > 0 && time.Now().UnixNano() > item.Expiration {
				delete(c.items, key)
			}
		}
		c.Unlock()
	}
}
