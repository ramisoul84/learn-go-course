# Episode 5: Mini-Project: Number Guessing Game

Module 1 · Basics

▶️ Watch on YouTube: [Learn Go — full course playlist](https://www.youtube.com/playlist?list=PLPDHYEltiPAI)

## What you'll learn

- Reading user input with `bufio.Scanner`
- Converting text to numbers with `strconv.Atoi`
- Random numbers with `math/rand/v2`
- Structuring a small program with functions
- Handling bad input without crashing

## Play the game

```bash
go run .            # the game from the video
go run ./solution   # the exercise solution: difficulty levels and a guess limit
```

```text
Guess my number from 1 to 100.
Your guess: 50
Too low!
Your guess: 75
Too high!
Your guess: sixty
Please type a whole number.
Your guess: 68
Correct! 3 tries.
```

## How it works

| Step | Code |
| --- | --- |
| Pick a secret number from 1 to 100 | `secret := rand.IntN(100) + 1` (package `math/rand/v2`) |
| Read a line from the keyboard | `scanner := bufio.NewScanner(os.Stdin)`, then `scanner.Scan()` and `scanner.Text()` |
| Remove spaces around the input | `strings.TrimSpace(text)` |
| Turn text into a number | `strconv.Atoi(text)` returns `(int, error)` |
| Ask again after bad input | `if err != nil { ...; continue }` |
| Give hints | `switch { case guess < secret: ... case guess > secret: ... default: ... }` |
| End the game | `return` from `main` |

## Exercise

Add difficulty levels (easy 1–10, medium 1–100, hard 1–1000) and a limit on the number of guesses.

Levels: easy is 1–10, medium 1–100, hard 1–1000. When the player runs out of guesses, show the secret number.
Solution: [`solution/main.go`](./solution/main.go).
