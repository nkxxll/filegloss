package filegloss

import (
	"github.com/Binject/debug/elf"
)

func (f *File) GetFileInfo(filename string) {
    elf_file, err := elf.Open(filename)
    if err != nil {
        panic("Error while trying to open the specified file")
    }

    // get the exports
    exports, err := elf_file.Exports()
    if err != nil {
        panic("Exports failed")
    }
    for _, exp := range exports {
        f.Exports = append(f.Exports, exp.Name)
    }

    // get the sections
    f.Sections = make(map[string][]string)
    sections := elf_file.Sections
    for _, s := range sections {
        f.Sections[s.Name] = []string{s.Type.GoString(), s.Flags.GoString()}
    }

    // Symbols
    symbols, err := elf_file.Symbols()
    if err != nil {
        panic("Error getting symbols")
    }
    for _, s := range symbols {
        f.Symbols = append(f.Symbols, s.Name)
    }

    // Symbols
    DynamicSymbols, err := elf_file.DynamicSymbols()
    if err != nil {
        panic("Error getting DynamicSymbols")
    }
    for _, s := range DynamicSymbols {
        f.DynamicSymbols = append(f.DynamicSymbols, s.Name)
    }


    // Symbols
    ImportedSymbols, err := elf_file.ImportedSymbols()
    if err != nil {
        panic("Error getting ImportedSymbols")
    }
    for _, s := range ImportedSymbols {
        f.ImportedSymbols = append(f.ImportedSymbols, s.Name)
    }



    f.HeaderType = elf_file.FileHeader.Type.GoString()
    f.DataString = elf_file.FileHeader.Data.GoString()
    f.OSABI = elf_file.FileHeader.OSABI.GoString()
	f.Version = elf_file.FileHeader.Version.GoString()
	f.Machine = elf_file.FileHeader.Machine.GoString()
}
