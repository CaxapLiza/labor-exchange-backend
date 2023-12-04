package internal

type Employer struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Specialty string `json:"specialty"`
	AccountID int    `json:"account_id"`
}
