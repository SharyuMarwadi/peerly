package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"joshsoftware/peerly/apperrors"
	"joshsoftware/peerly/db"
	"joshsoftware/peerly/pkg/dto"
	corevalues "joshsoftware/peerly/service/coreValues"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func listCoreValuesHandler(coreValueSvc corevalues.Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		organisationID := vars["organisation_id"]

		coreValues, err := coreValueSvc.ListCoreValues(req.Context(), organisationID)
		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			apperrors.ErrorResp(rw, err)
			return
		}

		dto.Repsonse(rw, http.StatusOK, dto.SuccessResponse{Data: coreValues})
	})
}

func getCoreValueHandler(coreValueSvc corevalues.Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		coreValue, err := coreValueSvc.GetCoreValue(req.Context(), vars["organisation_id"], vars["id"])
		if err != nil {

			apperrors.ErrorResp(rw, err)
			return
		}

		dto.Repsonse(rw, http.StatusOK, dto.SuccessResponse{Data: coreValue})
	})
}

func createCoreValueHandler(coreValueSvc corevalues.Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		const userId int64 = 1
		var coreValue db.CoreValue
		err := json.NewDecoder(req.Body).Decode(&coreValue)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error while decoding request data")
			err = apperrors.JSONParsingErrorReq
			apperrors.ErrorResp(rw, err)
			return
		}

		fmt.Println("orgId (vars) = ", vars["organisation_id"])
		resp, err := coreValueSvc.CreateCoreValue(req.Context(), vars["organisation_id"], userId, coreValue)
		if err != nil {

			apperrors.ErrorResp(rw, err)
			return
		}

		dto.Repsonse(rw, http.StatusCreated, dto.SuccessResponse{Data: resp})
	})
}

func deleteCoreValueHandler(coreValueSvc corevalues.Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		err := coreValueSvc.DeleteCoreValue(req.Context(), vars["organisation_id"], vars["id"])
		if err != nil {

			apperrors.ErrorResp(rw, err)
			return
		}

		dto.Repsonse(rw, http.StatusOK, nil)
	})
}

func updateCoreValueHandler(coreValueSvc corevalues.Service) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)

		var updateReq dto.UpdateQueryRequest
		err := json.NewDecoder(req.Body).Decode(&updateReq)
		if err != nil {
			logger.WithField("err", err.Error()).Error("Error while decoding request data")
			err = apperrors.JSONParsingErrorReq
			apperrors.ErrorResp(rw, err)
			return
		}

		resp, err := coreValueSvc.UpdateCoreValue(req.Context(), vars["organisation_id"], vars["id"], updateReq)
		if err != nil {
			apperrors.ErrorResp(rw, err)
			return
		}

		dto.Repsonse(rw, http.StatusOK, dto.SuccessResponse{Data: resp})
	})
}
