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

func (t *Task) CompletedDuration() *time.Duration {
	if !t.Completed {
		return nil
	}

	if t.CompletedAt == nil {
		return nil
	}

	duration := t.CompletedAt.Sub(t.CreatedAt)

	return &duration
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

type TaskPatch struct {
	Title       Nulladble[string]
	Description Nulladble[string]
	Completed   Nulladble[bool]
}

func NewTaskPatch(
	title Nulladble[string],
	description Nulladble[string],
	completed Nulladble[bool],
) TaskPatch {
	return TaskPatch{
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}

func (p *TaskPatch) Validate() error {
	if p.Title.Set && p.Title.Value == nil {
		return fmt.Errorf(
			"Title ant be putched to null: %w",
			core_errors.ErrInvalidArgumnet,
		)
	}

	if p.Completed.Set && p.Completed.Value == nil {
		return fmt.Errorf(
			"Completed cant be patched to NULL: %w",
			core_errors.ErrInvalidArgumnet,
		)
	}

	return nil
}

func (t *Task) ApplyPatch(patch TaskPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("validate task patch: %w", err)
	}

	tmp := *t

	if patch.Title.Set {
		tmp.Title = *patch.Title.Value
	}

	if patch.Description.Set {
		tmp.Description = patch.Description.Value
	}

	if patch.Completed.Set {
		tmp.Completed = *patch.Completed.Value

		if tmp.Completed {
			completedAt := time.Now()
			tmp.CompletedAt = &completedAt
		} else {
			tmp.CompletedAt = nil
		}
	}

	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("validate pathced task: %w", err)
	}

	*t = tmp

	return nil
}
