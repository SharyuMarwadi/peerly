package corevalues

import (
	"context"

	"joshsoftware/peerly/apperrors"
	"joshsoftware/peerly/db"
	"joshsoftware/peerly/pkg/dto"

	logger "github.com/sirupsen/logrus"
)

type service struct {
	coreValuesRepo db.CoreValueStorer
}

type Service interface {
	ListCoreValues(ctx context.Context, organisationID string) (coreValues []db.CoreValue, err error)
	GetCoreValue(ctx context.Context, organisationID string, coreValueID string) (coreValue db.CoreValue, err error)
	CreateCoreValue(ctx context.Context, organisationID string, userId int64, coreValue db.CoreValue) (resp db.CoreValue, err error)
	DeleteCoreValue(ctx context.Context, organisationID string, coreValueID string) (err error)
	UpdateCoreValue(ctx context.Context, organisationID string, coreValueID string, coreValue dto.UpdateQueryRequest) (resp db.CoreValue, err error)
}

func NewService(coreValuesRepo db.CoreValueStorer) Service {
	return &service{
		coreValuesRepo: coreValuesRepo,
	}
}

func (cs *service) ListCoreValues(ctx context.Context, organisationID string) (coreValues []db.CoreValue, err error) {

	orgId, err := VarsStringToInt(organisationID, "organisationId")
	if err != nil {
		return
	}

	err = cs.coreValuesRepo.CheckOrganisation(ctx, orgId)
	if err != nil {
		return
	}

	coreValues, err = cs.coreValuesRepo.ListCoreValues(ctx, orgId)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error while fetching data")
		err = apperrors.InernalServerError
	}

	return

}

func (cs *service) GetCoreValue(ctx context.Context, organisationID string, coreValueID string) (coreValue db.CoreValue, err error) {
	// orgId, err := strconv.ParseInt(organisationID, 10, 64)
	// if err != nil {
	// 	logger.WithField("err", err.Error()).Error("Error while parsing organisation_id from url")
	// 	msgObj = dto.MessageObject{
	// 		Message: "Internal server error",
	// 	}
	// 	return
	// }

	orgId, err := VarsStringToInt(organisationID, "organisationId")
	if err != nil {
		return
	}

	err = cs.coreValuesRepo.CheckOrganisation(ctx, orgId)
	if err != nil {
		return
	}

	// coreValId, err := strconv.ParseInt(coreValueID, 10, 64)
	// if err != nil {
	// 	logger.WithField("err", err.Error()).Error("Error while parsing coreValueId from url")
	// 	msgObj = dto.MessageObject{
	// 		Message: "Internal server error",
	// 	}
	// 	return
	// }

	coreValId, err := VarsStringToInt(coreValueID, "coreValueId")
	if err != err {
		return
	}

	coreValue, err = cs.coreValuesRepo.GetCoreValue(ctx, orgId, coreValId)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error while fetching data")
		return
	}

	return
}

func (cs *service) CreateCoreValue(ctx context.Context, organisationID string, userId int64, coreValue db.CoreValue) (resp db.CoreValue, err error) {

	orgId, err := VarsStringToInt(organisationID, "organisationId")
	if err != nil {
		return
	}

	err = cs.coreValuesRepo.CheckOrganisation(ctx, orgId)
	if err != nil {
		return
	}

	err = Validate(ctx, coreValue, cs.coreValuesRepo, orgId)
	if err != nil {
		return
	}

	resp, err = cs.coreValuesRepo.CreateCoreValue(ctx, orgId, userId, coreValue)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error while creating core value")
		err = apperrors.InernalServerError
		return
	}

	return
}

func (cs *service) DeleteCoreValue(ctx context.Context, organisationID string, coreValueID string) (err error) {

	// organisationID, err := strconv.ParseInt(vars["organisation_id"], 10, 64)
	// if err != nil {
	// 	logger.WithField("err", err.Error()).Error("Error while parsing organisation_id from url")
	// 	rw.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	orgId, err := VarsStringToInt(organisationID, "organisationId")
	if err != nil {
		return
	}

	err = cs.coreValuesRepo.CheckOrganisation(ctx, orgId)
	if err != nil {
		return
	}

	// coreValueID, err := strconv.ParseInt(vars["id"], 10, 64)
	// if err != nil {
	// 	logger.WithField("err", err.Error()).Error("Error while parsing core value id from url")
	// 	rw.WriteHeader(http.StatusBadRequest)
	// 	return
	// }

	coreValId, err := VarsStringToInt(coreValueID, "coreValueId")
	if err != nil {
		return
	}

	_, err = cs.coreValuesRepo.GetCoreValue(ctx, orgId, coreValId)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error while fetching data")
		return
	}

	err = cs.coreValuesRepo.DeleteCoreValue(ctx, orgId, coreValId)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error while deleting core value")
		err = apperrors.InernalServerError

		return
	}

	return
}

func (cs *service) UpdateCoreValue(ctx context.Context, organisationID string, coreValueID string, reqData dto.UpdateQueryRequest) (resp db.CoreValue, err error) {

	orgId, err := VarsStringToInt(organisationID, "organisationId")
	if err != nil {
		return
	}

	coreValId, err := VarsStringToInt(coreValueID, "coreValueId")
	if err != nil {
		return
	}

	//validate organisation
	err = cs.coreValuesRepo.CheckOrganisation(ctx, orgId)
	if err != nil {
		return
	}

	//validate corevalue
	//get data
	coreValue, err := cs.coreValuesRepo.GetCoreValue(ctx, orgId, coreValId)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error while fetching data")
		return
	}

	//set empty fields
	if reqData.Text == "" {
		reqData.Text = coreValue.Text
	}
	if reqData.Description == "" {
		reqData.Description = coreValue.Description
	}
	if reqData.ThumbnailUrl == "" {
		reqData.ThumbnailUrl = *coreValue.ThumbnailURL
	}

	resp, err = cs.coreValuesRepo.UpdateCoreValue(ctx, orgId, coreValId, reqData)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Error while updating core value")
		err = apperrors.InernalServerError

		return
	}

	return
}
