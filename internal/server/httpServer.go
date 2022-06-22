package server

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"goDemoApp/deployments/logging"
	"goDemoApp/internal/api"
	"net/http"
)

func InitHttpServer() {
	log := logging.GetLogger()
	defer func() {
		if r := recover(); r != nil {
			log.Errorf("Error setting up http server %s", r)
			panic(r)
		}
	}()

	r := mux.NewRouter()
	addRoutes(r)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func addRoutes(r *mux.Router) {
	r.Handle("/metric", promhttp.Handler())
	r.HandleFunc("/health-check", api.HealthCheckHandler)
	r.HandleFunc("/home", api.HomeHandler)
}