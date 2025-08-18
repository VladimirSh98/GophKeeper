package database

import (
	"github.com/pressly/goose/v3"
)

// UpgradeMigrations applies migrations to database
func (db *DBConnectionStruct) UpgradeMigrations() error {
	err := goose.Up(db.Conn, db.Cfg.MigrationsDir)
	if err != nil {
		return err
	}
	return nil
}
