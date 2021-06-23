package postgresql

import (
	"database/sql"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/gopaytech/go-commons/pkg/db"
	"github.com/gopaytech/go-commons/pkg/db/migration"
)

func New(config db.Config, path string) (mgrn migration.Migration, err error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%v/%s", config.Username, config.Password, config.Host, config.Port, config.DatabaseName)
	m, err := migrate.New(path, dsn)
	if err != nil {
		return
	}

	mgrn = &migration.Migrator{
		Migrate: m,
	}

	return
}

func WithInstance(db *sql.DB, path string) (mgrn migration.Migration, err error) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	m, _ := migrate.NewWithDatabaseInstance(
		path,
		"mysql",
		driver,
	)

	if err != nil {
		return
	}

	mgrn = &migration.Migrator{
		Migrate: m,
	}

	return
}
