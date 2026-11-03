package main

import "fmt"

// Notifier is anything that can send a message.
type Notifier interface {
	Notify(msg string) string
}

// Email, SMS and Slack never mention Notifier. Having the method is enough.
type Email struct{ To string }

func (e Email) Notify(msg string) string {
	return "Email to " + e.To + ": " + msg
}

type SMS struct{ Phone string }

func (s SMS) Notify(msg string) string {
	return "SMS to " + s.Phone + ": " + msg
}

type Slack struct{ Channel string }

func (s Slack) Notify(msg string) string {
	return "Slack " + s.Channel + ": " + msg
}

// sendAll accepts any Notifier, including types written after it.
func sendAll(ns []Notifier, msg string) {
	for _, n := range ns {
		fmt.Println(n.Notify(msg))
	}
}

// Temp implements fmt.Stringer, so fmt prints it with String.
type Temp float64

func (t Temp) String() string {
	return fmt.Sprintf("%.1f°C", float64(t))
}

// describe accepts any value, and uses a type switch to find out what it is.
func describe(v any) string {
	switch x := v.(type) {
	case int:
		return fmt.Sprintf("int %d", x)
	case string:
		return "string " + x
	default:
		return fmt.Sprintf("something else: %v", x)
	}
}

func main() {
	// The program from the start of the episode.
	notifiers := []Notifier{
		Email{To: "alex@example.com"},
		SMS{Phone: "+1-555-0100"},
		Slack{Channel: "#orders"},
	}
	sendAll(notifiers, "Your order shipped")

	// fmt.Stringer
	fmt.Println(Temp(21.5))

	// any and type switches
	fmt.Println(describe(42))
	fmt.Println(describe("hi"))
	fmt.Println(describe(3.5))

	// A type assertion with comma-ok.
	var v any = "hello"
	s, ok := v.(string)
	fmt.Println(s, ok)
	n, ok := v.(int)
	fmt.Println(n, ok)
}
