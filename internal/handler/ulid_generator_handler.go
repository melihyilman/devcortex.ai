package handler

import (
	"math/rand"
	"net/http"
	"time"

	"devcortex.ai/internal/view"
	"github.com/oklog/ulid/v2"
)

func ULIDGeneratorTool(w http.ResponseWriter, r *http.Request) {
	data := &view.PageData{
		Title: "ULID Generator",
	}

	if r.Method == http.MethodPost {
		t := time.Now()
		entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
		newULID := ulid.MustNew(ulid.Timestamp(t), entropy)

		data.ToolSpecificData = map[string]interface{}{
			"GeneratedULID": newULID.String(),
		}
	}

	view.Render(w, r, "ulid-generator.html", data)
}
