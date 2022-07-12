package server

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"goDemoApp/internal/api"
	"goDemoApp/internal/config"
	"goDemoApp/internal/logger"
	"goDemoApp/internal/utils"
	"net/http"
	"time"
)

func InitHTTPServer(widgetHandler api.WidgetHandler, exitHandler func(http.ResponseWriter, *http.Request)) *http.Server {
	log := logger.GetLogger()
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Error setting up http server %s", r)
			panic(r)
		}
	}()

	r := mux.NewRouter()
	addRoutes(r, widgetHandler, exitHandler)

	http.Handle("/", r)

	httpServer := http.Server{
		Addr:        "0.0.0.0:" + config.GetConfig().HTTPServerConfig.Port,
		IdleTimeout: time.Duration(utils.ConvertStringToInt(config.GetConfig().HTTPServerConfig.IdleTimeoutSeconds, 1000)) * time.Second,
	}
	return &httpServer
}

func StartHTTPServer(httpServer *http.Server) {
	log := logger.GetLogger()
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Errorf("Error in starting server: %s", err.Error())
	}
}

func addRoutes(r *mux.Router, widgetHandler api.WidgetHandler, exitHandler func(http.ResponseWriter, *http.Request)) {
	r.Handle("/metric", promhttp.Handler())
	r.HandleFunc("/health-check", api.HealthCheckHandler)
	if isDevMode() {
		r.HandleFunc("/kill", exitHandler)
	}
	r.HandleFunc("/home", api.HomeHandler)
	r.HandleFunc("/widget", widgetHandler.CreateWidget).Methods("POST")
}

func isDevMode() bool {
	env := config.GetConfig().Environment
	return env == "dev" || env == "slt"
}
