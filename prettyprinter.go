package filegloss

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	columnWidth = 30
)

const maxListItems = 10

type Color int

// ansi colors for lipgloss
const (
	Black         Color = iota
	Red           Color = iota
	Green         Color = iota
	Yellow        Color = iota
	Blue          Color = iota
	Magenta       Color = iota
	Cyan          Color = iota
	White         Color = iota
	Gray          Color = iota
	BrightRed     Color = iota
	BrightGreen   Color = iota
	BrightYellow  Color = iota
	BrightBlue    Color = iota
	BrightMagenta Color = iota
	BrightCyan    Color = iota
	BrightWhite   Color = iota
)

// return the number of the color
func (c Color) String() string {
	return fmt.Sprintf("%d", c)
}

var (
	listStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, true, false, false).
			BorderForeground(lipgloss.Color(Gray.String())).
			MarginRight(2).
			Height(8).
			Width(columnWidth + 1)
	listUnderStyle = lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, true, false, false).
			BorderForeground(lipgloss.Color(Gray.String())).
			MarginRight(2).
			Height(8).
			Width(columnWidth + 1 - 4)
	titleStyle = lipgloss.
			NewStyle().
			Foreground(lipgloss.Color(BrightWhite.String())).
			Background(lipgloss.Color(Blue.String()))
	valueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(BrightWhite.String()))
	genericListHeaderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(BrightWhite.String())).
				BorderStyle(lipgloss.NormalBorder()).
				PaddingLeft(5).
				PaddingRight(5).
				BorderBottom(true).
				BorderForeground(lipgloss.Color(Gray.String()))
	listUnderHeaderStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color(BrightWhite.String())).
				Background(lipgloss.Color(Magenta.String())).
				MarginLeft(2).
				MarginRight(2).
				PaddingLeft(5).
				PaddingRight(5)
	listItemStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color(BrightWhite.String()))
	listUnderItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(BrightWhite.String())).MarginLeft(2).MarginRight(2)
)

var currentColor Color = 0

func listHeaderStyle() lipgloss.Style {
	currentColor += 1 % 16
	if currentColor == BrightWhite || currentColor == White {
		currentColor += 1 % 16
	}
	return genericListHeaderStyle.Background(lipgloss.Color(currentColor.String()))
}

func (f *File) PrettyPrint() string {
	basicInfo := renderList("File Info", []string{f.Name, f.HeaderType, f.Machine, f.Version, f.OSABI})
	fileSymbols := renderList("Symbols", f.Symbols)
	fileDynamicSymbols := renderList("DynamicSymbols", f.DynamicSymbols)
	fileImportedSymbols := renderList("ImportedSymbols", f.ImportedSymbols)
	fileExprots := renderList("Exports", f.Exports)
	fileSections := renderMapList("Sections", f.Sections)
    lfs := strings.Repeat("\n", maxListItems - 1)
    rightBorder := lipgloss.NewStyle().Border(lipgloss.NormalBorder(), false, false, false, true).BorderForeground(lipgloss.Color(Gray.String())).Width(1).Render(lfs)
	s := lipgloss.JoinHorizontal(lipgloss.Top,
        rightBorder,
		basicInfo,
		fileSymbols,
		fileDynamicSymbols,
		fileImportedSymbols,
		fileSections,
		fileExprots,
	)
	return s
}

// this should be a header for a list of a header for one word with lipgloss
func renderTitle(s string) string {
	return titleStyle.Render(s)
}

func renderValue(h string) string {
	return valueStyle.Render(h)
}

func renderListHeader(header string) string {
	return listHeaderStyle().Render(header)
}

func renderUnderListHeader(header string) string {
	return listUnderHeaderStyle.Render(header)
}

func renderListValue(value string) string {
	return listItemStyle.Render(value)
}

func renderUnderListValue(value string) string {
	return listUnderItemStyle.Render(value)
}

func renderList(header string, list []string) string {
	if len(list) == 0 {
		return ""
	}
	h := renderListHeader(header)
	l := fmt.Sprintf("- %s", renderListValue(list[0]))
	var mli int
	if len(list) < maxListItems {
		mli = len(list) - 1
	} else {
		mli = maxListItems
	}
	for _, i := range list[1:mli] {
		l = lipgloss.JoinVertical(lipgloss.Left,
			l,
			fmt.Sprintf("- %s", renderListValue(i)),
		)
	}
	resList := listStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			h,
			l,
		),
	)
	return resList
}

func renderUnderList(header string, list []string) string {
	if len(list) == 0 {
		return ""
	}
	h := renderUnderListHeader(header)
	l := fmt.Sprintf("- %s", (list[0]))
	var mli int
	if len(list) < maxListItems {
		mli = len(list) - 1
	} else {
		mli = maxListItems
	}
	for _, i := range list[1:mli] {
		l = lipgloss.JoinVertical(lipgloss.Left,
			l,
			fmt.Sprintf("- %s", renderUnderListValue(i)),
		)
	}
	resList := listStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			h,
			l,
		),
	)
	return resList
}

func renderMapList(header string, list map[string][]string) string {
	if len(list) == 0 {
		return ""
	}
	var mli int
	if len(list) < maxListItems {
		mli = len(list)
	} else {
		mli = maxListItems
	}
	h := renderListHeader(header)
    var l string
	count := 0
	for k, v := range list {
		if count > mli {
			l = lipgloss.JoinVertical(lipgloss.Left,
				l,
				"...",
			)
			break
		}
        if (len(l) == 0) {
            l = renderKeyVal(k, v)
            continue
        }
		l = lipgloss.JoinVertical(lipgloss.Left,
			l,
			renderKeyVal(k, v),
		)
		count += 1
	}
	resList := listStyle.Render(
		lipgloss.JoinVertical(lipgloss.Left,
			h,
			l,
		),
	)
	return resList
}

func renderKeyVal(k string, v []string) string {
	if len(v) == 0 {
		return ""
	}
	s := k + ": "
	for e := range v {
		s += v[e]
		if len(s) > columnWidth-10 {
			s += "..."
			break
		}
		s += ", "
	}
	return s
}
