package geecache

import (
	"fmt"
	"sync"
)

type Getter interface {
	Get(key string) ([]byte, error)
}

type GetterFunc func(key string) ([]byte, error)

func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}

type Group struct {
	name      string
	getter    Getter
	mainCache *cache
}

var (
	mu sync.RWMutex
	groups map[string]*Group
)

func NewGroup(name string, cacheBytes int64, getter Getter) *Group {
	if getter == nil {
		panic("nil Getter")
	}
	group := &Group{
		name: name,
		getter: getter,
		mainCache: &cache{cacheBytes: cacheBytes},
	}
	mu.Lock()
	defer mu.Unlock()
	if groups == nil {
		groups = make(map[string]*Group)
	}
	groups[name] = group

	return group
}

func GetGroup(name string) *Group {
	mu.RLock()
	defer mu.RUnlock()
	g := groups[name]
	return g
}

func (g *Group) Get(key string) (ByteView, error) {
	if key == "" {
		return ByteView{}, fmt.Errorf("key is empty")
	}

	if v, ok := g.mainCache.get(key); ok {
		return v, nil
	}
	
	return g.load(key)
}

// 加载，先看看能不能 从别的节点获取，不能才使用本地加载
func (g *Group) load(key string) (ByteView, error) {
	return g.getLocally(key)
}

// 本地加载，获得数据后，要存入缓存里面哦
func (g *Group) getLocally(key string) (ByteView, error) {
	b, err := g.getter.Get(key)
	if err != nil {
		return ByteView{}, err
	}
	value := ByteView{b}
	g.populateCache(key, value)
	return value, err
} 

// 将数据存入缓存 
func (g *Group) populateCache(key string, value ByteView) {
	g.mainCache.add(key, value)
}

