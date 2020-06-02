package service

import (
	"fmt"
	"net/http"

	"joshsoftware/peerly/config"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

const (
	versionHeader = "Accept"
	authHeader    = "X-Auth-Token"
)

// InitRouter -  The routing mechanism. Mux helps us define handler functions and the access methods
func InitRouter(deps Dependencies) (router *mux.Router) {
	router = mux.NewRouter()

	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return config.JWTKey(), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	// No version requirement for /ping
	router.HandleFunc("/ping", pingHandler).Methods(http.MethodGet)
	// Version 1 API management
	v1 := fmt.Sprintf("application/vnd.%s.v1", config.AppName())

	//core values
	router.HandleFunc("/organisations/{organisation_id:[0-9]+}/core_values/{id:[0-9]+}", getCoreValueHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/organisations/{organisation_id:[0-9]+}/core_values", listCoreValuesHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/organisations/{organisation_id:[0-9]+}/core_values", createCoreValueHandler(deps)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/organisations/{organisation_id:[0-9]+}/core_values/{id:[0-9]+}", deleteCoreValueHandler(deps)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/organisations/{organisation_id:[0-9]+}/core_values/{id:[0-9]+}", updateCoreValueHandler(deps)).Methods(http.MethodPut).Headers(versionHeader, v1)

	//users
	router.HandleFunc("/users", listUsersHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/users/{email}", getUserByEmailHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)

	// Basic logout
	router.Handle("/logout", negroni.New(
		negroni.HandlerFunc(jwtMiddleware.HandlerWithNext),
		negroni.Wrap(http.HandlerFunc(handleLogout(deps))),
	)).Methods(http.MethodDelete).Headers(versionHeader, v1)

	// TODO: Finish login system
	router.HandleFunc("/auth/google", handleAuth(deps)).Methods(http.MethodGet)

	router.HandleFunc("/organizations/{domainName}", getOrganizationByDomainNameHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)

	router.HandleFunc("/organizations", listOrganizationHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/organizations/{id:[0-9]+}", getOrganizationHandler(deps)).Methods(http.MethodGet).Headers(versionHeader, v1)
	router.HandleFunc("/organizations", createOrganizationHandler(deps)).Methods(http.MethodPost).Headers(versionHeader, v1)
	router.HandleFunc("/organizations/{id:[0-9]+}", deleteOrganizationHandler(deps)).Methods(http.MethodDelete).Headers(versionHeader, v1)
	router.HandleFunc("/organizations/{id:[0-9]+}", updateOrganizationHandler(deps)).Methods(http.MethodPut).Headers(versionHeader, v1)
	return
}
