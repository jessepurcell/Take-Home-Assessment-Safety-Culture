package folder_test

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/georgechieng-sc/interns-2022/folder"
	"github.com/gofrs/uuid"
)

func Test_folder_GetAllFolders(t *testing.T) {
	// Setup test parameters
	//folders := folder.GetDataFromFile("get_test_data.json")

	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//f := folder.NewDriver(tt.folders)
			//get := f.GetFoldersByOrgID(tt.orgID)
			//assert.Equal(t, tt.want, get)
		})
	}
}

// feel free to change how the unit test is structured
func Test_folder_GetFoldersByOrgID(t *testing.T) {
	// Setup test parameters
	folders := folder.GetDataFromFile("get_test_data.json")

	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{"org_test_1", uuid.FromStringOrNil("c1556e17-b7c0-45a3-a6ae-9546248fb17a"), folders, folder.GetDataFromFile("get_org_1.json")},
		{"org_test_2", uuid.FromStringOrNil("b504a1fc-6979-4a93-ab77-400ce86de1e5"), folders, folder.GetDataFromFile("get_org_2.json")},
		{"org_test_3", uuid.FromStringOrNil(""), folders, []folder.Folder{}},
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
		{"charlie", orgID, folders, []folder.Folder{}, ""},
		{"echo", orgID, folders, []folder.Folder{}, ""},
		{"invalid_folder", orgID, folders, []folder.Folder{}, "folder does not exist"},
		{"foxtrot", orgID, folders, []folder.Folder{}, "folder does not exist in the specified organization"},
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
