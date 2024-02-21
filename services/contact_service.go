package services

import (
	"context"

	"github.com/nozzlium/fiber_nvim_crud/params"
	"github.com/nozzlium/fiber_nvim_crud/responsebody"
)

type ContactService interface {
	Create(ctx context.Context, param params.Contact) (responsebody.Contact, error)
	Find(ctx context.Context, param params.Contact) (responsebody.Contacts, error)
	FindById(ctx context.Context, param params.Contact) (responsebody.Contact, error)
	Update(ctx context.Context, param params.Contact) (responsebody.Contact, error)
	Delete(ctx context.Context, param params.Contact) (responsebody.Contact, error)
}
