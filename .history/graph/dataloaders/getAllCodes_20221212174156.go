package dataloaders

import (
	"sync"
	"time"

	"github.com/sRRRs-7/loose_style.git/graph/model"
)

type GetAllCodesConfig struct {
	// Fetch is a method to provide the data for the loader
	Fetch func(limit []int64) ([][]*model.Code, []error)
	// batch wait time
	Wait time.Duration
	// maximum number of keys to send in one batch, 0 = no limit
	MaxBatch int
}

type GetAllCodesLoader struct {
	// provide the data for the loader
	fetch func(limit []int64) ([][]*model.Code, []error)
	// batch wait time
	wait time.Duration
	// maximum number of keys to send in one batch, 0 = no limit
	maxBatch int
	// create cache
	cache map[int64][]*model.Code
	// execute batch
	batch *GetAllCodesBatch
	// prevent race
	mu sync.Mutex
}

type GetAllCodesBatch struct {
	keys    []int64
	data    [][]*model.Code
	error   []error
	closing bool
	done    chan struct{}
}

func NewGetAllCodesLoader(config GetAllCodesConfig) *GetAllCodesLoader {
	return &GetAllCodesLoader{
		fetch:    config.Fetch,
		wait:     config.Wait,
		maxBatch: config.MaxBatch,
	}
}

// which will call method ? (LoadThunk, LoadAll, LoadAllThunk, Prime, Clear)
func (l *GetAllCodesLoader) Load(key int64) ([]*model.Code, error) {
	return l.LoadThunk(key)()
}

// LoadThunk is function that block wait for a GetProducts
func (l *GetAllCodesLoader) LoadThunk(key int64) func() ([]*model.Code, error) {
	l.mu.Lock()
	if it, ok := l.cache[key]; ok {
		l.mu.Unlock()
		return func() ([]*model.Code, error) {
			return it, nil
		}
	}

	if l.batch == nil {
		l.batch = &GetAllCodesBatch{done: make(chan struct{})}
	}

	batch := l.batch
	pos := batch.keyIndex(l, key)
	l.mu.Unlock()

	return func() ([]*model.Code, error) {
		<-batch.done

		var data []*model.Code
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
func (l *GetAllCodesLoader) LoadAll(keys []int64) ([][]*model.Code, []error) {
	results := make([]func() ([]*model.Code, error), len(keys))

	for i, key := range keys {
		results[i] = l.LoadThunk(key)
	}

	products := make([][]*model.Code, len(keys))
	errors := make([]error, len(keys))
	for i, thunk := range results {
		products[i], errors[i] = thunk()
	}

	return products, errors
}

// wait block function
func (l *GetAllCodesLoader) LoadAllThunk(keys []int64) func() ([][]*model.Code, []error) {
	results := make([]func() ([]*model.Code, error), len(keys))
	for i, key := range keys {
		results[i] = l.LoadThunk(key)
	}

	return func() ([][]*model.Code, []error) {
		products := make([][]*model.Code, len(keys))
		errors := make([]error, len(keys))
		for i, thunk := range results {
			products[i], errors[i] = thunk()
		}
		return products, errors
	}
}

// prime the cache provide key and data
func (l *GetAllCodesLoader) Prime(key int64, data []*model.Code) bool {
	l.mu.Lock()
	var found bool
	if _, found = l.cache[key]; found {
		cpy := make([]*model.Code, len(data))
		l.unsafeSet(key, cpy)
	}
	l.mu.Unlock()
	return found
}

// cache clear if exists
func (l *GetAllCodesLoader) Clear(key int64) {
	l.mu.Lock()
	delete(l.cache, key)
	l.mu.Unlock()
}

// set cache
func (l *GetAllCodesLoader) unsafeSet(key int64, data []*model.Code) {
	if l.cache == nil {
		l.cache = map[int64][]*model.Code{}
	}
	l.cache[key] = data
}

// get key index
func (b *GetAllCodesBatch) keyIndex(l *GetAllCodesLoader, key int64) int {
	for i, exitKey := range b.keys {
		if exitKey == key {
			return i
		}
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

func (b *GetAllCodesBatch) startTimer(l *GetAllCodesLoader) {
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

func (b *GetAllCodesBatch) end(l *GetAllCodesLoader) {
	b.data, b.error = l.fetch(b.keys)
	close(b.done)
}
