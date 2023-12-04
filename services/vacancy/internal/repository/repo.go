package repository

import (
	"github.com/student/labor-exchange-backend/services/common"
	"github.com/student/labor-exchange-backend/services/vacancy/internal"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList() ([]internal.Vacancy, error) {
	query := "SELECT * FROM vacancy"
	rows, err := ir.DB.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var vacancies []internal.Vacancy
	for rows.Next() {
		var vacancy internal.Vacancy
		if err := rows.Scan(&vacancy.ID, &vacancy.Title, &vacancy.Salary,
			&vacancy.Description, &vacancy.EmployerID); err != nil {
			return nil, err
		}
		vacancies = append(vacancies, vacancy)
	}

	return vacancies, nil
}

func (ir *Repository) Get(id int) (*internal.Vacancy, error) {
	query := "SELECT * FROM vacancy WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var vacancy internal.Vacancy
	if err := row.Scan(&vacancy.ID, &vacancy.Title, &vacancy.Salary,
		&vacancy.Description, &vacancy.EmployerID); err != nil {
		return nil, err
	}

	return &vacancy, nil
}

func (ir *Repository) Create(newVacancy *internal.Vacancy) error {
	query := "INSERT INTO vacancy (title, salary, description, employer_id) VALUES ($1, $2, $3) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newVacancy.Title, newVacancy.Salary, newVacancy.Description, newVacancy.EmployerID).Scan(&newVacancy.ID)
	return err
}

func (ir *Repository) Update(id int, title string, salary string, description string, employerId int) error {
	query := "UPDATE vacancy SET title = $1, salary = $2, description = $3, employer_id = $4 WHERE id = $5"
	_, err := ir.DB.Connection.Exec(query, title, salary, description, employerId, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM vacancy WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
