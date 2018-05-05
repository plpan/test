package main

import (
	"database/sql"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestQuerySingleRow(t *testing.T) {
	t.Parallel()
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(1).
		AddRow(2)
	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	var id int
	if err := db.QueryRow("SELECT").Scan(&id); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{"id"}))
	if err := db.QueryRow("SELECT").Scan(&id); err != sql.ErrNoRows {
		t.Fatal("expected sql no rows error")
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Fatal(err)
	}
}
