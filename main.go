package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

type fileTransferStatus struct {
        errorMsg string
        customMsg string
}

func main() {
	err := readPlaylist()
    if err != nil {
        log.Fatal(err)
    }
}

func readPlaylist() error {
    file, err := os.Open("/Users/dennis/Desktop/lovedTracks.m3u")
    if err != nil {
        return fmt.Errorf("Problem opening the playlist: '%s' \n", err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    buf := make([]byte, 0, 64*1024) // Buffering size 1MB
    scanner.Buffer(buf, 1024*1024)

    // Looping through the playlist 
    for scanner.Scan() {
        orgLocation := scanner.Text()
        if orgLocation[0:1] == "/" {

            err = setupPath(fileDest(orgLocation))
            if err != nil {
                return err
            }
            err = copyFile(fileDest(orgLocation), orgLocation)
            if err != nil {
                return err
            }
        }
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return nil
}

func setupPath(fileDest string) error {
    // Creating folder path if it dosn't exist
    _, err := os.Stat(getFolderPath(fileDest))
    if os.IsNotExist(err) {
        err = os.MkdirAll(getFolderPath(fileDest), os.ModePerm)
        if err != nil {
            return err
        }
    }

    return nil
}

func copyFile(fileDest, orgLocation string) error{
    fmt.Printf("Copying: \"%s\" to \"%s\"\n", orgLocation, fileDest)

    // Checks to see if the file already exists 
    _, err := os.Stat(fileDest)
    if os.IsNotExist(err) {
         // Open the original file
        original, err := os.Open(orgLocation)
        if err != nil {
            return err
        }
        defer original.Close()

           new, err := os.Create(fileDest)
        if err != nil {
            return err
        }
        defer new.Close()

        bytesWritten, err := io.Copy(new, original)
        if err != nil {
            return err
        }

        fmt.Printf("Bytes Written: %d\n", bytesWritten)
        return nil
    }

    fmt.Println("File already exists")
    return nil
}

func fileDest(orgLocation string) string {
    stringToRemove := "/Users/dennis/Music/iTunes/iTunes Media/Music"
    destination := "/Users/dennis/Desktop/test"

    return destination + orgLocation[len(stringToRemove): ]
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
