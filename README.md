# ASCII Art in Go

## Description

This is a simple and efficient ASCII Art generator written in Go. The program takes a string input and converts it into ASCII characters using the "standard" banner font.

## Features

- Converts plain text to ASCII art using the standard font.
- Supports multi-line input with `\n` and actual newlines.
- Preserves formatting and spacing.
- Lightweight: Uses only Go’s standard library.
- CLI-based: Simple command-line interface for fast execution.

## Compatibility

This program is designed for Linux/macOS.  
Windows users should use WSL (Windows Subsystem for Linux).

## Installation

### Prerequisites

- [Go](https://go.dev/) installed on your system.

### Clone the repository

```sh
git clone https://github.com/MarinaDrlr/ascii-art
cd ascii-art
```

### Run the program

```sh
go run . "Hello World"
```

### To check for special characters like newlines in the output, you can use:

```sh
go run . "Hello World" | cat -e
```

## Usage

The program accepts a single string argument which will be converted to ASCII Art.  
It uses the "standard.txt" font file included in the project.

### Examples & Output:

```sh
go run . ""
           # nothing printed
```

```sh
go run . "\n" | cat -e
$          # a single blank line printed
```

```sh
go run . "Hello" | cat -e
 _    _          _   _          $
| |  | |        | | | |         $
| |__| |   ___  | | | |   ___   $
|  __  |  / _ \ | | | |  / _ \  $
| |  | | |  __/ | | | | | (_) | $
|_|  |_|  \___| |_| |_|  \___/  $
                                $
                                $
```

## Character Support

This program supports:

- Numbers (0-9)
- Latin letters (A-Z, a-z)
- Spaces
- Special characters (e.g., !@#$%^&*()-_=+[]{}|;:'",.<>?/)
- Newline character (`\n`) for multi-line formatting

For characters that have special meaning in the shell (e.g., `"`, `'`, `\\`, `&`, `|`, `*`), you must escape them using a backslash (`\\`) so they appear as literal characters.

Example:

```sh
go run . "Hello \"World\" |" | cat -e
```

## Error Handling

If no input is provided, the program displays a helpful usage message:

```sh
go run .
Error: No input was given.
Usage: go run . [STRING]
EX: go run . "something"
```

If the standard banner file is missing, empty, or corrupted, an appropriate error message will be shown:

```sh
go run . "Test"
Error: Banner file standard.txt is corrupted.
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

## License

This project is open-source and licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

