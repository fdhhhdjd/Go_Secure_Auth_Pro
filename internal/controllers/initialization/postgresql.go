package initialization

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/internal/models"
	_ "github.com/lib/pq"
)

// ConnectPG establishes a connection to a PostgreSQL database using the provided configuration.
// It takes a `cfg` parameter of type `configs.Config` which contains the necessary database connection details.
// It returns a pointer to `sql.DB` and an error. If the connection is successful, the error will be `nil`.
// Otherwise, it will return an error indicating the reason for the failure.
func ConnectPG(cfg models.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", cfg.Database.Username, cfg.Database.Name, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port)

	var db *sql.DB
	var err error

	maxRetries := 10
	for i := 0; i < maxRetries; i++ {
		db, err = sql.Open("postgres", connStr)
		if err != nil {
			fmt.Printf("Error connecting to database: %v\n", err)
			fmt.Println("Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}

		err = db.Ping()
		if err != nil {
			fmt.Printf("Error pinging database: %v\n", err)
			fmt.Println("Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}

		fmt.Println("SUCCESSFULLY CONNECTED POSTGRESQL 🐘!")
		return db, nil
	}

	return nil, fmt.Errorf("could not connect to database after %d attempts", maxRetries)
}
