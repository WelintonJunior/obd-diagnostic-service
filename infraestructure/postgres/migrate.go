package infraestructure

import (
	"errors"

	"github.com/WelintonJunior/obd-diagnostic-service/types"
	"gorm.io/gorm"
)

type PostgresMigrateService struct {
	db *gorm.DB
}

func NewPostgresMigrateService(db *gorm.DB) (*PostgresMigrateService, error) {
	var service PostgresMigrateService

	if db == nil {
		return nil, errors.New("No gorm db passes")
	}

	service.db = db
	return &service, nil
}

func (r *PostgresMigrateService) MigrateApply() error {
	return r.db.AutoMigrate(
		&types.OBDReading{},
		&types.DTCLog{},
		&types.ReadingSession{},
	)
}

func (r *PostgresMigrateService) MigrateRevert() error {
	return r.db.Migrator().DropTable(
		&types.OBDReading{},
		&types.DTCLog{},
		&types.ReadingSession{},
	)
}
