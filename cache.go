package cache

type Cache map[string]any

func New() Cache {
	return Cache(make(map[string]any))
}

func (c Cache) Set(key string, value any) {
	c[key] = value
}

func (c Cache) Get(key string) (any, bool) {
	v, ok := c[key]
	return v, ok
}

func (c Cache) Delete(key string) {
	delete(c, key)
}
