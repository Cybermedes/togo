package main

import "time"

type item struct {
	Task        string    `json:"task"`
	IsCompleted bool      `json:"isCompleted"`
	StartedAt   time.Time `json:"startedAt"`
	FinishedAt  time.Time `json:"finishedAt"`
}

type Todos []item
