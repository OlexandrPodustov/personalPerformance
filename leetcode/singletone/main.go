package main

import "sync"

type connector struct{}

var (
	lock     sync.Mutex
	instance *connector
)

func GetInstance() *connector {
	if instance == nil {
		lock.Lock()
		if instance == nil {
			instance = &connector{}
		}
		lock.Unlock()
	}

	return instance
}
