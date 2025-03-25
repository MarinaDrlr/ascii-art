# ASCII Art Generator in Go

## Description

This is a simple ASCII Art generator written in Go.  
It takes a string as input and prints it using ASCII characters in a graphical representation.

The program supports:
- Latin letters (A-Z, a-z)
- Numbers (0-9)
- Spaces and common punctuation marks
- Newlines (`\n`) for multi-line text

## Usage

### Basic example:

```sh
go run . "Hello World"
```

### Example with newline character:

```sh
go run . "Hello\nWorld"
```

You can use `cat -e` to visualize line endings:

```sh
go run . "Hello World" | cat -e
```


### Example:

```sh
go run . "Hello World" | cat -e
 _    _            _    _                   __          __                    _        _   $
| |  | |          | |  | |                  \ \        / /                   | |      | |  $
| |__| |    ___   | |  | |    ___            \ \  /\  / /     ___     _ __   | |    __| |  $
|  __  |   / _ \  | |  | |   / _ \            \ \/  \/ /     / _ \   | '__|  | |   / _` |  $
| |  | |  |  __/  | |  | |  | (_) |            \  /\  /     | (_) |  | |     | |  | (_| |  $
|_|  |_|   \___|  |_|  |_|   \___/              \/  \/       \___/   |_|     |_|   \__,_|  $
                                                                                           $
                                                                                           $
```

If the input is an empty string or contains only newlines, nothing or just blank lines will be printed.

## Compatibility

This program is designed for Linux/macOS.  
Windows users should use WSL (Windows Subsystem for Linux).

## Installation

### Prerequisites

- [Go](https://go.dev/) installed on your system.

### Clone the repository

```sh
git clone https://github.com/yourusername/ascii-art
cd ascii-art
```

### Run the program

```sh
go run . "Hello World"
```

## Testing

Unit tests are included to validate the ASCII rendering. To run tests:

```sh
go test
```
or
```sh
go test -v
```
for verbose output with detailed information about each test case.

### Sample test output

```sh
=== RUN   TestNormalizeInput
--- PASS: TestNormalizeInput (0.00s)
=== RUN   TestNormalizeInputNewlines
--- PASS: TestNormalizeInputNewlines (0.00s)
=== RUN   TestLoadBanner
--- PASS: TestLoadBanner (0.00s)
=== RUN   TestLoadCorruptedBanner
--- PASS: TestLoadCorruptedBanner (0.00s)
=== RUN   TestGenerateASCIIArt
--- PASS: TestGenerateASCIIArt (0.00s)
=== RUN   TestGenerateASCIIArtUnsupportedChar
--- PASS: TestGenerateASCIIArtUnsupportedChar (0.00s)
PASS
ok      ascii-art       0.003s
```

## License

This project is open-source and licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

