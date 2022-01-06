package storage

import "time"

type Storage interface {
	Set(key string, value string, expiresAt time.Duration) error
	Get(key string) (string, error)
}
