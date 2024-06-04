package helpers

import (
	"database/sql"
	"strconv"
)

// NullStringToString converts a sql.NullString to a string.
// If the sql.NullString is valid, it returns the string value.
// Otherwise, it returns an empty string.
func NullStringToString(ns sql.NullString) string {
	if ns.Valid {
		return ns.String
	}
	return ""
}

// NullInt16ToString converts a sql.NullInt16 to an int.
// If the sql.NullInt16 is valid, it returns the int value.
// Otherwise, it returns 0.
func NullInt16ToString(ni sql.NullInt16) int {
	if ni.Valid {
		result, _ := strconv.ParseInt(strconv.FormatInt(int64(ni.Int16), 10), 10, 64)
		return int(result) // Convert result to int before returning
	}
	return 0
}
