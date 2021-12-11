package internal_test

import (
	"testing"
	"time"

	"github.com/terrytay/godo/internal"
)

func TestPriority_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   internal.Priority
		withErr bool
	}{
		{
			"OK: PriorityNone",
			internal.PriorityNone,
			false,
		},
		{
			"OK: PriorityLow",
			internal.PriorityLow,
			false,
		},
		{
			"OK: PriorityMedium",
			internal.PriorityMedium,
			false,
		},
		{
			"OK: PriorityHigh",
			internal.PriorityHigh,
			false,
		},
		{
			"ERR: unknown value",
			internal.Priority(-1),
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actualErr := tt.input.Validate()
			if (actualErr != nil) != tt.withErr {
				t.Fatalf("expected error %t, got %s", tt.withErr, actualErr)
			}

		})
	}

}

func TestDates_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   internal.Dates
		withErr bool
	}{
		{
			"OK: Start.IsZero",
			internal.Dates{
				Due: time.Now(),
			},
			false,
		},
		{
			"OK: Due.IsZero",
			internal.Dates{
				Start: time.Now(),
			},
			false,
		},
		{
			"OK: Start < Due",
			internal.Dates{
				Start: time.Now(),
				Due:   time.Now().Add(2 * time.Hour),
			},
			false,
		},
		{
			"ERR: Start > Due",
			internal.Dates{
				Start: time.Now().Add(2 * time.Hour),
				Due:   time.Now(),
			},
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actualErr := tt.input.Validate()
			if (actualErr != nil) != tt.withErr {
				t.Fatalf("expected error %t, got %s", tt.withErr, actualErr)
			}
		})
	}
}

func TestTask_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   internal.Task
		withErr bool
	}{
		{
			"OK",
			internal.Task{
				Description: "eat my poopoo",
				Dates: internal.Dates{
					Start: time.Now(),
					Due:   time.Now().Add(2 * time.Hour),
				},
				Priority: internal.PriorityHigh,
			},
			false,
		},
		{
			"ERR: Description",
			internal.Task{
				Dates: internal.Dates{
					Start: time.Now(),
					Due:   time.Now().Add(2 * time.Hour),
				},
				Priority: internal.PriorityHigh,
			},
			true,
		},
		{
			"ERR: Dates",
			internal.Task{
				Description: "eat my poopoo again",
				Dates: internal.Dates{
					Start: time.Now().Add(2 * time.Hour),
					Due:   time.Now(),
				},
				Priority: internal.PriorityHigh,
			},
			true,
		},
		{
			"ERR: Priority",
			internal.Task{
				Description: "eat my poopoo",
				Dates: internal.Dates{
					Start: time.Now(),
					Due:   time.Now().Add(2 * time.Hour),
				},
				Priority: internal.Priority(-1),
			},
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actualErr := tt.input.Validate()
			if (actualErr != nil) != tt.withErr {
				t.Fatalf("expected error %t, got %s", tt.withErr, actualErr)
			}
		})
	}
}
