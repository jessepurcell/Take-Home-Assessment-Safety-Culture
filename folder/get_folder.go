package folder

import (
	"github.com/gofrs/uuid"
	"strings"
)

func GetAllFolders() []Folder {
	return GetSampleData()
}

func (f *driver) GetFoldersByOrgID(orgID uuid.UUID) []Folder {
	folders := f.folders

	res := []Folder{}
	for _, f := range folders {
		if f.OrgId == orgID {
			res = append(res, f)
		}
	}

	return res
}

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) []Folder {
	// Your code here...
	folders := GetAllFolders()

	var childFolders []Folder
	for _, f := range folders {
		folderPaths := strings.Split(f.Paths, ".")
		for _, path := range folderPaths {
			// Check if the path contains the folder name
			// And check if the folder isn't the last one in the tree as this would not be a child folder
			if path == name && path != folderPaths[len(folderPaths)-1] {
				childFolders = append(childFolders, f)
			}
		}
	}
	return childFolders
}
