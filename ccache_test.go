package cache

import (
	"fmt"
	"testing"
	"time"
)

func Test_ccache(t *testing.T) {
	cache := NewCCache(100)

	data := NewData("abc").SetKey("key")
	data.ToCCache(cache, time.Second*1)
	time.Sleep(time.Second * 2)
	key := NewKey("key")
	res, err := key.FetchFromCCache(cache)
	fmt.Println("err=", err)

	fmt.Println(res)

}
