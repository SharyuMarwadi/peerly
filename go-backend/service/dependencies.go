package service

import (
	// "joshsoftware/peerly/aws"
	"joshsoftware/peerly/db"
	DB "joshsoftware/peerly/db"
	corevalues "joshsoftware/peerly/service/coreValues"

	"github.com/jmoiron/sqlx"
)

// Dependencies - Stuff we need for the service package
type Dependencies struct {
	Store DB.Storer
	// AWSStore aws.AWSStorer
	// define other service dependencies
	CoreValueService corevalues.Service
}

func NewServices(db *sqlx.DB, store db.Storer) Dependencies {
	coreValueRepo := DB.NewCoreValueRepo(db)
	coreValueService := corevalues.NewService(coreValueRepo)

	return Dependencies{
		Store:            store,
		CoreValueService: coreValueService,
	}
}
