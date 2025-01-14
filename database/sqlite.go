package database

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gobackup/gobackup/helper"
	"github.com/gobackup/gobackup/logger"
)

// SQLite database
//
// type: sqlite
// path:
type SQLite struct {
	Base
	path     string
	database string

	_dumpFilePath string
}

func (db *SQLite) init() error {
	viper := db.viper

	db.path = helper.ExplandHome(viper.GetString("path"))
	db.database = strings.TrimSuffix(filepath.Base(db.path), filepath.Ext(db.path))

	db._dumpFilePath = filepath.Join(db.dumpPath, db.database+".sql")

	return nil
}

func (db *SQLite) buildArgs() []string {
	args := []string{
		db.path,
		fmt.Sprintf(".dump > %s", db._dumpFilePath),
	}
	return args
}

func (db *SQLite) perform() error {
	logger := logger.Tag("SQLite")

	logger.Info("-> Dumping SQLite...")
	if _, err := helper.Exec("sqlite3", db.buildArgs()...); err != nil {
		return err
	}

	logger.Info("dump path:", db._dumpFilePath)
	return nil
}
