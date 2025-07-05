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

	toolData := make(map[string]interface{})

	if r.Method == http.MethodPost {
		action := r.FormValue("action")

		if action == "explain" {
			expression := strings.TrimSpace(r.FormValue("cron_expression"))
			toolData["Expression"] = expression

			cronService := tool.NewCronExplainer()
			explanation, err := cronService.Explain(expression)
			if err != nil {
				toolData["Error"] = err.Error()
			} else {
				toolData["Explanation"] = explanation
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
			toolData["GeneratedCron"] = cronString
			toolData["Input"] = map[string]string{
				"minute":       minute,
				"hour":         hour,
				"day_of_month": dayOfMonth,
				"month":        month,
				"day_of_week":  dayOfWeek,
			}
		}
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
