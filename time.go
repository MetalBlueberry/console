package console

import (
	"sync"
	"time"
)

func Time(label ...string)    { DefaultConsole.Time(label...) }
func TimeLog(label ...string) { DefaultConsole.TimeLog(label...) }
func TimeEnd(label ...string) { DefaultConsole.TimeEnd(label...) }

func (c *Console) Time(label ...string) {
	key := c.keyFromLabels(label...)
	ok := c.timer.Time(key)
	if !ok {
		c.Log("Timer '%s' already exist", key)
	}
}

func (c *Console) TimeEnd(label ...string) {
	key := c.keyFromLabels(label...)
	duration, ok := c.timer.TimeEnd(key)
	if !ok {
		c.Log("Timer '%s' does not exist", key)
		return
	}

	c.Log("%s: %d ms", key, duration.Milliseconds())
}

func (c *Console) TimeLog(label ...string) {
	key := c.keyFromLabels(label...)
	duration, ok := c.timer.TimeLog(key)
	if !ok {
		c.Log("Timer '%s' does not exist", key)
		return
	}
	c.Log("%s: %d ms", key, duration.Milliseconds())
}

type Timer struct {
	sync.RWMutex
	m map[string]time.Time
}

func (c *Timer) init() {
	if c.m == nil {
		c.m = map[string]time.Time{}
	}
}

func (c *Timer) Time(key string) bool {
	c.Lock()
	defer c.Unlock()
	c.init()
	_, ok := c.m[key]
	if ok {
		return false
	}

	c.m[key] = time.Now().UTC()
	return true
}

func (c *Timer) TimeEnd(key string) (time.Duration, bool) {
	c.Lock()
	defer c.Unlock()
	c.init()

	start, ok := c.m[key]
	if !ok {
		return time.Duration(0), false
	}

	delete(c.m, key)
	return time.Now().UTC().Sub(start), true
}

func (c *Timer) TimeLog(key string) (time.Duration, bool) {
	c.RLock()
	defer c.RUnlock()
	c.init()

	start, ok := c.m[key]
	if !ok {
		return time.Duration(0), false
	}

	return time.Now().UTC().Sub(start), true
}
