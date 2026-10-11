package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Fprintln(os.Stderr, "Использование: task-05 <входной-файл.jsonl> <все-каналы.txt> <проводные-каналы.txt>")
		os.Exit(2)
	}

	if err := run(os.Args[1], os.Args[2], os.Args[3], os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка:", err)
		os.Exit(1)
	}
}

func run(inputPath, outputPath, wiredPath string, stdout io.Writer) error {
	channels, err := LoadChannels(inputPath)
	if err != nil {
		return err
	}

	output, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("создать выходной файл: %w", err)
	}

	writeErr := WriteChannels(io.MultiWriter(stdout, output), channels)
	closeErr := output.Close()
	if writeErr != nil {
		return fmt.Errorf("вывести каналы: %w", writeErr)
	}
	if closeErr != nil {
		return fmt.Errorf("закрыть выходной файл: %w", closeErr)
	}

	wiredOutput, err := os.Create(wiredPath)
	if err != nil {
		return fmt.Errorf("создать файл проводных каналов: %w", err)
	}
	writeErr = WriteWiredChannels(wiredOutput, channels)
	closeErr = wiredOutput.Close()
	if writeErr != nil {
		return fmt.Errorf("вывести проводные каналы: %w", writeErr)
	}
	if closeErr != nil {
		return fmt.Errorf("закрыть файл проводных каналов: %w", closeErr)
	}
	return nil
}
