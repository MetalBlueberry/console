package console

import (
	"strings"
	"sync"
)

const (
	defaultLabel = "default"
)

func Count(label ...string) {
	DefaultConsole.Count(label...)
}

func CountReset(label ...string) {
	DefaultConsole.CountReset(label...)
}

func (c *Console) Count(label ...string) {
	key := c.keyFromLabels(label...)
	c.counter.Inc(key)

	c.Log("%s: %d", key, c.counter.Get(key))
}

func (c *Console) CountReset(label ...string) {
	key := c.keyFromLabels(label...)
	c.counter.Reset(key)

	c.Log("%s: %d", key, c.counter.Get(key))
}

func (c *Console) keyFromLabels(label ...string) string {
	key := strings.Join(label, "")
	if key == "" {
		key = defaultLabel
	}
	return key
}

type Counter struct {
	sync.RWMutex
	m map[string]int
}

func (c *Counter) init() {
	if c.m == nil {
		c.m = map[string]int{}
	}
}

func (c *Counter) Inc(key string) {
	c.Lock()
	c.init()
	defer c.Unlock()
	c.m[key]++
}
func (c *Counter) Reset(key string) {
	c.Lock()
	c.init()
	defer c.Unlock()
	c.m[key] = 0
}

func (c *Counter) Get(key string) int {
	c.RLock()
	c.init()
	defer c.RUnlock()
	return c.m[key]
}
