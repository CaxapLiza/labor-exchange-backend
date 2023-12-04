package repository

import (
	"github.com/student/labor-exchange-backend/services/account/internal"
	"github.com/student/labor-exchange-backend/services/common"
)

type Repository struct {
	DB *common.Database
}

func NewRepository(db *common.Database) *Repository {
	return &Repository{DB: db}
}

func (ir *Repository) GetList() ([]internal.Account, error) {
	query := "SELECT * FROM account"
	rows, err := ir.DB.Connection.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var accounts []internal.Account
	for rows.Next() {
		var account internal.Account
		if err := rows.Scan(&account.ID, &account.Login, &account.Password, &account.Role); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (ir *Repository) Get(id int) (*internal.Account, error) {
	query := "SELECT * FROM account WHERE id = $1"
	row := ir.DB.Connection.QueryRow(query, id)

	var account internal.Account
	if err := row.Scan(&account.ID, &account.Login, &account.Password, &account.Role); err != nil {
		return nil, err
	}

	return &account, nil
}

func (ir *Repository) Create(newAccount *internal.Account) error {
	query := "INSERT INTO account (login, password, role) VALUES ($1, $2, $3) RETURNING id"
	err := ir.DB.Connection.QueryRow(query, newAccount.Login, newAccount.Password, newAccount.Role).Scan(&newAccount.ID)
	return err
}

func (ir *Repository) Update(id int, login string, password string, role string) error {
	query := "UPDATE account SET login = $1, password = $2, role = $3 WHERE id = $4"
	_, err := ir.DB.Connection.Exec(query, login, password, role, id)
	return err
}

func (ir *Repository) Delete(id int) error {
	query := "DELETE FROM account WHERE id = $1"
	_, err := ir.DB.Connection.Exec(query, id)
	return err
}

func (ir *Repository) Authenticate(login, password string) (*internal.Account, error) {
	query := "SELECT id, role FROM account WHERE login = $1 AND password = $2"
	row := ir.DB.Connection.QueryRow(query, login, password)

	var account internal.Account
	if err := row.Scan(&account.ID, &account.Role); err != nil {
		return nil, err
	}

	return &account, nil
}
