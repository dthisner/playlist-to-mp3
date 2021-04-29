package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	readFile()
}

func readFile() {
	    file, err := os.Open("/Users/dennis/Desktop/All Loved tracks.m3u")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    buf := make([]byte, 0, 64*1024) // Buffering size 1MB
    scanner.Buffer(buf, 1024*1024)

    for scanner.Scan() {
        fileLocation := scanner.Text()
        if fileLocation[0:1] == "/" {
            copyFile(fileLocation)
        }
        
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
}

func copyFile(fileLocation string) {
    stringToRemove := "/Users/dennis/Music/iTunes/iTunes Media/Music"
    destination := "/Volumes/Hugin/Music"

    musicLocation := fileLocation[len(stringToRemove): ]
    destination = destination + musicLocation

    fmt.Printf("Copying: %s to %s\n", fileLocation, destination)

    original, err := os.Open(fileLocation)
    if err != nil {
        log.Fatal(err)
    }
    defer original.Close()

    err = os.MkdirAll(getFolderPath(destination), os.ModePerm)
    if err != nil {
        log.Fatal(err)
    }

    new, err := os.Create(destination)
    if err != nil {
        log.Fatal(err)
    }
    defer new.Close()

    bytesWritten, err := io.Copy(new, original)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Bytes Written: %d\n", bytesWritten)
}

// This removes the name of the music file from the string
// This is to be able to create the folders before copying the file
func getFolderPath(path string) string {
    trailingSlash := 0

    for i, letter := range path {
        if string(letter) == "/" {                
            trailingSlash= i
        }
    }
    return path[0:trailingSlash]
}
