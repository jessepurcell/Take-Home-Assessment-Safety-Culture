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
	folders := folder.GetAllFolders()

	var test1Want = []folder.Folder{
		{"clear-supergran", orgID, "stunning-horridus.sacred-moonstar.nearby-maestro.dashing-forearm.clear-supergran"},
		{"related-kitty", orgID, "stunning-horridus.sacred-moonstar.nearby-maestro.dashing-forearm.related-kitty"},
		{"organic-hulk", orgID, "stunning-horridus.sacred-moonstar.nearby-maestro.dashing-forearm.organic-hulk"},
	}

	t.Parallel()
	tests := [...]struct {
		name    string
		orgID   uuid.UUID
		folders []folder.Folder
		want    []folder.Folder
	}{
		{"dashing-forearm", orgID, folders, test1Want},             // Find 3 subfolders
		{"dashing-forearm", uuid.Must(uuid.NewV4()), folders, nil}, // Subfolder in wrong org
		{"does_not_exist", orgID, folders, nil},                    // Folder does not exist
		{"creative-scalphunter", orgID, folders, nil},              // Top level folder
		{"merry-mega-man", orgID, folders, nil},                    // Bottom level folder
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := folder.NewDriver(tt.folders)
			//orgFolder := f.GetFoldersByOrgID(tt.orgID)
			subFolders, _ := f.GetAllChildFolders(tt.orgID, tt.name)
			assert.Equal(t, tt.want, subFolders)
		})
	}
}
