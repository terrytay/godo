package internal_test

import (
	"testing"

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
