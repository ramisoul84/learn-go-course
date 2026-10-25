# Episode 15: Generics

Module 3 · Writing idiomatic Go

▶️ Watch on YouTube: [Learn Go — full course playlist](PLAYLIST_LINK)

## What you'll learn

- Why generics were added to Go
- Type parameters on functions
- Constraints, `any`, `comparable` and `cmp.Ordered`
- Generic types, such as a `Stack[T]`
- When not to use generics

## Code

Code for this episode will be added when the video is released.

## Exercise

Write a generic `Filter[T any](items []T, keep func(T) bool) []T` and use it with ints and strings.
