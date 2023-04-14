package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	return true
}

func Move(src string, dest string, tree string) error {
	log.Printf("Moving %s to %s", src, dest)
	err := os.Rename(filepath.Join(tree, src), filepath.Join(tree, dest))
	if err != nil {
		return fmt.Errorf("Failed to copy: %w", err)
	}
	return nil
}
