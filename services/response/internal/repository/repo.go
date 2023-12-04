package repository

import (
	"github.com/student/labor-exchange-backend/services/common"
	"github.com/student/labor-exchange-backend/services/response/internal"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList(id int) ([]internal.Response, error) {
	query := "SELECT * FROM response WHERE job_seeker_id = $1"
	rows, err := ir.DB.Connection.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responses []internal.Response
	for rows.Next() {
		var response internal.Response
		if err := rows.Scan(&response.ID, &response.JobSeekerID, &response.VacancyID); err != nil {
			return nil, err
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (ir *Repository) Create(newResponse *internal.Response) error {
	query := "INSERT INTO response (job_seeker_id, vacancy_id) VALUES ($1, $2) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newResponse.JobSeekerID, newResponse.VacancyID).Scan(&newResponse.ID)
	return err
}

func (ir *Repository) Update(id int, jobSeekerId int, employerId int) error {
	query := "UPDATE response SET job_seeker_id = $1, vacancy_id = $2 WHERE id = $3"
	_, err := ir.DB.Connection.Exec(query, jobSeekerId, employerId, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM response WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
