package repositories

import (
	"context"
	"database/sql"
	"errors"

	"github.com/nozzlium/fiber_nvim_crud/entities"
	"github.com/nozzlium/fiber_nvim_crud/params"
)

type ContactRepositoryImpl struct{}

func NewContactRepository() *ContactRepositoryImpl {
	return &ContactRepositoryImpl{}
}

func (repository *ContactRepositoryImpl) Create(
	tx *sql.Tx,
	contact entities.Contact,
) (entities.Contact, error) {
	query := `
    INSERT INTO contact (
      first_name,
      last_name,
      phone,
      is_vip
    ) values (
      ?, ?, ?, ?
    )
  `
	result, err := tx.Exec(
		query,
		contact.FirstName,
		contact.LastName,
		contact.Phone,
		contact.IsVip,
	)
	if err != nil {
		return entities.Contact{}, err
	}

	userId, err := result.LastInsertId()
	contact.ID = uint(userId)
	return contact, err
}

func (repository *ContactRepositoryImpl) Find(
	ctx context.Context,
	db *sql.DB,
	params params.Contact,
) ([]entities.Contact, error) {
	query := `
    SELECT 
      id,
      first_name,
      last_name,
      phone,
      is_vip
    FROM contact
    OFFSET ?
    LIMIT ?
  `
	rows, err := db.QueryContext(
		ctx,
		query,
		(params.PageNo-1)*params.PageSize,
		params.PageSize,
	)
	if err != nil {
		return nil, err
	}
	contacts := []entities.Contact{}
	for rows.Next() {
		var id uint
		var firstName, lastName, phone string
		var isVip bool
		err := rows.Scan(
			&id,
			&firstName,
			&lastName,
			&phone,
			&isVip,
		)
		if err != nil {
			return nil, err
		}
		contacts = append(
			contacts,
			entities.Contact{
				ID:        id,
				FirstName: firstName,
				LastName:  lastName,
				Phone:     phone,
				IsVip:     isVip,
			},
		)
	}
	defer rows.Close()
	return contacts, nil
}

func (repository *ContactRepositoryImpl) FindById(
	ctx context.Context,
	db *sql.DB,
	params params.Contact,
) (entities.Contact, error) {
	query := `
    SELECT
      id,
      first_name,
      last_name,
      phone,
      is_vip
    FROM contact
    WHERE id = ?
    LIMIT 1
  `
	row := db.QueryRowContext(
		ctx,
		query,
		params.Contact.ID,
	)
	var id uint
	var firstName, lastName, phone string
	var isVip bool
	err := row.Scan(
		&id,
		&firstName,
		&lastName,
		&phone,
		&isVip,
	)
	return entities.Contact{
		ID:        id,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		IsVip:     isVip,
	}, err
}

func (repository *ContactRepositoryImpl) Edit(
	tx *sql.Tx,
	contact entities.Contact,
) (entities.Contact, error) {
	query := `
    UPDATE contact
    SET
      first_name = ?,
      last_name = ?,
      phone = ?,
      is_vip = ?
    WHERE id = ?
  `
	result, err := tx.Exec(
		query,
		contact.FirstName,
		contact.LastName,
		contact.Phone,
		contact.IsVip,
	)
	if err != nil {
		return entities.Contact{}, err
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		return entities.Contact{}, err
	}
	if rowCount != 1 {
		return entities.Contact{}, errors.New("unknown error, update unsuccessful")
	}
	return contact, nil
}

func (repository *ContactRepositoryImpl) Delete(
	tx *sql.Tx,
	contact entities.Contact,
) (entities.Contact, error) {
	query := `
    DELETE 
    FROM contact
    WHERE id = ?
  `
	result, err := tx.Exec(
		query,
		contact.ID,
	)
	if err != nil {
		return entities.Contact{}, err
	}
	rowCount, err := result.RowsAffected()
	if err != nil {
		return entities.Contact{}, err
	}
	if rowCount != 1 {
		return entities.Contact{}, errors.New("unknown error, deletion unsuccessful")
	}
	return contact, nil
}
