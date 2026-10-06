package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "Использование: task-01 <входной-файл.jsonl> <выходной-файл.txt>")
		os.Exit(2)
	}

	fmt.Fprintln(os.Stderr, "начало")
	os.Exit(1)
}
