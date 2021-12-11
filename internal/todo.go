package internal

import (
	"errors"
	"time"
)

var ErrInvalidPriority = errors.New("parsed unknown value")
var ErrInvalidStartAndDueDates = errors.New("start dates should be before end dates")
var ErrDescriptionEmpty = errors.New("Task description cannot be empty")

const (
	// PriorityNone indicates the task needs to be prioritized
	PriorityNone Priority = iota

	PriorityLow

	PriorityMedium

	PriorityHigh
)

type Priority int8

// Validate validates the priority to be either none, low, medium or high
func (p Priority) Validate() error {
	switch p {
	case PriorityNone, PriorityLow, PriorityMedium, PriorityHigh:
		return nil
	}
	return ErrInvalidPriority
}

// Category is used to organize the tasks
type Category string

// Date is used to indicate the start/completion of tasks. It is not enforced on Tasks
type Dates struct {
	Start time.Time
	Due   time.Time
}

// Validation validates the start and due date to be practical/legit
func (d Dates) Validate() error {
	if !d.Start.IsZero() && !d.Due.IsZero() && d.Start.After(d.Due) {
		return ErrInvalidStartAndDueDates
	}
	return nil
}

// Task is an activity that needs to be completed within a period of time
type Task struct {
	IsDone      bool
	Priority    Priority
	ID          string
	Description string
	Dates       Dates
	SubTasks    []Task
	Categories  []Category
}

// Validate validates fields of Task struct
func (t Task) Validate() error {
	if t.Description == "" {
		return ErrDescriptionEmpty
	}

	if err := t.Priority.Validate(); err != nil {
		return err
	}

	if err := t.Dates.Validate(); err != nil {
		return err
	}

	return nil
}
