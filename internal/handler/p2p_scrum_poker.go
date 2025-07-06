package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func P2PScrumPokerHandler(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title:            "P2P Scrum Poker",
		ToolSpecificData: make(map[string]interface{}),
	}
	view.Render(w, r, "p2p-scrum-poker.html", data)
}
