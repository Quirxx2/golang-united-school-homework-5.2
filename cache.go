package cache

import "time"

type Stored struct {
	Val      string
	Deadline time.Time
}

type Cache struct {
	Cached map[string]Stored
}

func NewCache() Cache {
	dt := make(map[string]Stored)
	return Cache{Cached: dt}
}

func (ch Cache) Get(key string) (string, bool) {
	val, ok := ch.Cached[key]
	if time.Now().After(val.Deadline) {
		return "", false
	}
	return val.Val, ok
}

func (ch Cache) Put(key, value string) {
	st := Stored{Val: value, Deadline: time.Now().Add(10 * time.Minute)}
	ch.Cached[key] = st
}

func (ch Cache) Keys() []string {
	keys := []string{}
	now := time.Now()
	for key, dt := range ch.Cached {
		if now.Before(dt.Deadline) {
			keys = append(keys, key)
		}
	}
	return keys
}

func (ch Cache) PutTill(key, value string, deadline time.Time) {
	dt := Stored{Val: value, Deadline: deadline}
	ch.Cached[key] = dt
}