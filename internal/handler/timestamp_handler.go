package handler

import (
	"net/http"
	"strconv"
	"time"

	"devcortex.ai/internal/view"
)

func TimestampTool(w http.ResponseWriter, r *http.Request) {
	pageData := &view.PageData{
		Title: "Date/Time Converter",
	}

	// Provide a list of common timezones to the template
	timezones := []string{"UTC", "America/New_York", "Europe/London", "Asia/Tokyo", "Australia/Sydney"}

	// Default view data
	toolData := map[string]interface{}{
		"Timezones": timezones,
		"Now":       time.Now().Unix(),
	}

	if r.Method == http.MethodPost {
		action := r.FormValue("action")
		timezoneStr := r.FormValue("timezone")
		loc, err := time.LoadLocation(timezoneStr)
		if err != nil {
			loc = time.UTC // Default to UTC on error
		}

		if action == "to_date" {
			tsStr := r.FormValue("timestamp")
			toolData["TimestampInput"] = tsStr
			ts, err := strconv.ParseInt(tsStr, 10, 64)
			if err != nil {
				toolData["Error"] = "Please enter a valid numeric timestamp."
			} else {
				t := time.Unix(ts, 0).In(loc)
				toolData["ReadableDate"] = t.Format("2006-01-02 15:04:05 MST")
			}
		} else if action == "to_timestamp" {
			dateStr := r.FormValue("date_string")
			toolData["DateStringInput"] = dateStr
			// Use a more flexible layout that doesn't assume timezone
			const layout = "2006-01-02 15:04:05"
			t, err := time.ParseInLocation(layout, dateStr, loc)
			if err != nil {
				toolData["Error"] = "Please enter a valid date in 'YYYY-MM-DD HH:MM:SS' format."
			} else {
				toolData["TimestampResult"] = t.Unix()
			}
		}
		toolData["SelectedTimezone"] = timezoneStr
	}

	pageData.ToolSpecificData = toolData
	view.Render(w, r, "timestamp.html", pageData)
}
