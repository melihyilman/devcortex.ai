package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func VimCheatsheetTool(w http.ResponseWriter, r *http.Request) {
	cheatsheetData := map[string][]Command{
		"Essential Commands": {
			{Command: "i", Description: "Enter Insert mode."},
			{Command: "Esc", Description: "Exit Insert mode and return to Normal mode."},
			{Command: ":w", Description: "Write (save) the file."},
			{Command: ":q", Description: "Quit."},
			{Command: ":q!", Description: "Quit without saving changes."},
			{Command: ":wq", Description: "Write and quit."},
		},
		"Navigation (Normal Mode)": {
			{Command: "h", Description: "Move left."},
			{Command: "j", Description: "Move down."},
			{Command: "k", Description: "Move up."},
			{Command: "l", Description: "Move right."},
			{Command: "w", Description: "Move to the start of the next word."},
			{Command: "b", Description: "Move to the start of the previous word."},
			{Command: "gg", Description: "Go to the first line of the file."},
			{Command: "G", Description: "Go to the last line of the file."},
		},
		"Editing (Normal Mode)": {
			{Command: "x", Description: "Delete the character under the cursor."},
			{Command: "dd", Description: "Delete the current line."},
			{Command: "yy", Description: "Yank (copy) the current line."},
			{Command: "p", Description: "Paste the copied/deleted text after the cursor."},
			{Command: "u", Description: "Undo the last change."},
			{Command: "Ctrl + r", Description: "Redo the last change."},
		},
		"Search and Replace": {
			{Command: "/pattern", Description: "Search forward for 'pattern'."},
			{Command: "?pattern", Description: "Search backward for 'pattern'."},
			{Command: "n", Description: "Repeat the search in the same direction."},
			{Command: "N", Description: "Repeat the search in the opposite direction."},
			{Command: ":%s/old/new/g", Description: "Replace all occurrences of 'old' with 'new' in the entire file."},
			{Command: ":%s/old/new/gc", Description: "Replace with confirmation for each occurrence."},
		},
	}

	data := &view.PageData{
		Title:            "Vim Cheatsheet",
		ToolSpecificData: cheatsheetData,
	}
	view.Render(w, r, "vim-cheatsheet.html", data)
}
