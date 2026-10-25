# Episode 12: Error Handling

Module 3 · Writing idiomatic Go

▶️ Watch on YouTube: [Learn Go — full course playlist](PLAYLIST_LINK)

## What you'll learn

- Why Go returns errors instead of throwing exceptions
- Creating errors with `errors.New` and `fmt.Errorf`
- Wrapping errors with `%w`
- Checking errors with `errors.Is` and `errors.As`
- Custom error types, and when to use `panic`

## Code

Code for this episode will be added when the video is released.

## Exercise

Write a `ParseAge(s string) (int, error)` function that returns a different, checkable error for "not a number" and "out of range".
