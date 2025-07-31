package entity

import "time"

type Experience struct {
	ID          int       `json:"id"`
	CompanyName string    `json:"company_name"`
	PhotoURL    *string   `json:"photo_url,omitempty"`
	Description string    `json:"description"`
	Role        string    `json:"role"`
	CategoryID  *int      `json:"category_id,omitempty"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}
