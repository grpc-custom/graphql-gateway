package cache

import (
	"fmt"
	"sync"
	"time"

	"github.com/golang/protobuf/proto"
	lru "github.com/hashicorp/golang-lru"
)

type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{}, ttl time.Duration)
}

type item struct {
	key      string
	value    interface{}
	ttl      time.Duration
	expireAt time.Time
}

func newItem(key string, value interface{}, ttl time.Duration) *item {
	item := &item{
		key:   key,
		value: value,
		ttl:   ttl,
	}
	item.setExpireAt()
	return item
}

func (i *item) setExpireAt() {
	if i.ttl <= 0 {
		return
	}
	i.expireAt = time.Now().Add(i.ttl)
}

func (i *item) expired() bool {
	if i.ttl <= 0 {
		return false
	}
	return i.expireAt.Before(time.Now())
}

type cache struct {
	mutex sync.RWMutex
	data  *lru.Cache
}

var _ Cache = (*cache)(nil)

func New(size int) Cache {
	c, err := lru.New(size)
	if err != nil {
		panic("unexpected error creating cache: " + err.Error())
	}
	return &cache{
		mutex: sync.RWMutex{},
		data:  c,
	}
}

func (c *cache) getItem(key string) (*item, bool) {
	value, ok := c.data.Get(key)
	if !ok {
		return nil, false
	}
	v, ok := value.(*item)
	if !ok || v.expired() {
		return nil, false
	}
	return v, true
}

func (c *cache) setItem(item *item) {
	c.data.Add(item.key, item)
}

func (c *cache) Get(key string) (interface{}, bool) {
	item, ok := c.getItem(key)
	if !ok {
		return nil, false
	}
	return item.value, ok
}

func (c *cache) Set(key string, value interface{}, ttl time.Duration) {
	item, ok := c.getItem(key)
	if ok {
		item.value = value
		item.ttl = ttl
		item.setExpireAt()
	} else {
		item = newItem(key, value, ttl)
	}
	c.setItem(item)
}

func GenerateKey(name string, msg proto.Message) string {
	const format = "%s%s"
	return fmt.Sprintf(format, name, msg.String())
}
