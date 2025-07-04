package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

type Command struct {
	Command     string
	Description string
}

func GitCheatsheetTool(w http.ResponseWriter, r *http.Request) {
	cheatsheetData := map[string][]Command{
		"Setup": {
			{Command: "git config --global user.name \"[name]\"", Description: "Sets the name you want attached to your commit transactions."},
			{Command: "git config --global user.email \"[email address]\"", Description: "Sets the email you want attached to your commit transactions."},
			{Command: "git config --global color.ui auto", Description: "Enables helpful colorization of command line output."},
		},
		"Creating Projects": {
			{Command: "git init [project-name]", Description: "Creates a new local repository with the specified name."},
			{Command: "git clone [url]", Description: "Downloads a project and its entire version history."},
		},
		"Making Changes": {
			{Command: "git status", Description: "Lists all new or modified files to be committed."},
			{Command: "git add [file]", Description: "Snapshots the file in preparation for versioning."},
			{Command: "git reset [file]", Description: "Unstages the file, but preserves its contents."},
			{Command: "git diff", Description: "Shows file differences not yet staged."},
			{Command: "git diff --staged", Description: "Shows file differences between staging and the last file version."},
			{Command: "git commit -m \"[descriptive message]\"", Description: "Records file snapshots permanently in version history."},
		},
		"Branching": {
			{Command: "git branch", Description: "Lists all local branches in the current repository."},
			{Command: "git branch [branch-name]", Description: "Creates a new branch."},
			{Command: "git checkout [branch-name]", Description: "Switches to the specified branch and updates the working directory."},
			{Command: "git merge [branch-name]", Description: "Combines the specified branch’s history into the current branch."},
			{Command: "git branch -d [branch-name]", Description: "Deletes the specified branch."},
		},
		"Updating & Publishing": {
			{Command: "git remote add [alias] [url]", Description: "Adds a new remote Git repository as a shortname."},
			{Command: "git fetch [alias]", Description: "Downloads all history from the remote repository."},
			{Command: "git pull [alias] [branch]", Description: "Fetches and merges the specified remote branch into the current branch."},
			{Command: "git push [alias] [branch]", Description: "Uploads all local branch commits to the remote repository."},
		},
	}

	data := &view.PageData{
		Title:            "Git Cheatsheet",
		ToolSpecificData: cheatsheetData,
	}
	view.Render(w, r, "git-cheatsheet.html", data)
}
