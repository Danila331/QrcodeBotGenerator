package store

import (
	"database/sql"
	"fmt"
	"github/Danila331/testlyceumbot/models"

	_ "modernc.org/sqlite"
)

type DB struct {
	database *sql.DB
}

func NewDB(dataSourceName string) (*DB, error) {
	db, err := sql.Open("sqlite", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &DB{
		database: db,
	}, nil
}

func (db *DB) Close() error {
	return db.database.Close()
}

func (db *DB) SearchByChatid(chatid int) (models.User, error) {
	var column3 string
	var column1, column2 int
	sqliteQuery := fmt.Sprintf("SELECT * FROM users WHERE chatid=%d", chatid)
	err := db.database.QueryRow(sqliteQuery).Scan(&column1, &column2, &column3)
	if err != nil {
		return models.User{}, err
	}
	return models.User{
		Id:       column1,
		ChatId:   column2,
		UserName: column3,
	}, nil
}

func (db *DB) CreateUser(chatid int, username string) error {
	sqliteQuery := fmt.Sprintf("INSERT INTO users(chatid, username) VALUES (%d, '%s')", chatid, username)
	_, err := db.database.Exec(sqliteQuery)
	if err != nil {
		return err
	}
	return nil
}
