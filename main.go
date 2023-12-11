package main

import (
	"fmt"
	"github/Danila331/testlyceumbot/server"
)

func main() {
	var botToken string
	fmt.Println("Введите свой токенбот для начала работы")
	fmt.Scan(&botToken)
	server.StartServer(botToken)
}

// package main

// import (
// 	"database/sql"
// 	"fmt"
// 	"log"

// 	_ "modernc.org/sqlite"
// )

// func main() {
// 	// Open a connection to the SQLite database
// 	conn, err := sql.Open("sqlite", "telegram.db")
// 	if err != nil {
// 		log.Fatal("Error opening the database:", err)
// 	}
// 	defer conn.Close()

// 	// Create a table in the database using an SQL query
// 	createTableSQL := `
//         CREATE TABLE IF NOT EXISTS users (
//             id INTEGER PRIMARY KEY AUTOINCREMENT,
//             chatid INTEGER UNIQUE NOT NULL,
//             username TEXT
//         );
//     `
// 	_, err = conn.Exec(createTableSQL)
// 	if err != nil {
// 		log.Fatal("Error creating the table:", err)
// 	}

// 	fmt.Println("Table created successfully!")
// }
