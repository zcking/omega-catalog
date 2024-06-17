package internal

import (
	omega_catalogv1 "github.com/zcking/omega-catalog/gen/go/omega_catalog/v1"
	"github.com/zcking/omega-catalog/internal/database"
)

type Impl struct {
	omega_catalogv1.UnimplementedOmegaCatalogServiceServer

	db *database.Database
}

func NewImpl(db *database.Database) (*Impl, error) {
	i := &Impl{
		db: db,
	}
	return i, nil
}

func (i *Impl) Close() error {
	return i.db.Close()
}
