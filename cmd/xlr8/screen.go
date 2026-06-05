package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var selectedStyle = lipgloss.NewStyle().Bold(true).
	Foreground(lipgloss.Color("15")).Background(lipgloss.Color("34")).Padding(0, 1)

var normalStyle = lipgloss.NewStyle().Padding(0, 1)

var fileBoxStyle = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("10")).Padding(0, 1)

type model struct {
	width     int
	height    int
	fileList  []os.DirEntry
	pointer   int
	scrollTop int
	dir       string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m *model) updateScroll() {
	cellWidth := 30
	columns := m.width / cellWidth
	if columns < 1 {
		columns = 1
	}
	headerHeight := 8
	footerHeight := 8
	visibleRows := m.height - headerHeight - footerHeight
	if visibleRows < 3 {
		visibleRows = 3
	}
	currentRow := m.pointer / columns
	m.scrollTop = currentRow - visibleRows/2
	if m.scrollTop < 0 {
		m.scrollTop = 0
	}
	totalRows := (len(m.fileList) + columns - 1) / columns
	maxScroll := totalRows - visibleRows
	if maxScroll < 0 {
		maxScroll = 0
	}
	if m.scrollTop > maxScroll {
		m.scrollTop = maxScroll
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "left":
			if m.pointer > 0 {
				m.pointer--
			}
		case "right":
			if m.pointer < len(m.fileList)-1 {
				m.pointer++
			}
		case "up":
			columns := m.width / 30
			if columns < 1 {
				columns = 1
			}
			if m.pointer-columns >= 0 {
				m.pointer -= columns
			}
		case "down":
			columns := m.width / 30
			if columns < 1 {
				columns = 1
			}
			if m.pointer+columns < len(m.fileList) {
				m.pointer += columns
			}
		case "enter":
			if len(m.fileList) == 0 {
				break
			}
			selected := m.fileList[m.pointer]
			if selected.IsDir() {
				m.dir = filepath.Join(
					m.dir,
					selected.Name(),
				)
				m.fileList = GetFoldersList(m.dir)
				m.pointer = 0
				m.scrollTop = 0
			}
		case "backspace", "backspace2":
			parent := filepath.Dir(m.dir)
			if parent != m.dir {
				m.dir = parent
				m.fileList = GetFoldersList(m.dir)
				m.pointer = 0
				m.scrollTop = 0
			}
		}
	}
	m.updateScroll()
	return m, nil
}

func (m model) View() string {
	var b strings.Builder
	cellWidth := 30
	columns := m.width / cellWidth
	if columns < 1 {
		columns = 1
	}
	headerHeight := 8
	footerHeight := 8
	visibleRows := m.height - headerHeight - footerHeight
	if visibleRows < 3 {
		visibleRows = 3
	}
	startRow := m.scrollTop
	endRow := startRow + visibleRows
	startIndex := startRow * columns
	endIndex := endRow * columns
	if endIndex > len(m.fileList) {
		endIndex = len(m.fileList)
	}
	b.WriteString("📂 Current Directory  : ")
	b.WriteString(m.dir)
	b.WriteString(" - Files Count : ")
	b.WriteString(strconv.Itoa(len(m.fileList)))
	b.WriteString("\n\n")
	if startRow > 0 {
		b.WriteString("▲ More files above\n")
	}
	var grid strings.Builder
	for index := startIndex; index < endIndex; index++ {
		item := m.fileList[index]
		name := item.Name()
		if len(name) > 24 {
			name = name[:21] + "..."
		}
		if item.IsDir() {
			name = "📁 " + name
		} else {
			name = "📄 " + name
		}
		var cell string
		if index == m.pointer {
			cell = selectedStyle.
				Width(cellWidth - 2).
				Render(name)
		} else {
			cell = normalStyle.
				Width(cellWidth - 2).
				Render(name)
		}
		grid.WriteString(cell)
		if (index-startIndex+1)%columns == 0 {
			grid.WriteString("\n")
		}
	}
	b.WriteString(
		fileBoxStyle.Render(grid.String()),
	)
	if endIndex < len(m.fileList) {
		b.WriteString("\n▼ More files below")
	}
	b.WriteString("\n\n")
	if len(m.fileList) > 0 {
		currentRow := m.pointer/columns + 1
		totalRows := (len(m.fileList) + columns - 1) / columns
		selected := m.fileList[m.pointer]
		b.WriteString(
			fmt.Sprintf(
				"Selected : %d/%d | Row %d/%d | %s",
				m.pointer+1,
				len(m.fileList),
				currentRow,
				totalRows,
				selected.Name(),
			),
		)
	}

	b.WriteString("\n\n")
	b.WriteString("'↑ ↓ ← →' Navigate  ,'Enter' Open Folder ,'Backspace' Parent Folder, 'q or ctrl + c' Quit")
	return b.String()
}

func StartScreen() {
	dir, err := GetCurrentDir()
	if err != nil {
		panic(err)
	}

	m := model{
		dir:      dir,
		fileList: GetFoldersList(dir),
	}

	p := tea.NewProgram(m, tea.WithAltScreen())

	finalModel, err := p.Run()
	if err != nil {
		panic(err)
	}

	selected := finalModel.(model)

	fmt.Println("Target Path :", selected.dir)

	tmpFile := filepath.Join(os.TempDir(), "xlr8-cwd")

	err = os.WriteFile(
		tmpFile,
		[]byte(selected.dir),
		0644,
	)
	if err != nil {
		panic(err)
	}
}
