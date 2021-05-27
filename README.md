# M3U-to-drive

This is to copy files from your iTunes playlist into your music player. Avoding copying music that you don't really fancy or just from specific playlists.

## Usage

1. Change the location variables to match yours.
1. Export your playlist from iTunes/Music
1. Open the file and remove the top row and delete it (this is to make sure everything is in rows and not just 1 long line)
1. Run the script and it will start copying.

## ToDo

1. [ ] Run in parallel
1. [ ] Config file for variables
1. [ ] Unit Tests
1. [ ] Ability to pass in variables
1. [x] Check to see if the file already exists and don't copy new if it does
1. [ ] Overwrite to copy over if the file exist already
1. [x] Check to see if folder already exist to skip the creating folder step again
1. [ ] Setup using UI as a desktop app
1. [x] Write out a summary of Time it took, files copied
1. [ ] List files that wasn't able to be copied
1. [ ] Implement Cobra

## Bug

1. [ ] Needs to open and save m3u file before using it
