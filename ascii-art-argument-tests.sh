go run . "" | cat -e

go run . "\n" | cat -e

go run . "Hello\n" | cat -e

go run . "hello" | cat -e

go run . "HeLlO HuMaN" | cat -e

go run . "Hello There" | cat -e

go run . "1Hello 2There" | cat -e

go run . "{Hello There}" | cat -e

go run . "Hello\nThere" | cat -e

go run . "Hello\n\nThere" | cat -e

go run . "1a\"#FdwHywR&/()=" | cat -e

go run . "{|}~" | cat -e

go run . "[\]^_ 'a" | cat -e

go run . "1a\":;<=>?@" | cat -e

go run . '\!" #$%&'"'"'()*+,-./' | cat -e

go run . "ABCDEFGHIJKLMNOPQRSTUVWXYZ" | cat -e

go run . "abcdefghijklmnopqrstuvwxyz" | cat -e