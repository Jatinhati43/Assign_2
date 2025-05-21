# Word Count – Count Words in a Line of Text

This repository contains a Go program that reads a line of text from the user and counts the number of words in that line.

The project follows the [golang-standards/project-layout](https://github.com/golang-standards/project-layout) for organizing Go projects.


- `cmd/main.go`: The entry point of the application.
- `internal/wordcount/counter.go`: Contains the logic to count words.
- `internal/wordcount/counter_test.go`: Unit tests for the word counting function.


## 🧠 Functionality

- Reads a line of text from standard input.
- Counts the number of words separated by spaces.
- Prints the total word count.

Example:

Enter a line of text: Hello, how are you doing today?
Word count: 6

### Run the Application

go run ./cmd/wordcount



🧪 Running Tests
Run unit tests with:

go test ./internal/...

✅ Run Tests with Coverage

go test -cover ./internal/...
