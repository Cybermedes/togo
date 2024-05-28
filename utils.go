package main

import (
	"errors"
	"time"
)

// Add a task to your TODO list
func (t *Todos) Add(task string) {

	todo := item{
		Task:        task,
		IsCompleted: false,
		StartedAt:   time.Now(),
		FinishedAt:  time.Time{},
	}

	*t = append(*t, todo)
}

// Delete a task from the list, regardless if it is finished or not
func (t *Todos) Delete(index int) error {
	list := *t
	if index <= 0 || index > len(list) {
		return errors.New("invalid index/ index out of range")
	}

	*t = append(list[:index-1], list[index:]...)

	return nil
}

// Set the status of a task to "completed" and update time
func (t *Todos) Complete(index int) error {
	list := *t
	if index <= 0 || index > len(list) {
		return errors.New("invalid index/ index out of range")
	}

	list[index-1].FinishedAt = time.Now()
	list[index-1].IsCompleted = true

	return nil
}
