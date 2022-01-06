package storage_test

import (
	"testing"
	"time"

	"github.com/santosdavid/go-api-v2/infrastructure/ui/gin/auth/storage"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T) {
	storage := storage.Inmemory{make(map[string]string)}

	err := storage.Set("1", "test", time.Duration(1))

	assert.Nil(t, err)
	assert.Equal(t, "test", storage.Items["1"])
}

func TestGetError(t *testing.T) {
	storage := storage.Inmemory{map[string]string{}}

	_, err := storage.Get("1")

	assert.Error(t, err)
}

func TestGet(t *testing.T) {
	storage := storage.Inmemory{map[string]string{"1": "test"}}

	value, err := storage.Get("1")

	assert.Nil(t, err)
	assert.Equal(t, "test", value)
}
