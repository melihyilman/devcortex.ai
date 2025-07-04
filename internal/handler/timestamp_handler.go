package handler

import (
	"net/http"
	"strconv"
	"time"

	"devcortex.ai/internal/view"
)

type TimestampData struct {
	TimestampInput  string
	DateStringInput string
	ReadableDate    string
	TimestampResult int64
	Error           string
}

func TimestampTool(w http.ResponseWriter, r *http.Request) {
	pageData := &view.PageData{
		Title: "Unix Timestamp Converter",
		ToolSpecificData:  map[string]interface{}{},
	}

	if r.Method == http.MethodPost {
		action := r.FormValue("action")
		data := make(map[string]interface{})

		if action == "to_date" {
			tsStr := r.FormValue("timestamp")
			data["TimestampInput"] = tsStr
			ts, err := strconv.ParseInt(tsStr, 10, 64)
			if err != nil {
				data["Error"] = "Please enter a valid numeric timestamp."
			} else {
				t := time.Unix(ts, 0)
				data["ReadableDate"] = t.Format("02-01-2006 15:04:05")
			}
		} else if action == "to_timestamp" {
			dateStr := r.FormValue("date_string")
			data["DateStringInput"] = dateStr
			const layout = "2006-01-02 15:04:05"
			t, err := time.Parse(layout, dateStr)
			if err != nil {
				data["Error"] = "Please enter a valid date in 'YYYY-MM-DD HH:MM:SS' format."
			} else {
				data["TimestampResult"] = t.Unix()
			}
		}
		pageData.ToolSpecificData = data
	}

	view.Render(w, r, "timestamp.html", pageData)
}