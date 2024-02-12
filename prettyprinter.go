package filegloss

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

const columnWidth = 20

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
	titleStyle = lipgloss.
			NewStyle().
			Foreground(lipgloss.Color(BrightWhite.String())).
			Background(lipgloss.Color(Blue.String()))
	valueStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(BrightWhite.String()))
	listHeaderStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color(BrightWhite.String())).
			Background(lipgloss.Color(Cyan.String())).BorderStyle(lipgloss.NormalBorder()).
			BorderBottom(true).
			BorderForeground(lipgloss.Color(Gray.String()))
	listItemStyle = lipgloss.NewStyle().Foreground(lipgloss.Color(BrightWhite.String()))
)

func (f *File) PrettyPrint() string {
	return renderList("File", f.DynamicSymbols)
}

// this should be a header for a list of a header for one word with lipgloss
func renderTitle(s string) string {
	return titleStyle.Render(s)
}

func renderValue(h string) string {
	return valueStyle.Render(h)
}

func renderListHeader(header string) string {
	return listHeaderStyle.Render(header)
}

func renderListValue(value string) string {
	return listItemStyle.Render(value)
}

func renderList(header string, list []string) string {
	h := renderTitle(header)
	var l string
	for _, l := range list {
		l = lipgloss.JoinHorizontal(lipgloss.Top,
			l,
			renderListValue(l),
		)
	}
	resList := listStyle.Render(
		lipgloss.JoinHorizontal(lipgloss.Top,
			h,
			l,
		),
	)
	return resList
}
