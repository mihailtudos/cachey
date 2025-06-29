package cache_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/mihailtudos/cachey/cache"
)

func TestCache(t *testing.T) {
	c := cache.New()
	key := "test"
	value := 1

	t.Run("Create new cache", func(t *testing.T) {
		tp := reflect.TypeOf(c)
		if tp.Kind() != reflect.Ptr {
			t.Fatal("returned type must be a pointer")
		}

		expectedType := reflect.TypeOf((*cache.Cache)(nil))
		if tp != expectedType {
			t.Fatal("returned type must be a map")
		}
	})

	t.Run("Set and check value exists", func(t *testing.T) {
		c.Set(key, value, time.Second*10)
		v, ok := c.Get(key)
		if !ok {
			t.Fatalf("expected true got: %v", ok)
		}

		if val, ok := v.(int); ok {
			if val != value {
				t.Fatalf("expected value 1 got: %v", val)
			}
		} else {
			t.Fatalf("expected int value, got: %T", v)
		}
	})

	t.Run("deleting key that doesn't exists", func(t *testing.T) {
		c.Delete("fake")
	})

	t.Run("deleting key that exists", func(t *testing.T) {
		c.Delete(key)
	})

	t.Run("key should not be found after deletion", func(t *testing.T) {
		_, ok := c.Get(key)
		if ok {
			t.Fatal("key and value should have been removed")
		}
	})
}
