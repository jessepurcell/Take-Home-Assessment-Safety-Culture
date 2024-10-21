package folder

import (
	"errors"
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

func (f *driver) GetAllChildFolders(orgID uuid.UUID, name string) ([]Folder, error) {
	folderExists := false
	folders := f.folders
	var childFolders []Folder
	for _, folder := range folders {
		folderPaths := strings.Split(folder.Paths, ".")
		for _, folderPath := range folderPaths {
			if folderPath == name {
				folderExists = true
				if folder.OrgId != orgID {
					return nil, errors.New("folder does not exist in the specified organization")
				}
				if folderPath == folderPaths[len(folderPaths)-1] {
					// Folder is the last node in the tree
					break
				}
				childFolders = append(childFolders, folder)
			}
		}
	}
	if folderExists {
		return childFolders, nil
	}
	return nil, errors.New("folder does not exist")
}
