package utils

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// BorderedBoxWithTitle creates a border around some content and adds a title
// to the top border.
// `style` should be the style used to create the bordered content without the title.
// `style` can have padding, but no margins (won't look right).
func BorderedBoxWithTitle(style lipgloss.Style, content, title string) string {
	style = style.UnsetBorderTop()

	border := style.GetBorderStyle()
	box := style.Render(content)

	minBorderTop := fmt.Sprintf(
		"%s%s %s %s%s", // e.g. ╭─ Config ─╮
		border.TopLeft,
		border.Top,
		title,
		border.Top,
		border.TopRight,
	)

	minBorderTopLen := lipgloss.Width(minBorderTop)
	boxLen := lipgloss.Width(box)

	if minBorderTopLen > boxLen {
		// TODO: add only enough padding to meet the min length
		padding := minBorderTopLen - boxLen
		// Add padding to right side so it meets the min length, and re-render
		style = style.PaddingRight(style.GetHorizontalPadding() + padding)
		box = style.Render(content)
		boxLen = lipgloss.Width(box)
	}

	repeatCount := boxLen - minBorderTopLen + 1 // I don't know why +1 but it works ...

	finalBorderTop := fmt.Sprintf(
		"%s%s %s %s%s",
		border.TopLeft,
		border.Top,
		title,
		strings.Repeat(border.Top, repeatCount),
		border.TopRight,
	)

	return lipgloss.JoinVertical(
		lipgloss.Left,
		finalBorderTop,
		box,
	)
}

func Max(values ...int) int {
	max := values[0]

	for _, v := range values {
		if v > max {
			max = v
		}
	}

	return max
}

func IntToString(v int) string {
	return strconv.Itoa(v)
}

func StringToInt(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

// RemoveSliceItem removes the first occurrence of a specific element
// from a slice.
func RemoveSliceItem[T comparable](slice []T, d T) []T {
	for i, v := range slice {
		if v == d {
			if i == len(slice)-1 {
				return slice[:i]
			} else {
				return append(slice[:i], slice[i+1:]...)
			}
		}
	}

	// No changes
	return slice
}
