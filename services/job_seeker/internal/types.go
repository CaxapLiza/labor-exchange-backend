package internal

import "time"

type JobSeeker struct {
	ID            int       `json:"id"`
	FullName      string    `json:"full_name"`
	BirthDate     time.Time `json:"birth_date"`
	ContactDetail string    `json:"contact_details"`
	Position      string    `json:"position"`
	Education     string    `json:"education"`
	Experience    string    `json:"experience"`
	AccountID     int       `json:"account_id"`
}
