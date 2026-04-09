package main

import (
	"flag"
	"fmt"
	"os"
)

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
 
    fmt.Println("===== GUNTER =====")
    fmt.Printf("Program has ran and finished with code %v\n", code)
    fmt.Printf("-database %v\n", *database)
    fmt.Printf("-file %v\n", *file)
    fmt.Printf("-dir %v\n", *dir)
    fmt.Println("==================")
    os.Exit(code)
}
