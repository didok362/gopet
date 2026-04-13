package domain

import (
	"fmt"
	core_errors "gopet/internal/core/errors"
	"time"
)

type Task struct {
	ID      int
	Version int

	Title       string
	Description *string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time

	AuthorUserID int
}

func NewTask(
	id int,
	version int,
	tittle string,
	description *string,
	completed bool,
	createdAt time.Time,
	completedAt *time.Time,
	authorUserId int,
) Task {
	return Task{
		ID:           id,
		Version:      version,
		Title:        tittle,
		Description:  description,
		Completed:    completed,
		CreatedAt:    createdAt,
		CompletedAt:  completedAt,
		AuthorUserID: authorUserId,
	}
}

func NewTaskUninitialized(
	title string,
	desription *string,
	authorUserID int,
) Task {
	return NewTask(
		UninitializedID,
		UninitializedVersion,
		title,
		desription,
		false,
		time.Now(),
		nil,
		authorUserID,
	)
}

func (t *Task) Validate() error {
	taitlen := len([]rune(t.Title))
	if taitlen < 1 || taitlen > 100 {
		return fmt.Errorf(
			"invaluid title len: %d: %w",
			taitlen,
			core_errors.ErrInvalidArgumnet,
		)
	}
	if t.Description != nil {
		descriptionlen := len([]rune(*t.Description))
		if descriptionlen < 1 || descriptionlen > 1000 {
			return fmt.Errorf(
				"invaluid description len: %d: %w",
				descriptionlen,
				core_errors.ErrInvalidArgumnet,
			)
		}
	}
	if (t.Completed == true && t.CompletedAt == nil) || (t.Completed == true && t.CompletedAt.Before(t.CreatedAt) || (t.CompletedAt != nil && t.Completed == false)) {
		return fmt.Errorf(
			"invaluid CompletedAt: %w",
			core_errors.ErrInvalidArgumnet,
		)
	}
	return nil
}
