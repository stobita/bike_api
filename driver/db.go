package driver

import (
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stobita/bike_api/adapter/gateway"
)

type SqlHandler struct {
	Conn *sqlx.DB
}

func NewDBConn() gateway.SqlHandler {
	driver := "mysql"
	dbUser := os.Getenv("BIKE_DB_USER")
	dbPassword := os.Getenv("BIKE_DB_PASS")
	dbHost := os.Getenv("BIKE_DB_HOST")
	dbName := os.Getenv("BIKE_DB_NAME")
	conn, err := sqlx.Open(driver, fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName))
	// conn, err := sqlx.Connect("mysql", "user=gouser dbname=bike_api")
	if err != nil {
		log.Fatalln(err)
	}
	sqlHandler := new(SqlHandler)
	sqlHandler.Conn = conn
	return sqlHandler
}

func (handler *SqlHandler) Query(query string, args ...interface{}) (gateway.Rows, error) {

	rows, err := handler.Conn.Queryx(query, args...)
	if err != nil {
		return new(SqlRow), err
	}
	sqlRow := new(SqlRow)
	sqlRow.Rows = rows
	return sqlRow, nil
}

type SqlRow struct {
	Rows *sqlx.Rows
}

func (r SqlRow) Next() bool {
	return r.Rows.Next()
}

func (r SqlRow) StructScan(dest interface{}) error {
	return r.Rows.StructScan(dest)
}
