package services

import (
	"context"
	"database/sql"

	"github.com/nozzlium/fiber_nvim_crud/entities"
	"github.com/nozzlium/fiber_nvim_crud/libs"
	"github.com/nozzlium/fiber_nvim_crud/params"
	"github.com/nozzlium/fiber_nvim_crud/repositories"
	"github.com/nozzlium/fiber_nvim_crud/responsebody"
)

type ContactServiceImpl struct {
	ContactRepository repositories.ContactRepository
	DB                *sql.DB
}

func NewContactService(
	contactRepository repositories.ContactRepository,
	db *sql.DB,
) *ContactServiceImpl {
	return &ContactServiceImpl{
		ContactRepository: contactRepository,
		DB:                db,
	}
}

func (service *ContactServiceImpl) Create(
	ctx context.Context,
	param params.Contact,
) (responsebody.Contact, error) {
	// tx, err := service.DB.BeginTx(
	// 	ctx,
	// 	nil,
	// )
	// if err != nil {
	// 	err = tx.Rollback()
	// 	return responsebody.Contact{}, err
	// }
	// err = tx.Commit()
	// if err != nil {
	// 	return responsebody.Contact{}, err
	// }

	contact := param.Contact
	result, err := service.ContactRepository.Create(nil, entities.Contact{
		FirstName: contact.FirstName,
		LastName:  contact.LastName,
		Phone:     contact.Phone,
		IsVip:     contact.IsVip,
	})
	return libs.ContactEntityToResponse(result), err
}

func (service *ContactServiceImpl) Find(
	ctx context.Context,
	param params.Contact,
) (responsebody.Contacts, error) {
	contacts, err := service.ContactRepository.Find(
		ctx,
		service.DB,
		param,
	)
	return responsebody.Contacts{
		PageNo:   param.PageNo,
		PageSize: uint(len(contacts)),
		Contacts: libs.ContactEntitiesToResponses(contacts),
	}, err
}

func (service *ContactServiceImpl) FindById(
	ctx context.Context,
	param params.Contact,
) (responsebody.Contact, error) {
	contact, err := service.ContactRepository.FindById(
		ctx,
		service.DB,
		param,
	)
	return libs.ContactEntityToResponse(contact), err
}

func (service *ContactServiceImpl) Update(
	ctx context.Context,
	param params.Contact,
) (responsebody.Contact, error) {
	saved, err := service.ContactRepository.FindById(
		ctx,
		service.DB,
		param,
	)
	if err != nil {
		return responsebody.Contact{}, err
	}
	if param.Contact.FirstName != "" {
		saved.FirstName = param.Contact.FirstName
	}
	if param.Contact.LastName != "" {
		saved.LastName = param.Contact.LastName
	}
	if param.Contact.Phone != "" {
		saved.Phone = param.Contact.Phone
	}
	// tx, err := service.DB.BeginTx(
	// 	ctx,
	// 	nil,
	// )
	// if err != nil {
	// 	return responsebody.Contact{}, err
	// }
	result, err := service.ContactRepository.Edit(nil, saved)
	// if err != nil {
	// 	err = tx.Rollback()
	// 	return responsebody.Contact{}, err
	// }
	// err = tx.Commit()
	// if err != nil {
	// 	return responsebody.Contact{}, err
	// }
	return libs.ContactEntityToResponse(result), err
}

func (service *ContactServiceImpl) Delete(
	ctx context.Context,
	param params.Contact,
) (responsebody.Contact, error) {
	// tx, err := service.DB.BeginTx(ctx, nil)
	// if err != nil {
	// 	return responsebody.Contact{}, err
	// }
	result, err := service.ContactRepository.Delete(nil, param.Contact)
	// if err != nil {
	// 	err = tx.Rollback()
	// 	return responsebody.Contact{}, err
	// }
	// err = tx.Commit()
	if err != nil {
		return responsebody.Contact{}, err
	}
	return libs.ContactEntityToResponse(result), err
}
