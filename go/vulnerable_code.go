package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		// ユーザー入力を取得
		userID := r.URL.Query().Get("id")

		// データベース接続を開く
		db, err := sql.Open("mysql", "user:password@/dbname")
		if err != nil {
			http.Error(w, "Database connection failed", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		// SQLインジェクションの脆弱性を含むクエリ
		query := fmt.Sprintf("SELECT name FROM users WHERE id = '%s'", userID)
		rows, err := db.Query(query)
		if err != nil {
			http.Error(w, "Query failed", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var name string
			if err := rows.Scan(&name); err != nil {
				http.Error(w, "Error scanning rows", http.StatusInternalServerError)
				return
			}
			fmt.Fprintf(w, "Name: %s\n", name)
		}
	})

	http.ListenAndServe(":8080", nil)
}
