package tasks_postgres_repositroy

import (
	"gopet/internal/core/domain"
	"time"
)

type TaskModel struct {
	ID           int
	Version      int
	Title        string
	Description  *string
	Completed    bool
	CreatedAt    time.Time
	CompletedAt  *time.Time
	AuthorUserID int
}

func tasksDomainsFromModels(taskModels []TaskModel) []domain.Task {
	domains := make([]domain.Task, len(taskModels))

	for i, model := range taskModels {
		domains[i] = domain.NewTask(
			model.ID,
			model.Version,
			model.Title,
			model.Description,
			model.Completed,
			model.CreatedAt,
			model.CompletedAt,
			model.AuthorUserID,
		)
	}

	return domains
}
