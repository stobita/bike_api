package model

import (
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var engine *xorm.Engine
var err error

func init() {
	driver := "mysql"
	dbUser := os.Getenv("BIKE_DB_USER")
	dbPassword := os.Getenv("BIKE_DB_PASS")
	dbHost := os.Getenv("BIKE_DB_HOST")
	dbName := os.Getenv("BIKE_DB_NAME")
	engine, err = xorm.NewEngine(driver, fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPassword, dbHost, dbName))
	if err != nil {
		panic("failed to connect database")
	}
	engine.ShowSQL(true)
	defer engine.Clone()
}

// Model Base Model Struct
type Model struct {
	ID              int64  `xorm:"id pk" json:"id"`
	TableNameString string `xorm:"extends"`
}

// ModelInterface Common Model Interface
type modelInterface interface {
	TableName() string
}

// TableName Get TableName
func (m Model) TableName() string {
	return m.TableNameString
}

// Create Common Create
func (m *Model) Create() error {
	_, err := engine.Insert(m)
	if err != nil {
		return err
	}
	return nil
}
