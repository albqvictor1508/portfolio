package entity

import "time"

type Experience struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	CompanyName string    `json:"company_name"`
	Role        string    `json:"role"`
	DemoURL     string    `json:"demo_url"`
	Description string    `json:"description"`
	CategoryID  *int      `json:"category_id,omitempty"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
}
