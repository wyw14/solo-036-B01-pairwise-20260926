package model

import (
	"github.com/google/uuid"
)

func NewRevision() string { return uuid.NewString() }
