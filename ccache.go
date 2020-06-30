package cache

//  内存lru cache

import (
	"errors"
	"fmt"
	gccache "github.com/karlseguin/ccache"
	"time"
)

const MAXSIZE = 1000000
const DEFAULTSIZE = 1000

type CCache struct {
	c *gccache.Cache
}

func NewCCache(maxSize int) *CCache {
	if maxSize > MAXSIZE {
		maxSize = MAXSIZE
	}
	if maxSize <= 0 {
		maxSize = DEFAULTSIZE
	}

	cache := new(CCache)
	count := uint32(maxSize/10 + 1)
	cache.c = gccache.New(gccache.Configure().MaxSize(int64(maxSize)).ItemsToPrune(count))
	return cache
}

func (cd Data) ToCCache(source *CCache, duration time.Duration) {
	if cd.key == "" {
		fmt.Println("toccache ： 没有key")
		return
	}
	source.c.Set(cd.key, cd.data, duration)
}

func (k DataKey) FetchFromCCache(source *CCache) (interface{}, error) {
	key := string(k)
	item := source.c.Get(key)
	if item != nil {
		if item.TTL().Seconds() > 0 {
			return item.Value(), nil
		}
	}
	return nil, errors.New("fetchfromccache: no data")
}
