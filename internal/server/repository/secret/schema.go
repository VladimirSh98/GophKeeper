package secret

import (
	"encoding/json"
	"time"
)

// DataType type enum
type DataType int

const (
	LoginPassword DataType = iota
	TextData
	BinaryData
	BankCard
)

// Secret struct
type Secret struct {
	ID        int
	UserID    int
	Archived  bool
	CreatedAt time.Time
	UpdatedAt time.Time
	DataType  DataType
	Content   []byte
	Metadata  json.RawMessage
}
