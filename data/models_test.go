package data

import (
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	orm "github.com/upper/db/v4"
)

func TestNew(t *testing.T) {
	fakeDB, _, _ := sqlmock.New()
	defer fakeDB.Close()

	_ = os.Setenv("DATABASE_TYPE", "postgres")
	m := New(fakeDB)

	if fmt.Sprintf("%T", m) != "data.Models" {
		t.Error("Wrong Type", fmt.Sprintf("%T", m))

	}

	_ = os.Setenv("DATABASE_TYPE", "mysql")
	m = New(fakeDB)

	if fmt.Sprintf("%T", m) != "data.Models" {
		t.Error("Wrong Type", fmt.Sprintf("%T", m))

	}

}

func TestGestInsertID(t *testing.T) {
	var id orm.ID
	id = int64(1)

	returnedID := getInsertID(id)
	if fmt.Sprintf("%T", returnedID) != "int" {
		t.Error("wrong type returned")
	}

	id = 1

	returnedID = getInsertID(id)
	if fmt.Sprintf("%T", returnedID) != "int" {
		t.Error("wrong type returned")
	}
}
