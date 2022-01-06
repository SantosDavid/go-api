package domain

import "github.com/google/uuid"

type ID = uuid.UUID

func New() ID {
	return ID(uuid.New())
}
