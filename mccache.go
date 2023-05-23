package mccache

import (
	"fmt"
	"log"
	"sync"
)

// 为指定的key加载数据
type Getter interface {
	Get(key string) ([]byte, error)
}

// 实现了Getter接口的函数类型
type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

var (
	mu     sync.RWMutex
	groups = make(map[string]*Group)
)

// Group是缓存的命名空间和相关数据，它们被加载到一台或多台机器组成的组中
type Group struct {
	name      string
	getter    Getter
	mainCache cache
}

// NewGroup在组中创建一个新实例
func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	mu.Lock()
	defer mu.Unlock()
	g := &Group{
		name:      name,
		getter:    getter,
		mainCache: cache{cacheBytes: cacheBytes},
	}
	groups[name] = g
	return g
}

// GetGroup返回先前通过NewGroup创建的组，没有该名称的组则返回nil
func GetGroup(name string) *Group {
	mu.RLock()
	g := groups[name]
	mu.RUnlock()
	return g
}

// 从缓存中获取一个键的值
func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is empty")
	}
	if v, ok := g.mainCache.get(key); ok {
		log.Println("[McCache] hit")
		return v, nil
	}
	return g.load(key)
}

// 缓存未命中，加载源数据
func (g *Group) load(key string) (value ByteView, err error) {
	return g.getLocally(key)
}

// getLocally调用用户回调函数获取源数据
func (g *Group) getLocally(key string) (ByteView, error) {
	bytes, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView{b: cloneBytes(bytes)}
	g.populateCache(key, value)
	return value, nil
}

func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}
