package sqlite

import "database/sql"

//	判断表是否存在
func TableExists(db *sql.DB, name string) (bool, error) {

	stmt, err := db.Prepare("SELECT name FROM sqlite_master WHERE type='table' AND name=?")
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
