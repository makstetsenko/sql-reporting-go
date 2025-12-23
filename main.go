package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"strings"

	"github.com/makstetsenko/sql-reporting-go/args_parser"
	"github.com/makstetsenko/sql-reporting-go/sql_db"

	_ "github.com/microsoft/go-mssqldb"
)

func main() {
	args := args_parser.Parse()
	connections := sql_db.GetDbConnections(&args.DbConnectionsConfig)

	for _, c := range connections {
		db := sql_db.ConnectToDb(c)
		defer db.Close()

		reportsPath := args.SqlScriptsPath + "/reports/" + c.Env

		if err := os.MkdirAll(reportsPath, 0755); err != nil {
			log.Fatalf("Cannot create reports directory %v", path.Clean(reportsPath))
		}

		log.Printf("Connected to %v, DB: %v\n", c.Server, c.Database)

		sqlFiles := sql_db.GetSqlScriptsPathList(&args.SqlScriptsPath)
		reportTables := make([]string, len(*sqlFiles))

		for i, sqlPath := range *sqlFiles {
			log.Println(sqlPath)

			sqlQuery := sql_db.ReadSql(&sqlPath)

			execResult, err := sql_db.ExecSqlQuery(db, sqlQuery)

			if err != nil {
				log.Printf("Cannot execute query %v\n%v\n", sqlPath, err)
				continue
			}

			_, sqlFileName := path.Split(sqlPath)
			reportTables[i] = fmt.Sprintf("# %v\n\n%v\n", sqlFileName, *execResult.DrawTable())

			csvFilePath := reportsPath + "/" + sqlFileName + ".csv"

			if err := os.WriteFile(csvFilePath, []byte(*execResult.DrawCsv()), 0644); err != nil {
				log.Printf("Cannot write %v: %v\n", csvFilePath, err)
			}

		}

		res := strings.Join(reportTables, fmt.Sprintf("\n\n\n\n%v\n\n\n\n", strings.Repeat("-", 80)))

		if err := os.WriteFile(reportsPath+"/report.txt", []byte(res), 0644); err != nil {
			log.Printf("Cannot write report.txt: %v\n", err)
		}

		err := db.Close()
		if err != nil {
			log.Printf("Cannot close connection: %v", err)
			continue
		}
	}
}
