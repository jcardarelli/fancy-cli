package test

import (
	"testing"

	db "github.com/jcardarelli/fancy-cli/database"
)

func TestSqlInit(t *testing.T) {
	dbFile := "data/testDb.sql"
	err := db.InitSqlDatabase(dbFile, "../sql/create-table.sql")
	if err != nil {
		t.Fatal("failed to initialize sqlite database")
	}
}
