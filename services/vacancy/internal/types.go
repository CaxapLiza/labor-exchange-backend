package internal

type Vacancy struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Salary      string `json:"salary"`
	Description string `json:"description"`
	EmployerID  int    `json:"employer_id"`
}
