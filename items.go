package main

import "time"

type item struct {
	task        string
	isCompleted bool
	startedAt   time.Time
	finishedAt  time.Time
}

type Todos []item
