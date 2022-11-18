package main

import "sync"

//
//### Задача "создай структуру-обертку для словаря"
//1) Сначала просто создай структуру и реализуй у нее методы:
//- Get(key string) (interface{}, bool)
//- Set(key string, value interface{})
//2) Потом сделай ее потокобезопасной (на это задание еще и пытаются вопросами сами навести, типа "а что тут не так")

// MyDict is an interface for dictionary.
type MyDict interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

// myThreadSaveDict is a thread-safe dictionary.
type myThreadSaveDict struct {
	dict map[string]interface{}
	// using sync.RWMutex instead of sync.Mutex because we have Get and Set methods that are read-only and read-write respectively
	mutex sync.RWMutex
}

// NewMyDict returns a new instance of MyDict.
func NewMyDict() MyDict {
	return &myThreadSaveDict{dict: make(map[string]interface{})}
}

// Get is thread-safe and returns (value, true) if key exists.
func (d *myThreadSaveDict) Get(key string) (interface{}, bool) {
	d.mutex.RLock()
	defer d.mutex.RUnlock()

	value, ok := d.dict[key]
	return value, ok
}

// Set is thread-safe and overwrites value if key exists.
func (d *myThreadSaveDict) Set(key string, value interface{}) {
	d.mutex.Lock()
	defer d.mutex.Unlock()

	d.dict[key] = value
}
