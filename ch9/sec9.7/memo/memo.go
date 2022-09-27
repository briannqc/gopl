package memo

import "sync"

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

type Memo struct {
	f     Func
	mu    sync.Mutex
	cache map[string]*entry
}

type Func func(string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{
		f:     f,
		cache: make(map[string]*entry),
	}
}

func (m *Memo) Get(key string) (interface{}, error) {
	m.mu.Lock()
	e := m.cache[key]

	if e == nil {
		// This is the first request for this key.
		// This goroutine becomes responsible for computing
		// the value abd broadcasting the ready condition.
		e = &entry{
			ready: make(chan struct{}),
		}
		m.cache[key] = e
		m.mu.Unlock()

		e.res.value, e.res.err = m.f(key)
		close(e.ready)
	} else {
		m.mu.Unlock()
		<-e.ready
	}
	return e.res.value, e.res.err
}

type request struct {
	key      string
	response chan result
}

// V2 uses a monitoring goroutine to cache and retrieve cached values
type V2 struct {
	requests chan request
}

func NewV2(f Func) *V2 {
	m := &V2{requests: make(chan request)}
	go m.server(f)
	return m
}

func (m *V2) Close() {
	close(m.requests)
}

func (m *V2) Get(key string) (interface{}, error) {
	req := request{
		key:      key,
		response: make(chan result),
	}
	m.requests <- req
	resp := <-req.response
	return resp.value, resp.err
}

func (m *V2) server(f Func) {
	cache := make(map[string]*entry)
	for req := range m.requests {
		e := cache[req.key]
		if e == nil {
			e = &entry{ready: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key)
		}
		go e.deliver(req.response)
	}
}

func (e *entry) call(f Func, key string) {
	e.res.value, e.res.err = f(key)
	close(e.ready)
}

func (e *entry) deliver(response chan result) {
	<-e.ready
	response <- e.res
}
