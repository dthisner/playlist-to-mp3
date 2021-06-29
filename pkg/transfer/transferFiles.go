/*
Copyright © 2021 Dennis Thisner <dthisner@protonmail.com>
*/

package transfer

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/hako/durafmt"
	"github.com/spf13/viper"

	constans "github.com/dthisner/m3u-to-drive/constants"
)

var copied, excisted int
var startTime time.Time
var M3U_LOCATION, DESTINATION, ORIGIN string

func StartTransfer() error {
	M3U_LOCATION = viper.GetString(constans.M3uLocation)
	DESTINATION = viper.GetString(constans.Dest)
	ORIGIN = viper.GetString(constans.Origin)

	startTime = time.Now()
	err := readPlaylist()
	if err != nil {
		return err
	}
	duration := durafmt.Parse(time.Since(startTime)).String()
	fmt.Printf("It took %s\nand copied %d files and %d was already on target location", duration, copied, excisted)
	return nil
}

func readPlaylist() error {
	file, err := os.Open(M3U_LOCATION)
	if err != nil {
		return fmt.Errorf("problem opening the playlist: '%s' \n ", err)
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

func copyFile(fileDest, orgLocation string) error {
	// Checks to see if the file already exists
	_, err := os.Stat(fileDest)
	if os.IsNotExist(err) {
		fmt.Printf("Copying: \"%s\" \nTo: \"%s\"\n", orgLocation, fileDest)

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

		_, err = io.Copy(new, original)
		if err != nil {
			return err
		}

		fmt.Printf("Successful\n")
		copied++
		return nil
	}

	fmt.Printf("Skipping, Already exists - \"%s\"\n", fileDest)
	excisted++
	return nil
}

func fileDest(orgLocation string) string {
	return DESTINATION + orgLocation[len(ORIGIN):]
}

// This removes the name of the music file from the string
// This is to be able to create the folders before copying the file
func getFolderPath(path string) string {
	trailingSlash := 0

	for i, letter := range path {
		if string(letter) == "/" {
			trailingSlash = i
		}
	}
	return path[0:trailingSlash]
}
