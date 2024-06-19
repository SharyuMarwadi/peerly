package corevalues

import (
	"context"
	"fmt"
	"joshsoftware/peerly/apperrors"
	"joshsoftware/peerly/db"
	"joshsoftware/peerly/pkg/dto"
	"strconv"

	logger "github.com/sirupsen/logrus"
)

func validateParentCoreValue(ctx context.Context, storer db.CoreValueStorer, organisationID, coreValueID int64) (ok bool) {
	coreValue, err := storer.GetCoreValue(ctx, organisationID, coreValueID)
	if err != nil {
		logger.WithField("err", err.Error()).Error("Parent core value id not present")
		return
	}

	if coreValue.ParentID != nil || coreValue.SoftDelete {
		logger.Error("Invalid parent core value id")
		return
	}

	return true
}

// Validate - ensures the core value object has all the info it needs
// func Validate(ctx context.Context, coreValue db.CoreValue, storer db.CoreValueStorer, organisationID int64) (valid bool, errFields map[string]string) {
// 	errFields = make(map[string]string)

// 	if coreValue.Text == "" {
// 		errFields["text"] = "Can't be blank"
// 	}
// 	if coreValue.Description == "" {
// 		errFields["description"] = "Can't be blank"
// 	}
// 	if coreValue.ParentID != nil {
// 		if !validateParentCoreValue(ctx, storer, organisationID, *coreValue.ParentID) {
// 			errFields["parent_id"] = "Invalid parent core value"
// 		}
// 	}

// 	if len(errFields) == 0 {
// 		valid = true
// 	}
// 	return
// }

func Validate(ctx context.Context, coreValue dto.CreateCoreValueReq, storer db.CoreValueStorer, organisationID int64) (err error) {

	if coreValue.Text == "" {
		err = apperrors.TextFieldBlank
	}
	if coreValue.Description == "" {
		err = apperrors.DescFieldBlank
	}
	if coreValue.ParentID != nil {
		if !validateParentCoreValue(ctx, storer, organisationID, *coreValue.ParentID) {
			err = apperrors.InvalidParentValue
		}
	}

	return
}

func VarsStringToInt(inp string, label string) (result int64, err error) {

	if len(inp) <= 0 {
		err = apperrors.InvalidOrgId
		return
	}
	result, err = strconv.ParseInt(inp, 10, 64)
	if err != nil {
		logger.WithField("err", err.Error()).Error(fmt.Scanf("Error while parsing %s from url", label))
		err = apperrors.InternalServerError
		return

	}

	return
}
