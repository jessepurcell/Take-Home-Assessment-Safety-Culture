package folder_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	// Setup test parameters
	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		// TODO: your tests here
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			get := f.GetFoldersByOrgID(tt.orgID)
			assert.Equal(t, tt.want, get)
		})
	}
}

func Test_folder_GetAllChildFolders(t *testing.T) {
	// Setup test parameters
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)
	folders := folder.GetDataFromFile("get_test_data.json")

	t.Parallel()
	tests := [...]struct {
		name        string
		orgID       uuid.UUID
		folders     []folder.Folder
		want        []folder.Folder
		errorOutput string
	}{
		{"alpha", orgID, folders, folder.GetDataFromFile("get_folder_test_1.json"), ""},
		{"bravo", orgID, folders, folder.GetDataFromFile("get_folder_test_2.json"), ""},
		{"charlie", orgID, folders, []folder.Folder(nil), ""},
		{"echo", orgID, folders, []folder.Folder(nil), ""},
		{"invalid_folder", orgID, folders, []folder.Folder(nil), "folder does not exist"},
		{"foxtrot", orgID, folders, []folder.Folder(nil), "folder does not exist in the specified organization"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			//orgFolder := f.GetFoldersByOrgID(tt.orgID)
			subFolders, err := f.GetAllChildFolders(tt.orgID, tt.name)
			if err != nil {
				assert.Equal(t, tt.errorOutput, err.Error())
			}
			assert.Equal(t, tt.want, subFolders)
		})
	}
}
