package postgresql

import (
	"context"

	"github.com/terrytay/godo/internal"
	"github.com/terrytay/godo/internal/postgresql/db"
)

// Task represents the repository used for interacting with Task records
type Task struct {
	q *db.Queries
}

// NewTask instantiates the Task Repository
func NewTask(d db.DBTX) *Task {
	return &Task{
		q: db.New(d),
	}
}

// Create inserts a new Task record
func (t *Task) Create(ctx context.Context, params internal.CreateParams) (internal.Task, error) {

	newID, err := t.q.InsertTask(ctx, db.InsertTaskParams{
		Description: params.Description,
		Priority:    newPriority(params.Priority),
		StartDate:   newNullTime(params.Dates.Start),
		DueDate:     newNullTime(params.Dates.Due),
	})

	if err != nil {
		return internal.Task{}, internal.WrapErrorf(err, internal.ErrorCodeUnknown, "insert task")
	}

	return internal.Task{
		ID:          newID.String(),
		Description: params.Description,
		Priority:    params.Priority,
		Dates:       params.Dates,
	}, nil

}
