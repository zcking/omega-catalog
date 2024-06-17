package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Catalog struct {
	ID        uuid.UUID `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4()"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime `gorm:"index"`

	Name       string    `json:"name" gorm:"primaryKey"`
	Comment    *string   `json:"comment"`
	Properties StringMap `json:"properties" gorm:"type:jsonb"`
}

func (Catalog) TableName() string {
	return "catalogs"
}

type StringMap map[string]string

// Scan implements the Scanner interface.
func (m *StringMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, m)
}

// Value implements the Valuer interface.
func (m StringMap) Value() (driver.Value, error) {
	return json.Marshal(m)
}
