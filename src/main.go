package main

import (
	"sync"
	"time"

	"github.com/syndtr/goleveldb/leveldb"
)

// App struct and methods
type App struct {
	db    *leveldb.DB
	mlock sync.Mutex
	lock  map[string]struct{}

	// params
	uploadIDs  map[string]bool
	volumes    []string
	fallback   string
	replicas   int
	subVolumes int
	protect    bool
	md5sum     bool
	volTimeout time.Duration
}

func (a *App) UnlockKey(key []byte) {
	a.mlock.Lock()
	defer a.mlock.Unlock()
	delete(a.lock, string(key))
}

func (a *App) LockKey(key []byte) bool {
	a.mlock.Lock()
	defer a.mlock.Unlock()
	if _, prs := a.lock[string(key)]; prs {
		return false
	}
	a.lock[string(key)] = struct{}{}

	return true
}
