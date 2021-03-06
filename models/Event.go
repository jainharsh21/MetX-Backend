package models

import "time"

type Event struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	ImgUrl        string    `json:"event_img_url"`
	Description   string    `json:"description"`
	Summary       string    `json:"summary"`
	EventAt       string    `json:"event_at"`
	EventAtOg     time.Time `json:"eventAtOg,omitEmpty"`
	Location      string    `json:"location"`
	Fees          int       `json:"fees"`
	StudentChapID string    `json:"student_chapter_id"`
	Attendees     []string  `json:"attendees"`
	Tags          []string  `json:"tags"`
}
