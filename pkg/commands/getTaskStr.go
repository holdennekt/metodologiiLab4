package commands

import (
	"fmt"

	dataproviders "github.com/holdennekt/metodologiiLab4/pkg/dataProviders"
)

func getTaskStr(t *dataproviders.Task) string {
	var details string
	if t.Details.Valid {
		details = t.Details.String
	} else {
		details = "no details"
	}
	var deadline string
	if t.Deadline.Valid {
		year := t.Deadline.Time.Year()
		month := t.Deadline.Time.Month().String()
		day := t.Deadline.Time.Day()
		deadline = fmt.Sprintf("%v %v %v", day, month, year)
	} else {
		deadline = "no deadline"
	}
	var completedAt string
	if t.CompletedAt.Valid {
		year := t.CompletedAt.Time.Year()
		month := t.CompletedAt.Time.Month().String()
		day := t.CompletedAt.Time.Day()
		completedAt = fmt.Sprintf("%v %v %v", day, month, year)
	} else {
		completedAt = "wasn't completed"
	}
	taskStr := fmt.Sprintf("id: %v, title: %v\ndetails: %v\ndeadline: %v, expired: %v\ncompleted: %v, completedAt: %v\n", t.Id, t.Title, details, deadline, t.Expired, t.Completed, completedAt)
	return taskStr
}
