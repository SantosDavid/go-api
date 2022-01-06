package storage

import (
	"errors"
	"time"
)

type Inmemory struct {
	Items map[string]string
}

func (i *Inmemory) Set(key string, value string, expiresAt time.Duration) error {
	i.Items[key] = value

	return nil
}

func (i *Inmemory) Get(key string) (string, error) {
	value, exists := i.Items[key]

	if exists {
		return value, nil
	}

	return "", errors.New("not found")
}
