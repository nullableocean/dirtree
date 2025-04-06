package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	ENTRY_PREFIX      = "├───"
	LAST_ENTRY_PREFIX = "└───"
	INDENT            = "│\t"
	LAST_INDENT       = "\t"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: COMMAND [-f] [directory]")
		os.Exit(1)
	}

	path := os.Args[len(os.Args)-1]
	printFiles := len(os.Args) > 2 && os.Args[1] == "-f"

	err := dirTree(os.Stdout, path, printFiles)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func dirTree(out io.Writer, path string, printFiles bool) error {
	return walkDir(out, path, printFiles, "")
}

func walkDir(out io.Writer, path string, printFiles bool, prefix string) error {
	entries, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	filteredEntries := make([]os.DirEntry, 0, len(entries))
	for _, entry := range entries {
		if entry.IsDir() || printFiles {
			filteredEntries = append(filteredEntries, entry)
		}
	}

	for i, entry := range filteredEntries {
		isLast := i == len(filteredEntries)-1

		var linePrefix string
		if isLast {
			linePrefix = prefix + LAST_ENTRY_PREFIX
		} else {
			linePrefix = prefix + ENTRY_PREFIX
		}

		fmt.Fprintf(out, "%s%s", linePrefix, entry.Name())

		if !entry.IsDir() {
			info, err := entry.Info()
			if err != nil {
				return err
			}
			size := info.Size()
			if size == 0 {
				fmt.Fprint(out, " (empty)")
			} else {
				fmt.Fprintf(out, " (%db)", size)
			}
		}
		fmt.Fprintln(out)

		if entry.IsDir() {
			var newPrefix string
			if isLast {
				newPrefix = prefix + LAST_INDENT
			} else {
				newPrefix = prefix + INDENT
			}

			fullPath := filepath.Join(path, entry.Name())
			err := walkDir(out, fullPath, printFiles, newPrefix)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
