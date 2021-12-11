package internal

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

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
	return NewErrorf(ErrorCodeInvalidArgument, "unknown value")
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
		return NewErrorf(ErrorCodeInvalidArgument, "start date should be before due date")
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
	if err := validation.ValidateStruct(&t,
		validation.Field(&t.Description, validation.Required),
		validation.Field(&t.Priority),
		validation.Field(&t.Dates),
	); err != nil {
		return WrapErrorf(err, ErrorCodeInvalidArgument, "invalid values")
	}

	return nil
}
