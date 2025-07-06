package handler

import (
	"net/http"

	"devcortex.ai/internal/view"
)

func PeerToPeerScrumPokerHandler(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title:            "Peer-to-Peer Scrum Poker",
		Description:      "Real-time, serverless Scrum Poker for agile teams using PeerJS for direct P2P communication.",
		ToolSpecificData: make(map[string]interface{}),
	}
	view.Render(w, r, "peer-to-peer-scrum-poker.html", data)
}
