package main

import (
	"fmt"
	"time"

	"github.com/alexeyco/simpletable"
)

func (t *Todos) PrintTable() {
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignRight, Text: "#"},
			{Align: simpletable.AlignRight, Text: "Task"},
			{Align: simpletable.AlignRight, Text: "Done"},
			{Align: simpletable.AlignRight, Text: "Started at"},
			{Align: simpletable.AlignRight, Text: "Finished at"},
		},
	}

	var cells [][]*simpletable.Cell

	for idx, item := range *t {
		idx++
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", idx)},
			{Text: item.Task},
			{Text: fmt.Sprintf("%t", item.IsCompleted)},
			{Text: item.StartedAt.Format(time.RFC822)},
			{Text: item.FinishedAt.Format(time.RFC822)},
		})
	}

	table.Body = &simpletable.Body{Cells: cells}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: "Your todos are here"},
		},
	}

	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}
