package mysql

import (
	"embed"
	migrate "github.com/rubenv/sql-migrate"
	"gorm.io/gorm"
)

//go:embed scripts/*.sql
var migrationScripts embed.FS

type IMigrator interface {
	Up() (int, error)
	Down() (int, error)
}

type migratorImpl struct {
	migrations migrate.MigrationSource
	db         *gorm.DB
}

func NewMigrator(db *gorm.DB, migrationScriptsRoot string) (IMigrator, error) {
	migrations := &migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationScripts,
		Root:       migrationScriptsRoot,
	}

	return &migratorImpl{
		migrations: migrations,
		db:         db,
	}, nil
}

func (m *migratorImpl) Up() (int, error) {
	return m.migrate(migrate.Up)
}

func (m *migratorImpl) Down() (int, error) {
	return m.migrate(migrate.Down)
}

func (m *migratorImpl) migrate(direction migrate.MigrationDirection) (int, error) {
	db, _ := m.db.DB()
	n, err := migrate.Exec(db, "mysql", m.migrations, direction)

	if err != nil {
		return 0, err
	}

	return n, nil
}
