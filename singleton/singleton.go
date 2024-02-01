package singleton

import "sync"

type singleton struct {
	Name    string
	Surname string
}

var (
	once     sync.Once
	instance *singleton
)

func Get() *singleton {
	once.Do(func() {
		instance = &singleton{}
	})
	return instance
}
