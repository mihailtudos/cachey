# Cachey

_Cachey_ is a lightweight, simple and powerful in-memory cache library built with Golang.

## Description

This library provides basic cache functionality as described below

1. Writing values in cache by key:

```go
    Set(key string, value interface{})
```

2. Retrieving a value by key:

```go
    Get(key string)
```

3. Removing a value from cache by key:

```go
    Delete(key)
```

## Example

```go
    cache := cache.New()

    cache.Set("userId", 42)
    userId := cache.Get("userId")

    fmt.Println(userId)

    cache.Delete("userId")
    userId := cache.Get("userId")

    fmt.Println(userId)
```
