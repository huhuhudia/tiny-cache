package cache

import (
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"
)

func TestCache_Get(t *testing.T) {
	cache := New()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func(){
		for i:=0; i<10000000; i++{
			cache.Set("this", "jiu shi", int64(time.Second))
			if _, exist := cache.Get("this"); exist{
				//print(el.(string))
			}
			cache.Set("this", "jiu", int64(time.Second))
			if _, exist := cache.Get("this"); exist{
				//print(el.(string))
			}
		}
		wg.Done()
		//print("quit")
	}()
	go func() {
		for i:=0; i<10000000; i++{
			cache.Get("this")
		}
		wg.Done()
	}()
	cache.Set("this", "jiu shi", int64(time.Second*10))
	//print(int(time.Second))
	time.Sleep(time.Second*10)
	val ,exist:= cache.Get("this")
	if exist{
		fmt.Println(reflect.TypeOf(val))
	}

	wg.Wait()
}

type Student struct {
	name string

}
func BenchmarkCache_Get(b *testing.B) {
	cache := New()
	for i:=0; i<b.N; i++{
		cache.Set("this", 1, int64(time.Second))
		if _, exist := cache.Get("this"); exist{
			//print(el.(string))
		}
	}
}