package handler

import (
	"fmt"
	"net/http"
	"strings"

	"devcortex.ai/internal/tool"
	"devcortex.ai/internal/view"
)

func CronTool(w http.ResponseWriter, r *http.Request) {
	pageData := &view.PageData{
		Title: "Cron Job Explainer & Generator",
	}

	toolData := map[string]interface{}{
		"Expression":    r.URL.Query().Get("Expression"),
		"Error":         r.URL.Query().Get("Error"),
		"Explanation":   r.URL.Query().Get("Explanation"),
		"GeneratedCron": r.URL.Query().Get("GeneratedCron"),
		"Input": map[string]string{
			"minute":       r.URL.Query().Get("minute"),
			"hour":         r.URL.Query().Get("hour"),
			"day_of_month": r.URL.Query().Get("day_of_month"),
			"month":        r.URL.Query().Get("month"),
			"day_of_week":  r.URL.Query().Get("day_of_week"),
		},
	}

	if r.Method == http.MethodPost {
		action := r.FormValue("action")
		redirectData := make(map[string]string)

		if action == "explain" {
			expression := strings.TrimSpace(r.FormValue("cron_expression"))
			redirectData["Expression"] = expression

			cronService := tool.NewCronExplainer()
			explanation, err := cronService.Explain(expression)
			if err != nil {
				redirectData["Error"] = err.Error()
			} else {
				redirectData["Explanation"] = explanation
			}
		} else if action == "generate" {
			minute := r.FormValue("minute")
			hour := r.FormValue("hour")
			dayOfMonth := r.FormValue("day_of_month")
			month := r.FormValue("month")
			dayOfWeek := r.FormValue("day_of_week")

			cronString := fmt.Sprintf("%s %s %s %s %s",
				getPart(minute),
				getPart(hour),
				getPart(dayOfMonth),
				getPart(month),
				getPart(dayOfWeek),
			)
			redirectData["GeneratedCron"] = cronString
			redirectData["minute"] = minute
			redirectData["hour"] = hour
			redirectData["day_of_month"] = dayOfMonth
			redirectData["month"] = month
			redirectData["day_of_week"] = dayOfWeek
		}
		redirectToPageWithData(w, r, redirectData)
		return
	}

	pageData.ToolSpecificData = toolData
	view.Render(w, r, "cron.html", pageData)
}

func getPart(value string) string {
	if value == "" {
		return "*"
	}
	return value
}
