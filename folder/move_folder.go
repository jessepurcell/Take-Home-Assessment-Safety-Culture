package folder

import (
	"errors"
	"strings"
)

func (f *driver) MoveFolder(name string, dst string) ([]Folder, error) {

	folders := GetAllFolders()

	var sourceFolder, destinationFolder *Folder
	for i := range folders {
		if folders[i].Name == name {
			sourceFolder = &folders[i]
		}
		if folders[i].Name == dst {
			destinationFolder = &folders[i]
		}
	}

	if sourceFolder == nil {
		return nil, errors.New("source folder does not exist")
	}
	if destinationFolder == nil {
		return nil, errors.New("destination folder does not exist")
	}
	if sourceFolder == destinationFolder {
		return nil, errors.New("cannot move a folder to itself")
	}
	if sourceFolder.OrgId != destinationFolder.OrgId {
		return nil, errors.New("cannot move a folder to a different organization")
	}
	if strings.Contains(destinationFolder.Paths, sourceFolder.Name) {
		return nil, errors.New("cannot move a folder to a child of itself")
	}

	var newFolders []Folder

	for _, folder := range folders {
		if strings.HasPrefix(folder.Paths, sourceFolder.Paths) {
			relativePath := strings.TrimPrefix(folder.Paths, sourceFolder.Paths)
			newPath := destinationFolder.Paths + "." + sourceFolder.Name + relativePath
			movedFolder := Folder{
				Name:  folder.Name,
				Paths: newPath,
				OrgId: folder.OrgId,
			}
			newFolders = append(newFolders, movedFolder)
		} else {
			newFolders = append(newFolders, folder)
		}
	}
	return newFolders, nil
}
