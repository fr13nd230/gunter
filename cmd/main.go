package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"flag"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// LoadDatabse(root string) ([]string, error)
// Walks recursively any dir and read all txt files that contains hashes.
func LoadDatabase(root string) ([]string, error) {
	var data []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !d.IsDir() && filepath.Ext(path) == ".txt" {
			file, err := os.OpenFile(path, os.O_RDONLY, 0755)
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

func getFileHash(path string) (*FileData, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0755)
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
		md5Hash:    string(hMD5.Sum(nil)),
		sha1Hash:   string(hSHA1.Sum(nil)),
		sha256Hash: string(hSHA256.Sum(nil)),
	}, nil
}

// Usage: gunter [args]
// Example: gunter -file path/to/file -database /opt/database
func main() {
	code := 1

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Usage: gunter [args]")
		fmt.Println("Example: gunter --file path/to/file")
		fmt.Println("List of supported args:")
		fmt.Println("\t -database [PATH] provide a path to a directory containing hash signature files.")
		fmt.Println("\t -file [PATH] provide the path to scan a single file.")
		fmt.Println("\t -dir [PATH] provide a path to recursively scan a dir.")
		os.Exit(code)
	}

	database := flag.String("database", "database/", "-database [PATH] provide a path to a directory containing hash singnature files.")
	dir := flag.String("dir", "", "-dir [PATH] provide a path to recursively scan a dir.")
	file := flag.String("file", "", "-file [PATH] provide a path to scan a single file.")

	flag.Parse()

	hashes, err := LoadDatabase(*database)
	if err != nil {
		code = -1
		fmt.Printf("Provided database couldn't be loaded, %v", err)
		os.Exit(code)
	}

	var fh *FileData
	if strings.TrimSpace(*file) != " " && dir == nil {
		fh, err = getFileHash(*file)
		if err != nil {
			code = -1
			fmt.Printf("Provided file couldn't be loaded, %v", err)
			os.Exit(code)
		}
	}

	fmt.Println("===== GUNTER =====")
	fmt.Printf("Program has ran and finished with code %v\n", code)
	fmt.Printf("-database %v\n", len(hashes))
	fmt.Printf("-file %v and md5 %v\n", *file, fh.md5Hash)
	fmt.Printf("-dir %v\n", *dir)
	fmt.Println("==================")
	os.Exit(code)
}
