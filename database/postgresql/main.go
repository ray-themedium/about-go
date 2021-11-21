package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   "localhost:5432",
		User:   url.UserPassword("leo", "9036"),
		Path:   "golang_pg",
	}

	q := dsn.Query()
	q.Add("sslmode", "disable")
	dsn.RawQuery = q.Encode()

	db, err := sql.Open("pgx", dsn.String())
	if err != nil {
		fmt.Println("sql.Open", err)
		return
	}
	defer func() {
		_ = db.Close()
		fmt.Println("closed")
	}()

	if err := db.PingContext(context.Background()); err != nil {
		fmt.Println("db.PingContext", err)
		return
	}

	// row := db.QueryRow("SELECT birth_year FROM users WHERE name = ?", "leo")
	row := db.QueryRowContext(context.Background(), "SELECT birth_year FROM users WHERE name = 'leo'")
	if row.Err() != nil {
		fmt.Println("db.QueryRow", row.Err().Error())
		return
	}

	var birthYear int64
	if err := row.Scan(&birthYear); err != nil {
		fmt.Println("row.Scan", err)
		return
	}
	fmt.Println(birthYear)

	//-

	name := "bbak"
	birthYear = 1903
	_, err = db.ExecContext(context.Background(), "INSERT INTO users(name, birth_year) VALUES($1, $2)", name, birthYear)
	if err != nil {
		fmt.Println("db.ExecContext", err)
		return
	}

	//-
	rows, err := db.QueryContext(context.Background(), "SELECT * FROM users")
	if err != nil {
		fmt.Println("db.QueryContext", err)
		return
	}
	defer func() {
		_ = rows.Close()
	}()

	if rows.Err() != nil {
		fmt.Println("rows.Err", err)
		return
	}

	for rows.Next() {
		var name string
		var birthYear int64

		if err := rows.Scan(&name, &birthYear); err != nil {
			fmt.Println("row.Scan", err)
			return
		}
		fmt.Println("name:", name, "birth_year:", birthYear)
	}
}
