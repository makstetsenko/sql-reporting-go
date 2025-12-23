package sql_db

import (
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type SqlExecResult struct {
	Cols []string
	Rows [][]interface{}
}

func (s *SqlExecResult) DrawTable() *string {
	t := getTable(s)

	(*t).AppendSeparator()

	colConfigs := make([]table.ColumnConfig, len(s.Cols))

	for i, c := range s.Cols {
		colConfigs[i] = table.ColumnConfig{
			Name:        c,
			AlignHeader: text.AlignLeft,
			Align:       text.AlignLeft,
		}
	}

	(*t).SetColumnConfigs(colConfigs)

	(*t).SetStyle(table.Style{
		Name: "MinimalNoBorder",
		Box: table.BoxStyle{
			MiddleHorizontal: "=",
			MiddleVertical:   "    ",
		},
		Color: table.ColorOptions{
			Header: text.Colors{},
			Row:    text.Colors{},
		},
		Options: table.Options{
			DrawBorder:      false,
			SeparateColumns: true,
			SeparateHeader:  true,
			SeparateRows:    false,
		},
	})

	res := (*t).Render()
	return &res
}

func (s *SqlExecResult) DrawCsv() *string {
	t := getTable(s)

	res := (*t).RenderCSV()
	return &res
}

func getTable(s *SqlExecResult) *table.Writer {
	t := table.NewWriter()

	headerRow := make([]interface{}, len(s.Cols))
	for i, c := range s.Cols {
		headerRow[i] = c
	}

	t.AppendHeader(headerRow)

	for _, row := range s.Rows {
		t.AppendRow(row)
	}

	if len(s.Rows) == 0 {
		emptyRow := make([]interface{}, len(s.Cols))
		emptyRow[0] = "NO DATA"
		t.AppendRow(emptyRow)
	}
	return &t
}
