// Package grade turns a test score into a letter.
package grade

// Letter returns the grade for a score from 0 to 100.
func Letter(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	default:
		return "F"
	}
}
