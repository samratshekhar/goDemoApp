package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"goDemoApp/internal/api"
	"goDemoApp/internal/config"
	"goDemoApp/internal/datastore/dao"
	"goDemoApp/internal/datastore/db/nosql"
	"goDemoApp/internal/logger"
	"goDemoApp/internal/server"
	"net/http"
	"time"
)

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS products
(
    id SERIAL,
    name TEXT NOT NULL,
    price NUMERIC(10,2) NOT NULL DEFAULT 0.00,
    CONSTRAINT products_pkey PRIMARY KEY (id)
)`

func ensureTableExists(db *sql.DB) {
	if _, err := db.Exec(tableCreationQuery); err != nil {
		//log.Fatal(err)
	}
}

func Initialize(user, password, dbname string) *sql.DB {
	connectionString :=
		fmt.Sprintf("%s:%s@tcp(%s)/demo", user, password, dbname)

	var err error
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		//log.Fatal(err)
	}
	return db
}

var httpServer *http.Server

func main() {
	setupLogger()
	widgetHandler := api.NewWidgetHandler(setupDatabase())
	httpServer = server.InitHTTPServer(widgetHandler, ExitHandler)
	server.StartHTTPServer(httpServer)
}

func setupDatabase() dao.WidgetDAO {
	cfg := config.GetConfig()
	return nosql.InitDDB(cfg.Dynamo)
}

func setupLogger() {
	cfg := config.GetConfig()
	var env logger.Option
	var logLevel logger.Option
	if cfg.Environment == "prod" {
		env = logger.ProdEnv()
	} else {
		env = logger.NonProdEnv()
	}
	logLevel = logger.LogLevel(cfg.Loglevel)
	logger.GetLogger(env, logLevel)
}

func ExitHandler(w http.ResponseWriter, r *http.Request) {
	go kill()
	w.WriteHeader(http.StatusOK)
}

func kill() {
	log := logger.GetLogger()
	log.Infof("Stopping server in 2s")
	<-time.After(time.Second * time.Duration(2))
	log.Info("Sending msg on chan")
	err := httpServer.Close()
	if err != nil {
		return
	}
}
