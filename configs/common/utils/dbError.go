package utils

import (
	"log"

	"github.com/fdhhhdjd/Go_Secure_Auth_Pro/configs/common/constants"
	"github.com/lib/pq"
)

func HandleDBError(err error) string {
	if pqErr, ok := err.(*pq.Error); ok {
		log.Printf("PostgreSQL Error Code: %s", pqErr.Code)

		switch pqErr.Code {
		case constants.ForeignKeyViolation:
			return "Foreign key violation"
		case constants.UniqueViolation:
			return "Unique violation"
		case constants.NotNullViolation:
			return "Not null violation"
		default:
			return "Unknown error"
		}

	}
	return ""
}
