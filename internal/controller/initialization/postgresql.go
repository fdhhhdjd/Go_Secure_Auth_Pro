package initialization

import (
	"database/sql"
	"fmt"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs"
	_ "github.com/lib/pq"
)

// ConnectPG establishes a connection to a PostgreSQL database using the provided configuration.
// It takes a `cfg` parameter of type `configs.Config` which contains the necessary database connection details.
// It returns a pointer to `sql.DB` and an error. If the connection is successful, the error will be `nil`.
// Otherwise, it will return an error indicating the reason for the failure.
func ConnectPG(cfg configs.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", cfg.Database.Username, cfg.Database.Name, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("SUCCESSFULLY CONNECTED POSTGRESQL 🐘!")
	return db, nil
}
