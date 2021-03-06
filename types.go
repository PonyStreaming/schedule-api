package main

import (
	"strings"
	"time"
)

type ErrArgumentNil struct {
	ParamName string
	FuncName  string
}

func (e ErrArgumentNil) Error() string {
	return "Argument " + e.ParamName + " cannot be nil in " + e.FuncName
}

type Event struct {
	ID          string    `json:"id"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
	Title       string    `json:"title"`
	Panelists   string    `json:"panelists"`
	Description string    `json:"description"`
	IsZoom      bool      `json:"isZoom"`
}

func (this *Event) IsValid() bool {
	return this != nil &&
		this.ID != "" &&
		!this.StartTime.IsZero() &&
		!this.EndTime.IsZero() &&
		this.Title != ""
}

func (this *Event) SharesID(other *Event) bool {
	return this != nil &&
		other != nil &&
		this.ID != "" &&
		strings.EqualFold(this.ID, other.ID)
}

func (this *Event) HasID(id string) bool {
	return this != nil &&
		id != "" &&
		this.ID != "" &&
		strings.EqualFold(this.ID, id)
}

type EventSchedule map[string][]Event

type Response struct {
	Rooms     EventSchedule `json:"rooms"`
	RoomOrder []string      `json:"roomOrder"`
	Times     []time.Time   `json:"times"`
}
