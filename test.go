package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/bet365/go-mssqldb-auth-krb5"
	_ "github.com/microsoft/go-mssqldb"
)

func test() {
	var (
		connectionString string
	)

	flag.StringVar(&connectionString, "connString", "", "Connection string")
	flag.Parse()

	// when the connection is opened it will use the krb5 Auth Provider created above.
	db, err := sql.Open("sqlserver", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	sql := "select 1234"

	var value string
	err = db.QueryRowContext(ctx, sql).Scan(&value)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(value)
}
