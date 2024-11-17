package repository

import (
	"database/sql"
	"fmt"
	"quiz-sb-pekan3/structs"

	_ "github.com/lib/pq"
)

// GetUserByUsername fetches a user from the database by username
func GetUserByUsername(db *sql.DB, username string) (*structs.User, error) {
	var user structs.User
	err := db.QueryRow("SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("no user found with that username")
		}
		return nil, err
	}
	return &user, nil
}
