package model

import (
	"database/sql"
	"log"
)

func CheckUser(con *sql.DB, username string, password string) bool {
	query := `SELECT COUNT(*) FROM users WHERE username = ? AND password = ?`
	//query = //$username alanına username gelicek. $password alanına da paswword gelecek

	// Sorguyu çalıştırın ve sonucu kontrol edin
	var count int
	err := con.QueryRow(query, username, password).Scan(&count)
	if err != nil {
		log.Println("Error executing query:", err)
		return false
	}

	// Kullanıcı bulunmuşsa count 1 olacaktır
	return count > 0
}
