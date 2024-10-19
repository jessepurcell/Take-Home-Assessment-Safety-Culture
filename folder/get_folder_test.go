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
	orgID := uuid.FromStringOrNil(folder.DefaultOrgID)
	allFolders := folder.GetAllFolders()

	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{"creative", orgID, allFolders, []folder.Folder(nil)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			//orgFolder := f.GetFoldersByOrgID(tt.orgID)
			subFolders := f.GetAllChildFolders(tt.orgID, tt.name)
			assert.Equal(t, tt.want, subFolders)
		})
	}
}
