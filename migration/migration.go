package migration

import (
	"gorm.io/gorm"
	"time"
)

var Db *gorm.DB

type Migration struct {
	gorm.Model
	Version    string     `json:"version"`
	ExecutedAt *time.Time `json:"executed_at"`
}

func MigrateUp(db *gorm.DB) []Migration {
	if db.Migrator().HasTable(Migration{}) {
		versions, _ := GetVersions(db)
		return versions
	}
	db.AutoMigrate(Migration{})
	return nil
}

func GetVersions(db *gorm.DB) ([]Migration, error) {
	var migrations []Migration
	results := db.Model(&Migration{}).Find(&migrations)

	return migrations, results.Error
}

func VersionExists(currentVersions []Migration, version string) bool {
	for _, currentVersion := range currentVersions {
		if currentVersion.Version == version {
			return true
		}
	}
	return false
}

func AddVersion(db *gorm.DB, version string) error {
	now := time.Now()
	migration := Migration{}
	migration.Version = version
	migration.ExecutedAt = &now
	results := db.Create(&migration)

	return results.Error

}
