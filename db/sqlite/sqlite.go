package sqlite

import "database/sql"

// IsExists 判断表或者索引是否存在
func IsExists(db *sql.DB, name string) (bool, error) {

	stmt, err := db.Prepare("SELECT name FROM sqlite_master WHERE type IN ('table', 'index') and name=?")
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(name)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	return rows.Next(), nil
}
