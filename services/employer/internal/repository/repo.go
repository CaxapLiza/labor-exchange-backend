package repository

import (
	"github.com/student/labor-exchange-backend/services/common"
	"github.com/student/labor-exchange-backend/services/employer/internal"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList() ([]internal.Employer, error) {
	query := "SELECT * FROM employer"
	rows, err := ir.DB.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var employers []internal.Employer
	for rows.Next() {
		var employer internal.Employer
		if err := rows.Scan(&employer.ID, &employer.Name, &employer.Specialty, &employer.AccountID); err != nil {
			return nil, err
		}
		employers = append(employers, employer)
	}

	return employers, nil
}

func (ir *Repository) Get(id int) (*internal.Employer, error) {
	query := "SELECT * FROM employer WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var employer internal.Employer
	if err := row.Scan(&employer.ID, &employer.Name, &employer.Specialty, &employer.AccountID); err != nil {
		return nil, err
	}

	return &employer, nil
}

func (ir *Repository) Create(newEmployer *internal.Employer) error {
	query := "INSERT INTO employer (name, specialty, account_id) VALUES ($1, $2, $3) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newEmployer.Name, newEmployer.Specialty, newEmployer.AccountID).Scan(&newEmployer.ID)
	return err
}

func (ir *Repository) Update(id int, name string, specialty string, accountId int) error {
	query := "UPDATE employer SET name = $1, specialty = $2, account_id = $3 WHERE id = $4"
	_, err := ir.DB.Connection.Exec(query, name, specialty, accountId, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM employer WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
