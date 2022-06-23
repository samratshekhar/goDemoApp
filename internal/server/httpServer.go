package server

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"goDemoApp/deployments/logging"
	"goDemoApp/internal/api"
	"goDemoApp/internal/config"
	"goDemoApp/internal/utils"
	"net/http"
	"time"
)

func InitHttpServer(exitHandler func(http.ResponseWriter, *http.Request)) *http.Server {
	log := logging.GetLogger()
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Error setting up http server %s", r)
			panic(r)
		}
	}()

	r := mux.NewRouter()
	addRoutes(r, exitHandler)

	http.Handle("/", r)

	httpServer := http.Server{
		Addr:        "0.0.0.0:" + config.GetConfig().HttpServerConfig.Port,
		IdleTimeout: time.Duration(utils.ConvertStringToInt(config.GetConfig().HttpServerConfig.IdleTimeoutSeconds, 1000)) * time.Second,
	}
	return &httpServer
}

func StartHttpServer(httpServer *http.Server) {
	log := logging.GetLogger()
	err := httpServer.ListenAndServe()
	if err != nil {
		log.Errorf("Error in starting server: %s", err.Error())
	}
}

func addRoutes(r *mux.Router, exitHandler func(http.ResponseWriter, *http.Request)) {
	r.Handle("/metric", promhttp.Handler())
	r.HandleFunc("/health-check", api.HealthCheckHandler)
	if isDevMode() {
		r.HandleFunc("/kill", exitHandler)
	}
	r.HandleFunc("/home", api.HomeHandler)
}

func isDevMode() bool {
	env := config.GetConfig().Env
	return env == "dev" || env == "slt"
}
