package internal

type Favorite struct {
	ID          int `json:"id"`
	JobSeekerID int `json:"job_seeker_id"`
	EmployerID  int `json:"employer_id"`
}
