package pkg

import "database/sql"

// truncateUsers as truncate table users
func TruncateUsers(db *sql.DB) {
	db.Exec("TRUNCATE users")
}
