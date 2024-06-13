package service

import (
	// "encoding/json"
	"net/http"

	// "joshsoftware/peerly/db"
	// "joshsoftware/peerly/pkg/dto"
	corevalues "joshsoftware/peerly/service/coreValues"
	// "github.com/gorilla/mux"
	// logger "github.com/sirupsen/logrus"
)

func listCoreValuesHandler(coreValueSvc corevalues.Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// vars := mux.Vars(req)
		// organisationID := vars["organisation_id"]

		// coreValues, msgObj := coreValueSvc.ListCoreValues(req.Context(), organisationID)
		// if (msgObj != dto.MessageObject{}) {
		// 	rw.WriteHeader(http.StatusInternalServerError)
		// 	repsonse(rw, http.StatusInternalServerError, errorResponse{
		// 		Error: msgObj,
		// 	})
		// 	return
		// }

		repsonse(rw, http.StatusOK, successResponse{})
	})
}

func getCoreValueHandler(coreValueSvc corevalues.Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// vars := mux.Vars(req)

		// coreValue, msgObj := coreValueSvc.GetCoreValue(req.Context(), vars["organisation_id"], vars["id"])
		// if (msgObj != dto.MessageObject{}) {

		// 	repsonse(rw, http.StatusInternalServerError, errorResponse{
		// 		Error: msgObj,
		// 	})
		// 	return
		// }

		repsonse(rw, http.StatusOK, successResponse{})
	})
}

func createCoreValueHandler(coreValueSvc corevalues.Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// vars := mux.Vars(req)
		// var coreValue db.CoreValue
		// err := json.NewDecoder(req.Body).Decode(&coreValue)
		// if err != nil {
		// 	logger.WithField("err", err.Error()).Error("Error while decoding request data")
		// 	repsonse(rw, http.StatusBadRequest, errorResponse{
		// 		Error: messageObject{
		// 			Message: "Invalid json request body",
		// 		},
		// 	})
		// 	return
		// }

		// resp, msgObj := coreValueSvc.CreateCoreValue(req.Context(), vars["organisation_id"], coreValue)
		// if (msgObj != dto.MessageObject{}) {

		// 	repsonse(rw, http.StatusInternalServerError, errorResponse{
		// 		Error: msgObj,
		// 	})
		// 	return
		// }

		// ok, errFields := coreValue.Validate(req.Context(), deps.Store, organisationID)
		// if !ok {
		// 	repsonse(rw, http.StatusBadRequest, errorResponse{
		// 		Error: errorObject{
		// 			Code:          "invalid-core-value",
		// 			Fields:        errFields,
		// 			messageObject: messageObject{"Invalid core value data"},
		// 		},
		// 	})
		// 	return
		// }

		// resp, err := deps.Store.CreateCoreValue(req.Context(), organisationID, coreValue)
		// if err != nil {
		// 	logger.WithField("err", err.Error()).Error("Error while creating core value")
		// 	repsonse(rw, http.StatusInternalServerError, errorResponse{
		// 		Error: messageObject{
		// 			Message: "Internal server error",
		// 		},
		// 	})
		// 	return
		// }

		repsonse(rw, http.StatusCreated, successResponse{})
	})
}

func deleteCoreValueHandler(coreValueSvc corevalues.Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// vars := mux.Vars(req)

		// msgObj := coreValueSvc.DeleteCoreValue(req.Context(), vars["organisation_id"], vars["id"])
		// if (msgObj != dto.MessageObject{}) {

		// 	repsonse(rw, http.StatusInternalServerError, errorResponse{
		// 		Error: msgObj,
		// 	})
		// 	return
		// }

		// organisationID, err := strconv.ParseInt(vars["organisation_id"], 10, 64)
		// if err != nil {
		// 	logger.WithField("err", err.Error()).Error("Error while parsing organisation_id from url")
		// 	rw.WriteHeader(http.StatusBadRequest)
		// 	return
		// }

		// coreValueID, err := strconv.ParseInt(vars["id"], 10, 64)
		// if err != nil {
		// 	logger.WithField("err", err.Error()).Error("Error while parsing core value id from url")
		// 	rw.WriteHeader(http.StatusBadRequest)
		// 	return
		// }

		// err = deps.Store.DeleteCoreValue(req.Context(), organisationID, coreValueID)
		// if err != nil {
		// 	logger.WithField("err", err.Error()).Error("Error while deleting core value")
		// 	repsonse(rw, http.StatusInternalServerError, errorResponse{
		// 		Error: messageObject{
		// 			Message: "Internal server error",
		// 		},
		// 	})
		// 	return
		// }

		repsonse(rw, http.StatusOK, nil)
	})
}

func updateCoreValueHandler(coreValueSvc corevalues.Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		// vars := mux.Vars(req)
		// organisationID, err := strconv.ParseInt(vars["organisation_id"], 10, 64)
		// if err != nil {
		// 	logger.WithField("err", err.Error()).Error("Error while parsing organisation_id from url")
		// 	rw.WriteHeader(http.StatusBadRequest)
		// 	return
		// }

		// coreValueID, err := strconv.ParseInt(vars["id"], 10, 64)
		// if err != nil {
		// 	logger.WithField("err", err.Error()).Error("Error while parsing core value id from url")
		// 	rw.WriteHeader(http.StatusBadRequest)
		// 	return
		// }

		// var coreValue db.CoreValue
		// err := json.NewDecoder(req.Body).Decode(&coreValue)
		// if err != nil {
		// 	logger.WithField("err", err.Error()).Error("Error while decoding request data")
		// 	repsonse(rw, http.StatusBadRequest, errorResponse{
		// 		Error: messageObject{
		// 			Message: "Invalid json request body",
		// 		},
		// 	})
		// 	return
		// }

		// ok, errFields := coreValue.Validate(req.Context(), deps.Store, organisationID)
		// if !ok {
		// 	repsonse(rw, http.StatusBadRequest, errorResponse{
		// 		Error: errorObject{
		// 			Code:          "invalid-core-value",
		// 			Fields:        errFields,
		// 			messageObject: messageObject{"Invalid core value data"},
		// 		},
		// 	})
		// 	return
		// }

		// resp, err := deps.Store.UpdateCoreValue(req.Context(), organisationID, coreValueID, coreValue)
		// if err != nil {
		// 	logger.WithField("err", err.Error()).Error("Error while updating core value")
		// 	repsonse(rw, http.StatusInternalServerError, errorResponse{
		// 		Error: messageObject{
		// 			Message: "Internal server error",
		// 		},
		// 	})
		// 	return
		// }

		// resp, msgObj := coreValueSvc.UpdateCoreValue(req.Context(), vars["organisation_id"], vars["id"], coreValue)
		// if (msgObj != dto.MessageObject{}) {

		// 	repsonse(rw, http.StatusInternalServerError, errorResponse{
		// 		Error: msgObj,
		// 	})
		// 	return
		// }

		repsonse(rw, http.StatusOK, successResponse{})
	})
}
