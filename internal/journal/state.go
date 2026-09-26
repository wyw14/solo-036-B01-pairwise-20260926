package journal

import (
	"github.com/wyw14/solo-036/internal/model"
	"sync"
	"time"
)

type FileRecord struct {
	Sequence int64       `json:"sequence"`
	Previous string      `json:"previous"`
	Checksum string      `json:"checksum"`
	Event    model.Event `json:"event"`
}

type FileSnapshot struct {
	Version      int                   `json:"version"`
	WrittenAt    time.Time             `json:"written_at"`
	LastSequence int64                 `json:"last_sequence"`
	LastChecksum string                `json:"last_checksum"`
	Control      model.ControlSnapshot `json:"control"`
}

type FileStore struct {
	mu           sync.Mutex
	directory    string
	eventsPath   string
	snapshotPath string
	sequence     int64
	checksum     string
}
