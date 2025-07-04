package tool

import (
	"fmt"
	"strings"

	"github.com/robfig/cron/v3"
)

// CronExplainer defines the interface for the cron explainer service.
type CronExplainer interface {
	Explain(expression string) (string, error)
}

// cronExplainer is the implementation of the CronExplainer interface.
type cronExplainer struct{}

// NewCronExplainer creates a new CronExplainer service.
func NewCronExplainer() CronExplainer {
	return &cronExplainer{}
}

// Explain validates and generates a human-readable explanation for a cron expression.
func (s *cronExplainer) Explain(expression string) (string, error) {
	if len(strings.Fields(expression)) != 5 {
		return "", fmt.Errorf("please enter a valid 5-field cron expression")
	}

	_, err := cron.ParseStandard(expression)
	if err != nil {
		return "", fmt.Errorf("invalid cron expression: %w", err)
	}

	return s.generateCronExplanation(expression), nil
}

func (s *cronExplainer) generateCronExplanation(expression string) string {
	fields := strings.Fields(expression)
	var parts []string

	// Minute
	minute := fields[0]
	if strings.HasPrefix(minute, "*/") {
		parts = append(parts, fmt.Sprintf("every %s minutes", minute[2:]))
	} else if minute == "*" {
		parts = append(parts, "every minute")
	} else {
		parts = append(parts, fmt.Sprintf("at minute %s", minute))
	}

	// Hour
	hour := fields[1]
	if strings.HasPrefix(hour, "*/") {
		parts = append(parts, fmt.Sprintf("every %s hours", hour[2:]))
	} else if hour == "*" {
		parts = append(parts, "every hour")
	} else {
		parts = append(parts, fmt.Sprintf("at hour %s", hour))
	}

	// Day of Month
	dayOfMonth := fields[2]
	if dayOfMonth != "*" {
		parts = append(parts, fmt.Sprintf("on day-of-month %s", dayOfMonth))
	}

	// Month
	month := fields[3]
	if month != "*" {
		parts = append(parts, fmt.Sprintf("in month %s", month))
	}

	// Day of Week
	dayOfWeek := fields[4]
	if dayOfWeek != "*" {
		dayMap := map[string]string{
			"0": "Sunday", "1": "Monday", "2": "Tuesday", "3": "Wednesday", "4": "Thursday", "5": "Friday", "6": "Saturday", "7": "Sunday",
		}
		if dayName, ok := dayMap[dayOfWeek]; ok {
			parts = append(parts, fmt.Sprintf("on %s", dayName))
		} else {
			parts = append(parts, fmt.Sprintf("on day-of-week %s", dayOfWeek))
		}
	}

	if fields[2] == "*" && fields[4] == "*" {
		parts = append(parts, "every day")
	}

	return "Runs " + strings.Join(parts, ", ") + "."
}
