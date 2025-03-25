package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/kirebyte/thd-project/settings"
	_ "modernc.org/sqlite"
)

func main() {
	cfg := settings.Load()

	//	When the database file does not exist...
	if _, err := os.Stat(cfg.DBPath); os.IsNotExist(err) {
		log.Println("📦 DB not found, initializing...")

		// Create the database file.
		file, err := os.Create(cfg.DBPath)
		if err != nil {
			log.Fatalf("❌ Failed to create DB file: %v", err)
		}
		file.Close()

		// Read the creation script from the file.
		script, err := os.ReadFile(cfg.CreationScript)
		if err != nil {
			log.Fatalf("❌ Failed to read creation script: %v", err)
		}

		// Open the database.
		db, err := sql.Open("sqlite", cfg.DBPath)
		if err != nil {
			log.Fatalf("❌ Failed to open DB: %v", err)
		}
		defer db.Close()

		// Execute the creation script.
		if _, err := db.Exec(string(script)); err != nil {
			log.Fatalf("❌ Failed to apply schema: %v", err)
		}

		// Log success.
		log.Println("✅ DB initialized successfully.")

		//	Exit the program.
		return
	}

	//	Log that the database already exists.
	log.Println("ℹ️  DB already exists. Skipping init.")
}
