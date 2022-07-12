package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"goDemoApp/internal/config"
	"sync"
)

var mysqlDB *sqlx.DB
var dbSync sync.Once

type Option func(*sqlx.DB, map[string]*sqlx.Stmt) error

func GetDB() *sqlx.DB {
	dbSync.Do(func() {
		dbConfig := config.GetConfig().MySQLConfig
		db, err := InitDB(dbConfig.UserName, dbConfig.Password, dbConfig.URL)
		if err != nil {
			panic(fmt.Sprintf("Error opening db connection %v", err))
		}
		mysqlDB = db
	})
	return mysqlDB
}

func InitDB(username, password, dbUrl string) (*sqlx.DB, error) {

	ds := dataSource(username, password, dbUrl)
	db, err := sql.Open("mysql", ds)

	if err != nil {
		return nil, err
	}
	return sqlx.NewDb(db, "mysql"), nil

}

//username:password@protocol(address)/dbname?param=value
func dataSource(username, password, dburl string) string {
	return fmt.Sprintf("%s:%s@tcp(%s)/demo?parseTime=true", username, password, dburl)
}
