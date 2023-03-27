package migration

import (
	"database/sql"

	"meta-aggregator/internal/pkg/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Migration interface {
	MigrateUp(up int) error
	MigrateDown(down int) error
}

type MySQLMigration struct {
	m *migrate.Migrate
}

func NewMigration(cfg *config.DBConfig) (Migration, error) {
	db, _ := sql.Open("mysql", cfg.MigrationSource())
	driver, error := mysql.WithInstance(db, &mysql.Config{})
	if error != nil {
		return nil, error
	}

	m, error := migrate.NewWithDatabaseInstance(
		"file://./migrations/mysql",
		"mysql",
		driver,
	)

	if error != nil {
		return nil, error
	}

	return &MySQLMigration{
		m,
	}, nil

}

func (t *MySQLMigration) MigrateUp(up int) error {
	if up == 0 {
		return t.m.Up()
	} else {
		return t.m.Steps(up)
	}
}

func (t *MySQLMigration) MigrateDown(down int) error {
	if down == 0 {
		return t.m.Down()
	} else {
		return t.m.Steps(-down)
	}
}
