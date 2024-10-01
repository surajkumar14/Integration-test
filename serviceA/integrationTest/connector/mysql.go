package connector

import (
	"database/sql"
)

func CreateRecord(db *sql.DB, id int, name string) error {
	query := "INSERT INTO users (id, name) VALUES (?, ?)"
	_, err := db.Exec(query, id, name)
	return err
}

func ReadRecord(db *sql.DB, id int) (string, error) {
	var name string
	query := "SELECT name FROM users WHERE id = ?"
	err := db.QueryRow(query, id).Scan(&name)
	return name, err
}

func UpdateRecord(db *sql.DB, id int, name string) error {
	query := "UPDATE users SET name = ? WHERE id = ?"
	_, err := db.Exec(query, name, id)
	return err
}

func DeleteRecord(db *sql.DB, id int) error {
	query := "DELETE FROM users WHERE id = ?"
	_, err := db.Exec(query, id)
	return err
}
