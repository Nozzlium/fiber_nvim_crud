package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/nozzlium/fiber_nvim_crud/entities"
	"github.com/nozzlium/fiber_nvim_crud/params"
)

var contacts = []entities.Contact{}
var id uint = 0

type ContactRepositoryMock struct {
}

func NewContactRepositoryMock() *ContactRepositoryMock {
	return &ContactRepositoryMock{}
}

func (repository *ContactRepositoryMock) Create(
	tx *sql.Tx,
	contact entities.Contact,
) (entities.Contact, error) {
	id++
	contact.ID = id
	contacts = append(contacts, contact)
	return contact, nil
}

func (repository *ContactRepositoryMock) Find(
	ctx context.Context,
	db *sql.DB,
	params params.Contact,
) ([]entities.Contact, error) {
	index := (params.PageNo - 1) * params.PageSize
	if index >= uint(len(contacts)) {
		return []entities.Contact{}, nil
	}
	max := (index + params.PageSize)
	if max >= uint(len(contacts)) {
		max = uint(len(contacts))
	}
	return contacts[index:max], nil
}

func (repository *ContactRepositoryMock) FindById(
	ctx context.Context,
	db *sql.DB,
	params params.Contact,
) (entities.Contact, error) {
	for _, contact := range contacts {
		if contact.ID == params.Contact.ID {
			return contact, nil
		}
	}
	return entities.Contact{}, errors.New("not found")
}

func (repository *ContactRepositoryMock) Edit(
	tx *sql.Tx,
	contact entities.Contact,
) (entities.Contact, error) {
	for i, cont := range contacts {
		if cont.ID == contact.ID {
			contacts = append(append(contacts[0:i], contact), contacts[i+1:]...)
			return contact, nil
		}
	}
	return entities.Contact{}, errors.New("not found")
}

func (repository *ContactRepositoryMock) Delete(
	tx *sql.Tx,
	contact entities.Contact,
) (entities.Contact, error) {
	for i, cont := range contacts {
		if cont.ID == contact.ID {
			contacts = append(contacts[0:i], contacts[i+1:len(contacts)]...)
			return cont, nil
		}
	}
	return entities.Contact{}, errors.New("not found")
}
