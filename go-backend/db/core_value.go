package db

import (
	"context"
	"joshsoftware/peerly/apperrors"
	"joshsoftware/peerly/pkg/dto"
	"time"

	"github.com/jmoiron/sqlx"
	logger "github.com/sirupsen/logrus"
)

type coreValueStore struct {
	DB *sqlx.DB
}

type CoreValueStorer interface {
	ListCoreValues(ctx context.Context, organisationID int64) (coreValues []CoreValue, err error)
	GetCoreValue(ctx context.Context, organisationID, coreValueID int64) (coreValue CoreValue, err error)
	CreateCoreValue(ctx context.Context, organisationID int64, userId int64, coreValue CoreValue) (resp CoreValue, err error)
	DeleteCoreValue(ctx context.Context, organisationID, coreValueID int64) (err error)
	UpdateCoreValue(ctx context.Context, organisationID, coreValueID int64, coreValue dto.UpdateQueryRequest) (resp CoreValue, err error)
	CheckOrganisation(ctx context.Context, organisationId int64) (err error)
}

func NewCoreValueRepo(db *sqlx.DB) CoreValueStorer {
	return &coreValueStore{
		DB: db,
	}
}

const (
	listCoreValuesQuery  = `SELECT id, orgId, text, description, parentId  FROM coreValues WHERE orgId = $1`
	getCoreValueQuery    = `SELECT id, orgId, text, description, parentId FROM coreValues WHERE orgId = $1 and id = $2`
	createCoreValueQuery = `INSERT INTO coreValues (orgId, text,
		description, parentId, thumbnailUrl,created_by, createdAt, updatedAt) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id, orgId, text, description, parentId, thumbnailUrl`
	// deleteSubCoreValueQuery = `DELETE FROM coreValues WHERE orgId = $1 and parentId = $2`
	deleteSubCoreValueQuery = `UPDATE coreValues SET soft_delete = true WHERE orgId = $1 and parentId = $2`
	// deleteCoreValueQuery    = `DELETE FROM coreValues WHERE orgId = $1 and id = $2`
	deleteCoreValueQuery = `UPDATE coreValues SET soft_delete = true WHERE orgId = $1 and id = $2`
	updateCoreValueQuery = `UPDATE coreValues SET (text, description, thumbnailurl, updatedAt) =
		($1, $2, $3, $4) where id = $5 and orgId = $6 RETURNING id, orgId, text, description, parentId`

	checkOrganisationQuery = `SELECT id from organizations WHERE id = $1`
)

// CoreValue - struct representing a core value object
type CoreValue struct {
	ID           int64     `db:"id" json:"id"`
	OrgID        int64     `db:"orgid" json:"org_id"`
	Text         string    `db:"text" json:"text"`
	Description  string    `db:"description" json:"description"`
	ParentID     *int64    `db:"parentid" json:"parent_id"`
	ThumbnailURL *string   `db:"thumbnailurl" json:"thumbnail_url"`
	CreatedAt    time.Time `db:"createdat" json:"-"`
	UpdatedAt    time.Time `db:"updatedat" json:"-"`
}

// // This function goes in service!
// func validateParentCoreValue(ctx context.Context, storer Storer, organisationID, coreValueID int64) (ok bool) {
// 	coreValue, err := storer.GetCoreValue(ctx, organisationID, coreValueID)
// 	if err != nil {
// 		logger.WithField("err", err.Error()).Error("Parent core value id not present")
// 		return
// 	}

// 	if coreValue.ParentID != nil {
// 		logger.Error("Invalid parent core value id")
// 		return
// 	}

// 	return true
// }

// // This function goes to the service helper helper
// // Validate - ensures the core value object has all the info it needs
// func (coreValue CoreValue) Validate(ctx context.Context, storer Storer, organisationID int64) (valid bool, errFields map[string]string) {
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

func (cs *coreValueStore) ListCoreValues(ctx context.Context, organisationID int64) (coreValues []CoreValue, err error) {
	coreValues = make([]CoreValue, 0)
	err = cs.DB.SelectContext(
		ctx,
		&coreValues,
		listCoreValuesQuery,
		organisationID,
	)

	if err != nil {
		logger.WithFields(logger.Fields{
			"err":   err.Error(),
			"orgId": organisationID,
		}).Error("Error while getting core values")
		return
	}

	return
}

func (cs *coreValueStore) GetCoreValue(ctx context.Context, organisationID, coreValueID int64) (coreValue CoreValue, err error) {
	err = cs.DB.GetContext(
		ctx,
		&coreValue,
		getCoreValueQuery,
		organisationID,
		coreValueID,
	)
	if err != nil {
		logger.WithFields(logger.Fields{
			"err":         err.Error(),
			"orgId":       organisationID,
			"coreValueId": coreValueID,
		}).Error("Error while getting core value")
		err = apperrors.InvalidCoreValueData
		return
	}

	return
}

func (cs *coreValueStore) CreateCoreValue(ctx context.Context, organisationID int64, userId int64, coreValue CoreValue) (resp CoreValue, err error) {
	now := time.Now()
	err = cs.DB.GetContext(
		ctx,
		&resp,
		createCoreValueQuery,
		organisationID,
		coreValue.Text,
		coreValue.Description,
		coreValue.ParentID,
		coreValue.ThumbnailURL,
		userId,
		now,
		now,
	)
	if err != nil {
		logger.WithFields(logger.Fields{
			"err":               err.Error(),
			"org_id":            organisationID,
			"core_value_params": coreValue,
		}).Error("Error while creating core value")
		return
	}

	return
}

func (cs *coreValueStore) DeleteCoreValue(ctx context.Context, organisationID, coreValueID int64) (err error) {
	_, err = cs.DB.ExecContext(
		ctx,
		deleteSubCoreValueQuery,
		organisationID,
		coreValueID,
	)
	if err != nil {
		logger.WithFields(logger.Fields{
			"err":         err.Error(),
			"orgId":       organisationID,
			"coreValueId": coreValueID,
		}).Error("Error while deleting sub core value")
		return
	}

	_, err = cs.DB.ExecContext(
		ctx,
		deleteCoreValueQuery,
		organisationID,
		coreValueID,
	)
	if err != nil {
		logger.WithFields(logger.Fields{
			"err":           err.Error(),
			"org_id":        organisationID,
			"core_value_id": coreValueID,
		}).Error("Error while deleting core value")
		return
	}

	return
}

func (cs *coreValueStore) UpdateCoreValue(ctx context.Context, organisationID int64, coreValueID int64, updateReq dto.UpdateQueryRequest) (resp CoreValue, err error) {
	now := time.Now()
	err = cs.DB.GetContext(
		ctx,
		&resp,
		updateCoreValueQuery,
		updateReq.Text,
		updateReq.Description,
		updateReq.ThumbnailUrl,
		now,
		coreValueID,
		organisationID,
	)
	if err != nil {
		logger.WithFields(logger.Fields{
			"err":           err.Error(),
			"org_id":        organisationID,
			"core_value_id": coreValueID,
		}).Error("Error while updating core value")
		return
	}

	return
}

func (cs *coreValueStore) CheckOrganisation(ctx context.Context, organisationId int64) (err error) {
	resp := []int64{}
	err = cs.DB.SelectContext(
		ctx,
		&resp,
		checkOrganisationQuery,
		organisationId,
	)

	if len(resp) <= 0 {
		err = apperrors.InvalidOrgId
	}

	if err != nil {
		logger.WithFields(logger.Fields{
			"err":    err.Error(),
			"org_id": organisationId,
		}).Error("Error while checking organisation")
		err = apperrors.InvalidOrgId
	}

	return
}
