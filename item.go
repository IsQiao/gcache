package gcache

import "time"

type GCacheItem[T interface{}] struct {
	expiredTime time.Time
	item        T
}

func (item *GCacheItem[T]) Expired() bool {
	return item.expiredTime.Before(time.Now())
}
