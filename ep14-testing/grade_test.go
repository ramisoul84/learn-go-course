package grade

import "testing"

func TestLetter(t *testing.T) {
	got := Letter(95)
	want := "A"
	if got != want {
		t.Errorf("Letter(95) = %q, want %q", got, want)
	}
}

func TestLetterTable(t *testing.T) {
	tests := []struct {
		name  string
		score int
		want  string
	}{
		{"top score", 100, "A"},
		{"lowest A", 90, "A"},
		{"highest B", 89, "B"},
		{"lowest C", 70, "C"},
		{"fail", 69, "F"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Letter(tt.score)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}
