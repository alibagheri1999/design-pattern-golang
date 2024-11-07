package main

import (
	"log"
	"sync"
)

type CacheManager struct {
	cache map[string]any
	mutex sync.RWMutex
}

var instance *CacheManager
var once sync.Once

func GetCacheManager() *CacheManager {
	once.Do(func() {
		instance = &CacheManager{
			cache: make(map[string]any),
		}
	})
	return instance
}

func (cm *CacheManager) Set(key string, value any) {
	cm.mutex.Lock()
	defer cm.mutex.Unlock()
	cm.cache[key] = value
}

func (cm *CacheManager) Get(key string) (any, bool) {
	cm.mutex.RLock()
	defer cm.mutex.RUnlock()
	value, ok := cm.cache[key]
	return value, ok
}

func GetData() {
	cache := GetCacheManager()
	item, ok := cache.Get("some_key")
	log.Println(item, ok)
}

func SetData() {
	cache := GetCacheManager()
	cache.Set("some_key", "some data")
}

func singleton() {
	SetData()
	GetData()
}