package database

import "database/sql"

func CheckHealth() error {
	if DB == nil {
		return sql.ErrConnDone
	}
	return DB.Ping()
}
