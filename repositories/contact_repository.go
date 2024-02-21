package repositories

import (
	"context"
	"database/sql"

	"github.com/nozzlium/fiber_nvim_crud/entities"
	"github.com/nozzlium/fiber_nvim_crud/params"
)

type ContactRepository interface {
	Create(tx *sql.Tx, contact entities.Contact) (entities.Contact, error)
	Find(ctx context.Context, db *sql.DB, params params.Contact) ([]entities.Contact, error)
	FindById(ctx context.Context, db *sql.DB, params params.Contact) (entities.Contact, error)
	Edit(tx *sql.Tx, contact entities.Contact) (entities.Contact, error)
	Delete(tx *sql.Tx, contact entities.Contact) (entities.Contact, error)
}
