package repository

import (
	"github.com/student/labor-exchange-backend/services/common"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetVacancyResponsesCount() (map[int]int, error) {
	query := "SELECT vacancy_id, COUNT(*) FROM response GROUP BY vacancy_id"
	rows, err := ir.DB.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	vacancyResponsesCount := make(map[int]int)
	for rows.Next() {
		var vacancyID, count int
		if err := rows.Scan(&vacancyID, &count); err != nil {
			return nil, err
		}
		vacancyResponsesCount[vacancyID] = count
	}

	return vacancyResponsesCount, nil
}

func (ir *Repository) GetEmployerResponsesCount() (map[int]int, error) {
	query := `
		SELECT e.id AS employer_id, COUNT(*) AS responses_count
		FROM response r
		JOIN vacancy v ON r.vacancy_id = v.id
		JOIN employer e ON v.employer_id = e.id
		GROUP BY e.id;
	`
	rows, err := ir.DB.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	employerResponsesCount := make(map[int]int)
	for rows.Next() {
		var employerID, count int
		if err := rows.Scan(&employerID, &count); err != nil {
			return nil, err
		}
		employerResponsesCount[employerID] = count
	}

	return employerResponsesCount, nil
}

func (ir *Repository) GetJobSeekerFavoritesCount() (map[int]int, error) {
	query := "SELECT job_seeker_id, COUNT(*) FROM favorite GROUP BY job_seeker_id"
	rows, err := ir.DB.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	jobSeekerFavoritesCount := make(map[int]int)
	for rows.Next() {
		var jobSeekerID, count int
		if err := rows.Scan(&jobSeekerID, &count); err != nil {
			return nil, err
		}
		jobSeekerFavoritesCount[jobSeekerID] = count
	}

	return jobSeekerFavoritesCount, nil
}
