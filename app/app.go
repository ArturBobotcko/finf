package app

import (
	"finf/jsondecoder"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
)

// returns a map of all file extensions from provided extensions.json file
func GetFileExtensions() map[string]jsondecoder.FileInfo {
	return jsondecoder.LoadFileExtensions()
}

// gets file extesion
func getFileExtension(name string) string {
	strs := filepath.Ext(name)

	if strs == "" {
		return "none"
	}
	return strs[1:] // Return without the dot
}

// convert POSIX symbolic notation for file permissions into string
func describeFilePermissions(mode os.FileMode) string {
	var permissions strings.Builder

	switch {
	case mode.IsRegular():
		permissions.WriteString("Regular file; ")
	case mode.IsDir():
		permissions.WriteString("Directory; ")
	case mode&os.ModeSymlink != 0:
		permissions.WriteString("Symbolic link; ")
	default:
		permissions.WriteString("Special file; ")
	}

	addPermission := func(cond bool, description string) {
		if cond {
			permissions.WriteString(description + "; ")
		}
	}

	// User permissions
	addPermission(mode&0400 != 0, "User can read")
	addPermission(mode&0200 != 0, "User can write")
	addPermission(mode&0100 != 0, "User can execute")

	// Group permissions
	addPermission(mode&0040 != 0, "Group can read")
	addPermission(mode&0020 != 0, "Group can write")
	addPermission(mode&0010 != 0, "Group can execute")

	// Others permissions
	addPermission(mode&0004 != 0, "Others can read")
	addPermission(mode&0002 != 0, "Others can write")
	addPermission(mode&0001 != 0, "Others can execute")

	return permissions.String()
}

func PrintFileInfo(filename string) error {
	dir := os.DirFS(".")
	fileStat, err := fs.Stat(dir, filename)
	if err != nil {
		return fmt.Errorf("error getting file information: %w", err)
	}

	fileExtension := strings.ToUpper(getFileExtension(filename))
	fileExts := GetFileExtensions()

	fmt.Println("Name:", fileStat.Name())
	if info, ok := fileExts[fileExtension]; ok {
		fmt.Print("File type: ")
		for i, description := range info.Descriptions {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%q", description)
		}
		fmt.Println()
	} else {
		fmt.Println("File type: unknown file type")
	}

	fmt.Println("Size:", fileStat.Size(), "Bytes")
	permissions := describeFilePermissions(fileStat.Mode())
	fmt.Printf("Access: %s\n", permissions)
	fmt.Println("Modified:", fileStat.ModTime())

	return nil
}
