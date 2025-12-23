package sql_db

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

func ReadSql(path *string) *string {
	data, err := os.ReadFile(*path)

	if err != nil {
		log.Fatalf("Cannot read sql script %v: %v", path, err)
	}

	content := string(data)
	return &content
}

func GetSqlScriptsPathList(dirPath *string) *[]string {
	path := strings.TrimFunc(*dirPath, unicode.IsSpace)
	path = strings.TrimSuffix(path, "/")

	files, err := filepath.Glob(path + "/*.sql")

	if err != nil {
		log.Fatalf("Cannot get list of sql files by path %v: %v", dirPath, err)
	}

	return &files
}
