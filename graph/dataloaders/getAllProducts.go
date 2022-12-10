package dataloaders

import (
	"sync"
	"time"

	"github.com/sRRRs-7/loose_style.git/graph/model"
)

type GetAllProductsConfig struct {
	// Fetch is a method to provide the data for the loader
	Fetch func(first []int64,) ([][]*model.Product, []error)
	// batch wait time
	Wait time.Duration
	// maximum number of keys to send in one batch, 0 = no limit
	MaxBatch int
}

type GetAllProductsLoader struct {
	// provide the data for the loader
	fetch func(first []int64) ([][]*model.Product, []error)
	// batch wait time
	wait time.Duration
	// maximum number of keys to send in one batch, 0 = no limit
	maxBatch int
	// create cache
	cache map[int64][]*model.Product
	// execute batch
	batch *GetAllProductsBatch
	// prevent race
	mu sync.Mutex
}

type GetAllProductsBatch struct {
	keys []int64
	data [][]*model.Product
	error []error
	closing bool
	done chan struct{}
}

func NewGetAllProductsLoader(config GetAllProductsConfig) *GetAllProductsLoader {
	return &GetAllProductsLoader{
		fetch: config.Fetch,
		wait: config.Wait,
		maxBatch: config.MaxBatch,
	}
}

// which will call method ? (LoadThunk, LoadAll, LoadAllThunk, Prime, Clear)
func (l *GetAllProductsLoader) Load(key int64) ([]*model.Product, error) {
	return l.LoadThunk(key)()
}

// LoadThunk is function that block wait for a GetProducts
func (l *GetAllProductsLoader) LoadThunk(key int64) (func() ([]*model.Product, error)) {
	l.mu.Lock()
	if it, ok := l.cache[key]; ok {
		l.mu.Unlock()
		return func() ([]*model.Product, error) {
			return it, nil
		}
	}

	if l.batch == nil {
		l.batch = &GetAllProductsBatch{done: make(chan struct{})}
	}

	batch := l.batch
	pos := batch.keyIndex(l, key)
	l.mu.Unlock()

	return func() ([]*model.Product, error) {
		<- batch.done

		var data []*model.Product
		if pos < len(batch.data) {
			data = batch.data[pos]
		}

		var err error
		if len(batch.error) == 1 {
			err = batch.error[0]
		} else if batch.error != nil {
			err = batch.error[pos]
		}

		if err == nil {
			l.mu.Lock()
			l.unsafeSet(key, data)
			l.mu.Unlock()
		}

		return data, err
	}
}

// fetch many key at once
func (l *GetAllProductsLoader) LoadAll(keys []int64) ([][]*model.Product, []error) {
	results := make([]func() ([]*model.Product, error), len(keys))

	for i, key := range keys {
		results[i] = l.LoadThunk(key)
	}

	products := make([][]*model.Product, len(keys))
	errors := make([]error, len(keys))
	for i, thunk := range results {
		products[i], errors[i] = thunk()
	}

	return products, errors
}

// wait block function
func (l *GetAllProductsLoader) LoadAllThunk(keys []int64) (func() ([][]*model.Product, []error)) {
	results := make([]func() ([]*model.Product, error), len(keys))
	for i, key := range keys {
		results[i] = l.LoadThunk(key)
	}

	return func() ([][]*model.Product, []error) {
		products := make([][]*model.Product, len(keys))
		errors := make([]error, len(keys))
		for i, thunk := range results {
			products[i], errors[i] = thunk()
		}
		return products, errors
	}
}

// prime the cache provide key and data
func (l *GetAllProductsLoader) Prime(key int64, data []*model.Product) bool {
	l.mu.Lock()
	var found bool
	if _, found = l.cache[key]; found {
		cpy := make([]*model.Product, len(data))
		l.unsafeSet(key, cpy)
	}
	l.mu.Unlock()
	return found
}

// cache clear if exists
func (l *GetAllProductsLoader) Clear(key int64) {
	l.mu.Lock()
	delete(l.cache, key)
	l.mu.Unlock()
}

// set cache
func (l *GetAllProductsLoader) unsafeSet(key int64, data []*model.Product) {
	if l.cache == nil {
		l.cache = map[int64][]*model.Product{}
	}
	l.cache[key] = data
}

// get key index
func (b *GetAllProductsBatch) keyIndex(l *GetAllProductsLoader, key int64) int {
	for i, exitKey := range b.keys {
		if exitKey == key { return i }
	}

	pos := len(b.keys)
	b.keys = append(b.keys, key)
	if pos == 0 {
		go b.startTimer(l)
	}

	if l.maxBatch != 0 && pos >= l.maxBatch-1 {
		if !b.closing {
			b.closing = true
			l.batch = nil
			go b.end(l)
		}
	}

	return pos
}


func (b *GetAllProductsBatch) startTimer(l *GetAllProductsLoader) {
	time.Sleep(l.wait)
	l.mu.Lock()

	// batch already finalize
	if b.closing {
		l.mu.Unlock()
		return
	}

	l.batch = nil
	l.mu.Unlock()

	b.end(l)
}

func (b *GetAllProductsBatch) end(l *GetAllProductsLoader) {
	b.data, b.error = l.fetch(b.keys)
	close(b.done)
}
