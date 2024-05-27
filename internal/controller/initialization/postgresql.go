package initialization

import (
	"database/sql"
	"fmt"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs"
	_ "github.com/lib/pq"
)

func Connect(cfg configs.Config) (*sql.DB, error) {
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable", cfg.Database.Username, cfg.Database.Name, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected! 🐘")
	return db, nil
}
