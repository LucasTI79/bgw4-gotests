package utils

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	"github.com/google/uuid"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql"
)

// 1,2,3,4,5,6,7,8,9
// 12403152-75b1-49c7-aa06-94d29778db1c // 36

func RegisterTxDbDatabase(t *testing.T) {
	t.Helper()

	err := godotenv.Load("../../.env")
	if err != nil {
		panic("error laoding .env file")
	}

	connectionString := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	txdb.Register("txdb", "mysql", connectionString)
}

func InitTxDbDatabase(t *testing.T) (*sql.DB, error) {
	t.Helper()
	db, err := sql.Open("txdb", uuid.New().String())
	if err != nil {
		return nil, err
	}

	return db, nil
}
