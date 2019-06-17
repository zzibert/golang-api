package main

import (
	"database/sql"
)

type user struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

func (u *user) createUser(db *sql.DB) error {
	err := db.QueryRow("INSERT INTO users(name, surname) VALUES($1, $2) RETURNING id", u.Name, u.Surname).Scan(&u.ID)

	if err != nil {
		return err
	}

	return nil
}

func getUsers(db *sql.DB, start, count int) ([]user, error) {
	rows, err := db.Query(
		"SELECT id, name,  surname FROM users LIMIT $1 OFFSET $2",
		count, start)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	users := []user{}

	for rows.Next() {
		var u user
		if err := rows.Scan(&u.ID, &u.Name, &u.Surname); err != nil {
			return nil, err
		}
		users = append(users, u)
	}

	return users, nil
}
