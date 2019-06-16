package main

import "database/sql"

type borrowing struct {
	ID     int `json:"id"`
	UserID int `json:"userID"`
	BookID int `json:"bookID"`
}

func (b *borrowing) borrow(db *sql.DB) error {
	// var exists bool
	// db.QueryRow("SELECT name FROM books WHERE id=$1 AND quantity > 0", b.BookID).Scan(&exists)
	// if !exists {
	// 	return sql.ErrNoRows
	// }

	err := db.QueryRow("INSERT INTO borrowings(userID, bookID) VALUES($1, $2) RETURNING id", b.UserID, b.BookID).Scan(&b.ID)

	if err != nil {
		return err
	}

	return nil
}

func getBorrowings(db *sql.DB, start, count int) ([]borrowing, error) {
	rows, err := db.Query(
		"SELECT id, userID, bookID FROM borrowings LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	borrowings := []borrowing{}

	for rows.Next() {
		var b borrowing
		if err := rows.Scan(&b.ID, &b.UserID, &b.BookID); err != nil {
			return nil, err
		}
		borrowings = append(borrowings, b)
	}

	return borrowings, nil
}
