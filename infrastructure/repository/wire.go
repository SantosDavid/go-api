package repository

import "github.com/google/wire"

var Wired = wire.NewSet(
	New,
)
