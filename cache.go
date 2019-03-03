package cache

import (
	"sync"
	"time"
)

type Cache struct {
	cache map[string] Element
	mu sync.RWMutex
}

type Element struct {
	object interface{}
	expire int64
}

func New()Cache{
	return Cache{make(map[string]Element),sync.RWMutex{}}
}


func (self *Cache) Get(key string)(interface{}, bool){
	self.mu.RLock()
	elem, exist := self.cache[key]
	if !exist{
		self.mu.RUnlock()
		return nil, false
	}
	if time.Now().UnixNano() > elem.expire{
		self.mu.RUnlock()
		return nil, false
	}
	self.mu.RUnlock()
	return elem.object,true
}

func (self *Cache) Set(key string, obj interface{},expired int64){
	self.mu.Lock()
	self.cache[key] = Element{obj,time.Now().UnixNano()+expired}
	self.mu.Unlock()
}




