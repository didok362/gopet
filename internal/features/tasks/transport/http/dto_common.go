package tasks_transport_http

import (
	"gopet/internal/core/domain"
	"time"
)

type TaskDTOResponse struct {
	ID           int        `json:"id"             example:"123" `
	Version      int        `json:"version"        example:"12"`
	Title        string     `json:"title"          example:"homework"`
	Description  *string    `json:"description"    example:"i need to do my homework"`
	Completed    bool       `json:"completed"      example:"false"`
	CreatedAt    time.Time  `json:"created_at"     example:"2026-02-26T10:30:00Z"`
	CompletedAt  *time.Time `json:"completed_at"   example:"null"`
	AuthorUserID int        `json:"author_user_id" example:"5"`
}

func taskDTOFromDomain(task domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
		AuthorUserID: task.AuthorUserID,
	}
}

func tasksDTOsFromDomains(tasks []domain.Task) []TaskDTOResponse {
	DTOs := make([]TaskDTOResponse, len(tasks))
	for i, task := range tasks {
		DTOs[i] = taskDTOFromDomain(task)
	}
	return DTOs
}
