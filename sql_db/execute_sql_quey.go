package sql_db

import (
	"database/sql"
	"log"
)

func ExecSqlQuery(db *sql.DB, sqlQuery *string) (*SqlExecResult, error) {
	rows, err := getSqlRows(db, sqlQuery)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	rowsData := getRowsData(rows, cols)

	sqlExecResult := SqlExecResult{
		Cols: cols,
		Rows: rowsData,
	}

	return &sqlExecResult, nil
}

func getRowsData(rows *sql.Rows, cols []string) [][]interface{} {
	var rowsData [][]interface{}

	for rows.Next() {
		values := make([]interface{}, len(cols))
		valuePointers := make([]interface{}, len(cols))

		for i := range cols {
			valuePointers[i] = &values[i]
		}

		if err := rows.Scan(valuePointers...); err != nil {
			log.Fatalf("Cannot read row data: %v\n", err)
		}

		for i, val := range values {
			if b, ok := val.([]byte); ok {
				values[i] = string(b)
			}
		}

		rowsData = append(rowsData, values)
	}
	return rowsData
}

func getSqlRows(db *sql.DB, sqlQuery *string) (*sql.Rows, error) {
	rows, err := db.Query(*sqlQuery)
	if err != nil {
		return nil, err
	}
	return rows, nil
}
