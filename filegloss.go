package filegloss

import (
	"fmt"
	"strings"
)

// File this struct has all information about the file that is needed to pretty print it or
// use it else where
type File struct {
	Name            string
	Exports         []string
	Sections        map[string][]string
	Symbols         []string
	DynamicSymbols  []string
	ImportedSymbols []string
	HeaderType      string
	DataString      string
	OSABI           string
	Version         string
	Machine         string
}

// PrintInfo prints all information about the file in ASCII
func (f *File) InfoString() string {
	var s string
	s += fmt.Sprintf("Name: %s\n", f.Name)
	s += fmt.Sprintf("HeaderType: %s\n", f.HeaderType)
	s += fmt.Sprintf("DataString: %s\n", f.DataString)
    s += fmt.Sprintf("OSABI: %s\n", f.OSABI)
    s += fmt.Sprintf("Version: %s\n", f.Version)
    s += fmt.Sprintf("Machine: %s\n", f.Machine)
    s += "Exports:\n"
    for _, ps := range f.Exports {
        s+= ps
    }
    s += "Sections:\n"
    for sname, sarray := range f.Sections {
        s+= "\tName: " + sname + "\n"
        s+= "\tType: " + sarray[0] + "\n"
        s+= "\tFlags: " + sarray[1] + "\n"
    }
    s += "\n"
    s += "Symbols:\n"
    for _, ps := range f.Symbols {
        s+= ps + "\n"
    }
    s += "\n"
    s += "DynamicSymbols:\n"
    for _, ps := range f.DynamicSymbols {
        s+= ps + "\n"
    }
    s += "\n"
    s += "ImportedSymbols:\n"
    for _, ps := range f.ImportedSymbols {
        s+= ps  + "\n"
    }
    return s
}

// PrettyPrintInfo prints all information about the file in a very pretty way
func (f *File) PrettyPrintInfo() {
}

func New(filename string) *File {
	// todo get the info
    var f File
    pathParts := strings.Split(filename, "/")
    f.Name = pathParts[len(pathParts) - 1]
    f.GetFileInfo(filename)
    return &f
}
