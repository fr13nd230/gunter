package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	// "strings"
)

// LoadDatabase(root string) ([]string, error)
// Walks recursively any dir and read all txt files that contains hashes.
func LoadDatabase(root string) ([]string, error) {
	var data []string
	secRoot, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	err = filepath.WalkDir(secRoot, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(path) == ".txt" {
			file, err := os.OpenFile(path, os.O_RDONLY, 0655)
			if err != nil {
				return err
			}
			defer file.Close()

			scanner := bufio.NewScanner(file)
			for scanner.Scan() {
				data = append(data, scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}
	return data, err
}

// FileData a struct that represents a file possible hashes.
type FileData struct {
	// File will have an md5 hash
	md5Hash string
	// File will have an sha1 hash
	sha1Hash string
	// File will have an sha256 hash
	sha256Hash string
}

// GetFileHash(path string) (*FileData, error)
// This function will open a file and hash it to three continous hashes.
func GetFileHash(path string) (*FileData, error) {
	secPath, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}
	file, err := os.OpenFile(secPath, os.O_RDONLY, 0644)
	if err == fs.ErrNotExist {
		return nil, errors.New("No such file exists.")
	}
	if err != nil {
		return nil, err
	}
	defer file.Close()

	hMD5 := md5.New()
	hSHA1 := sha1.New()
	hSHA256 := sha256.New()
	mw := io.MultiWriter(hMD5, hSHA1, hSHA256)

	if _, err := io.Copy(mw, file); err != nil {
		return nil, err
	}

	return &FileData{
		md5Hash:    fmt.Sprintf("%x", hMD5.Sum(nil)),
		sha1Hash:   fmt.Sprintf("%x", hSHA1.Sum(nil)),
		sha256Hash: fmt.Sprintf("%x", hSHA256.Sum(nil)),
	}, nil
}

// Usage: gunter [args]
// Example: gunter -file path/to/file -database /opt/database
func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: gunter [args]")
		fmt.Println("Example: gunter --file path/to/file")
		fmt.Println("List of supported args:")
		fmt.Println("\t -database [PATH] provide a path to a directory containing hash signature files.")
		fmt.Println("\t -file [PATH] provide the path to scan a single file.")
		fmt.Println("\t -dir [PATH] provide a path to recursively scan a dir.")
		os.Exit(2)
	}

	database := flag.String("database", "database/", "-database [PATH] provide a path to a directory containing hash signature files.")
	dir := flag.String("dir", "", "-dir [PATH] provide a path to recursively scan a dir.")
	file := flag.String("file", "", "-file [PATH] provide a path to scan a single file.")

	flag.Parse()

	hashes, err := LoadDatabase(*database)
	if err != nil {
		fmt.Printf("Provided database couldn't be loaded, %v", err)
		os.Exit(1)
	}

	var affected int
	// var scanned int

	if (*file != "" && *dir != "") || (*file == "" && *dir == "") {
		fmt.Println("Exactly one of -file or -dir must be provided, type gunter -h for usage help.")
		os.Exit(2)
	}

	if *file != "" {
		fh, err := GetFileHash(*file)
		if err != nil {
			fmt.Printf("Provided file couldn't be loaded: %v\n", err)
			os.Exit(1)
		}
		// scanned = 1
		for _, hash := range hashes {
			if hash == fh.md5Hash || hash == fh.sha1Hash || hash == fh.sha256Hash {
				affected++
			}
		}
	} else if *dir != "" {
		// err := filepath.WalkDir(*dir, func(path string, d fs.DirEntry, err error) error {
		// 	if err != nil {
		// 		return err
		// 	}
		// 	if !d.IsDir() && !strings.HasPrefix(filepath.Ext(path), ".") {
		// 		// Process all files (no extension filter for simplicity)
		// 		fh, err := GetFileHash(path)
		// 		if err != nil {
		// 			// Skip files we can't read, but continue scanning
		// 			return nil
		// 		}
		// 		scanned++
		// 		for _, hash := range hashes {
		// 			if hash == fh.md5Hash || hash == fh.sha1Hash || hash == fh.sha256Hash {
		// 				affected++
		// 			}
		// 		}
		// 	}
		// 	return nil
		// })
		// if err != nil {
		// 	fmt.Printf("Error scanning directory: %v\n", err)
		// 	os.Exit(1)
		// }
	}

	fmt.Println("========== GUNTER ==========")
	fmt.Printf("-database %v\n", len(hashes))
	fmt.Printf("-file %v\n", *file)
	fmt.Printf("-dir %v\n", *dir)
	fmt.Printf("Scanned files, %v has been affected.\n", affected)
	fmt.Println("============================")
	os.Exit(0)
}
