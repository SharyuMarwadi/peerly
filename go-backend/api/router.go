package api

import (
	"fmt"
	"net/http"

	"joshsoftware/peerly/config"
	"joshsoftware/peerly/service"

	"github.com/gorilla/mux"
)

const (
	versionHeader = "Accept-Version"
	authHeader    = "X-Auth-Token"
)

// InitRouter -  The routing mechanism. Mux helps us define handler functions and the access methods
func InitRouter(deps service.Dependencies) (router *mux.Router) {
	router = mux.NewRouter()

	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)
	// Version 1 API management
	v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())

	//core values
	router.Handle("/organisations/{organisation_id:[0-9]+}/core_values/{id:[0-9]+}", jwtAuthMiddleware(getCoreValueHandler(deps.CoreValueService))).Methods(http.MethodGet).Headers(versionHeader, v1)

	router.Handle("/organisations/{organisation_id:[0-9]+}/core_values", jwtAuthMiddleware(listCoreValuesHandler(deps.CoreValueService))).Methods(http.MethodGet).Headers(versionHeader, v1)

	router.Handle("/organisations/{organisation_id:[0-9]+}/core_values", jwtAuthMiddleware(createCoreValueHandler(deps.CoreValueService))).Methods(http.MethodPost).Headers(versionHeader, v1)

	router.Handle("/organisations/{organisation_id:[0-9]+}/core_values/{id:[0-9]+}", jwtAuthMiddleware(deleteCoreValueHandler(deps.CoreValueService))).Methods(http.MethodDelete).Headers(versionHeader, v1)

	router.Handle("/organisations/{organisation_id:[0-9]+}/core_values/{id:[0-9]+}", jwtAuthMiddleware(updateCoreValueHandler(deps.CoreValueService))).Methods(http.MethodPut).Headers(versionHeader, v1)

	return
}

// JWT token verification
func jwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		next.ServeHTTP(w, r)
	})
}
