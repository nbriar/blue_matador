# A Code Test for Blue Matador

## Running the Program
You have 3 decent options for running the program:

1. Copy the code into https://play.golang.org/
2. Compile the Go code into a binary and run it from the command line:
    `go build && ./blue_matador chucky-random-joke`
3. Just run the `main.go` file:
    `go run main.go chucky-category-joke science`
    
## Available Commands
```
chucky-category-joke <category> - (science)
chucky-search-joke <search-string> - (blue)
validate-integer-palindrome <input> - (4557554)
calculate-stats <input> - ("1 2 3 4 5")
```

## Examples

```
go run main.go chucky-category-joke science
go run main.go chucky-category-joke "not a valid category"
go run main.go chucky-search-joke "blue whale"
go run main.go chucky-search-joke "not a valid search"
go run main.go validate-integer-palindrome 4556565
go run main.go validate-integer-palindrome 445565544
go run main.go calculate-stats "1 2 3 4 5 6 7"

