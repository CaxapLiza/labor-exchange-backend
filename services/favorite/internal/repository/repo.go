package repository

import (
	"github.com/student/labor-exchange-backend/services/common"
	"github.com/student/labor-exchange-backend/services/favorite/internal"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList(id int) ([]internal.Favorite, error) {
	query := "SELECT * FROM favorite WHERE employer_id = $1"
	rows, err := ir.DB.Connection.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var favorites []internal.Favorite
	for rows.Next() {
		var favorite internal.Favorite
		if err := rows.Scan(&favorite.ID, &favorite.JobSeekerID, &favorite.EmployerID); err != nil {
			return nil, err
		}
		favorites = append(favorites, favorite)
	}

	return favorites, nil
}

func (ir *Repository) Create(newFavorite *internal.Favorite) error {
	query := "INSERT INTO favorite (job_seeker_id, employer_id) VALUES ($1, $2) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newFavorite.JobSeekerID, newFavorite.EmployerID).Scan(&newFavorite.ID)
	return err
}

func (ir *Repository) Update(id int, jobSeekerId int, employerId int) error {
	query := "UPDATE favorite SET job_seeker_id = $1, employer_id = $2 WHERE id = $3"
	_, err := ir.DB.Connection.Exec(query, jobSeekerId, employerId, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM favorite WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}
