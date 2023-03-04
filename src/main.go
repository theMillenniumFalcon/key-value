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
