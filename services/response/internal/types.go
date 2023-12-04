package internal

type Response struct {
	ID          int `json:"id"`
	JobSeekerID int `json:"job_seeker_id"`
	VacancyID   int `json:"vacancy_id"`
}
