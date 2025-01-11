package utils

import "sync"

// In Memory Mocks

type AutoIncrementInt struct {
	sync.Mutex
	id int
}

func (a *AutoIncrementInt) ID() (id int) {
	a.Lock()
	defer a.Unlock()

	id = a.id
	a.id++
	return
}
