package main

import "database/sql"

type book struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}

func (b *book) createBook(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO books(name, quantity) VALUES($1, $2) RETURNING id", b.Name, b.Quantity).Scan(&b.ID)

	if err != nil {
		return err
	}

	return nil
}

func getBooks(db *sql.DB, start, count int) ([]book, error) {
	rows, err := db.Query(
		"SELECT id, name, quantity FROM books LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	books := []book{}

	for rows.Next() {
		var b book
		if err := rows.Scan(&b.ID, &b.Name, &b.Quantity); err != nil {
			return nil, err
		}
		books = append(books, b)
	}

	return books, nil
}