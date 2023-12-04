package repository

import (
	"github.com/student/labor-exchange-backend/services/common"
	"github.com/student/labor-exchange-backend/services/job_seeker/internal"
	"time"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList() ([]internal.JobSeeker, error) {
	query := "SELECT * FROM job_seeker"
	rows, err := ir.DB.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobSeekers []internal.JobSeeker
	for rows.Next() {
		var job_seeker internal.JobSeeker
		if err := rows.Scan(&job_seeker.ID, &job_seeker.FullName, &job_seeker.BirthDate, &job_seeker.ContactDetail,
			&job_seeker.Position, &job_seeker.Education, &job_seeker.Experience, &job_seeker.AccountID); err != nil {
			return nil, err
		}
		jobSeekers = append(jobSeekers, job_seeker)
	}

	return jobSeekers, nil
}

func (ir *Repository) Get(id int) (*internal.JobSeeker, error) {
	query := "SELECT * FROM job_seeker WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var job_seeker internal.JobSeeker
	if err := row.Scan(&job_seeker.ID, &job_seeker.FullName, &job_seeker.BirthDate, &job_seeker.ContactDetail,
		&job_seeker.Position, &job_seeker.Education, &job_seeker.Experience, &job_seeker.AccountID); err != nil {
		return nil, err
	}

	return &job_seeker, nil
}

func (ir *Repository) Create(newJobSeeker *internal.JobSeeker) error {
	query := "INSERT INTO job_seeker (full_name, birth_date, contact_details, position, education, experience, account_id) VALUES ($1, $2, $3) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newJobSeeker.FullName, newJobSeeker.BirthDate, newJobSeeker.ContactDetail,
		newJobSeeker.Position, newJobSeeker.Education, newJobSeeker.Experience, newJobSeeker.AccountID).Scan(&newJobSeeker.ID)
	return err
}

func (ir *Repository) Update(id int, fullName string, birthDate time.Time, contactDetails string,
	position string, education string, experience string, accountId int) error {
	query := "UPDATE job_seeker SET full_name = $1, birth_date = $2, contact_details = $3, position = $4, education = $5, experience = $6, account_id = $7 WHERE id = $8"
	_, err := ir.DB.Connection.Exec(query, fullName, birthDate, contactDetails, position, education, experience, accountId, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM job_seeker WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
