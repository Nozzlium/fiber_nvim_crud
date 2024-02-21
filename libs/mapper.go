package libs

import (
	"github.com/nozzlium/fiber_nvim_crud/entities"
	"github.com/nozzlium/fiber_nvim_crud/responsebody"
)

func ContactEntityToResponse(entity entities.Contact) responsebody.Contact {
	return responsebody.Contact{
		ID:        entity.ID,
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
		Phone:     entity.Phone,
		IsVip:     entity.IsVip,
	}
}

func ContactEntitiesToResponses(entities []entities.Contact) []responsebody.Contact {
	responses := []responsebody.Contact{}
	for _, entity := range entities {
		responses = append(responses, ContactEntityToResponse(entity))
	}
	return responses
}
