/*
Copyright © 2021 Dennis Thisner <dthisner@protonmail.com>
*/
package transfer

import (
	"testing"
)

func TestFileDest(t *testing.T) {
	tests := []struct {
		name  string
		given string
		want  string
	}{
		{
			name:  "Artist / Album / Song",
			given: "/Users/dennis/Music/iTunes/iTunes Media/Music/Alestorm/Back Through Time/01 Back Through Time.m4a",
			want:  "/Volumes/Hugin/Music/Alestorm/Back Through Time/01 Back Through Time.m4a",
		},
		{
			name:  "Folder",
			given: "/Users/dennis/Music/iTunes/iTunes Media/Music/bob",
			want:  "/Volumes/Hugin/Music/bob",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val := fileDest(tt.given)

			if val != tt.want {
				t.Errorf("want %q; got %q", tt.want, val)
			}
		})
	}
}

func TestGetFolderPath(t *testing.T) {
	tests := []struct {
		name  string
		given string
		want  string
	}{
		{
			name:  "Long path",
			given: "/Users/dennis/Music/iTunes/iTunes Media/Music/Alestorm/Back Through Time/01 Back Through Time.m4a",
			want:  "/Users/dennis/Music/iTunes/iTunes Media/Music/Alestorm/Back Through Time",
		},
				{
			name:  "Short path",
			given: "/dennis/bob",
			want:  "/dennis",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			val := getFolderPath(tt.given)

			if val != tt.want {
				t.Errorf("want %q; got %q", tt.want, val)
			}
		})
	}
}