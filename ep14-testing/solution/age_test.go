// Solution to the episode 14 exercise. Run it with: go test ./solution
package age

import (
	"errors"
	"testing"
)

func TestParseAge(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr error
	}{
		{"normal age", "30", 30, nil},
		{"youngest", "0", 0, nil},
		{"oldest", "150", 150, nil},
		{"too old", "151", 0, ErrOutOfRange},
		{"negative", "-1", 0, ErrOutOfRange},
		{"not a number", "abc", 0, ErrNotNumber},
		{"empty", "", 0, ErrNotNumber},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAge(tt.input)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("ParseAge(%q) error = %v, want %v", tt.input, err, tt.wantErr)
			}
			if got != tt.want {
				t.Errorf("ParseAge(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
