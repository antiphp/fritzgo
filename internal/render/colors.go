package render

import "github.com/charmbracelet/lipgloss"

const blue = "#1577c4"

var (
	// Blue contains a blue similar or equal to the FRITZ!Box blue.
	Blue = lipgloss.NewStyle().Foreground(lipgloss.Color(blue)).Bold(true)
)
