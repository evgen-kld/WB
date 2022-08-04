package server

import (
	"Task2/develop/11/pkg/calendar"
	"time"
)

type EventModel struct {
	ID          string `json:"id"`
	Date        string `json:"date"`
	UserID      string `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (e EventModel) Validate() bool {
	if _, ok := time.Parse("2006-01-02", e.Date); ok != nil {
		return false
	}
	return !(e.UserID == "" || e.Title == "")
}

func (e *EventModel) ToEvent() *calendar.Event {
	event := calendar.NewEvent()
	event.ID = e.ID
	event.Title = e.Title
	event.UserID = e.UserID
	event.Description = e.Description
	event.Date, _ = time.Parse("2006-01-02", e.Date)
	return event
}
