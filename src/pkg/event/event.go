package event

import (
	"database/sql"
	"fmt"
	"time"
)

type Event struct {
	ID         int
	DB         *sql.DB
	Name       string    `json:"name"`
	Host       string    `json:"host"`
	Location   string    `json:"location"`
	Start      time.Time `json:"start"`
	End        time.Time `json:"end"`
	DressCode  string    `json:"dress-code"`
	Theme      string    `json:"theme"`
	Price      float64   `json:"price"`
	SignupLink string    `json:"signup-link"`
}

// Insert inserts event and all of its fields to database where event ID is created.
// It returns the created ID and an error.
func (e *Event) Insert() (int, error) {
	query := `
		INSERT INTO events (name, host, start, end, dc, theme, price, signup)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := e.DB.Exec(query, e.Name, e.Host, e.Start, e.End, e.DressCode, e.Theme, e.Price, e.SignupLink)
	if err != nil {
		return 0, fmt.Errorf("failed to execute insert query: %v", err)
	}

	eventID, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to retrieve last insert ID: %v", err)
	}

	e.ID = int(eventID)
	return e.ID, nil
}
