package main

import (
	"database/sql"
	"errors"
)

type borrowing struct {
	ID     int `json:"id"`
	UserID int `json:"userID"`
	BookID int `json:"bookID"`
}

type borrowingWithNames struct {
	ID          int    `json:"id"`
	UserName    string `json:"userName"`
	UserSurname string `json:"userSurname"`
	BookName    string `json:"bookName"`
}

func (b *borrowing) borrow(db *sql.DB) error {
	var exists string
	err := db.QueryRow("SELECT name FROM books WHERE id=$1 AND quantity > 0", b.BookID).Scan(&exists)
	if err != nil {

		return errors.New("This book is currently not available.")
	}
	_, err = db.Exec("UPDATE books SET quantity = quantity - 1 WHERE id=$1", b.BookID)
	err = db.QueryRow("INSERT INTO borrowings(userID, bookID) VALUES($1, $2) RETURNING id", b.UserID, b.BookID).Scan(&b.ID)

	if err != nil {
		return err
	}

	return nil
}

func (b *borrowing) unborrow(db *sql.DB) error {
	_, err := db.Exec("UPDATE books SET quantity = quantity + 1 WHERE id=$1", b.BookID)
	_, err = db.Exec("DELETE FROM borrowings WHERE userID=$1 AND bookID=$2", b.UserID, b.BookID)

	if err != nil {
		return err
	}

	return nil
}

func getBorrowings(db *sql.DB, start, count int) ([]borrowingWithNames, error) {
	rows, err := db.Query(
		"SELECT id, userID, bookID FROM borrowings LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	borrowings := []borrowingWithNames{}

	for rows.Next() {
		var b borrowing
		var bw borrowingWithNames
		if err := rows.Scan(&b.ID, &b.UserID, &b.BookID); err != nil {
			return nil, err
		}
		db.QueryRow("SELECT name, surname FROM users WHERE id=$1", b.UserID).Scan(&bw.UserName, &bw.UserSurname)
		db.QueryRow("SELECT name FROM books WHERE id=$1", b.BookID).Scan(&bw.BookName)
		bw.ID = b.ID

		borrowings = append(borrowings, bw)
	}

	return borrowings, nil
}
