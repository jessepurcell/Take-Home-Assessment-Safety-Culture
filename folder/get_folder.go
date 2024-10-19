package folder

import (
	"fmt"
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
	folderExists := false
	folders := GetAllFolders()
	var childFolders []Folder
	for _, folder := range folders {
		folderPaths := strings.Split(folder.Paths, ".")
		for _, folderPath := range folderPaths {
			if folderPath == name {
				folderExists = true
				if folder.OrgId != orgID {
					fmt.Println("Error: Folder does not exist in the specified organization")
					return nil
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
		return childFolders
	} else {
		fmt.Println("Error: Folder does not exist")
		return nil
	}
}
