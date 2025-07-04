package handler

import (
	"net/http"
	"strings"

	"devcortex.ai/internal/tool"
	"devcortex.ai/internal/view"
)

type CronData struct {
	Expression  string
	Explanation string
	Error       string
}

func CronTool(w http.ResponseWriter, r *http.Request) {
	data := CronData{}
	cronService := tool.NewCronExplainer()

	if r.Method == http.MethodPost {
		expression := strings.TrimSpace(r.FormValue("cron_expression"))
		data.Expression = expression

		explanation, err := cronService.Explain(expression)
		if err != nil {
			data.Error = err.Error()
		} else {
			data.Explanation = explanation
		}
	}

	view.Render(w, r, "cron.html", &view.PageData{
		Title:            "Cron Explainer",
		ToolSpecificData: data,
	})
}