package stdlib

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// ---- Writing files ----

// os.WriteFile: write all at once (simplest, loads all into memory)
func writeAll(path, content string) error {
	return os.WriteFile(path, []byte(content), 0644)
}

// os.Create + defer Close: fine-grained control
func writeFile(path string, lines []string) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("writeFile: %w", err)
	}
	defer f.Close()

	for _, line := range lines {
		if _, err := fmt.Fprintln(f, line); err != nil {
			return err
		}
	}
	return nil
}

// bufio.Writer: batches small writes for performance
func writeLarge(path string, lines []string) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush() // don't forget - unflushed writes are lost
}

// ---- Reading files ----

// os.ReadFile: read all at once (fine for small files)
func readAll(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", fmt.Errorf("readAll: %w", err)
	}
	return string(data), nil
}

// bufio.Scanner: efficient line-by-line reading
func readLines(path string) ([]string, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("readLines: %w", err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err() // check scanner error after loop
}

// bufio.Reader: read with fine-grained control
func readChunks(r io.Reader, chunkSize int) ([][]byte, error) {
	var chunks [][]byte
	buf := make([]byte, chunkSize)
	for {
		n, err := r.Read(buf)
		if n > 0 {
			chunk := make([]byte, n)
			copy(chunk, buf[:n])
			chunks = append(chunks, chunk)
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return chunks, nil
}

// ---- stdin / stdout / stderr ----
// os.Stdin, os.Stdout, os.Stderr implement io.Reader / io.Writer

func readStdin() []string {
	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

func printErr(format string, args ...any) {
	fmt.Fprintf(os.Stderr, format+"\n", args...)
}

// ---- os.File flags for OpenFile ----
//
// os.O_RDONLY  read only
// os.O_WRONLY  write only
// os.O_RDWR    read+write
// os.O_CREATE  create if not exists
// os.O_TRUNC   truncate on open
// os.O_APPEND  append on write
//
// f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)

/*
os.ReadFile / os.WriteFile are the simplest entry points (added in Go 1.16).
For anything large, stream through bufio to avoid loading into memory.
Always defer f.Close() immediately after a successful open/create.
Always call w.Flush() after a bufio.Writer.
*/
