package database

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"sort"
)

func RunMigrations() error {
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS migrations (
			id SERIAL PRIMARY KEY,
			filename VARCHAR(255) UNIQUE NOT NULL,
			executed_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create migrations table: %v", err)
	}

	files, err := filepath.Glob("migrations/*.sql")
	if err != nil {
		return fmt.Errorf("failed to read migration files: %v", err)
	}
	sort.Strings(files)

	for _, file := range files {
		filename := filepath.Base(file)

		var count int
		err := DB.QueryRow("SELECT COUNT(*) FROM migrations WHERE filename = $1", filename).Scan(&count)
		if err != nil {
			return fmt.Errorf("failed to check migration status: %v", err)
		}

		if count > 0 {
			continue
		}

		content, err := ioutil.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %v", filename, err)
		}

		_, err = DB.Exec(string(content))
		if err != nil {
			return fmt.Errorf("failed to execute migration %s: %v", filename, err)
		}

		_, err = DB.Exec("INSERT INTO migrations (filename) VALUES ($1)", filename)
		if err != nil {
			return fmt.Errorf("failed to record migration %s: %v", filename, err)
		}

		log.Printf("Migration %s executed", filename)
	}

	return nil
}
