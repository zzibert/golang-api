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
