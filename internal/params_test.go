package internal_test

import (
	"testing"
	"time"

	"github.com/terrytay/godo/internal"
)

func TestCreateParams_Validate(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		input   internal.CreateParams
		withErr bool
	}{
		{
			"OK",
			internal.CreateParams{
				Description: "hello world",
				Priority:    internal.PriorityHigh,
				Dates: internal.Dates{
					Start: time.Now(),
					Due:   time.Now().Add(2 * time.Hour),
				},
			},
			false,
		},
		{
			"ERR: Description",
			internal.CreateParams{
				Priority: internal.PriorityHigh,
				Dates: internal.Dates{
					Start: time.Now(),
					Due:   time.Now().Add(2 * time.Hour),
				},
			},
			true,
		},
		{
			"ERR",
			internal.CreateParams{},
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

func TestSearchParams_IsZero(t *testing.T) {
	t.Parallel()

	description := "blablabla"
	priority := internal.PriorityMedium
	isDone := true

	tests := []struct {
		name    string
		input   internal.SearchParams
		withErr bool
	}{
		{
			"OK",
			internal.SearchParams{
				Description: &description,
				Priority:    &priority,
				IsDone:      &isDone,
				From:        1,
				Size:        2,
			},
			false,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			actualErr := tt.input.IsZero()
			if (actualErr == true) != tt.withErr {
				t.Fatalf("expected error %t, got %t", tt.withErr, actualErr)
			}
		})
	}
}
