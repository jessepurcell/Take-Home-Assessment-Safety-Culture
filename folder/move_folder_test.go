package folder_test

import (
	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_folder_MoveFolder(t *testing.T) {
	folders := folder.GetDataFromFile("move_folder_test.json")

	t.Parallel()
	tests := [...]struct {
		name        string
		folders     []folder.Folder
		source      string
		destination string
		want        []folder.Folder
		errorOutput string
	}{
		{"bravo->delta", folders, "bravo", "delta", folder.GetDataFromFile("move_folder_test_1.json"), ""},
		{"bravo->golf", folders, "bravo", "golf", folder.GetDataFromFile("move_folder_test_2.json"), ""},
		{"bravo->charlie", folders, "bravo", "charlie", nil, "cannot move a folder to a child of itself"},
		{"bravo->bravo", folders, "bravo", "bravo", nil, "cannot move a folder to itself"},
		{"bravo->foxtrot", folders, "bravo", "foxtrot", nil, "cannot move a folder to a different organization"},
		{"invalid_folder->foxtrot", folders, "invalid_folder", "delta", nil, "source folder does not exist"},
		{"bravo->invalid_folder", folders, "bravo", "invalid_folder", nil, "destination folder does not exist"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			results, err := f.MoveFolder(tt.source, tt.destination)
			if err != nil {
				assert.Equal(t, tt.errorOutput, err.Error())
			}
			assert.Equal(t, tt.want, results)
		})
	}
}
