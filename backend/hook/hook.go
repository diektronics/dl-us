package hook

import (
	"fmt"
)

type Hook struct {
	name string
	ch   chan *Data
}

type Data struct {
	Files []string
	Extra interface{}
	Ch    chan error
}

var all map[string]*Hook
var names []string
var order map[string]int

func init() {
	all = make(map[string]*Hook)
	order = make(map[string]int)
}

func New(name string, worker func(int, *Hook), nWorkers int) *Hook {
	h := &Hook{name: name, ch: make(chan *Data, 100)}
	for i := 0; i < nWorkers; i++ {
		go worker(i, h)
	}
	return h
}

func Names() []string { return names }

func All() map[string]*Hook { return all }

func Order(h string) int { return order[h] }

func Validate(h string) error {
	if _, ok := order[h]; ok {
		return nil
	} else {
		return fmt.Errorf("%v is not a valid hook", h)
	}
}

func (h *Hook) Name() string { return h.name }

func (h *Hook) Channel() chan *Data { return h.ch }
